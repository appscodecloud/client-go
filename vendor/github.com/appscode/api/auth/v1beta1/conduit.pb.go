// Code generated by protoc-gen-go.
// source: conduit.proto
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

type ConduitRequest struct {
}

func (m *ConduitRequest) Reset()                    { *m = ConduitRequest{} }
func (m *ConduitRequest) String() string            { return proto.CompactTextString(m) }
func (*ConduitRequest) ProtoMessage()               {}
func (*ConduitRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ConduitWhoAmIResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	User   *ConduitUser   `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *ConduitWhoAmIResponse) Reset()                    { *m = ConduitWhoAmIResponse{} }
func (m *ConduitWhoAmIResponse) String() string            { return proto.CompactTextString(m) }
func (*ConduitWhoAmIResponse) ProtoMessage()               {}
func (*ConduitWhoAmIResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ConduitWhoAmIResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConduitWhoAmIResponse) GetUser() *ConduitUser {
	if m != nil {
		return m.User
	}
	return nil
}

type ConduitUsersResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	User   []*ConduitUser `protobuf:"bytes,2,rep,name=user" json:"user,omitempty"`
}

func (m *ConduitUsersResponse) Reset()                    { *m = ConduitUsersResponse{} }
func (m *ConduitUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*ConduitUsersResponse) ProtoMessage()               {}
func (*ConduitUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ConduitUsersResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConduitUsersResponse) GetUser() []*ConduitUser {
	if m != nil {
		return m.User
	}
	return nil
}

type ConduitUser struct {
	Phid         string       `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	UserName     string       `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	RealName     string       `protobuf:"bytes,3,opt,name=real_name,json=realName" json:"real_name,omitempty"`
	Image        string       `protobuf:"bytes,4,opt,name=image" json:"image,omitempty"`
	Uri          string       `protobuf:"bytes,5,opt,name=uri" json:"uri,omitempty"`
	Roles        []string     `protobuf:"bytes,6,rep,name=roles" json:"roles,omitempty"`
	PrimaryEmail string       `protobuf:"bytes,7,opt,name=primary_email,json=primaryEmail" json:"primary_email,omitempty"`
	Preferences  *Preferences `protobuf:"bytes,8,opt,name=preferences" json:"preferences,omitempty"`
}

func (m *ConduitUser) Reset()                    { *m = ConduitUser{} }
func (m *ConduitUser) String() string            { return proto.CompactTextString(m) }
func (*ConduitUser) ProtoMessage()               {}
func (*ConduitUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ConduitUser) GetPreferences() *Preferences {
	if m != nil {
		return m.Preferences
	}
	return nil
}

type Preferences struct {
	TimeZone    string `protobuf:"bytes,1,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`
	TimeFormat  string `protobuf:"bytes,2,opt,name=time_format,json=timeFormat" json:"time_format,omitempty"`
	DateFormate string `protobuf:"bytes,3,opt,name=date_formate,json=dateFormate" json:"date_formate,omitempty"`
}

func (m *Preferences) Reset()                    { *m = Preferences{} }
func (m *Preferences) String() string            { return proto.CompactTextString(m) }
func (*Preferences) ProtoMessage()               {}
func (*Preferences) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*ConduitRequest)(nil), "auth.v1beta1.ConduitRequest")
	proto.RegisterType((*ConduitWhoAmIResponse)(nil), "auth.v1beta1.ConduitWhoAmIResponse")
	proto.RegisterType((*ConduitUsersResponse)(nil), "auth.v1beta1.ConduitUsersResponse")
	proto.RegisterType((*ConduitUser)(nil), "auth.v1beta1.ConduitUser")
	proto.RegisterType((*Preferences)(nil), "auth.v1beta1.Preferences")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Conduit service

type ConduitClient interface {
	// This rpc is used to check a valid user from other applications.
	WhoAmI(ctx context.Context, in *ConduitRequest, opts ...grpc.CallOption) (*ConduitWhoAmIResponse, error)
	// appctl used this to validates the user token with phabricator.
	Users(ctx context.Context, in *ConduitRequest, opts ...grpc.CallOption) (*ConduitUsersResponse, error)
}

type conduitClient struct {
	cc *grpc.ClientConn
}

func NewConduitClient(cc *grpc.ClientConn) ConduitClient {
	return &conduitClient{cc}
}

func (c *conduitClient) WhoAmI(ctx context.Context, in *ConduitRequest, opts ...grpc.CallOption) (*ConduitWhoAmIResponse, error) {
	out := new(ConduitWhoAmIResponse)
	err := grpc.Invoke(ctx, "/auth.v1beta1.Conduit/WhoAmI", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conduitClient) Users(ctx context.Context, in *ConduitRequest, opts ...grpc.CallOption) (*ConduitUsersResponse, error) {
	out := new(ConduitUsersResponse)
	err := grpc.Invoke(ctx, "/auth.v1beta1.Conduit/Users", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Conduit service

type ConduitServer interface {
	// This rpc is used to check a valid user from other applications.
	WhoAmI(context.Context, *ConduitRequest) (*ConduitWhoAmIResponse, error)
	// appctl used this to validates the user token with phabricator.
	Users(context.Context, *ConduitRequest) (*ConduitUsersResponse, error)
}

func RegisterConduitServer(s *grpc.Server, srv ConduitServer) {
	s.RegisterService(&_Conduit_serviceDesc, srv)
}

func _Conduit_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConduitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1beta1.Conduit/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitServer).WhoAmI(ctx, req.(*ConduitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conduit_Users_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConduitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConduitServer).Users(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.v1beta1.Conduit/Users",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConduitServer).Users(ctx, req.(*ConduitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Conduit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1beta1.Conduit",
	HandlerType: (*ConduitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhoAmI",
			Handler:    _Conduit_WhoAmI_Handler,
		},
		{
			MethodName: "Users",
			Handler:    _Conduit_Users_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("conduit.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x93, 0xdf, 0x6a, 0x14, 0x31,
	0x14, 0xc6, 0x99, 0xfd, 0xbf, 0x67, 0xb6, 0xa5, 0x84, 0x2a, 0xe3, 0x5a, 0xb4, 0x4e, 0x41, 0xaa,
	0xe2, 0x0c, 0xad, 0xa0, 0x17, 0x5e, 0xa9, 0x28, 0x78, 0x23, 0x32, 0x22, 0x42, 0x6f, 0x96, 0x74,
	0xe7, 0x74, 0x27, 0xb0, 0x49, 0xc6, 0x24, 0xa3, 0xd4, 0x3b, 0x8b, 0x6f, 0xa0, 0xaf, 0xe2, 0x93,
	0xf8, 0x0a, 0xbe, 0x85, 0x37, 0x92, 0x3f, 0xd5, 0x29, 0x6c, 0x71, 0xc1, 0x9b, 0x21, 0xf9, 0x7d,
	0x5f, 0x72, 0x92, 0x33, 0x5f, 0x60, 0x63, 0x2e, 0x45, 0xd9, 0x30, 0x93, 0xd5, 0x4a, 0x1a, 0x49,
	0x26, 0xb4, 0x31, 0x55, 0xf6, 0xe1, 0xe0, 0x18, 0x0d, 0x3d, 0x98, 0xee, 0x2c, 0xa4, 0x5c, 0x2c,
	0x31, 0xa7, 0x35, 0xcb, 0xa9, 0x10, 0xd2, 0x50, 0xc3, 0xa4, 0xd0, 0xde, 0x3b, 0xbd, 0x41, 0xeb,
	0x5a, 0xcf, 0x65, 0x79, 0x99, 0x7e, 0xd5, 0xe2, 0xd2, 0x9c, 0xd6, 0xa8, 0x73, 0xf7, 0xf5, 0x3c,
	0xdd, 0x82, 0xcd, 0x67, 0xbe, 0x68, 0x81, 0xef, 0x1b, 0xd4, 0x26, 0x15, 0x70, 0x25, 0x90, 0x77,
	0x95, 0x7c, 0xc2, 0x5f, 0x16, 0xa8, 0x6b, 0x29, 0x34, 0x92, 0xdb, 0x30, 0xd0, 0x86, 0x9a, 0x46,
	0x27, 0xd1, 0x6e, 0xb4, 0x1f, 0x1f, 0x6e, 0x66, 0x7e, 0xbf, 0xec, 0x8d, 0xa3, 0x45, 0x50, 0xc9,
	0x7d, 0xe8, 0x35, 0x1a, 0x55, 0xd2, 0x71, 0xae, 0x6b, 0x59, 0xfb, 0x16, 0x59, 0xd8, 0xfa, 0xad,
	0x46, 0x55, 0x38, 0x5b, 0xca, 0x61, 0xbb, 0x05, 0xf5, 0x7f, 0x94, 0xeb, 0xae, 0x53, 0xee, 0x57,
	0x04, 0x71, 0x8b, 0x12, 0x02, 0xbd, 0xba, 0x62, 0xa5, 0x2b, 0x32, 0x2e, 0xdc, 0x98, 0x5c, 0x87,
	0xb1, 0xf5, 0xce, 0x04, 0xe5, 0xe8, 0xae, 0x31, 0x2e, 0x46, 0x16, 0xbc, 0xa2, 0x1c, 0xad, 0xa8,
	0x90, 0x2e, 0xbd, 0xd8, 0xf5, 0xa2, 0x05, 0x4e, 0xdc, 0x86, 0x3e, 0xe3, 0x74, 0x81, 0x49, 0xcf,
	0x09, 0x7e, 0x42, 0xb6, 0xa0, 0xdb, 0x28, 0x96, 0xf4, 0x1d, 0xb3, 0x43, 0xeb, 0x53, 0x72, 0x89,
	0x3a, 0x19, 0xec, 0x76, 0xad, 0xcf, 0x4d, 0xc8, 0x1e, 0x6c, 0xd4, 0x8a, 0x71, 0xaa, 0x4e, 0x67,
	0xc8, 0x29, 0x5b, 0x26, 0x43, 0xb7, 0x62, 0x12, 0xe0, 0x73, 0xcb, 0xc8, 0x63, 0x88, 0x6b, 0x85,
	0x27, 0xa8, 0x50, 0xcc, 0x51, 0x27, 0xa3, 0x55, 0x5d, 0x7e, 0xfd, 0xd7, 0x50, 0xb4, 0xdd, 0xa9,
	0x80, 0xb8, 0xa5, 0xd9, 0xbb, 0x18, 0xc6, 0x71, 0xf6, 0x49, 0x0a, 0x0c, 0x1d, 0x18, 0x59, 0x70,
	0x24, 0x05, 0x92, 0x9b, 0x10, 0x3b, 0xf1, 0x44, 0x2a, 0x4e, 0x4d, 0xe8, 0x03, 0x58, 0xf4, 0xc2,
	0x11, 0x72, 0x0b, 0x26, 0x25, 0x35, 0xe7, 0x86, 0xf3, 0x66, 0xc4, 0x96, 0x79, 0x07, 0x1e, 0x7e,
	0xeb, 0xc0, 0x30, 0x74, 0x9b, 0x7c, 0x89, 0x60, 0xe0, 0x23, 0x45, 0x76, 0x56, 0xfe, 0xa5, 0x90,
	0xc0, 0xe9, 0xde, 0x4a, 0xf5, 0x62, 0x1a, 0xd3, 0x47, 0x67, 0xdf, 0x93, 0xce, 0x28, 0x3a, 0xfb,
	0xf1, 0xf3, 0x6b, 0xe7, 0x1e, 0xb9, 0x93, 0x5f, 0xcc, 0x7f, 0x63, 0xaa, 0x3c, 0xac, 0xcf, 0xc3,
	0xa3, 0xca, 0x3f, 0x56, 0x92, 0x72, 0x46, 0x3e, 0x47, 0xd0, 0x77, 0x49, 0xfb, 0xc7, 0x29, 0xd2,
	0x4b, 0x93, 0xf4, 0x27, 0xa3, 0xe9, 0xc3, 0xd6, 0x21, 0xee, 0x92, 0xfd, 0x35, 0x0e, 0x61, 0x63,
	0xa4, 0x9f, 0x8e, 0x8f, 0x86, 0x41, 0x38, 0x1e, 0xb8, 0x77, 0xf8, 0xe0, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc0, 0xb8, 0x5b, 0xb4, 0xfc, 0x03, 0x00, 0x00,
}
