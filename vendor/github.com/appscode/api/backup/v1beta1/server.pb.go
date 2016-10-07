// Code generated by protoc-gen-go.
// source: server.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerCreateRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Credential string `protobuf:"bytes,2,opt,name=credential" json:"credential,omitempty"`
	BucketName string `protobuf:"bytes,3,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	Disk       string `protobuf:"bytes,4,opt,name=disk" json:"disk,omitempty"`
	// TODO will be removed
	Namespace string `protobuf:"bytes,5,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ServerCreateRequest) Reset()                    { *m = ServerCreateRequest{} }
func (m *ServerCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerCreateRequest) ProtoMessage()               {}
func (*ServerCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ServerDeleteRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ServerDeleteRequest) Reset()                    { *m = ServerDeleteRequest{} }
func (m *ServerDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerDeleteRequest) ProtoMessage()               {}
func (*ServerDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*ServerCreateRequest)(nil), "backup.v1beta1.ServerCreateRequest")
	proto.RegisterType((*ServerDeleteRequest)(nil), "backup.v1beta1.ServerDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Servers service

type ServersClient interface {
	Create(ctx context.Context, in *ServerCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *ServerDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type serversClient struct {
	cc *grpc.ClientConn
}

func NewServersClient(cc *grpc.ClientConn) ServersClient {
	return &serversClient{cc}
}

func (c *serversClient) Create(ctx context.Context, in *ServerCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.v1beta1.Servers/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serversClient) Delete(ctx context.Context, in *ServerDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.v1beta1.Servers/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Servers service

type ServersServer interface {
	Create(context.Context, *ServerCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *ServerDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterServersServer(s *grpc.Server, srv ServersServer) {
	s.RegisterService(&_Servers_serviceDesc, srv)
}

func _Servers_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.v1beta1.Servers/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServersServer).Create(ctx, req.(*ServerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Servers_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.v1beta1.Servers/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServersServer).Delete(ctx, req.(*ServerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Servers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup.v1beta1.Servers",
	HandlerType: (*ServersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Servers_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Servers_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("server.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xa6, 0x75, 0x6e, 0x2e, 0x8a, 0x87, 0x28, 0x52, 0xc6, 0x98, 0x52, 0x2f, 0xb2, 0x43, 0xc3,
	0xf4, 0xe6, 0x45, 0xd4, 0x5d, 0x27, 0x32, 0xc1, 0x83, 0x17, 0x49, 0xd3, 0x9f, 0x51, 0xda, 0x25,
	0x31, 0x49, 0x07, 0x22, 0x5e, 0xf6, 0x08, 0x8a, 0x2f, 0xe0, 0x3b, 0xe8, 0x8b, 0xf8, 0x0a, 0x3e,
	0x88, 0x2c, 0xe9, 0x70, 0x15, 0x07, 0x82, 0x97, 0xf2, 0xf7, 0xfb, 0xbe, 0x7e, 0xfd, 0xfe, 0x2f,
	0x41, 0x1b, 0x1a, 0xd4, 0x04, 0x54, 0x24, 0x95, 0x30, 0x02, 0x6f, 0xc6, 0x94, 0x65, 0x85, 0x8c,
	0x26, 0xbd, 0x18, 0x0c, 0xed, 0xb5, 0xda, 0x23, 0x21, 0x46, 0x39, 0x10, 0x2a, 0x53, 0x42, 0x39,
	0x17, 0x86, 0x9a, 0x54, 0x70, 0xed, 0xd4, 0xad, 0x0e, 0x95, 0x52, 0x33, 0x91, 0x2c, 0xe3, 0x77,
	0x66, 0x70, 0x62, 0xee, 0x25, 0x68, 0x62, 0x9f, 0x0e, 0x0f, 0x5f, 0x3d, 0xb4, 0x75, 0x65, 0x7f,
	0x7b, 0xae, 0x80, 0x1a, 0x18, 0xc2, 0x5d, 0x01, 0xda, 0xe0, 0x00, 0x35, 0x58, 0x5e, 0x68, 0x03,
	0x2a, 0xf0, 0xf6, 0xbc, 0x83, 0xe6, 0x70, 0xfe, 0x8a, 0x3b, 0x08, 0x31, 0x05, 0x09, 0x70, 0x93,
	0xd2, 0x3c, 0xf0, 0x2d, 0xb9, 0x80, 0xe0, 0x5d, 0xb4, 0x1e, 0x17, 0x2c, 0x03, 0x73, 0xcb, 0xe9,
	0x18, 0x82, 0x15, 0x27, 0x70, 0xd0, 0x05, 0x1d, 0x03, 0xc6, 0xa8, 0x96, 0xa4, 0x3a, 0x0b, 0x6a,
	0x96, 0xb1, 0x33, 0x6e, 0xa3, 0xe6, 0x4c, 0xad, 0x25, 0x65, 0x10, 0xac, 0x5a, 0xe2, 0x1b, 0x08,
	0x07, 0xf3, 0x8c, 0x7d, 0xc8, 0xe1, 0x2f, 0x19, 0x2b, 0x76, 0xfe, 0x0f, 0xbb, 0xc3, 0x77, 0x1f,
	0x35, 0x9c, 0x9f, 0xc6, 0x2f, 0x1e, 0xaa, 0xbb, 0xcd, 0xf1, 0x7e, 0x54, 0x6d, 0x3c, 0xfa, 0xa5,
	0x97, 0xd6, 0x76, 0xe4, 0x4a, 0x8c, 0xae, 0x45, 0x9a, 0x0c, 0x41, 0x4b, 0xc1, 0x35, 0x84, 0x97,
	0xd3, 0xb7, 0xc0, 0x5f, 0xf3, 0xa6, 0x1f, 0x9f, 0xcf, 0x7e, 0x3f, 0x3c, 0x21, 0x95, 0xc3, 0xc8,
	0x8a, 0x18, 0x14, 0x07, 0x03, 0x9a, 0x94, 0xce, 0xa4, 0xcc, 0xa9, 0xc9, 0x43, 0x39, 0x3d, 0xda,
	0x4f, 0x48, 0x4c, 0x59, 0x91, 0xd3, 0x63, 0xaf, 0x8b, 0x9f, 0x3c, 0x54, 0x77, 0xdb, 0x2e, 0xcb,
	0x55, 0xe9, 0x62, 0x49, 0xae, 0xc1, 0x42, 0xae, 0xd3, 0xee, 0x7f, 0x73, 0x9d, 0x35, 0x6f, 0x1a,
	0xa5, 0x32, 0xae, 0xdb, 0xeb, 0x73, 0xf4, 0x15, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x52, 0xd2, 0x6b,
	0xb4, 0x02, 0x00, 0x00,
}
