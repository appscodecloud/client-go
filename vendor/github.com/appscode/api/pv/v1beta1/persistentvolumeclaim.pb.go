// Code generated by protoc-gen-go.
// source: persistentvolumeclaim.proto
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

type PVCRegisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SizeGb    int64  `protobuf:"varint,3,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Namespace string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCRegisterRequest) Reset()                    { *m = PVCRegisterRequest{} }
func (m *PVCRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCRegisterRequest) ProtoMessage()               {}
func (*PVCRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type PVCUnregisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCUnregisterRequest) Reset()                    { *m = PVCUnregisterRequest{} }
func (m *PVCUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCUnregisterRequest) ProtoMessage()               {}
func (*PVCUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type PVCDescribeRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCDescribeRequest) Reset()                    { *m = PVCDescribeRequest{} }
func (m *PVCDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCDescribeRequest) ProtoMessage()               {}
func (*PVCDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type PVCInfo struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SizeGb      int64    `protobuf:"varint,2,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Namespace   string   `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Status      string   `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Volume      string   `protobuf:"bytes,5,opt,name=volume" json:"volume,omitempty"`
	AccessModes []string `protobuf:"bytes,6,rep,name=accessModes" json:"accessModes,omitempty"`
}

func (m *PVCInfo) Reset()                    { *m = PVCInfo{} }
func (m *PVCInfo) String() string            { return proto.CompactTextString(m) }
func (*PVCInfo) ProtoMessage()               {}
func (*PVCInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type PVCDescribeResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Pvc    *PVCInfo                `protobuf:"bytes,2,opt,name=pvc" json:"pvc,omitempty"`
}

func (m *PVCDescribeResponse) Reset()                    { *m = PVCDescribeResponse{} }
func (m *PVCDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*PVCDescribeResponse) ProtoMessage()               {}
func (*PVCDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *PVCDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PVCDescribeResponse) GetPvc() *PVCInfo {
	if m != nil {
		return m.Pvc
	}
	return nil
}

func init() {
	proto.RegisterType((*PVCRegisterRequest)(nil), "appscode.pv.v1beta1.PVCRegisterRequest")
	proto.RegisterType((*PVCUnregisterRequest)(nil), "appscode.pv.v1beta1.PVCUnregisterRequest")
	proto.RegisterType((*PVCDescribeRequest)(nil), "appscode.pv.v1beta1.PVCDescribeRequest")
	proto.RegisterType((*PVCInfo)(nil), "appscode.pv.v1beta1.PVCInfo")
	proto.RegisterType((*PVCDescribeResponse)(nil), "appscode.pv.v1beta1.PVCDescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PersistentVolumeClaims service

type PersistentVolumeClaimsClient interface {
	Describe(ctx context.Context, in *PVCDescribeRequest, opts ...grpc.CallOption) (*PVCDescribeResponse, error)
	Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type persistentVolumeClaimsClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumeClaimsClient(cc *grpc.ClientConn) PersistentVolumeClaimsClient {
	return &persistentVolumeClaimsClient{cc}
}

func (c *persistentVolumeClaimsClient) Describe(ctx context.Context, in *PVCDescribeRequest, opts ...grpc.CallOption) (*PVCDescribeResponse, error) {
	out := new(PVCDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.pv.v1beta1.PersistentVolumeClaims/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClaimsClient) Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.pv.v1beta1.PersistentVolumeClaims/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClaimsClient) Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.pv.v1beta1.PersistentVolumeClaims/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumeClaims service

type PersistentVolumeClaimsServer interface {
	Describe(context.Context, *PVCDescribeRequest) (*PVCDescribeResponse, error)
	Register(context.Context, *PVCRegisterRequest) (*appscode_dtypes.VoidResponse, error)
	Unregister(context.Context, *PVCUnregisterRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterPersistentVolumeClaimsServer(s *grpc.Server, srv PersistentVolumeClaimsServer) {
	s.RegisterService(&_PersistentVolumeClaims_serviceDesc, srv)
}

func _PersistentVolumeClaims_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVCDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeClaimsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.pv.v1beta1.PersistentVolumeClaims/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeClaimsServer).Describe(ctx, req.(*PVCDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumeClaims_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVCRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeClaimsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.pv.v1beta1.PersistentVolumeClaims/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeClaimsServer).Register(ctx, req.(*PVCRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumeClaims_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVCUnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeClaimsServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.pv.v1beta1.PersistentVolumeClaims/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeClaimsServer).Unregister(ctx, req.(*PVCUnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersistentVolumeClaims_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.pv.v1beta1.PersistentVolumeClaims",
	HandlerType: (*PersistentVolumeClaimsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _PersistentVolumeClaims_Describe_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _PersistentVolumeClaims_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PersistentVolumeClaims_Unregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("persistentvolumeclaim.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0x3d, 0x8e, 0xd3, 0x40,
	0x14, 0xd6, 0xc4, 0xbb, 0xf9, 0x99, 0x74, 0xb3, 0x28, 0x6b, 0x99, 0x80, 0xa2, 0x34, 0x84, 0x2d,
	0x3c, 0xda, 0xd0, 0x51, 0x12, 0x24, 0x44, 0x81, 0x14, 0x19, 0xe1, 0x82, 0x06, 0xc6, 0xce, 0x23,
	0x18, 0x12, 0xcf, 0xe0, 0x37, 0xb6, 0x14, 0x56, 0x69, 0xf6, 0x06, 0x88, 0x3b, 0x70, 0x03, 0x8e,
	0x01, 0x0d, 0xe2, 0x06, 0x1c, 0x04, 0x79, 0xec, 0x78, 0x37, 0xd9, 0x58, 0x5b, 0x20, 0xd2, 0x58,
	0xf3, 0xbe, 0x79, 0x7f, 0xdf, 0x9b, 0xef, 0x99, 0xde, 0x55, 0x90, 0x60, 0x84, 0x1a, 0x62, 0x9d,
	0xc9, 0x45, 0xba, 0x84, 0x70, 0x21, 0xa2, 0xa5, 0xab, 0x12, 0xa9, 0x25, 0x3b, 0x11, 0x4a, 0x61,
	0x28, 0x67, 0xe0, 0xaa, 0xcc, 0xcd, 0xce, 0x03, 0xd0, 0xe2, 0xdc, 0xe9, 0xcf, 0xa5, 0x9c, 0x2f,
	0x80, 0x0b, 0x15, 0x71, 0x11, 0xc7, 0x52, 0x0b, 0x1d, 0xc9, 0x18, 0x8b, 0x10, 0xe7, 0xfe, 0x26,
	0xa4, 0xe6, 0xbe, 0x97, 0xc3, 0x33, 0xbd, 0x52, 0x80, 0xdc, 0x7c, 0x0b, 0x7c, 0xb8, 0xa2, 0x6c,
	0xea, 0x4f, 0x3c, 0x98, 0xe7, 0xbd, 0x24, 0x1e, 0x7c, 0x4a, 0x01, 0x35, 0xb3, 0x69, 0x2b, 0x5c,
	0xa4, 0x39, 0x62, 0x93, 0x01, 0x19, 0x75, 0xbc, 0x8d, 0xc9, 0x18, 0x3d, 0x8a, 0xc5, 0x12, 0xec,
	0x86, 0x81, 0xcd, 0x99, 0x9d, 0xd2, 0x16, 0x46, 0x9f, 0xe1, 0xcd, 0x3c, 0xb0, 0xad, 0x01, 0x19,
	0x59, 0x5e, 0x33, 0x37, 0x9f, 0x05, 0xac, 0x4f, 0x3b, 0xb9, 0x03, 0x2a, 0x11, 0x82, 0x7d, 0x64,
	0x22, 0xae, 0x80, 0x61, 0x40, 0xef, 0x4c, 0xfd, 0xc9, 0xab, 0x38, 0xf9, 0xa7, 0xe2, 0x5b, 0x35,
	0xac, 0xdd, 0x1a, 0x6f, 0x0d, 0xbd, 0xa7, 0x80, 0x61, 0x12, 0x05, 0xf0, 0x3f, 0x2a, 0x7c, 0x23,
	0xb4, 0x35, 0xf5, 0x27, 0xcf, 0xe3, 0x77, 0xb2, 0x8a, 0x26, 0xfb, 0x87, 0xd3, 0xa8, 0x1f, 0xce,
	0x6e, 0x5a, 0xd6, 0xa3, 0x4d, 0xd4, 0x42, 0xa7, 0x58, 0xce, 0xad, 0xb4, 0x72, 0xbc, 0xd0, 0x8b,
	0x7d, 0x5c, 0xe0, 0x85, 0xc5, 0x06, 0xb4, 0x2b, 0xc2, 0x10, 0x10, 0x5f, 0xc8, 0x19, 0xa0, 0xdd,
	0x1c, 0x58, 0xa3, 0x8e, 0x77, 0x1d, 0x1a, 0x66, 0xf4, 0x64, 0x6b, 0x14, 0xa8, 0x64, 0x8c, 0xc0,
	0x78, 0x55, 0x28, 0xef, 0xba, 0x3b, 0x3e, 0x75, 0x2b, 0xf1, 0x15, 0x72, 0x71, 0x5f, 0x9a, 0xeb,
	0xaa, 0x03, 0x97, 0x5a, 0x2a, 0x0b, 0x0d, 0x99, 0xee, 0xb8, 0xef, 0xee, 0x91, 0xaa, 0x5b, 0xce,
	0xc3, 0xcb, 0x1d, 0xc7, 0x5f, 0x8e, 0x69, 0x6f, 0x5a, 0x89, 0xdd, 0x37, 0xed, 0x4e, 0x72, 0xb1,
	0x23, 0xfb, 0x4d, 0x68, 0x7b, 0xd3, 0x10, 0x7b, 0x50, 0x97, 0x6a, 0xe7, 0xf5, 0x9c, 0xd1, 0xed,
	0x8e, 0x05, 0xb7, 0x61, 0x76, 0xf9, 0xdd, 0x6e, 0xb4, 0xc9, 0xe5, 0xaf, 0x3f, 0x5f, 0x1b, 0x1f,
	0xd8, 0x7b, 0xbe, 0xb5, 0x23, 0x1f, 0xd3, 0x00, 0x92, 0x18, 0x34, 0x20, 0x2f, 0x73, 0xf0, 0x52,
	0x05, 0xc8, 0x2f, 0xca, 0xd3, 0x9a, 0x57, 0x2f, 0x82, 0xfc, 0xa2, 0x3a, 0xaf, 0xf9, 0xde, 0xed,
	0x2d, 0x5d, 0xd6, 0xec, 0x07, 0xa1, 0xed, 0xcd, 0x4a, 0xd5, 0xf3, 0xda, 0x59, 0x3a, 0xe7, 0xde,
	0x8d, 0xc9, 0xfb, 0x32, 0x9a, 0x55, 0x64, 0x56, 0xd7, 0xc8, 0x2c, 0x9d, 0x83, 0x91, 0x79, 0x4c,
	0xce, 0xd8, 0x4f, 0x42, 0xe9, 0xd5, 0x9e, 0xb2, 0x87, 0x75, 0x8c, 0x6e, 0xec, 0xf2, 0x6d, 0x9c,
	0xb6, 0x1e, 0xe8, 0xec, 0x60, 0x9c, 0x9e, 0x74, 0x5e, 0xb7, 0xca, 0x94, 0x41, 0xd3, 0xfc, 0x07,
	0x1f, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x7b, 0x3c, 0x55, 0x91, 0x05, 0x00, 0x00,
}
