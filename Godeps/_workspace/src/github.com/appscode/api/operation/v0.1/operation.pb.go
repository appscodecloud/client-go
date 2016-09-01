// Code generated by protoc-gen-go.
// source: operation.proto
// DO NOT EDIT!

/*
Package operation is a generated protocol buffer package.

It is generated from these files:
	operation.proto

It has these top-level messages:
	DescribeRequest
	DescribeResponse
	Auth
	Metadata
	Operation
	DataBucketDeleteRequest
	NamespaceAdminTaskRequest
*/
package operation

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"
import certificate "github.com/appscode/api/certificate/v0.1"
import ci "github.com/appscode/api/ci/v0.1"
import ci1 "github.com/appscode/api/ci/v0.1"
import kubernetes "github.com/appscode/api/kubernetes/v0.1"
import db "github.com/appscode/api/db/v0.1"
import namespace "github.com/appscode/api/namespace/v0.1"

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

// Next Id: 11
type OperationType int32

const (
	OperationType_UNKNOWN              OperationType = 0
	OperationType_CLUSTER_CREATE       OperationType = 1
	OperationType_CLUSTER_SCALE        OperationType = 2
	OperationType_CLUSTER_DELETE       OperationType = 3
	OperationType_CLUSTER_UPGRADE      OperationType = 4
	OperationType_CI_MASTER_CREATE     OperationType = 5
	OperationType_CI_MASTER_DELETE     OperationType = 6
	OperationType_CI_AGENT_CREATE      OperationType = 7
	OperationType_CI_AGENT_DELETE      OperationType = 8
	OperationType_DATA_BUCKET_DELETE   OperationType = 9
	OperationType_DATABASE_BACKUP      OperationType = 10
	OperationType_NAMESPACE_CREATE     OperationType = 11
	OperationType_NAMESPACE_ADMIN_TASK OperationType = 12
	OperationType_CHECK_DNS            OperationType = 13
)

var OperationType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CLUSTER_CREATE",
	2:  "CLUSTER_SCALE",
	3:  "CLUSTER_DELETE",
	4:  "CLUSTER_UPGRADE",
	5:  "CI_MASTER_CREATE",
	6:  "CI_MASTER_DELETE",
	7:  "CI_AGENT_CREATE",
	8:  "CI_AGENT_DELETE",
	9:  "DATA_BUCKET_DELETE",
	10: "DATABASE_BACKUP",
	11: "NAMESPACE_CREATE",
	12: "NAMESPACE_ADMIN_TASK",
	13: "CHECK_DNS",
}
var OperationType_value = map[string]int32{
	"UNKNOWN":              0,
	"CLUSTER_CREATE":       1,
	"CLUSTER_SCALE":        2,
	"CLUSTER_DELETE":       3,
	"CLUSTER_UPGRADE":      4,
	"CI_MASTER_CREATE":     5,
	"CI_MASTER_DELETE":     6,
	"CI_AGENT_CREATE":      7,
	"CI_AGENT_DELETE":      8,
	"DATA_BUCKET_DELETE":   9,
	"DATABASE_BACKUP":      10,
	"NAMESPACE_CREATE":     11,
	"NAMESPACE_ADMIN_TASK": 12,
	"CHECK_DNS":            13,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}
func (OperationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DescribeRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DescribeResponse struct {
	Status *dtypes.Status          `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Op     *Operation              `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	Log    []*DescribeResponse_Log `protobuf:"bytes,3,rep,name=log" json:"log,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeResponse) GetOp() *Operation {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *DescribeResponse) GetLog() []*DescribeResponse_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type DescribeResponse_Log struct {
	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *DescribeResponse_Log) Reset()                    { *m = DescribeResponse_Log{} }
func (m *DescribeResponse_Log) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse_Log) ProtoMessage()               {}
func (*DescribeResponse_Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Auth struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	TokenType string `protobuf:"bytes,4,opt,name=token_type,json=tokenType" json:"token_type,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Metadata holds other on request or operation specific data
// that could be used inside that operation.
// An welldefined message instead of a map is used
// so that the data fields can be explicitly defined with
// its own data type. Resolves data convertions.
type Metadata struct {
	// Contains PurchasePHID is this is a purchase request.
	PurchasePhid []string `protobuf:"bytes,1,rep,name=purchase_phid,json=purchasePhid" json:"purchase_phid,omitempty"`
	// PHID of the user who requested this operation.
	AuthorPhid string `protobuf:"bytes,2,opt,name=author_phid,json=authorPhid" json:"author_phid,omitempty"`
	AuthorName string `protobuf:"bytes,3,opt,name=author_name,json=authorName" json:"author_name,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Next Id: 16
type Operation struct {
	// Types that are valid to be assigned to Request:
	//	*Operation_ClusterCreateRequest
	//	*Operation_ClusterScaleRequest
	//	*Operation_ClusterDeleteRequest
	//	*Operation_ClusterUpgradeRequest
	//	*Operation_CiMasterCreateRequest
	//	*Operation_CiMasterDeleteRequest
	//	*Operation_CiAgentCreateRequest
	//	*Operation_CiAgentDeleteRequest
	//	*Operation_DataBucketDeleteRequest
	//	*Operation_DbBackupScheduleRequest
	//	*Operation_NamespaceCreateRequest
	//	*Operation_NamespaceAdminTaskRequest
	//	*Operation_DnsCheckRequest
	Request  isOperation_Request `protobuf_oneof:"request"`
	Type     OperationType       `protobuf:"varint,14,opt,name=type,enum=operation.OperationType" json:"type,omitempty"`
	Phid     string              `protobuf:"bytes,15,opt,name=phid" json:"phid,omitempty"`
	Auth     *Auth               `protobuf:"bytes,16,opt,name=auth" json:"auth,omitempty"`
	Metadata *Metadata           `protobuf:"bytes,17,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isOperation_Request interface {
	isOperation_Request()
}

type Operation_ClusterCreateRequest struct {
	ClusterCreateRequest *kubernetes.ClusterCreateRequest `protobuf:"bytes,1,opt,name=cluster_create_request,json=clusterCreateRequest,oneof"`
}
type Operation_ClusterScaleRequest struct {
	ClusterScaleRequest *kubernetes.ClusterScaleRequest `protobuf:"bytes,2,opt,name=cluster_scale_request,json=clusterScaleRequest,oneof"`
}
type Operation_ClusterDeleteRequest struct {
	ClusterDeleteRequest *kubernetes.ClusterDeleteRequest `protobuf:"bytes,3,opt,name=cluster_delete_request,json=clusterDeleteRequest,oneof"`
}
type Operation_ClusterUpgradeRequest struct {
	ClusterUpgradeRequest *kubernetes.ClusterUpgradeRequest `protobuf:"bytes,4,opt,name=cluster_upgrade_request,json=clusterUpgradeRequest,oneof"`
}
type Operation_CiMasterCreateRequest struct {
	CiMasterCreateRequest *ci1.MasterCreateRequest `protobuf:"bytes,5,opt,name=ci_master_create_request,json=ciMasterCreateRequest,oneof"`
}
type Operation_CiMasterDeleteRequest struct {
	CiMasterDeleteRequest *ci1.MasterDeleteRequest `protobuf:"bytes,6,opt,name=ci_master_delete_request,json=ciMasterDeleteRequest,oneof"`
}
type Operation_CiAgentCreateRequest struct {
	CiAgentCreateRequest *ci.AgentCreateRequest `protobuf:"bytes,7,opt,name=ci_agent_create_request,json=ciAgentCreateRequest,oneof"`
}
type Operation_CiAgentDeleteRequest struct {
	CiAgentDeleteRequest *ci.AgentDeleteRequest `protobuf:"bytes,8,opt,name=ci_agent_delete_request,json=ciAgentDeleteRequest,oneof"`
}
type Operation_DataBucketDeleteRequest struct {
	DataBucketDeleteRequest *DataBucketDeleteRequest `protobuf:"bytes,9,opt,name=data_bucket_delete_request,json=dataBucketDeleteRequest,oneof"`
}
type Operation_DbBackupScheduleRequest struct {
	DbBackupScheduleRequest *db.BackupScheduleRequest `protobuf:"bytes,10,opt,name=db_backup_schedule_request,json=dbBackupScheduleRequest,oneof"`
}
type Operation_NamespaceCreateRequest struct {
	NamespaceCreateRequest *namespace.CreateRequest `protobuf:"bytes,11,opt,name=namespace_create_request,json=namespaceCreateRequest,oneof"`
}
type Operation_NamespaceAdminTaskRequest struct {
	NamespaceAdminTaskRequest *NamespaceAdminTaskRequest `protobuf:"bytes,12,opt,name=namespace_admin_task_request,json=namespaceAdminTaskRequest,oneof"`
}
type Operation_DnsCheckRequest struct {
	DnsCheckRequest *certificate.DNSCheckRequest `protobuf:"bytes,13,opt,name=dns_check_request,json=dnsCheckRequest,oneof"`
}

func (*Operation_ClusterCreateRequest) isOperation_Request()      {}
func (*Operation_ClusterScaleRequest) isOperation_Request()       {}
func (*Operation_ClusterDeleteRequest) isOperation_Request()      {}
func (*Operation_ClusterUpgradeRequest) isOperation_Request()     {}
func (*Operation_CiMasterCreateRequest) isOperation_Request()     {}
func (*Operation_CiMasterDeleteRequest) isOperation_Request()     {}
func (*Operation_CiAgentCreateRequest) isOperation_Request()      {}
func (*Operation_CiAgentDeleteRequest) isOperation_Request()      {}
func (*Operation_DataBucketDeleteRequest) isOperation_Request()   {}
func (*Operation_DbBackupScheduleRequest) isOperation_Request()   {}
func (*Operation_NamespaceCreateRequest) isOperation_Request()    {}
func (*Operation_NamespaceAdminTaskRequest) isOperation_Request() {}
func (*Operation_DnsCheckRequest) isOperation_Request()           {}

func (m *Operation) GetRequest() isOperation_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Operation) GetClusterCreateRequest() *kubernetes.ClusterCreateRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterCreateRequest); ok {
		return x.ClusterCreateRequest
	}
	return nil
}

func (m *Operation) GetClusterScaleRequest() *kubernetes.ClusterScaleRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterScaleRequest); ok {
		return x.ClusterScaleRequest
	}
	return nil
}

func (m *Operation) GetClusterDeleteRequest() *kubernetes.ClusterDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterDeleteRequest); ok {
		return x.ClusterDeleteRequest
	}
	return nil
}

func (m *Operation) GetClusterUpgradeRequest() *kubernetes.ClusterUpgradeRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterUpgradeRequest); ok {
		return x.ClusterUpgradeRequest
	}
	return nil
}

func (m *Operation) GetCiMasterCreateRequest() *ci1.MasterCreateRequest {
	if x, ok := m.GetRequest().(*Operation_CiMasterCreateRequest); ok {
		return x.CiMasterCreateRequest
	}
	return nil
}

func (m *Operation) GetCiMasterDeleteRequest() *ci1.MasterDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_CiMasterDeleteRequest); ok {
		return x.CiMasterDeleteRequest
	}
	return nil
}

func (m *Operation) GetCiAgentCreateRequest() *ci.AgentCreateRequest {
	if x, ok := m.GetRequest().(*Operation_CiAgentCreateRequest); ok {
		return x.CiAgentCreateRequest
	}
	return nil
}

func (m *Operation) GetCiAgentDeleteRequest() *ci.AgentDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_CiAgentDeleteRequest); ok {
		return x.CiAgentDeleteRequest
	}
	return nil
}

func (m *Operation) GetDataBucketDeleteRequest() *DataBucketDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_DataBucketDeleteRequest); ok {
		return x.DataBucketDeleteRequest
	}
	return nil
}

func (m *Operation) GetDbBackupScheduleRequest() *db.BackupScheduleRequest {
	if x, ok := m.GetRequest().(*Operation_DbBackupScheduleRequest); ok {
		return x.DbBackupScheduleRequest
	}
	return nil
}

func (m *Operation) GetNamespaceCreateRequest() *namespace.CreateRequest {
	if x, ok := m.GetRequest().(*Operation_NamespaceCreateRequest); ok {
		return x.NamespaceCreateRequest
	}
	return nil
}

func (m *Operation) GetNamespaceAdminTaskRequest() *NamespaceAdminTaskRequest {
	if x, ok := m.GetRequest().(*Operation_NamespaceAdminTaskRequest); ok {
		return x.NamespaceAdminTaskRequest
	}
	return nil
}

func (m *Operation) GetDnsCheckRequest() *certificate.DNSCheckRequest {
	if x, ok := m.GetRequest().(*Operation_DnsCheckRequest); ok {
		return x.DnsCheckRequest
	}
	return nil
}

func (m *Operation) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Operation) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_ClusterCreateRequest)(nil),
		(*Operation_ClusterScaleRequest)(nil),
		(*Operation_ClusterDeleteRequest)(nil),
		(*Operation_ClusterUpgradeRequest)(nil),
		(*Operation_CiMasterCreateRequest)(nil),
		(*Operation_CiMasterDeleteRequest)(nil),
		(*Operation_CiAgentCreateRequest)(nil),
		(*Operation_CiAgentDeleteRequest)(nil),
		(*Operation_DataBucketDeleteRequest)(nil),
		(*Operation_DbBackupScheduleRequest)(nil),
		(*Operation_NamespaceCreateRequest)(nil),
		(*Operation_NamespaceAdminTaskRequest)(nil),
		(*Operation_DnsCheckRequest)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterCreateRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterCreateRequest); err != nil {
			return err
		}
	case *Operation_ClusterScaleRequest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterScaleRequest); err != nil {
			return err
		}
	case *Operation_ClusterDeleteRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterDeleteRequest); err != nil {
			return err
		}
	case *Operation_ClusterUpgradeRequest:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterUpgradeRequest); err != nil {
			return err
		}
	case *Operation_CiMasterCreateRequest:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiMasterCreateRequest); err != nil {
			return err
		}
	case *Operation_CiMasterDeleteRequest:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiMasterDeleteRequest); err != nil {
			return err
		}
	case *Operation_CiAgentCreateRequest:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiAgentCreateRequest); err != nil {
			return err
		}
	case *Operation_CiAgentDeleteRequest:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiAgentDeleteRequest); err != nil {
			return err
		}
	case *Operation_DataBucketDeleteRequest:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DataBucketDeleteRequest); err != nil {
			return err
		}
	case *Operation_DbBackupScheduleRequest:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DbBackupScheduleRequest); err != nil {
			return err
		}
	case *Operation_NamespaceCreateRequest:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NamespaceCreateRequest); err != nil {
			return err
		}
	case *Operation_NamespaceAdminTaskRequest:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NamespaceAdminTaskRequest); err != nil {
			return err
		}
	case *Operation_DnsCheckRequest:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DnsCheckRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Request has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 1: // request.cluster_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterCreateRequest{msg}
		return true, err
	case 2: // request.cluster_scale_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterScaleRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterScaleRequest{msg}
		return true, err
	case 3: // request.cluster_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterDeleteRequest{msg}
		return true, err
	case 4: // request.cluster_upgrade_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterUpgradeRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterUpgradeRequest{msg}
		return true, err
	case 5: // request.ci_master_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci1.MasterCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiMasterCreateRequest{msg}
		return true, err
	case 6: // request.ci_master_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci1.MasterDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiMasterDeleteRequest{msg}
		return true, err
	case 7: // request.ci_agent_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci.AgentCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiAgentCreateRequest{msg}
		return true, err
	case 8: // request.ci_agent_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci.AgentDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiAgentDeleteRequest{msg}
		return true, err
	case 9: // request.data_bucket_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DataBucketDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DataBucketDeleteRequest{msg}
		return true, err
	case 10: // request.db_backup_schedule_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(db.BackupScheduleRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DbBackupScheduleRequest{msg}
		return true, err
	case 11: // request.namespace_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(namespace.CreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_NamespaceCreateRequest{msg}
		return true, err
	case 12: // request.namespace_admin_task_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NamespaceAdminTaskRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_NamespaceAdminTaskRequest{msg}
		return true, err
	case 13: // request.dns_check_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(certificate.DNSCheckRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DnsCheckRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterCreateRequest:
		s := proto.Size(x.ClusterCreateRequest)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterScaleRequest:
		s := proto.Size(x.ClusterScaleRequest)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterDeleteRequest:
		s := proto.Size(x.ClusterDeleteRequest)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterUpgradeRequest:
		s := proto.Size(x.ClusterUpgradeRequest)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiMasterCreateRequest:
		s := proto.Size(x.CiMasterCreateRequest)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiMasterDeleteRequest:
		s := proto.Size(x.CiMasterDeleteRequest)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiAgentCreateRequest:
		s := proto.Size(x.CiAgentCreateRequest)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiAgentDeleteRequest:
		s := proto.Size(x.CiAgentDeleteRequest)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DataBucketDeleteRequest:
		s := proto.Size(x.DataBucketDeleteRequest)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DbBackupScheduleRequest:
		s := proto.Size(x.DbBackupScheduleRequest)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_NamespaceCreateRequest:
		s := proto.Size(x.NamespaceCreateRequest)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_NamespaceAdminTaskRequest:
		s := proto.Size(x.NamespaceAdminTaskRequest)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DnsCheckRequest:
		s := proto.Size(x.DnsCheckRequest)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DataBucketDeleteRequest struct {
	DataType  string `protobuf:"bytes,1,opt,name=data_type,json=dataType" json:"data_type,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Prefix    string `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *DataBucketDeleteRequest) Reset()                    { *m = DataBucketDeleteRequest{} }
func (m *DataBucketDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DataBucketDeleteRequest) ProtoMessage()               {}
func (*DataBucketDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type NamespaceAdminTaskRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *NamespaceAdminTaskRequest) Reset()                    { *m = NamespaceAdminTaskRequest{} }
func (m *NamespaceAdminTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*NamespaceAdminTaskRequest) ProtoMessage()               {}
func (*NamespaceAdminTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*DescribeRequest)(nil), "operation.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "operation.DescribeResponse")
	proto.RegisterType((*DescribeResponse_Log)(nil), "operation.DescribeResponse.Log")
	proto.RegisterType((*Auth)(nil), "operation.Auth")
	proto.RegisterType((*Metadata)(nil), "operation.Metadata")
	proto.RegisterType((*Operation)(nil), "operation.Operation")
	proto.RegisterType((*DataBucketDeleteRequest)(nil), "operation.DataBucketDeleteRequest")
	proto.RegisterType((*NamespaceAdminTaskRequest)(nil), "operation.NamespaceAdminTaskRequest")
	proto.RegisterEnum("operation.OperationType", OperationType_name, OperationType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Operations service

type OperationsClient interface {
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
}

type operationsClient struct {
	cc *grpc.ClientConn
}

func NewOperationsClient(cc *grpc.ClientConn) OperationsClient {
	return &operationsClient{cc}
}

func (c *operationsClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/operation.Operations/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Operations service

type OperationsServer interface {
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
}

func RegisterOperationsServer(s *grpc.Server, srv OperationsServer) {
	s.RegisterService(&_Operations_serviceDesc, srv)
}

func _Operations_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/operation.Operations/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "operation.Operations",
	HandlerType: (*OperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Operations_Describe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("operation.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x8d, 0x25, 0xd9, 0x16, 0xaf, 0x2d, 0x8b, 0x1e, 0x3b, 0x36, 0xad, 0xb8, 0xb0, 0xcb, 0xa6,
	0x41, 0x51, 0x14, 0x54, 0x92, 0x02, 0x05, 0xba, 0xe8, 0x82, 0xa2, 0x88, 0xc4, 0xb5, 0x2d, 0x1b,
	0x94, 0x84, 0x16, 0xe8, 0x82, 0x18, 0x92, 0x13, 0x99, 0xd0, 0x83, 0x0c, 0x1f, 0x45, 0x8b, 0xa0,
	0x9b, 0xee, 0xba, 0xee, 0x4f, 0x74, 0xd5, 0x8f, 0x69, 0x7f, 0xa1, 0x1f, 0xd2, 0x99, 0xe1, 0x9b,
	0x92, 0xb3, 0x11, 0x38, 0xe7, 0x9e, 0x39, 0x73, 0xe7, 0xcc, 0xd5, 0x9d, 0x81, 0xae, 0xe7, 0x93,
	0x00, 0x47, 0xae, 0xb7, 0x52, 0xfc, 0xc0, 0x8b, 0x3c, 0x24, 0xe4, 0x40, 0xef, 0x7c, 0xe6, 0x79,
	0xb3, 0x05, 0xe9, 0x63, 0xdf, 0xed, 0xe3, 0xd5, 0xca, 0x8b, 0x38, 0x1c, 0x26, 0xc4, 0xde, 0x09,
	0x83, 0x9d, 0xe8, 0x57, 0x9f, 0x84, 0x7d, 0xfe, 0x9b, 0xe2, 0x2f, 0x18, 0x6e, 0x93, 0x20, 0x72,
	0xdf, 0xb9, 0x36, 0x8e, 0x48, 0xff, 0xe7, 0x97, 0xca, 0xab, 0x32, 0x90, 0xf2, 0x4e, 0x39, 0xcf,
	0x4d, 0xc2, 0x78, 0x46, 0x56, 0x51, 0x1a, 0x90, 0xca, 0x81, 0x25, 0x0e, 0x23, 0x12, 0xa4, 0x11,
	0x99, 0x45, 0xe6, 0xb1, 0x45, 0x82, 0x15, 0x89, 0xe8, 0xb2, 0x89, 0xf2, 0x22, 0x66, 0x94, 0x6c,
	0xf9, 0x1e, 0x4f, 0xcb, 0x4a, 0x62, 0x0e, 0x8e, 0xb0, 0x85, 0x43, 0x52, 0x9e, 0xbf, 0xc2, 0x4b,
	0x12, 0xfa, 0xd8, 0x4e, 0x13, 0xcb, 0x87, 0x09, 0x47, 0xfe, 0x1c, 0xba, 0x43, 0x12, 0xda, 0x81,
	0x6b, 0x11, 0x83, 0xbc, 0x8f, 0x49, 0x18, 0x21, 0x04, 0x2d, 0xff, 0xc1, 0x75, 0xa4, 0xad, 0xcb,
	0xad, 0x2f, 0x04, 0x83, 0x7f, 0xcb, 0xff, 0x6c, 0x81, 0x58, 0xf0, 0x42, 0x9f, 0xfa, 0x42, 0xd0,
	0x0b, 0xd8, 0x09, 0xa9, 0x49, 0x71, 0xc8, 0xa9, 0x7b, 0xaf, 0x0f, 0x94, 0xc4, 0x1f, 0x65, 0xcc,
	0x51, 0x23, 0x8d, 0xa2, 0xe7, 0xd0, 0xf0, 0x7c, 0xa9, 0xc1, 0x39, 0xc7, 0x4a, 0x71, 0x02, 0x77,
	0xd9, 0x97, 0x41, 0xe3, 0xe8, 0x15, 0x34, 0x17, 0xde, 0x4c, 0x6a, 0x5e, 0x36, 0x29, 0xed, 0xa2,
	0x44, 0xab, 0xaf, 0xab, 0xdc, 0x78, 0x33, 0x83, 0x71, 0x7b, 0xdf, 0x41, 0x93, 0x7e, 0xa3, 0x73,
	0x10, 0x22, 0x97, 0x6e, 0x2b, 0xc2, 0x4b, 0x9f, 0xa7, 0xd2, 0x34, 0x0a, 0x00, 0x49, 0xb0, 0x4b,
	0x3f, 0x43, 0x6a, 0x39, 0x4f, 0x41, 0x30, 0xb2, 0xa1, 0x1c, 0x43, 0x4b, 0x8d, 0xa3, 0x07, 0x36,
	0x3f, 0xb7, 0x25, 0xdd, 0x75, 0x01, 0xa0, 0x1e, 0xb4, 0xe3, 0x90, 0x1e, 0x01, 0x05, 0x52, 0x81,
	0x7c, 0x8c, 0x8e, 0x61, 0x3b, 0xf2, 0xe6, 0x64, 0x45, 0xb3, 0x66, 0x81, 0x64, 0x80, 0x3e, 0x01,
	0xe0, 0x1f, 0x26, 0x73, 0x43, 0x6a, 0x25, 0x82, 0x1c, 0x99, 0x50, 0x40, 0x7e, 0x0f, 0xed, 0x5b,
	0x12, 0x61, 0x76, 0x58, 0xe8, 0x33, 0xe8, 0xf8, 0x71, 0x60, 0x3f, 0xd0, 0x43, 0x33, 0x53, 0xd3,
	0x9b, 0x94, 0xbd, 0x9f, 0x81, 0xf7, 0x14, 0x43, 0x17, 0xb0, 0x87, 0x69, 0x9e, 0x5e, 0x90, 0x50,
	0x92, 0x24, 0x20, 0x81, 0x6a, 0x04, 0x9e, 0x65, 0xb3, 0x4c, 0x18, 0x51, 0x44, 0xfe, 0x0b, 0x40,
	0xc8, 0xdd, 0x46, 0x3f, 0xc2, 0x49, 0x5a, 0x45, 0xa6, 0x1d, 0x10, 0x5a, 0xa2, 0x66, 0x90, 0x1c,
	0x7d, 0x7a, 0x8e, 0x97, 0x4a, 0x51, 0x74, 0x8a, 0x96, 0x30, 0x35, 0x4e, 0x4c, 0x4b, 0xe4, 0xed,
	0x13, 0xe3, 0xd8, 0xde, 0x80, 0xa3, 0x29, 0x3c, 0xcd, 0x94, 0x43, 0x1b, 0x2f, 0x0a, 0xe1, 0xe4,
	0xf0, 0x2f, 0x36, 0x08, 0x8f, 0x19, 0xaf, 0xd0, 0x3d, 0xb2, 0xd7, 0xe1, 0x72, 0xc2, 0x0e, 0x59,
	0x90, 0x52, 0xc2, 0xcd, 0x47, 0x13, 0x1e, 0x72, 0xe2, 0x7a, 0xc2, 0x15, 0x1c, 0xfd, 0x04, 0xa7,
	0x99, 0x72, 0xec, 0xcf, 0x02, 0xec, 0x14, 0xd2, 0x2d, 0x2e, 0xfd, 0xe9, 0x06, 0xe9, 0x69, 0xc2,
	0x2c, 0xb4, 0xb3, 0x4d, 0x57, 0x03, 0xc8, 0x00, 0xc9, 0x76, 0xcd, 0xe4, 0x2f, 0x5d, 0x77, 0x7a,
	0x9b, 0xab, 0x9f, 0x2a, 0xb6, 0xab, 0xdc, 0xe2, 0x4d, 0x06, 0x3f, 0xb5, 0xdd, 0x0d, 0x81, 0xaa,
	0x66, 0xcd, 0x8c, 0x9d, 0xba, 0x66, 0xdd, 0x83, 0x5c, 0xb3, 0x6a, 0xc2, 0x1d, 0x35, 0xc1, 0x35,
	0x79, 0x4f, 0xaa, 0xa7, 0xb9, 0xcb, 0x25, 0x4f, 0x98, 0xa4, 0xca, 0xe2, 0xeb, 0x65, 0xe0, 0xae,
	0xe3, 0x15, 0xc1, 0x5a, 0x8e, 0xed, 0x9a, 0xe0, 0xfa, 0x31, 0xb9, 0xeb, 0x38, 0xc2, 0xd0, 0x63,
	0x7f, 0x17, 0xd3, 0x8a, 0xed, 0x39, 0x59, 0xd3, 0x14, 0xb8, 0xa6, 0x5c, 0x6e, 0x19, 0x94, 0x3c,
	0xe0, 0xdc, 0xba, 0xfe, 0xa9, 0xb3, 0x39, 0x44, 0x6b, 0xac, 0xe7, 0x58, 0xa6, 0x85, 0xed, 0x79,
	0xec, 0xd3, 0xe2, 0x7d, 0x20, 0x4e, 0x5c, 0xaa, 0x5f, 0xe0, 0x4b, 0x9c, 0x29, 0x8e, 0xa5, 0x0c,
	0x38, 0x65, 0x9c, 0x32, 0xca, 0xca, 0xd6, 0xc6, 0x10, 0x9a, 0x80, 0x94, 0x77, 0x93, 0xba, 0xbf,
	0x7b, 0x5c, 0x57, 0x52, 0x8a, 0xb6, 0x5c, 0x77, 0xf8, 0x24, 0x0f, 0x55, 0x3d, 0x9e, 0xc1, 0x79,
	0xa1, 0x8a, 0x9d, 0xa5, 0x4b, 0xdb, 0x0d, 0x0e, 0xe7, 0xb9, 0xf2, 0x3e, 0x57, 0x7e, 0x5e, 0x32,
	0x65, 0x94, 0xd1, 0x55, 0xc6, 0x9e, 0x50, 0x72, 0xb1, 0xca, 0xd9, 0xea, 0xb1, 0x20, 0xfa, 0x1e,
	0x0e, 0x9d, 0x55, 0x68, 0xd2, 0x4d, 0xd9, 0x85, 0x7a, 0x87, 0xab, 0x9f, 0x2b, 0xe5, 0x7b, 0x6e,
	0x38, 0x1a, 0x6b, 0x8c, 0x54, 0xa8, 0x76, 0xe9, 0xc4, 0x32, 0x84, 0xbe, 0x82, 0x16, 0xef, 0x89,
	0x07, 0x74, 0xfa, 0x01, 0xdd, 0xf6, 0x86, 0xbb, 0x80, 0xb5, 0x48, 0x83, 0xb3, 0xf2, 0x8b, 0xa8,
	0x5b, 0x5c, 0x44, 0xb4, 0x61, 0xb6, 0x58, 0x5f, 0x93, 0x44, 0x9e, 0x40, 0xb7, 0xa4, 0xc0, 0x5a,
	0xb9, 0xc1, 0x83, 0xa8, 0x0f, 0xed, 0x65, 0xda, 0x61, 0xa5, 0x43, 0x4e, 0x3c, 0x2a, 0x11, 0xb3,
	0xe6, 0x6b, 0xe4, 0xa4, 0x81, 0x00, 0xbb, 0xe9, 0xce, 0xe4, 0x05, 0x9c, 0x3e, 0x52, 0x3d, 0xe8,
	0x19, 0x08, 0xbc, 0x0a, 0xf9, 0x16, 0x92, 0x7b, 0xa2, 0xcd, 0x00, 0x96, 0x72, 0xf5, 0x12, 0x69,
	0xd4, 0x2f, 0x91, 0x13, 0xd8, 0xf1, 0x03, 0xf2, 0xce, 0xfd, 0x25, 0x6d, 0xce, 0xe9, 0x48, 0xfe,
	0x16, 0xce, 0x1e, 0x3d, 0x96, 0x8f, 0xdf, 0x4b, 0x5f, 0xfe, 0xdd, 0x80, 0x4e, 0xc5, 0x35, 0xb4,
	0x07, 0xbb, 0xd3, 0xd1, 0xf5, 0xe8, 0xee, 0x87, 0x91, 0xf8, 0x84, 0x9a, 0x77, 0xa0, 0xdd, 0x4c,
	0xc7, 0x13, 0xdd, 0x30, 0x35, 0x43, 0x57, 0x27, 0xba, 0xb8, 0x85, 0x0e, 0xa1, 0x93, 0x61, 0x63,
	0x4d, 0xbd, 0xd1, 0xc5, 0x46, 0x99, 0x36, 0xd4, 0x6f, 0x74, 0x4a, 0x6b, 0xa2, 0x23, 0xe8, 0x66,
	0xd8, 0xf4, 0xfe, 0x8d, 0xa1, 0x0e, 0x75, 0xb1, 0x45, 0xaf, 0x3a, 0x51, 0xbb, 0x32, 0x6f, 0xd5,
	0xb2, 0xe2, 0x76, 0x15, 0x4d, 0x05, 0x76, 0xb8, 0xc0, 0x95, 0xa9, 0xbe, 0xd1, 0x47, 0x93, 0x8c,
	0xba, 0x5b, 0x01, 0x53, 0x66, 0x9b, 0xfa, 0x82, 0x86, 0xea, 0x44, 0x35, 0x07, 0x53, 0xed, 0x5a,
	0xcf, 0x71, 0x81, 0x91, 0x19, 0x3e, 0x50, 0xc7, 0xba, 0x39, 0x50, 0xb5, 0xeb, 0xe9, 0xbd, 0x08,
	0x6c, 0xb1, 0x91, 0x7a, 0xab, 0x8f, 0xef, 0x55, 0x4d, 0xcf, 0x74, 0xf7, 0xe8, 0xfd, 0x7e, 0x5c,
	0xa0, 0xea, 0xf0, 0xf6, 0x6a, 0x64, 0x4e, 0xd4, 0xf1, 0xb5, 0xb8, 0x8f, 0x3a, 0x20, 0x68, 0x6f,
	0x75, 0xed, 0xda, 0xa4, 0x95, 0x29, 0x76, 0x5e, 0xff, 0xb1, 0x05, 0x90, 0x1b, 0x16, 0xa2, 0x0f,
	0xd0, 0xce, 0x5e, 0x16, 0xa8, 0xb7, 0xf1, 0xb9, 0xc1, 0x4f, 0xa1, 0xf7, 0xec, 0x23, 0x4f, 0x11,
	0xf9, 0x9b, 0xdf, 0xff, 0xfd, 0xef, 0xcf, 0xc6, 0x4b, 0xa4, 0xd0, 0x57, 0xa3, 0x1f, 0xda, 0x9e,
	0x93, 0x3c, 0x1f, 0xf3, 0x19, 0xc9, 0xa3, 0x2b, 0x1f, 0x86, 0xfd, 0x0f, 0xac, 0x8a, 0x7f, 0xb3,
	0x76, 0xf8, 0xeb, 0xeb, 0xeb, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xe7, 0x1b, 0xeb, 0x90,
	0x0a, 0x00, 0x00,
}