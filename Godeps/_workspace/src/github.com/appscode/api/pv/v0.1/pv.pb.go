// Code generated by protoc-gen-go.
// source: pv.proto
// DO NOT EDIT!

package pv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PVRegisterRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Identifier string `protobuf:"bytes,3,opt,name=identifier" json:"identifier,omitempty"`
	Plugin     string `protobuf:"bytes,4,opt,name=plugin" json:"plugin,omitempty"`
	Size       int64  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Endpoint   string `protobuf:"bytes,6,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *PVRegisterRequest) Reset()                    { *m = PVRegisterRequest{} }
func (m *PVRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVRegisterRequest) ProtoMessage()               {}
func (*PVRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type PVUnregisterRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PVUnregisterRequest) Reset()                    { *m = PVUnregisterRequest{} }
func (m *PVUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVUnregisterRequest) ProtoMessage()               {}
func (*PVUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type PVDescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PVDescribeRequest) Reset()                    { *m = PVDescribeRequest{} }
func (m *PVDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*PVDescribeRequest) ProtoMessage()               {}
func (*PVDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type PVInfo struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size        int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Status      string   `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Volume      string   `protobuf:"bytes,4,opt,name=volume" json:"volume,omitempty"`
	Claim       string   `protobuf:"bytes,5,opt,name=claim" json:"claim,omitempty"`
	Plugin      string   `protobuf:"bytes,6,opt,name=plugin" json:"plugin,omitempty"`
	AccessModes []string `protobuf:"bytes,7,rep,name=accessModes" json:"accessModes,omitempty"`
}

func (m *PVInfo) Reset()                    { *m = PVInfo{} }
func (m *PVInfo) String() string            { return proto.CompactTextString(m) }
func (*PVInfo) ProtoMessage()               {}
func (*PVInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type PVDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Pv     *PVInfo        `protobuf:"bytes,2,opt,name=pv" json:"pv,omitempty"`
}

func (m *PVDescribeResponse) Reset()                    { *m = PVDescribeResponse{} }
func (m *PVDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*PVDescribeResponse) ProtoMessage()               {}
func (*PVDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *PVDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PVDescribeResponse) GetPv() *PVInfo {
	if m != nil {
		return m.Pv
	}
	return nil
}

func init() {
	proto.RegisterType((*PVRegisterRequest)(nil), "pv.PVRegisterRequest")
	proto.RegisterType((*PVUnregisterRequest)(nil), "pv.PVUnregisterRequest")
	proto.RegisterType((*PVDescribeRequest)(nil), "pv.PVDescribeRequest")
	proto.RegisterType((*PVInfo)(nil), "pv.PVInfo")
	proto.RegisterType((*PVDescribeResponse)(nil), "pv.PVDescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PersistentVolumes service

type PersistentVolumesClient interface {
	Describe(ctx context.Context, in *PVDescribeRequest, opts ...grpc.CallOption) (*PVDescribeResponse, error)
	Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type persistentVolumesClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumesClient(cc *grpc.ClientConn) PersistentVolumesClient {
	return &persistentVolumesClient{cc}
}

func (c *persistentVolumesClient) Describe(ctx context.Context, in *PVDescribeRequest, opts ...grpc.CallOption) (*PVDescribeResponse, error) {
	out := new(PVDescribeResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumes/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumesClient) Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumes/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumesClient) Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumes/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumes service

type PersistentVolumesServer interface {
	Describe(context.Context, *PVDescribeRequest) (*PVDescribeResponse, error)
	Register(context.Context, *PVRegisterRequest) (*dtypes.VoidResponse, error)
	Unregister(context.Context, *PVUnregisterRequest) (*dtypes.VoidResponse, error)
}

func RegisterPersistentVolumesServer(s *grpc.Server, srv PersistentVolumesServer) {
	s.RegisterService(&_PersistentVolumes_serviceDesc, srv)
}

func _PersistentVolumes_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PersistentVolumes/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Describe(ctx, req.(*PVDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumes_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumesServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PersistentVolumes/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Register(ctx, req.(*PVRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumes_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVUnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumesServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.PersistentVolumes/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Unregister(ctx, req.(*PVUnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersistentVolumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.PersistentVolumes",
	HandlerType: (*PersistentVolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _PersistentVolumes_Describe_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _PersistentVolumes_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PersistentVolumes_Unregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("pv.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6a, 0xd5, 0x40,
	0x14, 0x26, 0xb9, 0x6d, 0x7a, 0x7b, 0x2e, 0x08, 0x8e, 0xf5, 0x1a, 0x82, 0xc8, 0x25, 0x8b, 0x52,
	0x2a, 0x24, 0xf6, 0xba, 0x73, 0x27, 0xba, 0x71, 0x21, 0x5c, 0x22, 0x06, 0xb7, 0x69, 0x72, 0x7a,
	0x19, 0x4d, 0x67, 0xc6, 0xcc, 0x24, 0xa0, 0xa5, 0x1b, 0x5f, 0xc1, 0xa7, 0x70, 0xe7, 0xbb, 0xb8,
	0xf0, 0x05, 0x7c, 0x10, 0xe7, 0x27, 0x89, 0xf1, 0x8a, 0x50, 0x74, 0x13, 0xe6, 0x7c, 0x67, 0x66,
	0xbe, 0xf3, 0x7d, 0xf3, 0x05, 0xe6, 0xa2, 0x4b, 0x44, 0xc3, 0x15, 0x27, 0xbe, 0xe8, 0xa2, 0xfb,
	0x5b, 0xce, 0xb7, 0x35, 0xa6, 0x85, 0xa0, 0x69, 0xc1, 0x18, 0x57, 0x85, 0xa2, 0x9c, 0x49, 0xb7,
	0x23, 0x5a, 0x1a, 0xb8, 0x52, 0x1f, 0x04, 0xca, 0xd4, 0x7e, 0x1d, 0x1e, 0x7f, 0xf1, 0xe0, 0xf6,
	0x26, 0xcf, 0x70, 0x4b, 0xa5, 0xc2, 0x26, 0xc3, 0xf7, 0x2d, 0x4a, 0x45, 0x42, 0x38, 0x28, 0xeb,
	0xd6, 0x20, 0xa1, 0xb7, 0xf2, 0x4e, 0x0e, 0xb3, 0xa1, 0x24, 0x04, 0xf6, 0x58, 0x71, 0x89, 0xa1,
	0x6f, 0x61, 0xbb, 0x26, 0x0f, 0x00, 0x68, 0x85, 0x4c, 0xd1, 0x0b, 0xaa, 0x0f, 0xcc, 0x6c, 0x67,
	0x82, 0x90, 0x25, 0x04, 0xa2, 0x6e, 0xb7, 0x94, 0x85, 0x7b, 0xb6, 0xd7, 0x57, 0xe6, 0x2e, 0x49,
	0x3f, 0x62, 0xb8, 0xaf, 0xd1, 0x59, 0x66, 0xd7, 0x24, 0x82, 0x39, 0xb2, 0x4a, 0x70, 0xca, 0x54,
	0x18, 0xd8, 0xdd, 0x63, 0x1d, 0x3f, 0x83, 0x3b, 0x9b, 0xfc, 0x35, 0x6b, 0xfe, 0x67, 0xd8, 0xf8,
	0xa9, 0xd1, 0xfb, 0x1c, 0x65, 0xd9, 0xd0, 0x73, 0xfc, 0xb7, 0x2b, 0xbe, 0x7a, 0x10, 0x6c, 0xf2,
	0x17, 0xec, 0x82, 0x8f, 0x6d, 0x6f, 0x62, 0xc7, 0x20, 0xcb, 0x9f, 0xc8, 0xd2, 0x16, 0x48, 0xfd,
	0x20, 0xad, 0xec, 0xed, 0xe9, 0x2b, 0x83, 0x77, 0xbc, 0x6e, 0xf5, 0x0d, 0xbd, 0x35, 0xae, 0x22,
	0x47, 0xb0, 0x5f, 0xd6, 0x05, 0xbd, 0xb4, 0xde, 0x1c, 0x66, 0xae, 0x98, 0x18, 0x19, 0xfc, 0x66,
	0xe4, 0x0a, 0x16, 0x45, 0x59, 0xa2, 0x94, 0x2f, 0x79, 0x85, 0x32, 0x3c, 0x58, 0xcd, 0x74, 0x73,
	0x0a, 0xc5, 0x6f, 0x80, 0x4c, 0x55, 0x4b, 0xa1, 0x93, 0x81, 0xe4, 0x78, 0x9c, 0xca, 0xcc, 0xbf,
	0x58, 0xdf, 0x4a, 0x5c, 0x42, 0x92, 0x57, 0x16, 0x1d, 0xa7, 0x8c, 0x40, 0x07, 0xcc, 0xea, 0x59,
	0xac, 0x21, 0xd1, 0xa9, 0x73, 0xea, 0x33, 0x8d, 0xae, 0xbf, 0xfb, 0xda, 0x50, 0x6c, 0xa4, 0x79,
	0x12, 0xa6, 0x72, 0x3b, 0xbe, 0x24, 0x14, 0xe6, 0x03, 0x1b, 0xb9, 0xeb, 0x4e, 0xec, 0x78, 0x1e,
	0x2d, 0x77, 0x61, 0x37, 0x54, 0x9c, 0x7c, 0xfa, 0xf6, 0xe3, 0xb3, 0x7f, 0x42, 0x8e, 0x75, 0x92,
	0x85, 0x2c, 0xf5, 0xf8, 0x36, 0xd2, 0xa2, 0x4b, 0xbb, 0x47, 0xc9, 0x59, 0x7a, 0xd5, 0xbf, 0xcc,
	0x75, 0x7a, 0x65, 0xdc, 0xbe, 0x26, 0x6f, 0x61, 0x3e, 0xc4, 0x77, 0xa0, 0xda, 0x89, 0x73, 0x74,
	0x34, 0xe8, 0xca, 0x39, 0xad, 0x46, 0xa2, 0x33, 0x4b, 0xf4, 0x30, 0xba, 0x21, 0xd1, 0x13, 0xef,
	0x94, 0xbc, 0x03, 0xf8, 0x95, 0x3f, 0x72, 0xcf, 0xb1, 0xfd, 0x91, 0xc8, 0xbf, 0xf0, 0xf5, 0xc2,
	0x4e, 0x6f, 0xc8, 0x77, 0x1e, 0xd8, 0x3f, 0xf4, 0xf1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf,
	0xef, 0x22, 0x9c, 0xe7, 0x03, 0x00, 0x00,
}
