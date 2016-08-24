// Code generated by protoc-gen-go.
// source: build.proto
// DO NOT EDIT!

package ci

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BuildDescribeRequest struct {
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
	Number  int64  `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
}

func (m *BuildDescribeRequest) Reset()                    { *m = BuildDescribeRequest{} }
func (m *BuildDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildDescribeRequest) ProtoMessage()               {}
func (*BuildDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BuildDescribeResponse struct {
	Status            *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	JobName           string         `protobuf:"bytes,2,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Parents           string         `protobuf:"bytes,3,opt,name=parents" json:"parents,omitempty"`
	BuildNumber       int64          `protobuf:"varint,4,opt,name=build_number,json=buildNumber" json:"build_number,omitempty"`
	BaseUrl           string         `protobuf:"bytes,5,opt,name=base_url,json=baseUrl" json:"base_url,omitempty"`
	Building          bool           `protobuf:"varint,6,opt,name=building" json:"building,omitempty"`
	BuiltOn           string         `protobuf:"bytes,7,opt,name=built_on,json=builtOn" json:"built_on,omitempty"`
	Duration          int64          `protobuf:"varint,8,opt,name=duration" json:"duration,omitempty"`
	EstimatedDuration int64          `protobuf:"varint,9,opt,name=estimated_duration,json=estimatedDuration" json:"estimated_duration,omitempty"`
	FullDisplayName   string         `protobuf:"bytes,10,opt,name=full_display_name,json=fullDisplayName" json:"full_display_name,omitempty"`
	Result            string         `protobuf:"bytes,11,opt,name=result" json:"result,omitempty"`
	ConsoleOutput     string         `protobuf:"bytes,12,opt,name=console_output,json=consoleOutput" json:"console_output,omitempty"`
}

func (m *BuildDescribeResponse) Reset()                    { *m = BuildDescribeResponse{} }
func (m *BuildDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildDescribeResponse) ProtoMessage()               {}
func (*BuildDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *BuildDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type BuildListRequest struct {
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
}

func (m *BuildListRequest) Reset()                    { *m = BuildListRequest{} }
func (m *BuildListRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildListRequest) ProtoMessage()               {}
func (*BuildListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type BuildListResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	JobName string         `protobuf:"bytes,2,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Parents string         `protobuf:"bytes,3,opt,name=parents" json:"parents,omitempty"`
	Builds  []*Build       `protobuf:"bytes,4,rep,name=builds" json:"builds,omitempty"`
}

func (m *BuildListResponse) Reset()                    { *m = BuildListResponse{} }
func (m *BuildListResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildListResponse) ProtoMessage()               {}
func (*BuildListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BuildListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BuildListResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

type Build struct {
	Number    int64  `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*BuildDescribeRequest)(nil), "ci.BuildDescribeRequest")
	proto.RegisterType((*BuildDescribeResponse)(nil), "ci.BuildDescribeResponse")
	proto.RegisterType((*BuildListRequest)(nil), "ci.BuildListRequest")
	proto.RegisterType((*BuildListResponse)(nil), "ci.BuildListResponse")
	proto.RegisterType((*Build)(nil), "ci.Build")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Builds service

type BuildsClient interface {
	Describe(ctx context.Context, in *BuildDescribeRequest, opts ...grpc.CallOption) (*BuildDescribeResponse, error)
	List(ctx context.Context, in *BuildListRequest, opts ...grpc.CallOption) (*BuildListResponse, error)
}

type buildsClient struct {
	cc *grpc.ClientConn
}

func NewBuildsClient(cc *grpc.ClientConn) BuildsClient {
	return &buildsClient{cc}
}

func (c *buildsClient) Describe(ctx context.Context, in *BuildDescribeRequest, opts ...grpc.CallOption) (*BuildDescribeResponse, error) {
	out := new(BuildDescribeResponse)
	err := grpc.Invoke(ctx, "/ci.Builds/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) List(ctx context.Context, in *BuildListRequest, opts ...grpc.CallOption) (*BuildListResponse, error) {
	out := new(BuildListResponse)
	err := grpc.Invoke(ctx, "/ci.Builds/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Builds service

type BuildsServer interface {
	Describe(context.Context, *BuildDescribeRequest) (*BuildDescribeResponse, error)
	List(context.Context, *BuildListRequest) (*BuildListResponse, error)
}

func RegisterBuildsServer(s *grpc.Server, srv BuildsServer) {
	s.RegisterService(&_Builds_serviceDesc, srv)
}

func _Builds_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Builds/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).Describe(ctx, req.(*BuildDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builds_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Builds/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildsServer).List(ctx, req.(*BuildListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.Builds",
	HandlerType: (*BuildsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Builds_Describe_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Builds_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("build.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x71, 0xeb, 0x24, 0x4a, 0xd7, 0x2d, 0xa2, 0x2d, 0xaa, 0xe9, 0xc5, 0x66, 0xd8, 0x2f,
	0xcc, 0xee, 0xb2, 0x9b, 0x5d, 0x8f, 0xc0, 0x6e, 0x46, 0x0b, 0x1e, 0xbd, 0x36, 0xfe, 0xd1, 0x82,
	0x86, 0x63, 0x79, 0x96, 0x34, 0x08, 0xa5, 0x37, 0x7b, 0x85, 0x3e, 0xc0, 0x1e, 0x6a, 0xaf, 0xb0,
	0x9b, 0xbd, 0xc5, 0xa4, 0x23, 0xc5, 0x4b, 0xc2, 0x06, 0x83, 0x41, 0x6f, 0x42, 0xce, 0xf7, 0x1d,
	0x9d, 0xf3, 0xe9, 0x3b, 0x47, 0x46, 0x93, 0x42, 0xb1, 0xba, 0x8a, 0xdb, 0x8e, 0x4b, 0x8e, 0x07,
	0x25, 0x0b, 0xcf, 0x16, 0x9c, 0x2f, 0x6a, 0x9a, 0xe4, 0x2d, 0x4b, 0xf2, 0xa6, 0xe1, 0x32, 0x97,
	0x8c, 0x37, 0xc2, 0x66, 0x84, 0x27, 0x06, 0xae, 0xe4, 0xaa, 0xa5, 0x22, 0x81, 0x5f, 0x8b, 0x47,
	0x25, 0x3a, 0x7a, 0x6b, 0x0a, 0xcd, 0xa9, 0x28, 0x3b, 0x56, 0xd0, 0x94, 0x7e, 0x56, 0x54, 0x48,
	0x7c, 0x8a, 0x46, 0x9f, 0x78, 0x91, 0x35, 0xf9, 0x92, 0x12, 0xef, 0xa1, 0xf7, 0x6c, 0x9c, 0x0e,
	0x75, 0x7c, 0xa1, 0x43, 0x4c, 0xd0, 0xb0, 0xcd, 0x3b, 0xda, 0x48, 0x41, 0x06, 0x96, 0x71, 0x21,
	0x3e, 0x41, 0x41, 0xa3, 0x96, 0x05, 0xed, 0x88, 0xaf, 0x09, 0x3f, 0x75, 0x51, 0xf4, 0xcd, 0x47,
	0xc7, 0x3b, 0x5d, 0x44, 0xab, 0xb5, 0x51, 0xfc, 0x04, 0x05, 0x42, 0x0b, 0x55, 0x02, 0x9a, 0x4c,
	0x66, 0x87, 0xb1, 0xd5, 0x18, 0x7f, 0x00, 0x34, 0x75, 0xec, 0x96, 0x9c, 0xc1, 0x5f, 0xe5, 0xf8,
	0xdb, 0x72, 0x1e, 0xa1, 0x03, 0x30, 0x29, 0x73, 0xa2, 0xf6, 0x40, 0x94, 0x35, 0xee, 0x02, 0x20,
	0x53, 0xb7, 0xc8, 0x05, 0xcd, 0x54, 0x57, 0x93, 0x7d, 0x7b, 0xda, 0xc4, 0x57, 0x5d, 0x8d, 0x43,
	0x4d, 0x99, 0x4c, 0xd6, 0x2c, 0x48, 0xa0, 0xa9, 0x51, 0xda, 0xc7, 0x70, 0x4c, 0xff, 0x97, 0x19,
	0x6f, 0xc8, 0xd0, 0x1d, 0x33, 0xf1, 0x65, 0x63, 0x8e, 0x55, 0xaa, 0x03, 0xef, 0xc9, 0x08, 0x1a,
	0xf6, 0x31, 0x7e, 0x89, 0xb0, 0xf6, 0x96, 0x2d, 0x73, 0x49, 0xab, 0xac, 0xcf, 0x1a, 0x43, 0xd6,
	0xb4, 0x67, 0xe6, 0xeb, 0xf4, 0x17, 0x68, 0xfa, 0x51, 0xd5, 0x75, 0x56, 0x31, 0xd1, 0xd6, 0xf9,
	0xca, 0xde, 0x1e, 0x41, 0xbb, 0xfb, 0x86, 0x98, 0x5b, 0x1c, 0x5c, 0xd0, 0xd6, 0x77, 0x54, 0xa8,
	0x5a, 0x92, 0x09, 0x24, 0xb8, 0x08, 0x3f, 0x46, 0x87, 0xa5, 0x76, 0x9a, 0xd7, 0x34, 0xe3, 0x4a,
	0xb6, 0x4a, 0x92, 0x03, 0xe0, 0xef, 0x39, 0xf4, 0x12, 0xc0, 0xe8, 0x1d, 0x7a, 0x00, 0x03, 0x7a,
	0xcf, 0x84, 0xfc, 0x9f, 0x15, 0x88, 0x6e, 0x3d, 0x34, 0xdd, 0xa8, 0x74, 0x37, 0x63, 0x0e, 0x60,
	0x30, 0x42, 0x0f, 0xd8, 0xd7, 0xc5, 0xc7, 0x71, 0xc9, 0x62, 0xd0, 0x90, 0x3a, 0x22, 0xba, 0x42,
	0xfb, 0x00, 0x6c, 0x6c, 0xa8, 0xb7, 0xb9, 0xa1, 0x06, 0x77, 0x02, 0x6d, 0xdb, 0xb5, 0xa0, 0x33,
	0x34, 0xd6, 0x53, 0xd1, 0x6e, 0xe4, 0xcb, 0xd6, 0x2d, 0xf5, 0x6f, 0x60, 0xf6, 0xd3, 0x43, 0x01,
	0xd4, 0x15, 0x78, 0x85, 0x46, 0xeb, 0xe5, 0xc6, 0xa4, 0x17, 0xb0, 0xf3, 0xaa, 0xc2, 0xd3, 0x3f,
	0x30, 0xd6, 0xa2, 0xe8, 0xcd, 0xd7, 0xef, 0x3f, 0x6e, 0x07, 0x33, 0x7c, 0xae, 0x1f, 0x70, 0x2b,
	0x4a, 0x5e, 0xd9, 0x97, 0x5c, 0xb2, 0xe4, 0xcb, 0x79, 0xfc, 0x2a, 0xb1, 0x17, 0x49, 0xae, 0xd7,
	0xf6, 0xdc, 0x24, 0xd7, 0x56, 0xfa, 0x0d, 0xae, 0xd0, 0x9e, 0x31, 0x1b, 0x1f, 0xf5, 0xc5, 0x37,
	0xa6, 0x18, 0x1e, 0xef, 0xa0, 0xae, 0x5d, 0x02, 0xed, 0x9e, 0xe3, 0xa7, 0xff, 0xd8, 0xae, 0x08,
	0xe0, 0x7b, 0xf1, 0xfa, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x04, 0x11, 0xb5, 0x78, 0x04,
	0x00, 0x00,
}
