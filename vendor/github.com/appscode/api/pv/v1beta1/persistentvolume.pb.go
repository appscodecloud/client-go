// Code generated by protoc-gen-go.
// source: persistentvolume.proto
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

type PVRegisterRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Identifier string `protobuf:"bytes,3,opt,name=identifier" json:"identifier,omitempty"`
	Plugin     string `protobuf:"bytes,4,opt,name=plugin" json:"plugin,omitempty"`
	SizeGb     int64  `protobuf:"varint,5,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Endpoint   string `protobuf:"bytes,6,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *PVRegisterRequest) Reset()                    { *m = PVRegisterRequest{} }
func (m *PVRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVRegisterRequest) ProtoMessage()               {}
func (*PVRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PVUnregisterRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PVUnregisterRequest) Reset()                    { *m = PVUnregisterRequest{} }
func (m *PVUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVUnregisterRequest) ProtoMessage()               {}
func (*PVUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type PVDescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PVDescribeRequest) Reset()                    { *m = PVDescribeRequest{} }
func (m *PVDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*PVDescribeRequest) ProtoMessage()               {}
func (*PVDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type PVInfo struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SizeGb      int64    `protobuf:"varint,2,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Status      string   `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Volume      string   `protobuf:"bytes,4,opt,name=volume" json:"volume,omitempty"`
	Claim       string   `protobuf:"bytes,5,opt,name=claim" json:"claim,omitempty"`
	Plugin      string   `protobuf:"bytes,6,opt,name=plugin" json:"plugin,omitempty"`
	AccessModes []string `protobuf:"bytes,7,rep,name=access_modes,json=accessModes" json:"access_modes,omitempty"`
}

func (m *PVInfo) Reset()                    { *m = PVInfo{} }
func (m *PVInfo) String() string            { return proto.CompactTextString(m) }
func (*PVInfo) ProtoMessage()               {}
func (*PVInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type PVDescribeResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Pv     *PVInfo                 `protobuf:"bytes,2,opt,name=pv" json:"pv,omitempty"`
}

func (m *PVDescribeResponse) Reset()                    { *m = PVDescribeResponse{} }
func (m *PVDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*PVDescribeResponse) ProtoMessage()               {}
func (*PVDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *PVDescribeResponse) GetStatus() *appscode_dtypes.Status {
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
	proto.RegisterType((*PVRegisterRequest)(nil), "appscode.pv.v1beta1.PVRegisterRequest")
	proto.RegisterType((*PVUnregisterRequest)(nil), "appscode.pv.v1beta1.PVUnregisterRequest")
	proto.RegisterType((*PVDescribeRequest)(nil), "appscode.pv.v1beta1.PVDescribeRequest")
	proto.RegisterType((*PVInfo)(nil), "appscode.pv.v1beta1.PVInfo")
	proto.RegisterType((*PVDescribeResponse)(nil), "appscode.pv.v1beta1.PVDescribeResponse")
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
	Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type persistentVolumesClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumesClient(cc *grpc.ClientConn) PersistentVolumesClient {
	return &persistentVolumesClient{cc}
}

func (c *persistentVolumesClient) Describe(ctx context.Context, in *PVDescribeRequest, opts ...grpc.CallOption) (*PVDescribeResponse, error) {
	out := new(PVDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.pv.v1beta1.PersistentVolumes/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumesClient) Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.pv.v1beta1.PersistentVolumes/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumesClient) Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.pv.v1beta1.PersistentVolumes/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumes service

type PersistentVolumesServer interface {
	Describe(context.Context, *PVDescribeRequest) (*PVDescribeResponse, error)
	Register(context.Context, *PVRegisterRequest) (*appscode_dtypes.VoidResponse, error)
	Unregister(context.Context, *PVUnregisterRequest) (*appscode_dtypes.VoidResponse, error)
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
		FullMethod: "/appscode.pv.v1beta1.PersistentVolumes/Describe",
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
		FullMethod: "/appscode.pv.v1beta1.PersistentVolumes/Register",
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
		FullMethod: "/appscode.pv.v1beta1.PersistentVolumes/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Unregister(ctx, req.(*PVUnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersistentVolumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.pv.v1beta1.PersistentVolumes",
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
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("persistentvolume.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xd5, 0x38, 0xa9, 0x93, 0xdc, 0x7c, 0x9b, 0x4e, 0x3f, 0x52, 0xcb, 0x85, 0x12, 0xb2, 0x80,
	0xa8, 0x48, 0xb6, 0x1a, 0x76, 0x6c, 0x10, 0x05, 0x09, 0x21, 0x01, 0x32, 0x46, 0x64, 0x51, 0x09,
	0x55, 0xfe, 0xb9, 0x8d, 0x46, 0x24, 0x33, 0x83, 0x67, 0x1c, 0x09, 0xaa, 0x6e, 0xfa, 0x0a, 0xbc,
	0x07, 0x02, 0x89, 0x9f, 0x07, 0xe1, 0x15, 0x78, 0x10, 0xe4, 0x9f, 0xb8, 0x4e, 0x4a, 0x54, 0x09,
	0xba, 0xb1, 0xe6, 0xdc, 0x3b, 0x33, 0x3e, 0xe7, 0xdc, 0x7b, 0x07, 0x7a, 0x12, 0x13, 0xc5, 0x94,
	0x46, 0xae, 0xe7, 0x62, 0x9a, 0xce, 0xd0, 0x91, 0x89, 0xd0, 0x82, 0x6e, 0x05, 0x52, 0xaa, 0x48,
	0xc4, 0xe8, 0xc8, 0xb9, 0x33, 0xdf, 0x0f, 0x51, 0x07, 0xfb, 0xf6, 0xf5, 0x89, 0x10, 0x93, 0x29,
	0xba, 0x81, 0x64, 0x6e, 0xc0, 0xb9, 0xd0, 0x81, 0x66, 0x82, 0xab, 0xe2, 0x88, 0xbd, 0xbb, 0x38,
	0xb2, 0x26, 0x7f, 0x73, 0x29, 0x1f, 0xeb, 0xf7, 0x12, 0x95, 0x9b, 0x7f, 0x8b, 0x0d, 0x83, 0x4f,
	0x04, 0x36, 0xbd, 0xb1, 0x8f, 0x93, 0x8c, 0x4f, 0xe2, 0xe3, 0xbb, 0x14, 0x95, 0xa6, 0x16, 0xb4,
	0xa2, 0x69, 0x9a, 0x45, 0x2c, 0xd2, 0x27, 0xc3, 0x8e, 0xbf, 0x80, 0x94, 0x42, 0x93, 0x07, 0x33,
	0xb4, 0x8c, 0x3c, 0x9c, 0xaf, 0xe9, 0x2e, 0x00, 0x8b, 0x91, 0x6b, 0x76, 0xcc, 0x30, 0xb1, 0x1a,
	0x79, 0xa6, 0x16, 0xa1, 0x3d, 0x30, 0xe5, 0x34, 0x9d, 0x30, 0x6e, 0x35, 0xf3, 0x5c, 0x89, 0xe8,
	0x36, 0xb4, 0x14, 0xfb, 0x80, 0x47, 0x93, 0xd0, 0xda, 0xe8, 0x93, 0x61, 0xc3, 0x37, 0x33, 0xf8,
	0x24, 0xa4, 0x36, 0xb4, 0x91, 0xc7, 0x52, 0x30, 0xae, 0x2d, 0x33, 0x3f, 0x52, 0xe1, 0xc1, 0x23,
	0xd8, 0xf2, 0xc6, 0xaf, 0x79, 0xf2, 0x2f, 0x8c, 0x07, 0x0f, 0x33, 0xd1, 0x8f, 0x51, 0x45, 0x09,
	0x0b, 0xf1, 0xef, 0xae, 0xf8, 0x4e, 0xc0, 0xf4, 0xc6, 0x4f, 0xf9, 0xb1, 0xa8, 0xd2, 0xa4, 0xe6,
	0x49, 0x4d, 0x9b, 0xb1, 0xa4, 0xad, 0x07, 0xa6, 0xd2, 0x81, 0x4e, 0x55, 0x69, 0x54, 0x89, 0xb2,
	0x78, 0xd1, 0x0c, 0x0b, 0x93, 0x0a, 0x44, 0xff, 0x87, 0x8d, 0x68, 0x1a, 0xb0, 0x59, 0x6e, 0x51,
	0xc7, 0x2f, 0x40, 0xcd, 0x52, 0x73, 0xc9, 0xd2, 0x5b, 0xf0, 0x5f, 0x10, 0x45, 0xa8, 0xd4, 0xd1,
	0x4c, 0xc4, 0xa8, 0xac, 0x56, 0xbf, 0x31, 0xec, 0xf8, 0xdd, 0x22, 0xf6, 0x3c, 0x0b, 0x0d, 0x12,
	0xa0, 0x75, 0xed, 0x4a, 0x0a, 0xae, 0x90, 0xba, 0x15, 0xad, 0x4c, 0x45, 0x77, 0xb4, 0xed, 0x54,
	0xcd, 0x58, 0x74, 0x8d, 0xf3, 0x2a, 0x4f, 0x57, 0x7c, 0xef, 0x82, 0x21, 0xe7, 0xb9, 0xb6, 0xee,
	0x68, 0xc7, 0xf9, 0x43, 0xe7, 0x3a, 0x85, 0x3b, 0xbe, 0x21, 0xe7, 0xa3, 0xaf, 0x4d, 0xd8, 0xf4,
	0xaa, 0xa6, 0x1f, 0xe7, 0xca, 0x14, 0xfd, 0x41, 0xa0, 0xbd, 0x20, 0x42, 0x6f, 0xaf, 0xb9, 0x63,
	0xa5, 0x4a, 0xf6, 0x9d, 0x4b, 0xf7, 0x15, 0x8a, 0x06, 0x87, 0x67, 0xdf, 0x2c, 0xa3, 0x4d, 0xce,
	0x7e, 0xfe, 0xfa, 0x68, 0xbc, 0xa0, 0xcf, 0xdc, 0xa5, 0x49, 0x78, 0x9b, 0x86, 0x98, 0x70, 0xd4,
	0xa8, 0xdc, 0xf2, 0x0a, 0xb7, 0x2c, 0xb6, 0x72, 0x4f, 0xca, 0xd5, 0xa9, 0xbb, 0x3a, 0xa7, 0xca,
	0x3d, 0xc9, 0x8a, 0x7b, 0x4a, 0x3f, 0x13, 0x68, 0x2f, 0x66, 0x66, 0x2d, 0xf3, 0x95, 0xa1, 0xb2,
	0x6f, 0x5c, 0xb0, 0x74, 0x2c, 0x58, 0x5c, 0xf1, 0x7d, 0x53, 0xe3, 0xfb, 0xd2, 0xbe, 0x52, 0xbe,
	0xf7, 0xc9, 0x1e, 0xfd, 0x42, 0x00, 0xce, 0xc7, 0x86, 0x0e, 0xd7, 0x90, 0xbe, 0x30, 0x59, 0x97,
	0xd1, 0x5e, 0xb2, 0x79, 0xef, 0x4a, 0x69, 0x1f, 0x3c, 0x80, 0x9d, 0x48, 0xcc, 0xce, 0xff, 0x1f,
	0x48, 0x56, 0x63, 0x7b, 0x70, 0x6d, 0xb5, 0xa5, 0xbc, 0xec, 0x49, 0xf3, 0xc8, 0x61, 0xab, 0xdc,
	0x11, 0x9a, 0xf9, 0x23, 0x77, 0xef, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x26, 0x44, 0x59,
	0x72, 0x05, 0x00, 0x00,
}
