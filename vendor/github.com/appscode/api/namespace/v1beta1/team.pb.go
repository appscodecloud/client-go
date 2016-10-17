// Code generated by protoc-gen-go.
// source: team.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	team.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	GetRequest
	GetResponse
	IsAvailableRequest
	SubscriptionResponse
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Name             string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName      string   `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Email            string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	UserName         string   `protobuf:"bytes,4,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Password         string   `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	InviteEmails     []string `protobuf:"bytes,6,rep,name=invite_emails,json=inviteEmails" json:"invite_emails,omitempty"`
	SubscriptionType string   `protobuf:"bytes,7,opt,name=subscription_type,json=subscriptionType" json:"subscription_type,omitempty"`
	ClientIp         string   `protobuf:"bytes,8,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phid   string                  `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type IsAvailableRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *IsAvailableRequest) Reset()                    { *m = IsAvailableRequest{} }
func (m *IsAvailableRequest) String() string            { return proto.CompactTextString(m) }
func (*IsAvailableRequest) ProtoMessage()               {}
func (*IsAvailableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SubscriptionResponse struct {
	Status     *appscode_dtypes.Status       `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Product    *SubscriptionResponse_Product `protobuf:"bytes,2,opt,name=product" json:"product,omitempty"`
	AutoExtend bool                          `protobuf:"varint,3,opt,name=auto_extend,json=autoExtend" json:"auto_extend,omitempty"`
	Quota      string                        `protobuf:"bytes,4,opt,name=quota" json:"quota,omitempty"`
	DateStart  int64                         `protobuf:"varint,5,opt,name=date_start,json=dateStart" json:"date_start,omitempty"`
	DateEnd    int64                         `protobuf:"varint,6,opt,name=date_end,json=dateEnd" json:"date_end,omitempty"`
}

func (m *SubscriptionResponse) Reset()                    { *m = SubscriptionResponse{} }
func (m *SubscriptionResponse) String() string            { return proto.CompactTextString(m) }
func (*SubscriptionResponse) ProtoMessage()               {}
func (*SubscriptionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SubscriptionResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SubscriptionResponse) GetProduct() *SubscriptionResponse_Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type SubscriptionResponse_Product struct {
	Sku          string `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	Type         string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	DisplayName  string `protobuf:"bytes,3,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	PricingModel string `protobuf:"bytes,4,opt,name=pricing_model,json=pricingModel" json:"pricing_model,omitempty"`
}

func (m *SubscriptionResponse_Product) Reset()                    { *m = SubscriptionResponse_Product{} }
func (m *SubscriptionResponse_Product) String() string            { return proto.CompactTextString(m) }
func (*SubscriptionResponse_Product) ProtoMessage()               {}
func (*SubscriptionResponse_Product) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "appscode.namespace.v1beta1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "appscode.namespace.v1beta1.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "appscode.namespace.v1beta1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "appscode.namespace.v1beta1.GetResponse")
	proto.RegisterType((*IsAvailableRequest)(nil), "appscode.namespace.v1beta1.IsAvailableRequest")
	proto.RegisterType((*SubscriptionResponse)(nil), "appscode.namespace.v1beta1.SubscriptionResponse")
	proto.RegisterType((*SubscriptionResponse_Product)(nil), "appscode.namespace.v1beta1.SubscriptionResponse.Product")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Teams service

type TeamsClient interface {
	// Creates a new namespace, if name is valid and no namespace with same name exists.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Gets a namespace, if exists. This can be used to check existence of a namespace.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Check if a namespace name is available meaning name is valid and no namespace with same name exists.
	IsAvailable(ctx context.Context, in *IsAvailableRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	// Checks current subscription of a namespace
	Subscription(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error)
}

type teamsClient struct {
	cc *grpc.ClientConn
}

func NewTeamsClient(cc *grpc.ClientConn) TeamsClient {
	return &teamsClient{cc}
}

func (c *teamsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) IsAvailable(ctx context.Context, in *IsAvailableRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/IsAvailable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) Subscription(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*SubscriptionResponse, error) {
	out := new(SubscriptionResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/Subscription", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Teams service

type TeamsServer interface {
	// Creates a new namespace, if name is valid and no namespace with same name exists.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Gets a namespace, if exists. This can be used to check existence of a namespace.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Check if a namespace name is available meaning name is valid and no namespace with same name exists.
	IsAvailable(context.Context, *IsAvailableRequest) (*appscode_dtypes.VoidResponse, error)
	// Checks current subscription of a namespace
	Subscription(context.Context, *appscode_dtypes.VoidRequest) (*SubscriptionResponse, error)
}

func RegisterTeamsServer(s *grpc.Server, srv TeamsServer) {
	s.RegisterService(&_Teams_serviceDesc, srv)
}

func _Teams_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_IsAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).IsAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/IsAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).IsAvailable(ctx, req.(*IsAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_Subscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).Subscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/Subscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).Subscription(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Teams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.namespace.v1beta1.Teams",
	HandlerType: (*TeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Teams_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Teams_Get_Handler,
		},
		{
			MethodName: "IsAvailable",
			Handler:    _Teams_IsAvailable_Handler,
		},
		{
			MethodName: "Subscription",
			Handler:    _Teams_Subscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("team.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x4e, 0xd4, 0x4e,
	0x18, 0x4d, 0x77, 0xd9, 0x7f, 0xdf, 0x2e, 0xbf, 0xf0, 0x9b, 0x90, 0x58, 0x2b, 0x08, 0x96, 0xa8,
	0x88, 0xda, 0x0a, 0x1a, 0xa2, 0x78, 0x05, 0x86, 0x10, 0x2e, 0x34, 0xa4, 0x10, 0x2f, 0xbc, 0x69,
	0x66, 0xdb, 0xc9, 0x3a, 0xb1, 0xdb, 0x19, 0x3a, 0x53, 0x90, 0x18, 0x6f, 0x50, 0x5f, 0x40, 0x2f,
	0xbc, 0x37, 0xf1, 0x11, 0x7c, 0x12, 0x5f, 0xc1, 0x07, 0x31, 0x33, 0xd3, 0x5d, 0x4a, 0x56, 0x96,
	0x3f, 0x37, 0x9b, 0x99, 0x73, 0xbe, 0x6f, 0xbe, 0xb3, 0xa7, 0x67, 0x5a, 0x00, 0x49, 0x70, 0xdf,
	0xe3, 0x19, 0x93, 0x0c, 0x39, 0x98, 0x73, 0x11, 0xb1, 0x98, 0x78, 0x29, 0xee, 0x13, 0xc1, 0x71,
	0x44, 0xbc, 0x83, 0xe5, 0x2e, 0x91, 0x78, 0xd9, 0x99, 0xe9, 0x31, 0xd6, 0x4b, 0x88, 0x8f, 0x39,
	0xf5, 0x71, 0x9a, 0x32, 0x89, 0x25, 0x65, 0xa9, 0x30, 0x9d, 0xce, 0xcd, 0x41, 0xe7, 0x19, 0xfc,
	0xdc, 0x29, 0x3e, 0x96, 0x47, 0x9c, 0x08, 0x5f, 0xff, 0x9a, 0x02, 0xf7, 0x73, 0x05, 0x26, 0x5f,
	0x64, 0x04, 0x4b, 0x12, 0x90, 0xfd, 0x9c, 0x08, 0x89, 0x10, 0x4c, 0x28, 0x15, 0xb6, 0x35, 0x6f,
	0x2d, 0xb6, 0x02, 0xbd, 0x46, 0xb7, 0xa0, 0x13, 0x53, 0xc1, 0x13, 0x7c, 0x14, 0x6a, 0xae, 0xa2,
	0xb9, 0x76, 0x81, 0xbd, 0x52, 0x25, 0xd3, 0x50, 0x23, 0x7d, 0x4c, 0x13, 0xbb, 0xaa, 0x39, 0xb3,
	0x41, 0x37, 0xa0, 0x95, 0x0b, 0x92, 0x99, 0xae, 0x09, 0xcd, 0x34, 0x15, 0xa0, 0x5b, 0x1c, 0x68,
	0x72, 0x2c, 0xc4, 0x21, 0xcb, 0x62, 0xbb, 0x66, 0xb8, 0xc1, 0x1e, 0x2d, 0xc0, 0x24, 0x4d, 0x0f,
	0xa8, 0x24, 0xa1, 0x3e, 0x48, 0xd8, 0xf5, 0xf9, 0xea, 0x62, 0x2b, 0xe8, 0x18, 0x70, 0x53, 0x63,
	0xe8, 0x3e, 0xfc, 0x2f, 0xf2, 0xae, 0x88, 0x32, 0xca, 0xd5, 0x9f, 0x0e, 0xd5, 0x1f, 0xb3, 0x1b,
	0xfa, 0xa4, 0xa9, 0x32, 0xb1, 0x77, 0xc4, 0x89, 0x92, 0x12, 0x25, 0x94, 0xa4, 0x32, 0xa4, 0xdc,
	0x6e, 0x9a, 0x71, 0x06, 0xd8, 0xe6, 0xee, 0x3a, 0xfc, 0x37, 0x70, 0x41, 0x70, 0x96, 0x0a, 0x82,
	0x7c, 0xa8, 0x0b, 0x89, 0x65, 0x2e, 0xb4, 0x11, 0xed, 0x95, 0x6b, 0xde, 0xf0, 0x21, 0x19, 0x1b,
	0xbd, 0x5d, 0x4d, 0x07, 0x45, 0x99, 0x3b, 0x0f, 0xb0, 0x45, 0xe4, 0x18, 0x17, 0xdd, 0x00, 0xda,
	0xba, 0xe2, 0x8a, 0x13, 0xd4, 0x99, 0xfc, 0x2d, 0x8d, 0x0b, 0xf7, 0xf5, 0xda, 0x5d, 0x04, 0xb4,
	0x2d, 0xd6, 0x0f, 0x30, 0x4d, 0x70, 0x37, 0x19, 0xf7, 0x0c, 0xdd, 0x2f, 0x55, 0x98, 0xde, 0x2d,
	0x99, 0x72, 0x75, 0x1d, 0x01, 0x34, 0x78, 0xc6, 0xe2, 0x3c, 0x92, 0x5a, 0x4a, 0x7b, 0xe5, 0xa9,
	0x77, 0x76, 0x80, 0xbd, 0x7f, 0xcd, 0xf4, 0x76, 0x4c, 0x7f, 0x30, 0x38, 0x08, 0xcd, 0x41, 0x1b,
	0xe7, 0x92, 0x85, 0xe4, 0xbd, 0x24, 0x69, 0xac, 0x43, 0xd4, 0x0c, 0x40, 0x41, 0x9b, 0x1a, 0x51,
	0xf9, 0xda, 0xcf, 0x99, 0xc4, 0x45, 0x8a, 0xcc, 0x06, 0xcd, 0x02, 0xc4, 0x58, 0x92, 0x50, 0x48,
	0x9c, 0x49, 0x1d, 0xa2, 0x6a, 0xd0, 0x52, 0xc8, 0xae, 0x02, 0xd0, 0x75, 0x68, 0x6a, 0x5a, 0x1d,
	0x59, 0xd7, 0x64, 0x43, 0xed, 0x37, 0xd3, 0xd8, 0x39, 0x84, 0x46, 0x21, 0x02, 0x4d, 0x41, 0x55,
	0xbc, 0xcb, 0x0b, 0xb3, 0xd4, 0x52, 0xf9, 0xa7, 0xb3, 0x54, 0x38, 0xad, 0xd6, 0x23, 0x77, 0xa0,
	0x3a, 0x7a, 0x07, 0x16, 0x60, 0x92, 0x67, 0x34, 0xa2, 0x69, 0x2f, 0xec, 0xb3, 0x98, 0x24, 0x85,
	0xd6, 0x4e, 0x01, 0xbe, 0x54, 0xd8, 0xca, 0xa7, 0x1a, 0xd4, 0xf6, 0x08, 0xee, 0x0b, 0xf4, 0xdd,
	0x82, 0xba, 0x49, 0x1d, 0xba, 0x37, 0xce, 0xc1, 0x53, 0xf7, 0xd3, 0x59, 0xba, 0x48, 0xa9, 0xb1,
	0xd9, 0x5d, 0x3d, 0xfe, 0x65, 0x57, 0x9a, 0xd6, 0xf1, 0xef, 0x3f, 0xdf, 0x2a, 0x4b, 0xee, 0x6d,
	0xff, 0xd4, 0xdb, 0x60, 0xd8, 0xed, 0x17, 0xdd, 0xbe, 0x7a, 0x25, 0x89, 0x35, 0x6b, 0x09, 0x7d,
	0xb5, 0xa0, 0xba, 0x45, 0x24, 0xba, 0x33, 0x6e, 0xd6, 0x49, 0xda, 0x9d, 0xbb, 0xe7, 0xd6, 0x15,
	0x82, 0xd6, 0x4a, 0x82, 0x3c, 0xf4, 0xe0, 0x42, 0x82, 0xfc, 0x0f, 0x8a, 0xf8, 0x88, 0x7e, 0x5a,
	0xd0, 0x2e, 0x65, 0x1d, 0x79, 0xe3, 0x86, 0x8e, 0x5e, 0x0a, 0x67, 0x76, 0x24, 0xd7, 0xaf, 0x19,
	0x8d, 0x87, 0xd2, 0xb6, 0x4a, 0xd2, 0x9e, 0xa3, 0x67, 0x97, 0x91, 0xe6, 0x53, 0xf1, 0x10, 0x0f,
	0x75, 0xfd, 0xb0, 0xa0, 0x53, 0x0e, 0x3d, 0x9a, 0x39, 0x63, 0xb0, 0x91, 0xf5, 0xe8, 0xb2, 0x97,
	0xc7, 0xdd, 0x28, 0x29, 0x5d, 0x45, 0x4f, 0xce, 0x53, 0xda, 0xa5, 0x49, 0x42, 0xd3, 0x9e, 0x5f,
	0x7e, 0x23, 0x6e, 0xac, 0x81, 0x1b, 0xb1, 0xfe, 0xc9, 0x68, 0xcc, 0xe9, 0xe8, 0xf8, 0x8d, 0x96,
	0x0a, 0xea, 0x8e, 0xfa, 0x50, 0xec, 0x58, 0x6f, 0x1a, 0x05, 0xda, 0xad, 0xeb, 0x4f, 0xc7, 0xe3,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60, 0x53, 0xf9, 0xdd, 0xc3, 0x06, 0x00, 0x00,
}
