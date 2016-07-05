// Code generated by protoc-gen-go.
// source: quota.proto
// DO NOT EDIT!

package billing

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

type QuotaGetResponse struct {
	Status   *dtypes.Status   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Assigned map[string]int64 `protobuf:"bytes,2,rep,name=assigned" json:"assigned,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Used     map[string]int64 `protobuf:"bytes,3,rep,name=used" json:"used,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *QuotaGetResponse) Reset()                    { *m = QuotaGetResponse{} }
func (m *QuotaGetResponse) String() string            { return proto.CompactTextString(m) }
func (*QuotaGetResponse) ProtoMessage()               {}
func (*QuotaGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *QuotaGetResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QuotaGetResponse) GetAssigned() map[string]int64 {
	if m != nil {
		return m.Assigned
	}
	return nil
}

func (m *QuotaGetResponse) GetUsed() map[string]int64 {
	if m != nil {
		return m.Used
	}
	return nil
}

func init() {
	proto.RegisterType((*QuotaGetResponse)(nil), "billing.QuotaGetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Quota service

type QuotaClient interface {
	Get(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*QuotaGetResponse, error)
}

type quotaClient struct {
	cc *grpc.ClientConn
}

func NewQuotaClient(cc *grpc.ClientConn) QuotaClient {
	return &quotaClient{cc}
}

func (c *quotaClient) Get(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*QuotaGetResponse, error) {
	out := new(QuotaGetResponse)
	err := grpc.Invoke(ctx, "/billing.Quota/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Quota service

type QuotaServer interface {
	Get(context.Context, *dtypes.VoidRequest) (*QuotaGetResponse, error)
}

func RegisterQuotaServer(s *grpc.Server, srv QuotaServer) {
	s.RegisterService(&_Quota_serviceDesc, srv)
}

func _Quota_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Quota/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServer).Get(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Quota_serviceDesc = grpc.ServiceDesc{
	ServiceName: "billing.Quota",
	HandlerType: (*QuotaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Quota_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("quota.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x69, 0xeb, 0xa6, 0xfb, 0x0f, 0x65, 0x44, 0x91, 0x59, 0x3c, 0x8c, 0x0a, 0xda, 0x53,
	0xaa, 0xf3, 0x30, 0xd1, 0x93, 0x88, 0x78, 0x36, 0xa2, 0x57, 0xe9, 0x6c, 0xa8, 0xc1, 0x92, 0x74,
	0xfb, 0xa7, 0x83, 0x5e, 0x7d, 0x05, 0xdf, 0xc3, 0x97, 0xf1, 0x15, 0x7c, 0x10, 0xd3, 0xa4, 0x1b,
	0x28, 0xec, 0xe0, 0x25, 0x24, 0xff, 0x7c, 0xdf, 0x8f, 0x2f, 0xf9, 0xa0, 0x3f, 0xab, 0x94, 0x4e,
	0x69, 0x39, 0x57, 0x5a, 0x91, 0xcd, 0xa9, 0x28, 0x0a, 0x21, 0xf3, 0xf0, 0x30, 0x57, 0x2a, 0x2f,
	0x78, 0x92, 0x96, 0x22, 0x49, 0xa5, 0x34, 0x0a, 0x2d, 0x94, 0x44, 0x27, 0x0b, 0xf7, 0x9b, 0x71,
	0xa6, 0xeb, 0x92, 0x63, 0x62, 0x57, 0x37, 0x8f, 0x3e, 0x7d, 0x18, 0xdc, 0x37, 0xb8, 0x3b, 0xae,
	0x19, 0xc7, 0xd2, 0x38, 0x38, 0x39, 0x86, 0x2e, 0x1a, 0x7b, 0x85, 0x43, 0x6f, 0xe4, 0xc5, 0xfd,
	0xf1, 0x0e, 0x75, 0x4e, 0xfa, 0x60, 0xa7, 0xac, 0xbd, 0x25, 0x37, 0xb0, 0x95, 0x22, 0x8a, 0x5c,
	0xf2, 0x6c, 0xe8, 0x8f, 0x02, 0xa3, 0x3c, 0xa1, 0x6d, 0x1c, 0xfa, 0x17, 0x4a, 0xaf, 0x5b, 0xe5,
	0xad, 0xd4, 0xf3, 0x9a, 0xad, 0x8c, 0x64, 0x02, 0x1b, 0x15, 0x1a, 0x40, 0x60, 0x01, 0x47, 0xeb,
	0x01, 0x8f, 0xb8, 0x34, 0x5b, 0x43, 0x78, 0x05, 0xdb, 0xbf, 0x98, 0x64, 0x00, 0xc1, 0x1b, 0xaf,
	0x6d, 0xe6, 0x1e, 0x6b, 0xb6, 0x64, 0x0f, 0x3a, 0x8b, 0xb4, 0xa8, 0xb8, 0x49, 0xe7, 0xc5, 0x01,
	0x73, 0x87, 0x4b, 0xff, 0xc2, 0x0b, 0x27, 0xd0, 0x5b, 0xf1, 0xfe, 0x63, 0x1c, 0xbf, 0x42, 0xc7,
	0x26, 0x23, 0xcf, 0x10, 0x98, 0x74, 0x64, 0x77, 0xf9, 0x37, 0x4f, 0x4a, 0x64, 0x8c, 0xcf, 0x2a,
	0x8e, 0x3a, 0x3c, 0x58, 0xfb, 0x8a, 0x28, 0x7e, 0xff, 0xfa, 0xfe, 0xf0, 0x23, 0x32, 0x32, 0x45,
	0x95, 0xf8, 0xa2, 0x32, 0xd7, 0x58, 0xab, 0x4f, 0x16, 0xa7, 0xf4, 0x2c, 0xb1, 0xfd, 0x4e, 0xbb,
	0xb6, 0xa1, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x87, 0x98, 0x8c, 0xef, 0x01, 0x00,
	0x00,
}
