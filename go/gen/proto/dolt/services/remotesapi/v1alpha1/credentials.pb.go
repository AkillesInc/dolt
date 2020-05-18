// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dolt/services/remotesapi/v1alpha1/credentials.proto

package remotesapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WhoAmIRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoAmIRequest) Reset()         { *m = WhoAmIRequest{} }
func (m *WhoAmIRequest) String() string { return proto.CompactTextString(m) }
func (*WhoAmIRequest) ProtoMessage()    {}
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_24ed1d4c5faa311b, []int{0}
}

func (m *WhoAmIRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIRequest.Unmarshal(m, b)
}
func (m *WhoAmIRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIRequest.Marshal(b, m, deterministic)
}
func (m *WhoAmIRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIRequest.Merge(m, src)
}
func (m *WhoAmIRequest) XXX_Size() int {
	return xxx_messageInfo_WhoAmIRequest.Size(m)
}
func (m *WhoAmIRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIRequest proto.InternalMessageInfo

type WhoAmIResponse struct {
	// Ex: "bheni"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Ex: "Brian Hendriks"
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Ex: "brian@liquidata.co"
	EmailAddress         string   `protobuf:"bytes,3,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoAmIResponse) Reset()         { *m = WhoAmIResponse{} }
func (m *WhoAmIResponse) String() string { return proto.CompactTextString(m) }
func (*WhoAmIResponse) ProtoMessage()    {}
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_24ed1d4c5faa311b, []int{1}
}

func (m *WhoAmIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIResponse.Unmarshal(m, b)
}
func (m *WhoAmIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIResponse.Marshal(b, m, deterministic)
}
func (m *WhoAmIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIResponse.Merge(m, src)
}
func (m *WhoAmIResponse) XXX_Size() int {
	return xxx_messageInfo_WhoAmIResponse.Size(m)
}
func (m *WhoAmIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIResponse proto.InternalMessageInfo

func (m *WhoAmIResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *WhoAmIResponse) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *WhoAmIResponse) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*WhoAmIRequest)(nil), "dolt.services.remotesapi.v1alpha1.WhoAmIRequest")
	proto.RegisterType((*WhoAmIResponse)(nil), "dolt.services.remotesapi.v1alpha1.WhoAmIResponse")
}

func init() {
	proto.RegisterFile("dolt/services/remotesapi/v1alpha1/credentials.proto", fileDescriptor_24ed1d4c5faa311b)
}

var fileDescriptor_24ed1d4c5faa311b = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbf, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x55, 0x90, 0x2a, 0x30, 0x2d, 0x48, 0x9e, 0xaa, 0x4e, 0xb4, 0x2c, 0x2c, 0xd8, 0x84,
	0x8e, 0x4c, 0x85, 0x89, 0x85, 0xa1, 0x0c, 0x15, 0x2c, 0xd5, 0x35, 0x3e, 0x25, 0x96, 0xfc, 0xab,
	0x3e, 0xa7, 0x12, 0x33, 0xff, 0x38, 0xaa, 0x43, 0x88, 0x98, 0x2a, 0x46, 0x7f, 0x77, 0x9f, 0xf4,
	0xde, 0x99, 0x2d, 0x94, 0x37, 0x49, 0x12, 0xc6, 0xbd, 0x2e, 0x91, 0x64, 0x44, 0xeb, 0x13, 0x12,
	0x04, 0x2d, 0xf7, 0x05, 0x98, 0x50, 0x43, 0x21, 0xcb, 0x88, 0x0a, 0x5d, 0xd2, 0x60, 0x48, 0x84,
	0xe8, 0x93, 0xe7, 0xb3, 0x83, 0x24, 0x3a, 0x49, 0xf4, 0x92, 0xe8, 0xa4, 0xf9, 0x15, 0x1b, 0xaf,
	0x6b, 0xbf, 0xb4, 0x2f, 0x2b, 0xdc, 0x35, 0x48, 0x69, 0x9e, 0xd8, 0x65, 0x07, 0x28, 0x78, 0x47,
	0xc8, 0xa7, 0xec, 0xac, 0x21, 0x8c, 0x0e, 0x2c, 0x4e, 0x06, 0xd7, 0x83, 0xdb, 0xf3, 0xd5, 0xef,
	0x9b, 0xcf, 0xd8, 0x48, 0x69, 0x0a, 0x06, 0x3e, 0x37, 0x79, 0x7e, 0x92, 0xe7, 0x17, 0x3f, 0xec,
	0xf5, 0xb0, 0x72, 0xc3, 0xc6, 0x68, 0x41, 0x9b, 0x0d, 0x28, 0x15, 0x91, 0x68, 0x72, 0x9a, 0x77,
	0x46, 0x19, 0x2e, 0x5b, 0xf6, 0xf0, 0x35, 0x60, 0xfc, 0xb9, 0xcf, 0xff, 0xd6, 0x46, 0xe6, 0x96,
	0x0d, 0xdb, 0x30, 0xfc, 0x5e, 0x1c, 0xed, 0x22, 0xfe, 0x14, 0x99, 0x16, 0xff, 0x30, 0xda, 0xa6,
	0x4f, 0xef, 0x1f, 0xeb, 0x4a, 0xa7, 0xba, 0xd9, 0x8a, 0xd2, 0x5b, 0x69, 0xf4, 0xae, 0xd1, 0x0a,
	0x12, 0xdc, 0x69, 0x57, 0xca, 0x7c, 0xff, 0xca, 0xcb, 0x0a, 0x9d, 0xcc, 0xd7, 0x95, 0x47, 0x7f,
	0xe4, 0xb1, 0x67, 0xdb, 0x61, 0x76, 0x16, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x38, 0x85, 0x41,
	0x95, 0xc8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CredentialsServiceClient is the client API for CredentialsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CredentialsServiceClient interface {
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error)
}

type credentialsServiceClient struct {
	cc *grpc.ClientConn
}

func NewCredentialsServiceClient(cc *grpc.ClientConn) CredentialsServiceClient {
	return &credentialsServiceClient{cc}
}

func (c *credentialsServiceClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error) {
	out := new(WhoAmIResponse)
	err := c.cc.Invoke(ctx, "/dolt.services.remotesapi.v1alpha1.CredentialsService/WhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialsServiceServer is the server API for CredentialsService service.
type CredentialsServiceServer interface {
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error)
}

// UnimplementedCredentialsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCredentialsServiceServer struct {
}

func (*UnimplementedCredentialsServiceServer) WhoAmI(ctx context.Context, req *WhoAmIRequest) (*WhoAmIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}

func RegisterCredentialsServiceServer(s *grpc.Server, srv CredentialsServiceServer) {
	s.RegisterService(&_CredentialsService_serviceDesc, srv)
}

func _CredentialsService_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsServiceServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dolt.services.remotesapi.v1alpha1.CredentialsService/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsServiceServer).WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CredentialsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dolt.services.remotesapi.v1alpha1.CredentialsService",
	HandlerType: (*CredentialsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhoAmI",
			Handler:    _CredentialsService_WhoAmI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dolt/services/remotesapi/v1alpha1/credentials.proto",
}
