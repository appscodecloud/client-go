// Code generated by protoc-gen-go.
// source: volume.proto
// DO NOT EDIT!

package v1beta1

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

type VolumeListRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace    string `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,3,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
}

func (m *VolumeListRequest) Reset()                    { *m = VolumeListRequest{} }
func (m *VolumeListRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeListRequest) ProtoMessage()               {}
func (*VolumeListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type VolumeListResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Volumes []*Volume      `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *VolumeListResponse) Reset()                    { *m = VolumeListResponse{} }
func (m *VolumeListResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeListResponse) ProtoMessage()               {}
func (*VolumeListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *VolumeListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VolumeListResponse) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type Volume struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *Volume) Reset()                    { *m = Volume{} }
func (m *Volume) String() string            { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()               {}
func (*Volume) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func init() {
	proto.RegisterType((*VolumeListRequest)(nil), "glusterfs.v1beta1.VolumeListRequest")
	proto.RegisterType((*VolumeListResponse)(nil), "glusterfs.v1beta1.VolumeListResponse")
	proto.RegisterType((*Volume)(nil), "glusterfs.v1beta1.Volume")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Volumes service

type VolumesClient interface {
	List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error)
}

type volumesClient struct {
	cc *grpc.ClientConn
}

func NewVolumesClient(cc *grpc.ClientConn) VolumesClient {
	return &volumesClient{cc}
}

func (c *volumesClient) List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error) {
	out := new(VolumeListResponse)
	err := grpc.Invoke(ctx, "/glusterfs.v1beta1.Volumes/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Volumes service

type VolumesServer interface {
	List(context.Context, *VolumeListRequest) (*VolumeListResponse, error)
}

func RegisterVolumesServer(s *grpc.Server, srv VolumesServer) {
	s.RegisterService(&_Volumes_serviceDesc, srv)
}

func _Volumes_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.v1beta1.Volumes/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).List(ctx, req.(*VolumeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Volumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glusterfs.v1beta1.Volumes",
	HandlerType: (*VolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Volumes_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("volume.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0x53, 0x20, 0x10, 0x06, 0xfe, 0xe4, 0xcf, 0x1c, 0x0c, 0x12, 0x0e, 0xd8, 0x88, 0x21,
	0x31, 0xe9, 0x06, 0x78, 0x03, 0xbd, 0x1a, 0x0f, 0x35, 0xf1, 0xa0, 0x31, 0x66, 0x29, 0x2b, 0x12,
	0xa1, 0xbb, 0x30, 0x5b, 0x12, 0x43, 0xb8, 0x78, 0xf4, 0xea, 0xcb, 0xf8, 0x00, 0xbe, 0x81, 0xaf,
	0xe0, 0x83, 0x98, 0xee, 0x2e, 0xb5, 0x09, 0x31, 0x5e, 0x9a, 0xf6, 0x37, 0xdf, 0x7c, 0xfd, 0x66,
	0x06, 0xea, 0x6b, 0x39, 0x4f, 0x16, 0x22, 0x50, 0x2b, 0xa9, 0x25, 0x36, 0xa7, 0xf3, 0x84, 0xb4,
	0x58, 0x3d, 0x50, 0xb0, 0x1e, 0x8c, 0x85, 0xe6, 0x83, 0x76, 0x67, 0x2a, 0xe5, 0x74, 0x2e, 0x18,
	0x57, 0x33, 0xc6, 0xe3, 0x58, 0x6a, 0xae, 0x67, 0x32, 0x26, 0xdb, 0xd0, 0x3e, 0x48, 0xf1, 0x44,
	0x3f, 0x2b, 0x41, 0xcc, 0x3c, 0x2d, 0xf7, 0x5f, 0x3d, 0x68, 0x5e, 0x1b, 0xe7, 0x8b, 0x19, 0xe9,
	0x50, 0x2c, 0x13, 0x41, 0x1a, 0x8f, 0xa0, 0xfe, 0x94, 0x8c, 0xc5, 0x7d, 0x64, 0xff, 0xd2, 0xf2,
	0xba, 0x5e, 0xbf, 0x1a, 0xd6, 0x52, 0x76, 0x6e, 0x11, 0xf6, 0xa0, 0x61, 0x24, 0x31, 0x5f, 0x08,
	0x52, 0x3c, 0x12, 0xad, 0x82, 0x11, 0xfd, 0x4b, 0xe9, 0xe5, 0x0e, 0xe2, 0x29, 0xfc, 0x44, 0xcd,
	0xec, 0x8a, 0x46, 0xf9, 0x3f, 0x2b, 0x38, 0x4f, 0x7f, 0x09, 0x98, 0xcf, 0x42, 0x4a, 0xc6, 0x24,
	0xf0, 0x04, 0xca, 0xa4, 0xb9, 0x4e, 0xc8, 0xc4, 0xa8, 0x0d, 0x1b, 0x81, 0x9d, 0x23, 0xb8, 0x32,
	0x34, 0x74, 0x55, 0x1c, 0x41, 0xc5, 0xee, 0x88, 0x5a, 0x85, 0x6e, 0xb1, 0x5f, 0x1b, 0x1e, 0x06,
	0x7b, 0x5b, 0x0a, 0xac, 0x7f, 0xb8, 0x53, 0xfa, 0x1d, 0x28, 0x5b, 0x84, 0x08, 0x25, 0xc5, 0xf5,
	0xa3, 0x9b, 0xd5, 0xbc, 0x0f, 0x3f, 0x3c, 0xa8, 0xd8, 0x32, 0xe1, 0xbb, 0x07, 0xa5, 0x34, 0x17,
	0x1e, 0xff, 0x6a, 0x9b, 0x5b, 0x61, 0xbb, 0xf7, 0x87, 0xca, 0x0e, 0xe7, 0x47, 0x2f, 0x9f, 0x5f,
	0x6f, 0x85, 0x3b, 0xbc, 0x65, 0x5c, 0x29, 0x8a, 0xe4, 0xc4, 0x1e, 0x30, 0xeb, 0x65, 0xae, 0x97,
	0xb9, 0xc0, 0x6c, 0x93, 0x3f, 0xcc, 0xd6, 0x7d, 0x66, 0x47, 0xd8, 0xb2, 0xcd, 0xde, 0xba, 0xb7,
	0x67, 0xd5, 0x9b, 0x8a, 0xb3, 0x19, 0x97, 0xcd, 0xd9, 0x47, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3b, 0x8f, 0x2b, 0xed, 0x4f, 0x02, 0x00, 0x00,
}