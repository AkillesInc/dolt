// Copyright 2019 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This file incorporates work covered by the following copyright and
// permission notice:
//
// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package merge

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/dolthub/dolt/go/store/d"
	"github.com/dolthub/dolt/go/store/types"
)

func threeWayListMerge(ctx context.Context, a, b, parent types.List) (merged types.List, err error) {
	aSpliceChan, bSpliceChan := make(chan types.Splice), make(chan types.Splice)

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		defer close(aSpliceChan)
		return a.Diff(ctx, parent, aSpliceChan)
	})
	eg.Go(func() error {
		defer close(bSpliceChan)
		return b.Diff(ctx, parent, bSpliceChan)
	})

	// The algorithm below relies on determining whether one splice "comes
	// before" another, and whether the splices coming from the two diffs
	// remove/add precisely the same elements. Unfortunately, the Golang
	// zero-value for types.Splice (which is what gets read out of
	// a/bSpliceChan when they're closed) is actaually a valid splice,
	// albeit a meaningless one that indicates a no-op. It "comes before"
	// any other splice, so having it in play really gums up the logic
	// below. Rather than specifically checking for it all over the place,
	// swap the zero-splice out for one full of SPLICE_UNASSIGNED, which is
	// really the proper invalid splice value. That splice doesn't come
	// before ANY valid splice, so the logic below can flow more clearly.
	zeroSplice := types.Splice{}
	zeroToInvalid := func(sp types.Splice) types.Splice {
		if sp == zeroSplice {
			return types.Splice{SpAt: types.SPLICE_UNASSIGNED, SpRemoved: types.SPLICE_UNASSIGNED, SpAdded: types.SPLICE_UNASSIGNED, SpFrom: types.SPLICE_UNASSIGNED}
		}
		return sp
	}
	invalidSplice := zeroToInvalid(types.Splice{})

	merged = parent
	eg.Go(func() error {
		offset := uint64(0)
		aSplice, bSplice := invalidSplice, invalidSplice
		for {
			// Get the next splice from both a and b. If either diff(a,
			// parent) or diff(b, parent) is complete, aSplice or bSplice
			// will get an invalid types.Splice. Generally, though, this
			// allows us to proceed through both diffs in (index) order,
			// considering the "current" splice from both diffs at the same
			// time.
			if aSplice == invalidSplice {
				select {
				case a := <-aSpliceChan:
					aSplice = zeroToInvalid(a)
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			if bSplice == invalidSplice {
				select {
				case b := <-bSpliceChan:
					bSplice = zeroToInvalid(b)
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			// Both channels are producing zero values, so we're done.
			if aSplice == invalidSplice && bSplice == invalidSplice {
				break
			}
			if overlap(aSplice, bSplice) {
				if mergeable, err := canMerge(ctx, a, b, aSplice, bSplice); err != nil {
					return err
				} else if mergeable {
					splice := merge(aSplice, bSplice)
					merged, err = apply(ctx, a, merged, offset, splice)
					if err != nil {
						return err
					}

					offset += splice.SpAdded - splice.SpRemoved
					aSplice, bSplice = invalidSplice, invalidSplice
					continue
				}
				return newMergeConflict("Overlapping splices: %s vs %s", describeSplice(aSplice), describeSplice(bSplice))
			}
			if aSplice.SpAt < bSplice.SpAt {
				merged, err = apply(ctx, a, merged, offset, aSplice)
				if err != nil {
					return err
				}

				offset += aSplice.SpAdded - aSplice.SpRemoved
				aSplice = invalidSplice
				continue
			}
			merged, err = apply(ctx, b, merged, offset, bSplice)
			if err != nil {
				return err
			}

			offset += bSplice.SpAdded - bSplice.SpRemoved
			bSplice = invalidSplice
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return types.EmptyList, err
	}

	return merged, nil
}

func overlap(s1, s2 types.Splice) bool {
	earlier, later := s1, s2
	if s2.SpAt < s1.SpAt {
		earlier, later = s2, s1
	}
	return s1.SpAt == s2.SpAt || earlier.SpAt+earlier.SpRemoved > later.SpAt
}

// canMerge returns whether aSplice and bSplice can be merged into a single splice that can be applied to parent. Currently, we're only willing to do this if the two splices do _precisely_ the same thing -- that is, remove the same number of elements from the same starting index and insert the exact same list of new elements.
func canMerge(ctx context.Context, a, b types.List, aSplice, bSplice types.Splice) (bool, error) {
	if aSplice != bSplice {
		return false, nil
	}
	aIter, err := a.IteratorAt(ctx, aSplice.SpFrom)

	if err != nil {
		return false, err
	}

	bIter, err := b.IteratorAt(ctx, bSplice.SpFrom)

	if err != nil {
		return false, err
	}

	for count := uint64(0); count < aSplice.SpAdded; count++ {
		aVal, err := aIter.Next(ctx)

		if err != nil {
			return false, err
		}

		bVal, err := bIter.Next(ctx)

		if err != nil {
			return false, err
		}

		if aVal == nil || bVal == nil || !aVal.Equals(bVal) {
			return false, nil
		}
	}
	return true, nil
}

// Since merge() is only called when canMerge() is true, we know s1 and s2 are exactly equal.
func merge(s1, s2 types.Splice) types.Splice {
	return s1
}

func apply(ctx context.Context, source, target types.List, offset uint64, s types.Splice) (types.List, error) {
	toAdd := make([]types.Valuable, s.SpAdded)
	iter, err := source.IteratorAt(ctx, s.SpFrom)

	if err != nil {
		return types.EmptyList, err
	}

	for i := 0; uint64(i) < s.SpAdded; i++ {
		v, err := iter.Next(ctx)

		if err != nil {
			return types.EmptyList, err
		}

		if v == nil {
			d.Panic("List diff returned a splice that inserts a nonexistent element.")
		}
		toAdd[i] = v
	}
	return target.Edit().Splice(s.SpAt+offset, s.SpRemoved, toAdd...).List(ctx)
}

func describeSplice(s types.Splice) string {
	return fmt.Sprintf("%d elements removed at %d; adding %d elements", s.SpRemoved, s.SpAt, s.SpAdded)
}
