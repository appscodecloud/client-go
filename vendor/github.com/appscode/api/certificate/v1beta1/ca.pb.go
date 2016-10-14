// Code generated by protoc-gen-go.
// source: ca.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Use specific requests for protos
type CACreateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Csr  string `protobuf:"bytes,2,opt,name=csr" json:"csr,omitempty"`
}

func (m *CACreateRequest) Reset()                    { *m = CACreateRequest{} }
func (m *CACreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CACreateRequest) ProtoMessage()               {}
func (*CACreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type CACreateResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phid   string                  `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
	Cert   []byte                  `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
	Key    []byte                  `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Root   []byte                  `protobuf:"bytes,5,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *CACreateResponse) Reset()                    { *m = CACreateResponse{} }
func (m *CACreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CACreateResponse) ProtoMessage()               {}
func (*CACreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CACreateResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CACreateRequest)(nil), "appscode.certificate.v1beta1.CACreateRequest")
	proto.RegisterType((*CACreateResponse)(nil), "appscode.certificate.v1beta1.CACreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CAs service

type CAsClient interface {
	Create(ctx context.Context, in *CACreateRequest, opts ...grpc.CallOption) (*CACreateResponse, error)
}

type cAsClient struct {
	cc *grpc.ClientConn
}

func NewCAsClient(cc *grpc.ClientConn) CAsClient {
	return &cAsClient{cc}
}

func (c *cAsClient) Create(ctx context.Context, in *CACreateRequest, opts ...grpc.CallOption) (*CACreateResponse, error) {
	out := new(CACreateResponse)
	err := grpc.Invoke(ctx, "/appscode.certificate.v1beta1.CAs/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CAs service

type CAsServer interface {
	Create(context.Context, *CACreateRequest) (*CACreateResponse, error)
}

func RegisterCAsServer(s *grpc.Server, srv CAsServer) {
	s.RegisterService(&_CAs_serviceDesc, srv)
}

func _CAs_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CACreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.certificate.v1beta1.CAs/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAsServer).Create(ctx, req.(*CACreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CAs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.certificate.v1beta1.CAs",
	HandlerType: (*CAsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CAs_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("ca.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x51, 0x5d, 0x4a, 0x33, 0x31,
	0x14, 0x25, 0x6d, 0xbf, 0x7e, 0x6d, 0x14, 0x2c, 0x79, 0xd0, 0x50, 0x8a, 0x94, 0x01, 0xa1, 0x14,
	0x4c, 0x68, 0x05, 0x05, 0xdf, 0xea, 0xec, 0x60, 0x7c, 0xf3, 0x2d, 0xcd, 0x5c, 0xeb, 0xa0, 0x4e,
	0xe2, 0xe4, 0x56, 0xe8, 0x6b, 0x77, 0x20, 0x2e, 0xc0, 0x47, 0x57, 0xe0, 0x4a, 0xdc, 0x82, 0x0b,
	0x91, 0x24, 0xe3, 0x2f, 0x28, 0xbe, 0x0c, 0x87, 0x73, 0xcf, 0xb9, 0x67, 0x4e, 0x2e, 0xed, 0x68,
	0x25, 0x6c, 0x65, 0xd0, 0xb0, 0x81, 0xb2, 0xd6, 0x69, 0x93, 0x83, 0xd0, 0x50, 0x61, 0x71, 0x5e,
	0x68, 0x85, 0x20, 0x6e, 0x27, 0x73, 0x40, 0x35, 0xe9, 0x0f, 0x16, 0xc6, 0x2c, 0xae, 0x40, 0x2a,
	0x5b, 0x48, 0x55, 0x96, 0x06, 0x15, 0x16, 0xa6, 0x74, 0xd1, 0xdb, 0xdf, 0x7d, 0xf3, 0xfe, 0x30,
	0xdf, 0xf6, 0x74, 0x8e, 0x2b, 0x0b, 0x4e, 0x86, 0x6f, 0xe4, 0x93, 0x23, 0xba, 0x95, 0xce, 0xd2,
	0x0a, 0x14, 0x42, 0x06, 0x37, 0x4b, 0x70, 0xc8, 0x18, 0x6d, 0x95, 0xea, 0x1a, 0x38, 0x19, 0x92,
	0x51, 0x37, 0x0b, 0x98, 0xf5, 0x68, 0x53, 0xbb, 0x8a, 0x37, 0x02, 0xe5, 0x61, 0x72, 0x47, 0x68,
	0xef, 0xc3, 0xe9, 0xac, 0x29, 0x1d, 0x30, 0x49, 0xdb, 0x0e, 0x15, 0x2e, 0x5d, 0x30, 0x6f, 0x4c,
	0x77, 0xc4, 0x7b, 0xa5, 0x98, 0x2d, 0x4e, 0xc3, 0x38, 0xab, 0x65, 0x3e, 0xcb, 0x5e, 0x14, 0x79,
	0xbd, 0x38, 0x60, 0xcf, 0xf9, 0xfe, 0xbc, 0x39, 0x24, 0xa3, 0xcd, 0x2c, 0x60, 0x9f, 0x7f, 0x09,
	0x2b, 0xde, 0x0a, 0x94, 0x87, 0x5e, 0x55, 0x19, 0x83, 0xfc, 0x5f, 0x54, 0x79, 0x3c, 0x7d, 0x24,
	0xb4, 0x99, 0xce, 0x1c, 0x7b, 0x20, 0xb4, 0x1d, 0xff, 0x8c, 0xed, 0x8b, 0xdf, 0x1e, 0x55, 0x7c,
	0xeb, 0xde, 0x17, 0x7f, 0x95, 0xc7, 0xc2, 0xc9, 0xe1, 0xfa, 0x89, 0x37, 0x3a, 0x64, 0xfd, 0xfc,
	0x72, 0xdf, 0x18, 0x27, 0x7b, 0xf2, 0xcb, 0x15, 0x3e, 0x6d, 0x90, 0xf5, 0x06, 0xa9, 0x95, 0x3b,
	0x26, 0xe3, 0x93, 0xee, 0xd9, 0xff, 0x9a, 0x99, 0xb7, 0xc3, 0x21, 0x0e, 0x5e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xd7, 0x8a, 0x51, 0x80, 0x08, 0x02, 0x00, 0x00,
}
