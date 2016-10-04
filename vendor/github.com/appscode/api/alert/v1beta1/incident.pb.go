// Code generated by protoc-gen-go.
// source: incident.proto
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

type Incident struct {
	Phid                 string      `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	KubernetesCluster    string      `protobuf:"bytes,2,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string      `protobuf:"bytes,3,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string      `protobuf:"bytes,4,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string      `protobuf:"bytes,5,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	Alert                *dtypes.Uid `protobuf:"bytes,6,opt,name=alert" json:"alert,omitempty"`
	IcingaHost           string      `protobuf:"bytes,7,opt,name=icinga_host,json=icingaHost" json:"icinga_host,omitempty"`
	IcingaService        string      `protobuf:"bytes,8,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	Type                 string      `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
	State                string      `protobuf:"bytes,10,opt,name=state" json:"state,omitempty"`
	User                 *dtypes.Uid `protobuf:"bytes,11,opt,name=user" json:"user,omitempty"`
	// Timestamp of first reported event
	Reported int64 `protobuf:"varint,12,opt,name=reported" json:"reported,omitempty"`
	// Timestamp of first acknowledgement
	Acknowledged int64             `protobuf:"varint,13,opt,name=acknowledged" json:"acknowledged,omitempty"`
	Recovered    int64             `protobuf:"varint,14,opt,name=recovered" json:"recovered,omitempty"`
	IcingawebUrl string            `protobuf:"bytes,15,opt,name=icingaweb_url,json=icingawebUrl" json:"icingaweb_url,omitempty"`
	Events       []*Incident_Event `protobuf:"bytes,16,rep,name=events" json:"events,omitempty"`
}

func (m *Incident) Reset()                    { *m = Incident{} }
func (m *Incident) String() string            { return proto.CompactTextString(m) }
func (*Incident) ProtoMessage()               {}
func (*Incident) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Incident) GetAlert() *dtypes.Uid {
	if m != nil {
		return m.Alert
	}
	return nil
}

func (m *Incident) GetUser() *dtypes.Uid {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Incident) GetEvents() []*Incident_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Incident_Event struct {
	Type     string      `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	State    string      `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Reported int64       `protobuf:"varint,3,opt,name=reported" json:"reported,omitempty"`
	User     *dtypes.Uid `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Comment  string      `protobuf:"bytes,5,opt,name=comment" json:"comment,omitempty"`
}

func (m *Incident_Event) Reset()                    { *m = Incident_Event{} }
func (m *Incident_Event) String() string            { return proto.CompactTextString(m) }
func (*Incident_Event) ProtoMessage()               {}
func (*Incident_Event) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *Incident_Event) GetUser() *dtypes.Uid {
	if m != nil {
		return m.User
	}
	return nil
}

// Next Id: 6
type IncidentListRequest struct {
	KubernetesCluster    string   `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string   `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string   `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string   `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	State                []string `protobuf:"bytes,5,rep,name=state" json:"state,omitempty"`
}

func (m *IncidentListRequest) Reset()                    { *m = IncidentListRequest{} }
func (m *IncidentListRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentListRequest) ProtoMessage()               {}
func (*IncidentListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type IncidentListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Incidents []*Incident    `protobuf:"bytes,2,rep,name=incidents" json:"incidents,omitempty"`
}

func (m *IncidentListResponse) Reset()                    { *m = IncidentListResponse{} }
func (m *IncidentListResponse) String() string            { return proto.CompactTextString(m) }
func (*IncidentListResponse) ProtoMessage()               {}
func (*IncidentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *IncidentListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IncidentListResponse) GetIncidents() []*Incident {
	if m != nil {
		return m.Incidents
	}
	return nil
}

// Next Id: 2
type IncidentDescribeRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *IncidentDescribeRequest) Reset()                    { *m = IncidentDescribeRequest{} }
func (m *IncidentDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentDescribeRequest) ProtoMessage()               {}
func (*IncidentDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type IncidentDescribeResponse struct {
	Status   *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Incident *Incident      `protobuf:"bytes,2,opt,name=incident" json:"incident,omitempty"`
}

func (m *IncidentDescribeResponse) Reset()                    { *m = IncidentDescribeResponse{} }
func (m *IncidentDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*IncidentDescribeResponse) ProtoMessage()               {}
func (*IncidentDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *IncidentDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IncidentDescribeResponse) GetIncident() *Incident {
	if m != nil {
		return m.Incident
	}
	return nil
}

// Next Id: 9
type IncidentNotifyRequest struct {
	AlertPhid string `protobuf:"bytes,1,opt,name=alert_phid,json=alertPhid" json:"alert_phid,omitempty"`
	HostName  string `protobuf:"bytes,2,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	State     string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Output    string `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
	Time      string `protobuf:"bytes,6,opt,name=time" json:"time,omitempty"`
	Author    string `protobuf:"bytes,7,opt,name=author" json:"author,omitempty"`
	Comment   string `protobuf:"bytes,8,opt,name=comment" json:"comment,omitempty"`
}

func (m *IncidentNotifyRequest) Reset()                    { *m = IncidentNotifyRequest{} }
func (m *IncidentNotifyRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentNotifyRequest) ProtoMessage()               {}
func (*IncidentNotifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

// Next Id: 4
type IncidentEventCreateRequest struct {
	// Incident PHID
	Phid        string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Comment     string `protobuf:"bytes,2,opt,name=comment" json:"comment,omitempty"`
	Acknowledge bool   `protobuf:"varint,3,opt,name=acknowledge" json:"acknowledge,omitempty"`
}

func (m *IncidentEventCreateRequest) Reset()                    { *m = IncidentEventCreateRequest{} }
func (m *IncidentEventCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*IncidentEventCreateRequest) ProtoMessage()               {}
func (*IncidentEventCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func init() {
	proto.RegisterType((*Incident)(nil), "alert.v1beta1.Incident")
	proto.RegisterType((*Incident_Event)(nil), "alert.v1beta1.Incident.Event")
	proto.RegisterType((*IncidentListRequest)(nil), "alert.v1beta1.IncidentListRequest")
	proto.RegisterType((*IncidentListResponse)(nil), "alert.v1beta1.IncidentListResponse")
	proto.RegisterType((*IncidentDescribeRequest)(nil), "alert.v1beta1.IncidentDescribeRequest")
	proto.RegisterType((*IncidentDescribeResponse)(nil), "alert.v1beta1.IncidentDescribeResponse")
	proto.RegisterType((*IncidentNotifyRequest)(nil), "alert.v1beta1.IncidentNotifyRequest")
	proto.RegisterType((*IncidentEventCreateRequest)(nil), "alert.v1beta1.IncidentEventCreateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Incidents service

type IncidentsClient interface {
	List(ctx context.Context, in *IncidentListRequest, opts ...grpc.CallOption) (*IncidentListResponse, error)
	Describe(ctx context.Context, in *IncidentDescribeRequest, opts ...grpc.CallOption) (*IncidentDescribeResponse, error)
	Notify(ctx context.Context, in *IncidentNotifyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	CreateEvent(ctx context.Context, in *IncidentEventCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type incidentsClient struct {
	cc *grpc.ClientConn
}

func NewIncidentsClient(cc *grpc.ClientConn) IncidentsClient {
	return &incidentsClient{cc}
}

func (c *incidentsClient) List(ctx context.Context, in *IncidentListRequest, opts ...grpc.CallOption) (*IncidentListResponse, error) {
	out := new(IncidentListResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Incidents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) Describe(ctx context.Context, in *IncidentDescribeRequest, opts ...grpc.CallOption) (*IncidentDescribeResponse, error) {
	out := new(IncidentDescribeResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Incidents/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) Notify(ctx context.Context, in *IncidentNotifyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Incidents/Notify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incidentsClient) CreateEvent(ctx context.Context, in *IncidentEventCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/alert.v1beta1.Incidents/CreateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Incidents service

type IncidentsServer interface {
	List(context.Context, *IncidentListRequest) (*IncidentListResponse, error)
	Describe(context.Context, *IncidentDescribeRequest) (*IncidentDescribeResponse, error)
	Notify(context.Context, *IncidentNotifyRequest) (*dtypes.VoidResponse, error)
	CreateEvent(context.Context, *IncidentEventCreateRequest) (*dtypes.VoidResponse, error)
}

func RegisterIncidentsServer(s *grpc.Server, srv IncidentsServer) {
	s.RegisterService(&_Incidents_serviceDesc, srv)
}

func _Incidents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Incidents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).List(ctx, req.(*IncidentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Incidents/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).Describe(ctx, req.(*IncidentDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Incidents/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).Notify(ctx, req.(*IncidentNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incidents_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncidentEventCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncidentsServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alert.v1beta1.Incidents/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncidentsServer).CreateEvent(ctx, req.(*IncidentEventCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Incidents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alert.v1beta1.Incidents",
	HandlerType: (*IncidentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Incidents_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Incidents_Describe_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Incidents_Notify_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _Incidents_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("incident.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x41, 0x8f, 0xdb, 0x44,
	0x14, 0x96, 0x13, 0x27, 0x6b, 0x3f, 0xef, 0x06, 0x98, 0xa6, 0xed, 0xc8, 0xb4, 0x6a, 0x70, 0x69,
	0x1b, 0x56, 0xac, 0xcd, 0x26, 0x54, 0x42, 0x3d, 0x70, 0xa0, 0x20, 0x81, 0x84, 0x0a, 0x72, 0x29,
	0x07, 0x2e, 0x91, 0x63, 0x3f, 0xb2, 0xa6, 0x89, 0xc7, 0x78, 0xc6, 0x59, 0xad, 0x10, 0x1c, 0x38,
	0x71, 0x42, 0x42, 0xfc, 0x09, 0xfe, 0x0f, 0x47, 0xae, 0xf0, 0x1b, 0xb8, 0x22, 0x8f, 0x3d, 0x8e,
	0xd3, 0x8d, 0x43, 0xb7, 0x97, 0x55, 0xe6, 0x7b, 0xdf, 0x7b, 0xf3, 0xe6, 0x7b, 0x5f, 0x5e, 0x16,
	0x06, 0x71, 0x12, 0xc6, 0x11, 0x26, 0xc2, 0x4d, 0x33, 0x26, 0x18, 0x39, 0x0a, 0x96, 0x98, 0x09,
	0x77, 0x7d, 0x3a, 0x47, 0x11, 0x9c, 0xda, 0xb7, 0x16, 0x8c, 0x2d, 0x96, 0xe8, 0x05, 0x69, 0xec,
	0x05, 0x49, 0xc2, 0x44, 0x20, 0x62, 0x96, 0xf0, 0x92, 0x6c, 0xdf, 0x28, 0xe0, 0x48, 0x5c, 0xa4,
	0xc8, 0x3d, 0xf9, 0xb7, 0xc4, 0x9d, 0x7f, 0x7a, 0x60, 0x7c, 0x56, 0xd5, 0x25, 0x04, 0xf4, 0xf4,
	0x2c, 0x8e, 0xa8, 0x36, 0xd2, 0xc6, 0xa6, 0x2f, 0x3f, 0x93, 0x13, 0x20, 0xcf, 0xf3, 0x39, 0x66,
	0x09, 0x0a, 0xe4, 0xb3, 0x70, 0x99, 0x73, 0x81, 0x19, 0xed, 0x48, 0xc6, 0x1b, 0x9b, 0xc8, 0xe3,
	0x32, 0x40, 0x4e, 0x61, 0xd8, 0xa0, 0x27, 0xc1, 0x0a, 0x79, 0x1a, 0x84, 0x48, 0xbb, 0x32, 0xe1,
	0xda, 0x26, 0xf6, 0x44, 0x85, 0xc8, 0x14, 0xae, 0x37, 0x52, 0xd8, 0xfc, 0x3b, 0x0c, 0xc5, 0x57,
	0x17, 0x29, 0x52, 0x5d, 0xe6, 0x34, 0xea, 0x7d, 0x51, 0xc7, 0x76, 0x26, 0x15, 0x25, 0x69, 0x6f,
	0x77, 0x52, 0x11, 0x23, 0x6f, 0x41, 0x4f, 0x6a, 0x46, 0xfb, 0x23, 0x6d, 0x6c, 0x4d, 0x2c, 0xb7,
	0x14, 0xc4, 0x7d, 0x16, 0x47, 0x7e, 0x19, 0x21, 0x77, 0xc0, 0x8a, 0xc3, 0x38, 0x59, 0x04, 0xb3,
	0x33, 0xc6, 0x05, 0x3d, 0x90, 0xd5, 0xa0, 0x84, 0x3e, 0x65, 0x5c, 0x90, 0x7b, 0x30, 0xa8, 0x08,
	0x1c, 0xb3, 0x75, 0x1c, 0x22, 0x35, 0x24, 0xe7, 0xa8, 0x44, 0x9f, 0x96, 0x60, 0x21, 0x65, 0x51,
	0x9b, 0x9a, 0xa5, 0x94, 0xc5, 0x67, 0x32, 0x84, 0x1e, 0x17, 0x81, 0x40, 0x0a, 0x12, 0x2c, 0x0f,
	0xe4, 0x0e, 0xe8, 0x39, 0xc7, 0x8c, 0x5a, 0x97, 0x7b, 0x92, 0x01, 0x62, 0x83, 0x91, 0x61, 0xca,
	0x32, 0x81, 0x11, 0x3d, 0x1c, 0x69, 0xe3, 0xae, 0x5f, 0x9f, 0x89, 0x03, 0x87, 0x41, 0xf8, 0x3c,
	0x61, 0xe7, 0x4b, 0x8c, 0x16, 0x18, 0xd1, 0x23, 0x19, 0xdf, 0xc2, 0xc8, 0x2d, 0x30, 0x33, 0x0c,
	0xd9, 0x1a, 0x33, 0x8c, 0xe8, 0x40, 0x12, 0x36, 0x00, 0xb9, 0x0b, 0x55, 0xe7, 0xe7, 0x38, 0x9f,
	0xe5, 0xd9, 0x92, 0xbe, 0x26, 0x9b, 0x3b, 0xac, 0xc1, 0x67, 0xd9, 0x92, 0x3c, 0x84, 0x3e, 0xae,
	0x31, 0x11, 0x9c, 0xbe, 0x3e, 0xea, 0x8e, 0xad, 0xc9, 0x6d, 0x77, 0xcb, 0x7b, 0xae, 0x72, 0x90,
	0xfb, 0x49, 0xc1, 0xf2, 0x2b, 0xb2, 0xfd, 0x8b, 0x06, 0x3d, 0x89, 0xd4, 0x72, 0x68, 0xbb, 0xe4,
	0xe8, 0x34, 0xe5, 0x68, 0xbe, 0xb6, 0xfb, 0xc2, 0x6b, 0x95, 0x54, 0x7a, 0x9b, 0x54, 0x14, 0x0e,
	0x42, 0xb6, 0x5a, 0x61, 0x22, 0x2a, 0x1f, 0xa8, 0xa3, 0xf3, 0xaf, 0x06, 0xd7, 0x54, 0x97, 0x9f,
	0xc7, 0x5c, 0xf8, 0xf8, 0x7d, 0x8e, 0x5c, 0xb4, 0xd8, 0x5b, 0xbb, 0xaa, 0xbd, 0x3b, 0xaf, 0x60,
	0xef, 0xee, 0xab, 0xd8, 0x5b, 0xdf, 0x63, 0xef, 0x5a, 0xd0, 0xde, 0xa8, 0x5b, 0x0b, 0xea, 0xe4,
	0x30, 0xdc, 0x7e, 0x38, 0x4f, 0x59, 0xc2, 0x91, 0xdc, 0x87, 0x7e, 0x41, 0xc8, 0xb9, 0x7c, 0xad,
	0x35, 0x19, 0x28, 0x39, 0x9f, 0x4a, 0xd4, 0xaf, 0xa2, 0xe4, 0x21, 0x98, 0x6a, 0xf1, 0x70, 0xda,
	0x91, 0xe3, 0xbf, 0xd9, 0x32, 0x7e, 0x7f, 0xc3, 0x74, 0x4e, 0xe0, 0xa6, 0x82, 0x3f, 0x46, 0x1e,
	0x66, 0xf1, 0x1c, 0x95, 0xe6, 0x3b, 0xd6, 0x8c, 0x73, 0x0e, 0xf4, 0x32, 0xfd, 0x8a, 0x9d, 0x4e,
	0xc1, 0x50, 0xf7, 0xcb, 0x81, 0xec, 0x69, 0xb4, 0x26, 0x3a, 0x7f, 0x69, 0x70, 0x5d, 0xc1, 0x4f,
	0x98, 0x88, 0xbf, 0xbd, 0x50, 0x6d, 0xde, 0x06, 0x90, 0xd9, 0xb3, 0x46, 0xb3, 0xa6, 0x44, 0xbe,
	0x2c, 0x16, 0xe3, 0x9b, 0x60, 0x16, 0x2b, 0x42, 0x9a, 0xa0, 0x9a, 0xbf, 0x51, 0x00, 0x72, 0x14,
	0xca, 0xef, 0xdd, 0x5d, 0x7e, 0xd7, 0x9b, 0x7e, 0xbf, 0x01, 0x7d, 0x96, 0x8b, 0x34, 0x57, 0x8e,
	0xad, 0x4e, 0xb2, 0x42, 0xbc, 0x42, 0xb9, 0xaa, 0x8a, 0x0a, 0xf1, 0x4a, 0x72, 0x83, 0x5c, 0x9c,
	0xb1, 0xac, 0xda, 0x4b, 0xd5, 0xa9, 0x69, 0x7b, 0x63, 0xdb, 0xf6, 0x4b, 0xb0, 0xd5, 0xe3, 0xe4,
	0x17, 0xf1, 0x71, 0x86, 0x81, 0xd8, 0x37, 0x88, 0x66, 0xad, 0xce, 0x56, 0x2d, 0x32, 0x02, 0xab,
	0xb1, 0x57, 0xe4, 0xd3, 0x0c, 0xbf, 0x09, 0x4d, 0xfe, 0xd0, 0xc1, 0x54, 0xd7, 0x71, 0xf2, 0x13,
	0xe8, 0x85, 0xe1, 0x88, 0xd3, 0x32, 0x84, 0xc6, 0xd7, 0xd0, 0xbe, 0xbb, 0x97, 0x53, 0xfa, 0xc0,
	0x39, 0xf9, 0xf9, 0xcf, 0xbf, 0x7f, 0xef, 0x3c, 0x20, 0xf7, 0xbc, 0x20, 0x4d, 0x79, 0xc8, 0xa2,
	0xea, 0xc7, 0xae, 0xc8, 0xf4, 0xaa, 0x4c, 0xaf, 0x76, 0x20, 0xf9, 0x4d, 0x03, 0x43, 0x79, 0x89,
	0xdc, 0x6f, 0xb9, 0xe0, 0x05, 0x6f, 0xda, 0x0f, 0xfe, 0x97, 0x57, 0x35, 0xf3, 0xbe, 0x6c, 0xc6,
	0x25, 0xef, 0xbe, 0x54, 0x33, 0xde, 0x0f, 0x85, 0xb8, 0x3f, 0x92, 0x35, 0xf4, 0x4b, 0x93, 0x91,
	0xb7, 0x5b, 0x2e, 0xda, 0xf2, 0xa0, 0x3d, 0x54, 0x56, 0xff, 0x9a, 0xc5, 0x51, 0x7d, 0xf7, 0x7b,
	0xf2, 0xee, 0x63, 0xe7, 0xe5, 0x84, 0x78, 0xa4, 0x1d, 0x93, 0x5f, 0x35, 0xb0, 0xca, 0xd9, 0x97,
	0xfb, 0xf8, 0x9d, 0x96, 0xdb, 0x2f, 0x9b, 0xa4, 0xa5, 0x85, 0x0f, 0x65, 0x0b, 0x1f, 0x38, 0xd3,
	0xab, 0x3c, 0xdf, 0x2b, 0x7f, 0x17, 0x1e, 0x69, 0xc7, 0x1f, 0x99, 0xdf, 0x1c, 0x54, 0xac, 0x79,
	0x5f, 0xfe, 0x27, 0x32, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x01, 0xd0, 0x99, 0x92, 0xe0, 0x08,
	0x00, 0x00,
}