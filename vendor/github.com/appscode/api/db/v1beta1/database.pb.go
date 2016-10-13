// Code generated by protoc-gen-go.
// source: database.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	database.proto
	snapshot.proto

It has these top-level messages:
	Database
	DatabaseListRequest
	DatabaseListResponse
	DatabaseCreateRequest
	DatabaseScaleRequest
	DatabaseUpdateRequest
	DatabaseDescribeRequest
	SnapshotSummary
	DatabaseDescribeResponse
	DatabaseDeleteRequest
	DatabaseRecoverRequest
	DatabaseCheckRequest
	DatabaseInfo
	Snapshot
	SnapshotListRequest
	SnapshotListResponse
	BackupScheduleRequest
	BackupUnscheduleRequest
	SnapshotRestoreRequest
	SnapshotCheckRequest
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Next Id: 19
type Database struct {
	Phid             string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Cluster          string `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type             string `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	Sku              string `protobuf:"bytes,6,opt,name=sku" json:"sku,omitempty"`
	Version          string `protobuf:"bytes,8,opt,name=version" json:"version,omitempty"`
	AuthSecret       string `protobuf:"bytes,9,opt,name=auth_secret,json=authSecret" json:"auth_secret,omitempty"`
	ScheduleCronExpr string `protobuf:"bytes,10,opt,name=schedule_cron_expr,json=scheduleCronExpr" json:"schedule_cron_expr,omitempty"`
	PvSizeGb         int32  `protobuf:"varint,11,opt,name=pv_size_gb,json=pvSizeGb" json:"pv_size_gb,omitempty"`
	LastBackupAt     int64  `protobuf:"varint,12,opt,name=last_backup_at,json=lastBackupAt" json:"last_backup_at,omitempty"`
	Status           string `protobuf:"bytes,13,opt,name=status" json:"status,omitempty"`
	CreatedAt        int64  `protobuf:"varint,14,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	DeletedAt        int64  `protobuf:"varint,15,opt,name=deleted_at,json=deletedAt" json:"deleted_at,omitempty"`
	DoNotDelete      bool   `protobuf:"varint,17,opt,name=do_not_delete,json=doNotDelete" json:"do_not_delete,omitempty"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Next Id: 4
type DatabaseListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// List of status to get the agent filterd on the status
	// values in
	//   PENDING
	//   FAILED
	//   READY
	//   DELETING
	//   DELETED
	//   DESTROYED
	Status []string `protobuf:"bytes,3,rep,name=status" json:"status,omitempty"`
}

func (m *DatabaseListRequest) Reset()                    { *m = DatabaseListRequest{} }
func (m *DatabaseListRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseListRequest) ProtoMessage()               {}
func (*DatabaseListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Next Id: 3
type DatabaseListResponse struct {
	Status    *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Databases []*Database    `protobuf:"bytes,2,rep,name=databases" json:"databases,omitempty"`
}

func (m *DatabaseListResponse) Reset()                    { *m = DatabaseListResponse{} }
func (m *DatabaseListResponse) String() string            { return proto.CompactTextString(m) }
func (*DatabaseListResponse) ProtoMessage()               {}
func (*DatabaseListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DatabaseListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DatabaseListResponse) GetDatabases() []*Database {
	if m != nil {
		return m.Databases
	}
	return nil
}

// Next Id: 19
type DatabaseCreateRequest struct {
	Cluster          string                               `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name             string                               `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type             string                               `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Sku              string                               `protobuf:"bytes,4,opt,name=sku" json:"sku,omitempty"`
	Version          string                               `protobuf:"bytes,5,opt,name=version" json:"version,omitempty"`
	PvSizeGb         int32                                `protobuf:"varint,7,opt,name=pv_size_gb,json=pvSizeGb" json:"pv_size_gb,omitempty"`
	AuthSecretName   string                               `protobuf:"bytes,8,opt,name=auth_secret_name,json=authSecretName" json:"auth_secret_name,omitempty"`
	Size             int32                                `protobuf:"varint,9,opt,name=size" json:"size,omitempty"`
	SnapshotPhid     string                               `protobuf:"bytes,14,opt,name=snapshot_phid,json=snapshotPhid" json:"snapshot_phid,omitempty"`
	Hostname         string                               `protobuf:"bytes,15,opt,name=hostname" json:"hostname,omitempty"`
	StorageClass     string                               `protobuf:"bytes,16,opt,name=storage_class,json=storageClass" json:"storage_class,omitempty"`
	IgnoreValidation bool                                 `protobuf:"varint,17,opt,name=ignore_validation,json=ignoreValidation" json:"ignore_validation,omitempty"`
	InitialScript    *DatabaseCreateRequest_InitialScript `protobuf:"bytes,18,opt,name=initial_script,json=initialScript" json:"initial_script,omitempty"`
}

func (m *DatabaseCreateRequest) Reset()                    { *m = DatabaseCreateRequest{} }
func (m *DatabaseCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseCreateRequest) ProtoMessage()               {}
func (*DatabaseCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DatabaseCreateRequest) GetInitialScript() *DatabaseCreateRequest_InitialScript {
	if m != nil {
		return m.InitialScript
	}
	return nil
}

type DatabaseCreateRequest_InitialScript struct {
	ScriptPath string `protobuf:"bytes,1,opt,name=script_path,json=scriptPath" json:"script_path,omitempty"`
	// Types that are valid to be assigned to ScriptSource:
	//	*DatabaseCreateRequest_InitialScript_GitRepo_
	//	*DatabaseCreateRequest_InitialScript_Secret_
	//	*DatabaseCreateRequest_InitialScript_ConfigMap_
	ScriptSource isDatabaseCreateRequest_InitialScript_ScriptSource `protobuf_oneof:"ScriptSource"`
}

func (m *DatabaseCreateRequest_InitialScript) Reset()         { *m = DatabaseCreateRequest_InitialScript{} }
func (m *DatabaseCreateRequest_InitialScript) String() string { return proto.CompactTextString(m) }
func (*DatabaseCreateRequest_InitialScript) ProtoMessage()    {}
func (*DatabaseCreateRequest_InitialScript) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type isDatabaseCreateRequest_InitialScript_ScriptSource interface {
	isDatabaseCreateRequest_InitialScript_ScriptSource()
}

type DatabaseCreateRequest_InitialScript_GitRepo_ struct {
	GitRepo *DatabaseCreateRequest_InitialScript_GitRepo `protobuf:"bytes,2,opt,name=git_repo,json=gitRepo,oneof"`
}
type DatabaseCreateRequest_InitialScript_Secret_ struct {
	Secret *DatabaseCreateRequest_InitialScript_Secret `protobuf:"bytes,3,opt,name=secret,oneof"`
}
type DatabaseCreateRequest_InitialScript_ConfigMap_ struct {
	ConfigMap *DatabaseCreateRequest_InitialScript_ConfigMap `protobuf:"bytes,4,opt,name=config_map,json=configMap,oneof"`
}

func (*DatabaseCreateRequest_InitialScript_GitRepo_) isDatabaseCreateRequest_InitialScript_ScriptSource() {
}
func (*DatabaseCreateRequest_InitialScript_Secret_) isDatabaseCreateRequest_InitialScript_ScriptSource() {
}
func (*DatabaseCreateRequest_InitialScript_ConfigMap_) isDatabaseCreateRequest_InitialScript_ScriptSource() {
}

func (m *DatabaseCreateRequest_InitialScript) GetScriptSource() isDatabaseCreateRequest_InitialScript_ScriptSource {
	if m != nil {
		return m.ScriptSource
	}
	return nil
}

func (m *DatabaseCreateRequest_InitialScript) GetGitRepo() *DatabaseCreateRequest_InitialScript_GitRepo {
	if x, ok := m.GetScriptSource().(*DatabaseCreateRequest_InitialScript_GitRepo_); ok {
		return x.GitRepo
	}
	return nil
}

func (m *DatabaseCreateRequest_InitialScript) GetSecret() *DatabaseCreateRequest_InitialScript_Secret {
	if x, ok := m.GetScriptSource().(*DatabaseCreateRequest_InitialScript_Secret_); ok {
		return x.Secret
	}
	return nil
}

func (m *DatabaseCreateRequest_InitialScript) GetConfigMap() *DatabaseCreateRequest_InitialScript_ConfigMap {
	if x, ok := m.GetScriptSource().(*DatabaseCreateRequest_InitialScript_ConfigMap_); ok {
		return x.ConfigMap
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DatabaseCreateRequest_InitialScript) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DatabaseCreateRequest_InitialScript_OneofMarshaler, _DatabaseCreateRequest_InitialScript_OneofUnmarshaler, _DatabaseCreateRequest_InitialScript_OneofSizer, []interface{}{
		(*DatabaseCreateRequest_InitialScript_GitRepo_)(nil),
		(*DatabaseCreateRequest_InitialScript_Secret_)(nil),
		(*DatabaseCreateRequest_InitialScript_ConfigMap_)(nil),
	}
}

func _DatabaseCreateRequest_InitialScript_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DatabaseCreateRequest_InitialScript)
	// ScriptSource
	switch x := m.ScriptSource.(type) {
	case *DatabaseCreateRequest_InitialScript_GitRepo_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GitRepo); err != nil {
			return err
		}
	case *DatabaseCreateRequest_InitialScript_Secret_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Secret); err != nil {
			return err
		}
	case *DatabaseCreateRequest_InitialScript_ConfigMap_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConfigMap); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DatabaseCreateRequest_InitialScript.ScriptSource has unexpected type %T", x)
	}
	return nil
}

func _DatabaseCreateRequest_InitialScript_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DatabaseCreateRequest_InitialScript)
	switch tag {
	case 2: // ScriptSource.git_repo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DatabaseCreateRequest_InitialScript_GitRepo)
		err := b.DecodeMessage(msg)
		m.ScriptSource = &DatabaseCreateRequest_InitialScript_GitRepo_{msg}
		return true, err
	case 3: // ScriptSource.secret
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DatabaseCreateRequest_InitialScript_Secret)
		err := b.DecodeMessage(msg)
		m.ScriptSource = &DatabaseCreateRequest_InitialScript_Secret_{msg}
		return true, err
	case 4: // ScriptSource.config_map
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DatabaseCreateRequest_InitialScript_ConfigMap)
		err := b.DecodeMessage(msg)
		m.ScriptSource = &DatabaseCreateRequest_InitialScript_ConfigMap_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DatabaseCreateRequest_InitialScript_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DatabaseCreateRequest_InitialScript)
	// ScriptSource
	switch x := m.ScriptSource.(type) {
	case *DatabaseCreateRequest_InitialScript_GitRepo_:
		s := proto.Size(x.GitRepo)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DatabaseCreateRequest_InitialScript_Secret_:
		s := proto.Size(x.Secret)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DatabaseCreateRequest_InitialScript_ConfigMap_:
		s := proto.Size(x.ConfigMap)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DatabaseCreateRequest_InitialScript_GitRepo struct {
	Repository string `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Revision   string `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
	Directory  string `protobuf:"bytes,3,opt,name=directory" json:"directory,omitempty"`
}

func (m *DatabaseCreateRequest_InitialScript_GitRepo) Reset() {
	*m = DatabaseCreateRequest_InitialScript_GitRepo{}
}
func (m *DatabaseCreateRequest_InitialScript_GitRepo) String() string {
	return proto.CompactTextString(m)
}
func (*DatabaseCreateRequest_InitialScript_GitRepo) ProtoMessage() {}
func (*DatabaseCreateRequest_InitialScript_GitRepo) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0, 0}
}

type DatabaseCreateRequest_InitialScript_Secret struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DatabaseCreateRequest_InitialScript_Secret) Reset() {
	*m = DatabaseCreateRequest_InitialScript_Secret{}
}
func (m *DatabaseCreateRequest_InitialScript_Secret) String() string {
	return proto.CompactTextString(m)
}
func (*DatabaseCreateRequest_InitialScript_Secret) ProtoMessage() {}
func (*DatabaseCreateRequest_InitialScript_Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0, 1}
}

type DatabaseCreateRequest_InitialScript_ConfigMap struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DatabaseCreateRequest_InitialScript_ConfigMap) Reset() {
	*m = DatabaseCreateRequest_InitialScript_ConfigMap{}
}
func (m *DatabaseCreateRequest_InitialScript_ConfigMap) String() string {
	return proto.CompactTextString(m)
}
func (*DatabaseCreateRequest_InitialScript_ConfigMap) ProtoMessage() {}
func (*DatabaseCreateRequest_InitialScript_ConfigMap) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0, 2}
}

// Next Id: 4
type DatabaseScaleRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	Size    int32  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *DatabaseScaleRequest) Reset()                    { *m = DatabaseScaleRequest{} }
func (m *DatabaseScaleRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseScaleRequest) ProtoMessage()               {}
func (*DatabaseScaleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Next Id: 4
type DatabaseUpdateRequest struct {
	Cluster     string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid         string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	DoNotDelete bool   `protobuf:"varint,3,opt,name=do_not_delete,json=doNotDelete" json:"do_not_delete,omitempty"`
}

func (m *DatabaseUpdateRequest) Reset()                    { *m = DatabaseUpdateRequest{} }
func (m *DatabaseUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseUpdateRequest) ProtoMessage()               {}
func (*DatabaseUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Next Id: 3
type DatabaseDescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *DatabaseDescribeRequest) Reset()                    { *m = DatabaseDescribeRequest{} }
func (m *DatabaseDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseDescribeRequest) ProtoMessage()               {}
func (*DatabaseDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Next Id: 7
type SnapshotSummary struct {
	BackupAttempt  int32 `protobuf:"varint,3,opt,name=backup_attempt,json=backupAttempt" json:"backup_attempt,omitempty"`
	BackupSuccess  int32 `protobuf:"varint,4,opt,name=backup_success,json=backupSuccess" json:"backup_success,omitempty"`
	RestoreAttempt int32 `protobuf:"varint,5,opt,name=restore_attempt,json=restoreAttempt" json:"restore_attempt,omitempty"`
	RestoreSuccess int32 `protobuf:"varint,6,opt,name=restore_success,json=restoreSuccess" json:"restore_success,omitempty"`
}

func (m *SnapshotSummary) Reset()                    { *m = SnapshotSummary{} }
func (m *SnapshotSummary) String() string            { return proto.CompactTextString(m) }
func (*SnapshotSummary) ProtoMessage()               {}
func (*SnapshotSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// Next Id: 17
type DatabaseDescribeResponse struct {
	Status          *dtypes.Status   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	SnapshotSummary *SnapshotSummary `protobuf:"bytes,9,opt,name=snapshot_summary,json=snapshotSummary" json:"snapshot_summary,omitempty"`
	Database        *Database        `protobuf:"bytes,16,opt,name=database" json:"database,omitempty"`
}

func (m *DatabaseDescribeResponse) Reset()                    { *m = DatabaseDescribeResponse{} }
func (m *DatabaseDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DatabaseDescribeResponse) ProtoMessage()               {}
func (*DatabaseDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DatabaseDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DatabaseDescribeResponse) GetSnapshotSummary() *SnapshotSummary {
	if m != nil {
		return m.SnapshotSummary
	}
	return nil
}

func (m *DatabaseDescribeResponse) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

// Next Id: 4
type DatabaseDeleteRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	Destroy bool   `protobuf:"varint,3,opt,name=destroy" json:"destroy,omitempty"`
}

func (m *DatabaseDeleteRequest) Reset()                    { *m = DatabaseDeleteRequest{} }
func (m *DatabaseDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseDeleteRequest) ProtoMessage()               {}
func (*DatabaseDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Next Id: 3
type DatabaseRecoverRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid     string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *DatabaseRecoverRequest) Reset()                    { *m = DatabaseRecoverRequest{} }
func (m *DatabaseRecoverRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseRecoverRequest) ProtoMessage()               {}
func (*DatabaseRecoverRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// Next Id: 4
type DatabaseCheckRequest struct {
	Cluster      string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Uid          string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	PurchasePhid string `protobuf:"bytes,3,opt,name=purchase_phid,json=purchasePhid" json:"purchase_phid,omitempty"`
}

func (m *DatabaseCheckRequest) Reset()                    { *m = DatabaseCheckRequest{} }
func (m *DatabaseCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*DatabaseCheckRequest) ProtoMessage()               {}
func (*DatabaseCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*Database)(nil), "db.v1beta1.Database")
	proto.RegisterType((*DatabaseListRequest)(nil), "db.v1beta1.DatabaseListRequest")
	proto.RegisterType((*DatabaseListResponse)(nil), "db.v1beta1.DatabaseListResponse")
	proto.RegisterType((*DatabaseCreateRequest)(nil), "db.v1beta1.DatabaseCreateRequest")
	proto.RegisterType((*DatabaseCreateRequest_InitialScript)(nil), "db.v1beta1.DatabaseCreateRequest.InitialScript")
	proto.RegisterType((*DatabaseCreateRequest_InitialScript_GitRepo)(nil), "db.v1beta1.DatabaseCreateRequest.InitialScript.GitRepo")
	proto.RegisterType((*DatabaseCreateRequest_InitialScript_Secret)(nil), "db.v1beta1.DatabaseCreateRequest.InitialScript.Secret")
	proto.RegisterType((*DatabaseCreateRequest_InitialScript_ConfigMap)(nil), "db.v1beta1.DatabaseCreateRequest.InitialScript.ConfigMap")
	proto.RegisterType((*DatabaseScaleRequest)(nil), "db.v1beta1.DatabaseScaleRequest")
	proto.RegisterType((*DatabaseUpdateRequest)(nil), "db.v1beta1.DatabaseUpdateRequest")
	proto.RegisterType((*DatabaseDescribeRequest)(nil), "db.v1beta1.DatabaseDescribeRequest")
	proto.RegisterType((*SnapshotSummary)(nil), "db.v1beta1.SnapshotSummary")
	proto.RegisterType((*DatabaseDescribeResponse)(nil), "db.v1beta1.DatabaseDescribeResponse")
	proto.RegisterType((*DatabaseDeleteRequest)(nil), "db.v1beta1.DatabaseDeleteRequest")
	proto.RegisterType((*DatabaseRecoverRequest)(nil), "db.v1beta1.DatabaseRecoverRequest")
	proto.RegisterType((*DatabaseCheckRequest)(nil), "db.v1beta1.DatabaseCheckRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Databases service

type DatabasesClient interface {
	List(ctx context.Context, in *DatabaseListRequest, opts ...grpc.CallOption) (*DatabaseListResponse, error)
	Create(ctx context.Context, in *DatabaseCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Scale(ctx context.Context, in *DatabaseScaleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *DatabaseUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Describe(ctx context.Context, in *DatabaseDescribeRequest, opts ...grpc.CallOption) (*DatabaseDescribeResponse, error)
	Delete(ctx context.Context, in *DatabaseDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Recover(ctx context.Context, in *DatabaseRecoverRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type databasesClient struct {
	cc *grpc.ClientConn
}

func NewDatabasesClient(cc *grpc.ClientConn) DatabasesClient {
	return &databasesClient{cc}
}

func (c *databasesClient) List(ctx context.Context, in *DatabaseListRequest, opts ...grpc.CallOption) (*DatabaseListResponse, error) {
	out := new(DatabaseListResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Create(ctx context.Context, in *DatabaseCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Scale(ctx context.Context, in *DatabaseScaleRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Scale", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Update(ctx context.Context, in *DatabaseUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Describe(ctx context.Context, in *DatabaseDescribeRequest, opts ...grpc.CallOption) (*DatabaseDescribeResponse, error) {
	out := new(DatabaseDescribeResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Delete(ctx context.Context, in *DatabaseDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databasesClient) Recover(ctx context.Context, in *DatabaseRecoverRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.v1beta1.Databases/Recover", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Databases service

type DatabasesServer interface {
	List(context.Context, *DatabaseListRequest) (*DatabaseListResponse, error)
	Create(context.Context, *DatabaseCreateRequest) (*dtypes.VoidResponse, error)
	Scale(context.Context, *DatabaseScaleRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *DatabaseUpdateRequest) (*dtypes.VoidResponse, error)
	Describe(context.Context, *DatabaseDescribeRequest) (*DatabaseDescribeResponse, error)
	Delete(context.Context, *DatabaseDeleteRequest) (*dtypes.VoidResponse, error)
	Recover(context.Context, *DatabaseRecoverRequest) (*dtypes.VoidResponse, error)
}

func RegisterDatabasesServer(s *grpc.Server, srv DatabasesServer) {
	s.RegisterService(&_Databases_serviceDesc, srv)
}

func _Databases_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).List(ctx, req.(*DatabaseListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Create(ctx, req.(*DatabaseCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Scale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Scale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Scale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Scale(ctx, req.(*DatabaseScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Update(ctx, req.(*DatabaseUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Describe(ctx, req.(*DatabaseDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Delete(ctx, req.(*DatabaseDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Databases_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseRecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabasesServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.v1beta1.Databases/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabasesServer).Recover(ctx, req.(*DatabaseRecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Databases_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.v1beta1.Databases",
	HandlerType: (*DatabasesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Databases_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Databases_Create_Handler,
		},
		{
			MethodName: "Scale",
			Handler:    _Databases_Scale_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Databases_Update_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Databases_Describe_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Databases_Delete_Handler,
		},
		{
			MethodName: "Recover",
			Handler:    _Databases_Recover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("database.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x57, 0x4d, 0x6f, 0x1b, 0x45,
	0x18, 0xee, 0x66, 0x1b, 0x7f, 0xbc, 0x8e, 0x1d, 0x77, 0x28, 0x65, 0xb5, 0x84, 0xd6, 0x6c, 0x0b,
	0x44, 0x05, 0xd9, 0xd4, 0x48, 0x20, 0x2a, 0x21, 0xd4, 0x24, 0xa5, 0x41, 0x40, 0xd5, 0xae, 0x4b,
	0x84, 0x8a, 0x60, 0x35, 0xde, 0x1d, 0xec, 0x55, 0x9c, 0x9d, 0x65, 0x67, 0xd6, 0x6a, 0x5a, 0xf5,
	0xd2, 0x7f, 0x80, 0x40, 0x20, 0xce, 0xbd, 0x70, 0xe5, 0x80, 0xb8, 0xf0, 0x2f, 0xf8, 0x03, 0x1c,
	0x38, 0xf0, 0x33, 0xd0, 0x7c, 0xec, 0x87, 0x63, 0x1b, 0x9a, 0x34, 0x97, 0x68, 0xe6, 0x99, 0x77,
	0xdf, 0x79, 0xe6, 0x9d, 0xe7, 0x7d, 0xc6, 0x81, 0x56, 0x80, 0x39, 0x1e, 0x62, 0x46, 0xba, 0x71,
	0x42, 0x39, 0x45, 0x10, 0x0c, 0xbb, 0xd3, 0x6b, 0x43, 0xc2, 0xf1, 0x35, 0x7b, 0x63, 0x44, 0xe9,
	0x68, 0x42, 0x7a, 0x38, 0x0e, 0x7b, 0x38, 0x8a, 0x28, 0xc7, 0x3c, 0xa4, 0x11, 0x53, 0x91, 0xf6,
	0x45, 0x1c, 0xc7, 0xcc, 0xa7, 0xc1, 0xb2, 0xf5, 0x0b, 0x02, 0x0e, 0xf8, 0x61, 0x4c, 0x58, 0x4f,
	0xfe, 0x55, 0xb8, 0xf3, 0xb3, 0x09, 0xb5, 0x1d, 0xbd, 0x29, 0x42, 0x70, 0x36, 0x1e, 0x87, 0x81,
	0x65, 0x74, 0x8c, 0xcd, 0xba, 0x2b, 0xc7, 0xc8, 0x82, 0xaa, 0x3f, 0x49, 0x19, 0x27, 0x89, 0xb5,
	0x22, 0xe1, 0x6c, 0x2a, 0xa2, 0x23, 0x7c, 0x40, 0x2c, 0x53, 0x45, 0x8b, 0xb1, 0xc0, 0x44, 0x76,
	0x6b, 0x55, 0x61, 0x62, 0x8c, 0xda, 0x60, 0xb2, 0xfd, 0xd4, 0xaa, 0x48, 0x48, 0x0c, 0x45, 0xce,
	0x29, 0x49, 0x58, 0x48, 0x23, 0xab, 0xa6, 0x72, 0xea, 0x29, 0xba, 0x04, 0x0d, 0x9c, 0xf2, 0xb1,
	0xc7, 0x88, 0x9f, 0x10, 0x6e, 0xd5, 0xe5, 0x2a, 0x08, 0x68, 0x20, 0x11, 0xf4, 0x16, 0x20, 0xe6,
	0x8f, 0x49, 0x90, 0x4e, 0x88, 0xe7, 0x27, 0x34, 0xf2, 0xc8, 0x83, 0x38, 0xb1, 0x40, 0xc6, 0xb5,
	0xb3, 0x95, 0xed, 0x84, 0x46, 0x37, 0x1f, 0xc4, 0x09, 0xda, 0x00, 0x88, 0xa7, 0x1e, 0x0b, 0x1f,
	0x12, 0x6f, 0x34, 0xb4, 0x1a, 0x1d, 0x63, 0x73, 0xd5, 0xad, 0xc5, 0xd3, 0x41, 0xf8, 0x90, 0xdc,
	0x1a, 0xa2, 0x2b, 0xd0, 0x9a, 0x60, 0xc6, 0xbd, 0x21, 0xf6, 0xf7, 0xd3, 0xd8, 0xc3, 0xdc, 0x5a,
	0xeb, 0x18, 0x9b, 0xa6, 0xbb, 0x26, 0xd0, 0x2d, 0x09, 0xde, 0xe0, 0xe8, 0x02, 0x54, 0x18, 0xc7,
	0x3c, 0x65, 0x56, 0x53, 0xee, 0xa2, 0x67, 0xe8, 0x15, 0x00, 0x3f, 0x21, 0x98, 0x93, 0x40, 0x7c,
	0xd9, 0x92, 0x5f, 0xd6, 0x35, 0x72, 0x83, 0x8b, 0xe5, 0x80, 0x4c, 0x88, 0x5e, 0x5e, 0x57, 0xcb,
	0x1a, 0xb9, 0xc1, 0x91, 0x03, 0xcd, 0x80, 0x7a, 0x11, 0xe5, 0x9e, 0xc2, 0xac, 0x73, 0x1d, 0x63,
	0xb3, 0xe6, 0x36, 0x02, 0x7a, 0x9b, 0xf2, 0x1d, 0x09, 0x39, 0x5f, 0xc2, 0x0b, 0xd9, 0xd5, 0x7c,
	0x1a, 0x32, 0xee, 0x92, 0x6f, 0x53, 0xc2, 0x78, 0xf9, 0x46, 0x8c, 0xb9, 0x1b, 0x91, 0xd5, 0x5f,
	0x29, 0x55, 0xbf, 0xa0, 0x6f, 0x76, 0xcc, 0x82, 0xbe, 0x93, 0xc0, 0xf9, 0xd9, 0xe4, 0x2c, 0xa6,
	0x11, 0x23, 0xe8, 0xf5, 0x3c, 0x5e, 0x24, 0x6f, 0xf4, 0x5b, 0x5d, 0xa5, 0x9a, 0xee, 0x40, 0xa2,
	0xf9, 0xf1, 0xfb, 0x50, 0xcf, 0xc4, 0xca, 0xac, 0x95, 0x8e, 0xb9, 0xd9, 0xe8, 0x9f, 0xef, 0x16,
	0x72, 0xed, 0x66, 0xc9, 0xdd, 0x22, 0xcc, 0xf9, 0xbd, 0x0a, 0x2f, 0x66, 0xf8, 0xb6, 0xac, 0xd4,
	0x33, 0x9d, 0x49, 0xaa, 0x6c, 0x65, 0x81, 0xca, 0xcc, 0x79, 0x95, 0x9d, 0x5d, 0xa8, 0xb2, 0xd5,
	0x59, 0x95, 0xcd, 0xca, 0xa2, 0x7a, 0x44, 0x16, 0x9b, 0xd0, 0x2e, 0x69, 0xd0, 0x93, 0xbb, 0x2b,
	0x99, 0xb6, 0x0a, 0x21, 0xde, 0xd6, 0x3c, 0x44, 0x12, 0x29, 0xd3, 0x55, 0x57, 0x8e, 0xd1, 0x65,
	0x68, 0xb2, 0x08, 0xc7, 0x6c, 0x4c, 0xb9, 0x27, 0x9b, 0xa9, 0x25, 0x3f, 0x5d, 0xcb, 0xc0, 0x3b,
	0xa2, 0xa9, 0x6c, 0xa8, 0x8d, 0x29, 0xe3, 0x32, 0xf5, 0xba, 0x5c, 0xcf, 0xe7, 0x32, 0x01, 0xa7,
	0x09, 0x1e, 0x11, 0xcf, 0x9f, 0x60, 0xc6, 0xac, 0xb6, 0x4e, 0xa0, 0xc0, 0x6d, 0x81, 0xa1, 0x37,
	0xe1, 0x5c, 0x38, 0x8a, 0x68, 0x42, 0xbc, 0x29, 0x9e, 0x84, 0x81, 0x6c, 0x75, 0x2d, 0xa1, 0xb6,
	0x5a, 0xd8, 0xcb, 0x71, 0xb4, 0x07, 0xad, 0x30, 0x0a, 0x79, 0x88, 0x27, 0x1e, 0xf3, 0x93, 0x30,
	0xe6, 0x16, 0x92, 0x57, 0xdb, 0x5b, 0x74, 0x5f, 0x33, 0xf7, 0xd2, 0xfd, 0x58, 0x7d, 0x37, 0x90,
	0x9f, 0xb9, 0xcd, 0xb0, 0x3c, 0xb5, 0xff, 0x31, 0xa1, 0x39, 0x13, 0x20, 0xda, 0x57, 0xed, 0xe0,
	0xc5, 0x98, 0x8f, 0xf5, 0x55, 0x82, 0x82, 0xee, 0x60, 0x3e, 0x46, 0xf7, 0xa0, 0x36, 0x0a, 0xb9,
	0x97, 0x90, 0x98, 0xca, 0x1b, 0x6d, 0xf4, 0xdf, 0x3b, 0x26, 0x89, 0xee, 0xad, 0x90, 0xbb, 0x24,
	0xa6, 0xbb, 0x67, 0xdc, 0xea, 0x48, 0x0d, 0xd1, 0x1d, 0xa8, 0x68, 0xc3, 0x30, 0x65, 0xce, 0x77,
	0x8f, 0x9b, 0x53, 0xdd, 0xe9, 0xee, 0x19, 0x57, 0xe7, 0x41, 0xf7, 0x01, 0x7c, 0x1a, 0x7d, 0x13,
	0x8e, 0xbc, 0x03, 0x1c, 0x4b, 0x51, 0x35, 0xfa, 0xef, 0x1f, 0x37, 0xeb, 0xb6, 0xcc, 0xf0, 0x19,
	0x8e, 0x77, 0xcf, 0xb8, 0x75, 0x3f, 0x9b, 0xd8, 0x3e, 0x54, 0xf5, 0x19, 0xd0, 0x45, 0x00, 0x51,
	0x0a, 0x16, 0x72, 0x9a, 0x1c, 0x66, 0xe5, 0x2a, 0x10, 0xa1, 0x93, 0x84, 0x4c, 0x43, 0xa9, 0x61,
	0xd5, 0x00, 0xf9, 0x1c, 0x6d, 0x40, 0x3d, 0x08, 0x13, 0xe2, 0xcb, 0x4f, 0x55, 0x27, 0x14, 0x80,
	0xbd, 0x01, 0x15, 0xed, 0x98, 0x59, 0x03, 0x19, 0x45, 0x03, 0xd9, 0x97, 0xa0, 0x9e, 0x93, 0x5b,
	0x14, 0xb0, 0xd5, 0x82, 0x35, 0x75, 0x88, 0x01, 0x4d, 0x13, 0x9f, 0x38, 0x7b, 0x85, 0x5b, 0x0c,
	0x7c, 0x3c, 0x79, 0x86, 0xbe, 0x6d, 0x83, 0x99, 0x86, 0x81, 0x66, 0x2d, 0x86, 0x79, 0xb7, 0x98,
	0x45, 0xb7, 0x38, 0xa3, 0xc2, 0x10, 0x3e, 0x8f, 0x83, 0x67, 0x32, 0x84, 0xf9, 0xc4, 0x73, 0x5e,
	0x6a, 0xce, 0x7b, 0xe9, 0x4d, 0x78, 0x29, 0xdb, 0x68, 0x87, 0x08, 0x41, 0x0e, 0x4f, 0xb2, 0x95,
	0xf3, 0xab, 0x01, 0xeb, 0x03, 0xdd, 0xc9, 0x83, 0xf4, 0xe0, 0x00, 0x27, 0x87, 0xe8, 0x35, 0x68,
	0xe5, 0x2f, 0x08, 0x27, 0x07, 0x31, 0xd7, 0x27, 0x6c, 0x0e, 0xf5, 0x13, 0x22, 0xc1, 0x52, 0x18,
	0x4b, 0x7d, 0x9f, 0x30, 0x26, 0x65, 0x95, 0x87, 0x0d, 0x14, 0x88, 0xde, 0x80, 0xf5, 0x84, 0x88,
	0x5e, 0x27, 0x79, 0xba, 0x55, 0x19, 0xd7, 0xd2, 0x70, 0x96, 0xaf, 0x14, 0x98, 0x25, 0xac, 0xcc,
	0x04, 0xea, 0x8c, 0xce, 0x1f, 0x06, 0x58, 0xf3, 0x67, 0x3f, 0xa6, 0xdd, 0x7f, 0x04, 0xed, 0xdc,
	0xd6, 0x98, 0x3a, 0xb8, 0xb4, 0xbd, 0x46, 0xff, 0xe5, 0x72, 0x5b, 0x1c, 0xa9, 0x8d, 0xbb, 0xce,
	0x8e, 0x14, 0xeb, 0x6d, 0xa8, 0x65, 0xef, 0x81, 0x34, 0xb6, 0x65, 0xaf, 0x46, 0x1e, 0xe5, 0x7c,
	0x55, 0x48, 0x44, 0xdd, 0xe5, 0x49, 0x24, 0x62, 0x41, 0x35, 0x20, 0x8c, 0x27, 0xf4, 0x50, 0x8b,
	0x23, 0x9b, 0x3a, 0x3b, 0x70, 0x21, 0xdf, 0x94, 0xf8, 0x74, 0x4a, 0x92, 0x93, 0xe8, 0x62, 0x54,
	0xf4, 0xc7, 0xf6, 0x98, 0xf8, 0xfb, 0x27, 0xe1, 0x78, 0x19, 0x9a, 0x71, 0x9a, 0xf8, 0x63, 0xcc,
	0x88, 0x7a, 0x39, 0x54, 0x53, 0xaf, 0x65, 0xa0, 0x78, 0x39, 0xfa, 0x7f, 0xd5, 0xa0, 0x9e, 0xed,
	0xc4, 0xd0, 0x8f, 0x06, 0x9c, 0x15, 0xaf, 0x37, 0xba, 0xb4, 0xa8, 0x88, 0xa5, 0x1f, 0x0d, 0x76,
	0x67, 0x79, 0x80, 0x52, 0x82, 0xf3, 0xc9, 0x93, 0xdf, 0xac, 0x95, 0x9a, 0xf1, 0xe4, 0xcf, 0xbf,
	0xbf, 0x5f, 0xf9, 0x10, 0x7d, 0xd0, 0x9b, 0xf9, 0x41, 0xb9, 0x9f, 0x0e, 0x49, 0x12, 0x11, 0x4e,
	0x58, 0x4f, 0x27, 0xe9, 0xe9, 0x83, 0xb0, 0xde, 0x23, 0x3d, 0x7a, 0xdc, 0xcb, 0x5f, 0x7a, 0xf4,
	0x9d, 0x01, 0x15, 0x65, 0x8d, 0xe8, 0xd5, 0xff, 0xb5, 0x4d, 0xfb, 0x7c, 0x26, 0xba, 0x3d, 0x1a,
	0x06, 0x39, 0xa1, 0xdb, 0x25, 0x42, 0x5b, 0xce, 0xf3, 0x11, 0xba, 0x6e, 0x5c, 0x45, 0x4f, 0x0d,
	0x58, 0x95, 0xe6, 0x85, 0x16, 0x16, 0xa3, 0xec, 0x6b, 0x4b, 0x18, 0x0d, 0x4b, 0x8c, 0xf6, 0xec,
	0xbb, 0xcf, 0xc5, 0xa8, 0xf7, 0x28, 0x0d, 0x83, 0xc7, 0x3d, 0xec, 0xcb, 0x5f, 0xe9, 0x3d, 0x26,
	0x76, 0x17, 0x2c, 0x7f, 0x32, 0xa0, 0xa2, 0xac, 0x70, 0x71, 0xe5, 0x66, 0x6c, 0x72, 0x09, 0xcf,
	0x7b, 0x25, 0x9e, 0xbb, 0xf6, 0xf6, 0x29, 0xf0, 0x14, 0xcc, 0x7e, 0x31, 0xa0, 0x96, 0xf9, 0x07,
	0xba, 0xbc, 0x88, 0xdb, 0x11, 0x67, 0xb5, 0xaf, 0xfc, 0x77, 0x90, 0x66, 0x7b, 0xb7, 0xc4, 0xf6,
	0x26, 0x3a, 0x0d, 0xb6, 0xe8, 0x07, 0x03, 0x2a, 0xca, 0x2c, 0x16, 0x17, 0x71, 0xc6, 0x48, 0x96,
	0x14, 0x71, 0x86, 0xd6, 0xd5, 0x53, 0xa1, 0xf5, 0xd4, 0x80, 0xaa, 0x36, 0x19, 0xe4, 0x2c, 0xb4,
	0xbd, 0x19, 0x07, 0x5a, 0x42, 0xec, 0x6b, 0x49, 0xe9, 0x0b, 0x7b, 0x70, 0x9a, 0xfa, 0x4b, 0xd4,
	0xce, 0xd7, 0x8d, 0xab, 0x5b, 0xf5, 0xfb, 0x55, 0x9d, 0x60, 0x58, 0x91, 0xff, 0x24, 0xbe, 0xf3,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x07, 0x03, 0xe3, 0x98, 0x0e, 0x00, 0x00,
}
