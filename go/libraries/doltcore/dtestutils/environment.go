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

package dtestutils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dolthub/dolt/go/libraries/doltcore/doltdb"
	"github.com/dolthub/dolt/go/libraries/doltcore/env"
	"github.com/dolthub/dolt/go/libraries/doltcore/table"
	"github.com/dolthub/dolt/go/libraries/doltcore/table/editor"
	"github.com/dolthub/dolt/go/libraries/doltcore/table/typed/noms"
	"github.com/dolthub/dolt/go/libraries/utils/filesys"
	"github.com/dolthub/dolt/go/store/types"
)

const (
	TestHomeDir = "/user/bheni"
	WorkingDir  = "/user/bheni/datasets/states"
)

func testHomeDirFunc() (string, error) {
	return TestHomeDir, nil
}

func CreateTestEnv() *env.DoltEnv {
	const name = "billy bob"
	const email = "bigbillieb@fake.horse"
	initialDirs := []string{TestHomeDir, WorkingDir}
	fs := filesys.NewInMemFS(initialDirs, nil, WorkingDir)
	dEnv := env.Load(context.Background(), testHomeDirFunc, fs, doltdb.InMemDoltDB, "test")
	cfg, _ := dEnv.Config.GetConfig(env.GlobalConfig)
	cfg.SetStrings(map[string]string{
		env.UserNameKey:  name,
		env.UserEmailKey: email,
	})
	err := dEnv.InitRepo(context.Background(), types.Format_7_18, name, email)

	if err != nil {
		panic("Failed to initialize environment:" + err.Error())
	}

	return dEnv
}

func CreateEnvWithSeedData(t *testing.T) *env.DoltEnv {
	dEnv := CreateTestEnv()
	imt, sch := CreateTestDataTable(true)

	ctx := context.Background()
	vrw := dEnv.DoltDB.ValueReadWriter()
	rd := table.NewInMemTableReader(imt)
	wr := noms.NewNomsMapCreator(ctx, vrw, sch)

	_, _, err := table.PipeRows(ctx, rd, wr, false)
	require.NoError(t, err)
	err = rd.Close(ctx)
	require.NoError(t, err)
	err = wr.Close(ctx)
	require.NoError(t, err)

	wrSch := wr.GetSchema()
	wrSch.Indexes().Merge(sch.Indexes().AllIndexes()...)

	idxSch := sch.Indexes().GetByName(IndexName)
	idxRows, err := editor.RebuildIndexRowData(ctx, vrw, sch, wr.GetMap(), idxSch)
	require.NoError(t, err)

	idxRef, err := doltdb.WriteValAndGetRef(ctx, vrw, idxRows)
	require.NoError(t, err)

	indexMap, err := types.NewMap(ctx, vrw, types.String(IndexName), idxRef)
	require.NoError(t, err)

	err = dEnv.PutTableToWorking(ctx, wrSch, wr.GetMap(), indexMap, TableName)
	require.NoError(t, err)

	return dEnv
}
