// Code generated by protoc-gen-go.
// source: kvrpcpb.proto
// DO NOT EDIT!

/*
Package kvrpcpb is a generated protocol buffer package.

It is generated from these files:
	kvrpcpb.proto

It has these top-level messages:
	ResultType
	KeyAddress
	KvPair
	CmdGetRequest
	CmdGetResponse
	CmdScanRequest
	Item
	CmdScanResponse
	CmdPrewriteRequest
	CmdPrewriteResponse
	CmdCommitRequest
	CmdCommitResponse
	CmdCleanupRequest
	CmdCleanupResponse
	CmdRollbackThenGetRequest
	CmdRollbackThenGetResponse
	CmdCommitThenGetRequest
	CmdCommitThenGetResponse
	Request
	Response
*/
package kvrpcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import metapb "github.com/pingcap/kvproto/pkg/metapb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MessageType int32

const (
	MessageType_CmdGet      MessageType = 1
	MessageType_CmdScan     MessageType = 2
	MessageType_CmdPrewrite MessageType = 3
	MessageType_CmdCommit   MessageType = 4
	MessageType_CmdCleanup  MessageType = 5
	// Below types both use for Get failed. If Get failed, it may be locked.
	// So it tries to clean primary lock(CmdCleanup), and then server will return
	// either committed or rolled back. Finally, client will commit/rollback
	// primary lock and THEN Get again.
	MessageType_CmdRollbackThenGet MessageType = 6
	MessageType_CmdCommitThenGet   MessageType = 7
)

var MessageType_name = map[int32]string{
	1: "CmdGet",
	2: "CmdScan",
	3: "CmdPrewrite",
	4: "CmdCommit",
	5: "CmdCleanup",
	6: "CmdRollbackThenGet",
	7: "CmdCommitThenGet",
}
var MessageType_value = map[string]int32{
	"CmdGet":             1,
	"CmdScan":            2,
	"CmdPrewrite":        3,
	"CmdCommit":          4,
	"CmdCleanup":         5,
	"CmdRollbackThenGet": 6,
	"CmdCommitThenGet":   7,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ResultType_Type int32

const (
	ResultType_Ok         ResultType_Type = 1
	ResultType_Retryable  ResultType_Type = 2
	ResultType_Locked     ResultType_Type = 3
	ResultType_Committed  ResultType_Type = 4
	ResultType_Rolledback ResultType_Type = 5
	// Known result type add here.
	ResultType_Other ResultType_Type = 9
)

var ResultType_Type_name = map[int32]string{
	1: "Ok",
	2: "Retryable",
	3: "Locked",
	4: "Committed",
	5: "Rolledback",
	9: "Other",
}
var ResultType_Type_value = map[string]int32{
	"Ok":         1,
	"Retryable":  2,
	"Locked":     3,
	"Committed":  4,
	"Rolledback": 5,
	"Other":      9,
}

func (x ResultType_Type) Enum() *ResultType_Type {
	p := new(ResultType_Type)
	*p = x
	return p
}
func (x ResultType_Type) String() string {
	return proto.EnumName(ResultType_Type_name, int32(x))
}
func (x *ResultType_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ResultType_Type_value, data, "ResultType_Type")
	if err != nil {
		return err
	}
	*x = ResultType_Type(value)
	return nil
}
func (ResultType_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// ResultType use for many different result, so dispose each Type prudently.
type ResultType struct {
	Type             *ResultType_Type `protobuf:"varint,1,opt,name=type,enum=kvrpcpb.ResultType_Type,def=9" json:"type,omitempty"`
	Msg              *string          `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ResultType) Reset()                    { *m = ResultType{} }
func (m *ResultType) String() string            { return proto.CompactTextString(m) }
func (*ResultType) ProtoMessage()               {}
func (*ResultType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ResultType_Type ResultType_Type = ResultType_Other

func (m *ResultType) GetType() ResultType_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_ResultType_Type
}

func (m *ResultType) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

type KeyAddress struct {
	Key              []byte       `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	RegionId         *uint64      `protobuf:"varint,2,opt,name=region_id" json:"region_id,omitempty"`
	Peer             *metapb.Peer `protobuf:"bytes,3,opt,name=peer" json:"peer,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *KeyAddress) Reset()                    { *m = KeyAddress{} }
func (m *KeyAddress) String() string            { return proto.CompactTextString(m) }
func (*KeyAddress) ProtoMessage()               {}
func (*KeyAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeyAddress) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyAddress) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *KeyAddress) GetPeer() *metapb.Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type KvPair struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	Value            []byte      `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *KvPair) Reset()                    { *m = KvPair{} }
func (m *KvPair) String() string            { return proto.CompactTextString(m) }
func (*KvPair) ProtoMessage()               {}
func (*KvPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KvPair) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *KvPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type CmdGetRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	Version          *uint64     `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdGetRequest) Reset()                    { *m = CmdGetRequest{} }
func (m *CmdGetRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdGetRequest) ProtoMessage()               {}
func (*CmdGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CmdGetRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdGetRequest) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type CmdGetResponse struct {
	ResType          *ResultType `protobuf:"bytes,1,opt,name=res_type" json:"res_type,omitempty"`
	Value            []byte      `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	PrimaryLock      []byte      `protobuf:"bytes,3,opt,name=primary_lock" json:"primary_lock,omitempty"`
	LockVersion      *uint64     `protobuf:"varint,4,opt,name=lock_version" json:"lock_version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdGetResponse) Reset()                    { *m = CmdGetResponse{} }
func (m *CmdGetResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdGetResponse) ProtoMessage()               {}
func (*CmdGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CmdGetResponse) GetResType() *ResultType {
	if m != nil {
		return m.ResType
	}
	return nil
}

func (m *CmdGetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *CmdGetResponse) GetPrimaryLock() []byte {
	if m != nil {
		return m.PrimaryLock
	}
	return nil
}

func (m *CmdGetResponse) GetLockVersion() uint64 {
	if m != nil && m.LockVersion != nil {
		return *m.LockVersion
	}
	return 0
}

type CmdScanRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	Limit            *uint32     `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Version          *uint64     `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdScanRequest) Reset()                    { *m = CmdScanRequest{} }
func (m *CmdScanRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdScanRequest) ProtoMessage()               {}
func (*CmdScanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CmdScanRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdScanRequest) GetLimit() uint32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *CmdScanRequest) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type Item struct {
	ResType *ResultType `protobuf:"bytes,1,opt,name=res_type" json:"res_type,omitempty"`
	Key     []byte      `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value   []byte      `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	// primary_lock_key
	PrimaryLock      []byte  `protobuf:"bytes,4,opt,name=primary_lock" json:"primary_lock,omitempty"`
	LockVersion      *uint64 `protobuf:"varint,5,opt,name=lock_version" json:"lock_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Item) GetResType() *ResultType {
	if m != nil {
		return m.ResType
	}
	return nil
}

func (m *Item) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Item) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Item) GetPrimaryLock() []byte {
	if m != nil {
		return m.PrimaryLock
	}
	return nil
}

func (m *Item) GetLockVersion() uint64 {
	if m != nil && m.LockVersion != nil {
		return *m.LockVersion
	}
	return 0
}

type CmdScanResponse struct {
	// ok if !ok then retry.
	Ok               *bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Results          []*Item `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdScanResponse) Reset()                    { *m = CmdScanResponse{} }
func (m *CmdScanResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdScanResponse) ProtoMessage()               {}
func (*CmdScanResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CmdScanResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *CmdScanResponse) GetResults() []*Item {
	if m != nil {
		return m.Results
	}
	return nil
}

type CmdPrewriteRequest struct {
	Puts  []*KvPair     `protobuf:"bytes,1,rep,name=puts" json:"puts,omitempty"`
	Dels  []*KeyAddress `protobuf:"bytes,2,rep,name=dels" json:"dels,omitempty"`
	Locks []*KeyAddress `protobuf:"bytes,3,rep,name=locks" json:"locks,omitempty"`
	// primary_lock_key
	PrimaryLock      []byte  `protobuf:"bytes,4,opt,name=primary_lock" json:"primary_lock,omitempty"`
	StartVersion     *uint64 `protobuf:"varint,5,opt,name=start_version" json:"start_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdPrewriteRequest) Reset()                    { *m = CmdPrewriteRequest{} }
func (m *CmdPrewriteRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdPrewriteRequest) ProtoMessage()               {}
func (*CmdPrewriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CmdPrewriteRequest) GetPuts() []*KvPair {
	if m != nil {
		return m.Puts
	}
	return nil
}

func (m *CmdPrewriteRequest) GetDels() []*KeyAddress {
	if m != nil {
		return m.Dels
	}
	return nil
}

func (m *CmdPrewriteRequest) GetLocks() []*KeyAddress {
	if m != nil {
		return m.Locks
	}
	return nil
}

func (m *CmdPrewriteRequest) GetPrimaryLock() []byte {
	if m != nil {
		return m.PrimaryLock
	}
	return nil
}

func (m *CmdPrewriteRequest) GetStartVersion() uint64 {
	if m != nil && m.StartVersion != nil {
		return *m.StartVersion
	}
	return 0
}

type CmdPrewriteResponse struct {
	Ok *bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	// This Item doesn't contain value = 3
	Results          []*Item `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CmdPrewriteResponse) Reset()                    { *m = CmdPrewriteResponse{} }
func (m *CmdPrewriteResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdPrewriteResponse) ProtoMessage()               {}
func (*CmdPrewriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CmdPrewriteResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *CmdPrewriteResponse) GetResults() []*Item {
	if m != nil {
		return m.Results
	}
	return nil
}

type CmdCommitRequest struct {
	StartVersion     *uint64       `protobuf:"varint,1,opt,name=start_version" json:"start_version,omitempty"`
	KeysAddress      []*KeyAddress `protobuf:"bytes,2,rep,name=keys_address" json:"keys_address,omitempty"`
	CommitVersion    *uint64       `protobuf:"varint,3,opt,name=commit_version" json:"commit_version,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *CmdCommitRequest) Reset()                    { *m = CmdCommitRequest{} }
func (m *CmdCommitRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdCommitRequest) ProtoMessage()               {}
func (*CmdCommitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CmdCommitRequest) GetStartVersion() uint64 {
	if m != nil && m.StartVersion != nil {
		return *m.StartVersion
	}
	return 0
}

func (m *CmdCommitRequest) GetKeysAddress() []*KeyAddress {
	if m != nil {
		return m.KeysAddress
	}
	return nil
}

func (m *CmdCommitRequest) GetCommitVersion() uint64 {
	if m != nil && m.CommitVersion != nil {
		return *m.CommitVersion
	}
	return 0
}

type CmdCommitResponse struct {
	Ok               *bool  `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CmdCommitResponse) Reset()                    { *m = CmdCommitResponse{} }
func (m *CmdCommitResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdCommitResponse) ProtoMessage()               {}
func (*CmdCommitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CmdCommitResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

type CmdCleanupRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	StartVersion     *uint64     `protobuf:"varint,2,opt,name=start_version" json:"start_version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdCleanupRequest) Reset()                    { *m = CmdCleanupRequest{} }
func (m *CmdCleanupRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdCleanupRequest) ProtoMessage()               {}
func (*CmdCleanupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CmdCleanupRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdCleanupRequest) GetStartVersion() uint64 {
	if m != nil && m.StartVersion != nil {
		return *m.StartVersion
	}
	return 0
}

type CmdCleanupResponse struct {
	ResType          *ResultType `protobuf:"bytes,1,opt,name=res_type" json:"res_type,omitempty"`
	CommitVersion    *uint64     `protobuf:"varint,2,opt,name=commit_version" json:"commit_version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdCleanupResponse) Reset()                    { *m = CmdCleanupResponse{} }
func (m *CmdCleanupResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdCleanupResponse) ProtoMessage()               {}
func (*CmdCleanupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *CmdCleanupResponse) GetResType() *ResultType {
	if m != nil {
		return m.ResType
	}
	return nil
}

func (m *CmdCleanupResponse) GetCommitVersion() uint64 {
	if m != nil && m.CommitVersion != nil {
		return *m.CommitVersion
	}
	return 0
}

type CmdRollbackThenGetRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	LockVersion      *uint64     `protobuf:"varint,2,opt,name=lock_version" json:"lock_version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdRollbackThenGetRequest) Reset()                    { *m = CmdRollbackThenGetRequest{} }
func (m *CmdRollbackThenGetRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdRollbackThenGetRequest) ProtoMessage()               {}
func (*CmdRollbackThenGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *CmdRollbackThenGetRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdRollbackThenGetRequest) GetLockVersion() uint64 {
	if m != nil && m.LockVersion != nil {
		return *m.LockVersion
	}
	return 0
}

type CmdRollbackThenGetResponse struct {
	Ok               *bool  `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CmdRollbackThenGetResponse) Reset()                    { *m = CmdRollbackThenGetResponse{} }
func (m *CmdRollbackThenGetResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdRollbackThenGetResponse) ProtoMessage()               {}
func (*CmdRollbackThenGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *CmdRollbackThenGetResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *CmdRollbackThenGetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type CmdCommitThenGetRequest struct {
	KeyAddress       *KeyAddress `protobuf:"bytes,1,opt,name=key_address" json:"key_address,omitempty"`
	LockVersion      *uint64     `protobuf:"varint,2,opt,name=lock_version" json:"lock_version,omitempty"`
	CommitVersion    *uint64     `protobuf:"varint,3,opt,name=commit_version" json:"commit_version,omitempty"`
	GetVersion       *uint64     `protobuf:"varint,4,opt,name=get_version" json:"get_version,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CmdCommitThenGetRequest) Reset()                    { *m = CmdCommitThenGetRequest{} }
func (m *CmdCommitThenGetRequest) String() string            { return proto.CompactTextString(m) }
func (*CmdCommitThenGetRequest) ProtoMessage()               {}
func (*CmdCommitThenGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *CmdCommitThenGetRequest) GetKeyAddress() *KeyAddress {
	if m != nil {
		return m.KeyAddress
	}
	return nil
}

func (m *CmdCommitThenGetRequest) GetLockVersion() uint64 {
	if m != nil && m.LockVersion != nil {
		return *m.LockVersion
	}
	return 0
}

func (m *CmdCommitThenGetRequest) GetCommitVersion() uint64 {
	if m != nil && m.CommitVersion != nil {
		return *m.CommitVersion
	}
	return 0
}

func (m *CmdCommitThenGetRequest) GetGetVersion() uint64 {
	if m != nil && m.GetVersion != nil {
		return *m.GetVersion
	}
	return 0
}

type CmdCommitThenGetResponse struct {
	Ok               *bool  `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CmdCommitThenGetResponse) Reset()                    { *m = CmdCommitThenGetResponse{} }
func (m *CmdCommitThenGetResponse) String() string            { return proto.CompactTextString(m) }
func (*CmdCommitThenGetResponse) ProtoMessage()               {}
func (*CmdCommitThenGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *CmdCommitThenGetResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func (m *CmdCommitThenGetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Request struct {
	Type             *MessageType               `protobuf:"varint,1,opt,name=type,enum=kvrpcpb.MessageType" json:"type,omitempty"`
	CmdGetReq        *CmdGetRequest             `protobuf:"bytes,2,opt,name=cmd_get_req" json:"cmd_get_req,omitempty"`
	CmdScanReq       *CmdScanRequest            `protobuf:"bytes,3,opt,name=cmd_scan_req" json:"cmd_scan_req,omitempty"`
	CmdPrewriteReq   *CmdPrewriteRequest        `protobuf:"bytes,4,opt,name=cmd_prewrite_req" json:"cmd_prewrite_req,omitempty"`
	CmdCommitReq     *CmdCommitRequest          `protobuf:"bytes,5,opt,name=cmd_commit_req" json:"cmd_commit_req,omitempty"`
	CmdCleanupReq    *CmdCleanupRequest         `protobuf:"bytes,6,opt,name=cmd_cleanup_req" json:"cmd_cleanup_req,omitempty"`
	CmdRbGetReq      *CmdRollbackThenGetRequest `protobuf:"bytes,7,opt,name=cmd_rb_get_req" json:"cmd_rb_get_req,omitempty"`
	CmdCommitGetReq  *CmdCommitThenGetRequest   `protobuf:"bytes,8,opt,name=cmd_commit_get_req" json:"cmd_commit_get_req,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *Request) GetType() MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MessageType_CmdGet
}

func (m *Request) GetCmdGetReq() *CmdGetRequest {
	if m != nil {
		return m.CmdGetReq
	}
	return nil
}

func (m *Request) GetCmdScanReq() *CmdScanRequest {
	if m != nil {
		return m.CmdScanReq
	}
	return nil
}

func (m *Request) GetCmdPrewriteReq() *CmdPrewriteRequest {
	if m != nil {
		return m.CmdPrewriteReq
	}
	return nil
}

func (m *Request) GetCmdCommitReq() *CmdCommitRequest {
	if m != nil {
		return m.CmdCommitReq
	}
	return nil
}

func (m *Request) GetCmdCleanupReq() *CmdCleanupRequest {
	if m != nil {
		return m.CmdCleanupReq
	}
	return nil
}

func (m *Request) GetCmdRbGetReq() *CmdRollbackThenGetRequest {
	if m != nil {
		return m.CmdRbGetReq
	}
	return nil
}

func (m *Request) GetCmdCommitGetReq() *CmdCommitThenGetRequest {
	if m != nil {
		return m.CmdCommitGetReq
	}
	return nil
}

type Response struct {
	Type             *MessageType                `protobuf:"varint,1,opt,name=type,enum=kvrpcpb.MessageType" json:"type,omitempty"`
	CmdGetResp       *CmdGetResponse             `protobuf:"bytes,2,opt,name=cmd_get_resp" json:"cmd_get_resp,omitempty"`
	CmdScanResp      *CmdScanResponse            `protobuf:"bytes,3,opt,name=cmd_scan_resp" json:"cmd_scan_resp,omitempty"`
	CmdPrewriteResp  *CmdPrewriteResponse        `protobuf:"bytes,4,opt,name=cmd_prewrite_resp" json:"cmd_prewrite_resp,omitempty"`
	CmdCommitResp    *CmdCommitResponse          `protobuf:"bytes,5,opt,name=cmd_commit_resp" json:"cmd_commit_resp,omitempty"`
	CmdCleanupResp   *CmdCleanupResponse         `protobuf:"bytes,6,opt,name=cmd_cleanup_resp" json:"cmd_cleanup_resp,omitempty"`
	CmdRbGetResp     *CmdRollbackThenGetResponse `protobuf:"bytes,7,opt,name=cmd_rb_get_resp" json:"cmd_rb_get_resp,omitempty"`
	CmdCommitGetResp *CmdCommitThenGetResponse   `protobuf:"bytes,8,opt,name=cmd_commit_get_resp" json:"cmd_commit_get_resp,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *Response) GetType() MessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MessageType_CmdGet
}

func (m *Response) GetCmdGetResp() *CmdGetResponse {
	if m != nil {
		return m.CmdGetResp
	}
	return nil
}

func (m *Response) GetCmdScanResp() *CmdScanResponse {
	if m != nil {
		return m.CmdScanResp
	}
	return nil
}

func (m *Response) GetCmdPrewriteResp() *CmdPrewriteResponse {
	if m != nil {
		return m.CmdPrewriteResp
	}
	return nil
}

func (m *Response) GetCmdCommitResp() *CmdCommitResponse {
	if m != nil {
		return m.CmdCommitResp
	}
	return nil
}

func (m *Response) GetCmdCleanupResp() *CmdCleanupResponse {
	if m != nil {
		return m.CmdCleanupResp
	}
	return nil
}

func (m *Response) GetCmdRbGetResp() *CmdRollbackThenGetResponse {
	if m != nil {
		return m.CmdRbGetResp
	}
	return nil
}

func (m *Response) GetCmdCommitGetResp() *CmdCommitThenGetResponse {
	if m != nil {
		return m.CmdCommitGetResp
	}
	return nil
}

func init() {
	proto.RegisterType((*ResultType)(nil), "kvrpcpb.ResultType")
	proto.RegisterType((*KeyAddress)(nil), "kvrpcpb.KeyAddress")
	proto.RegisterType((*KvPair)(nil), "kvrpcpb.KvPair")
	proto.RegisterType((*CmdGetRequest)(nil), "kvrpcpb.CmdGetRequest")
	proto.RegisterType((*CmdGetResponse)(nil), "kvrpcpb.CmdGetResponse")
	proto.RegisterType((*CmdScanRequest)(nil), "kvrpcpb.CmdScanRequest")
	proto.RegisterType((*Item)(nil), "kvrpcpb.Item")
	proto.RegisterType((*CmdScanResponse)(nil), "kvrpcpb.CmdScanResponse")
	proto.RegisterType((*CmdPrewriteRequest)(nil), "kvrpcpb.CmdPrewriteRequest")
	proto.RegisterType((*CmdPrewriteResponse)(nil), "kvrpcpb.CmdPrewriteResponse")
	proto.RegisterType((*CmdCommitRequest)(nil), "kvrpcpb.CmdCommitRequest")
	proto.RegisterType((*CmdCommitResponse)(nil), "kvrpcpb.CmdCommitResponse")
	proto.RegisterType((*CmdCleanupRequest)(nil), "kvrpcpb.CmdCleanupRequest")
	proto.RegisterType((*CmdCleanupResponse)(nil), "kvrpcpb.CmdCleanupResponse")
	proto.RegisterType((*CmdRollbackThenGetRequest)(nil), "kvrpcpb.CmdRollbackThenGetRequest")
	proto.RegisterType((*CmdRollbackThenGetResponse)(nil), "kvrpcpb.CmdRollbackThenGetResponse")
	proto.RegisterType((*CmdCommitThenGetRequest)(nil), "kvrpcpb.CmdCommitThenGetRequest")
	proto.RegisterType((*CmdCommitThenGetResponse)(nil), "kvrpcpb.CmdCommitThenGetResponse")
	proto.RegisterType((*Request)(nil), "kvrpcpb.Request")
	proto.RegisterType((*Response)(nil), "kvrpcpb.Response")
	proto.RegisterEnum("kvrpcpb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("kvrpcpb.ResultType_Type", ResultType_Type_name, ResultType_Type_value)
}

var fileDescriptor0 = []byte{
	// 900 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0x4b, 0x6f, 0xeb, 0x44,
	0x14, 0x96, 0x63, 0xe7, 0x75, 0x1c, 0x27, 0xee, 0xe4, 0xd2, 0xe6, 0x86, 0xd7, 0xbd, 0x46, 0x48,
	0x17, 0xd0, 0x0d, 0x22, 0xe8, 0xaa, 0x52, 0x55, 0x90, 0xaa, 0x2e, 0x78, 0x14, 0xd4, 0xaa, 0xed,
	0x0a, 0x16, 0x91, 0x13, 0x8f, 0xd2, 0x28, 0x76, 0xec, 0x8e, 0x27, 0x41, 0xd9, 0x23, 0xfe, 0x05,
	0x1b, 0x36, 0xac, 0xf9, 0x87, 0x1c, 0xcf, 0xd8, 0x8e, 0x5f, 0x2d, 0x4d, 0xc5, 0xa6, 0x95, 0x9d,
	0xf3, 0x7d, 0xe7, 0x3b, 0xe7, 0x3b, 0x73, 0x3c, 0x60, 0x2c, 0x37, 0x2c, 0x98, 0x05, 0xd3, 0x51,
	0xc0, 0x7c, 0xee, 0x93, 0x66, 0xfc, 0x38, 0xec, 0x78, 0x94, 0xdb, 0xc9, 0x6b, 0xeb, 0x4f, 0x05,
	0xe0, 0x9a, 0x86, 0x6b, 0x97, 0xdf, 0x6e, 0x03, 0x4a, 0xde, 0x82, 0xc6, 0xf1, 0xff, 0x40, 0x79,
	0xa5, 0xbc, 0xe9, 0x8e, 0x07, 0xa3, 0x84, 0x63, 0x17, 0x32, 0x8a, 0xfe, 0x9c, 0xd4, 0x2f, 0xf9,
	0x1d, 0x65, 0x44, 0x07, 0xd5, 0x0b, 0xe7, 0x83, 0x1a, 0x46, 0xb7, 0xad, 0x1b, 0xd0, 0x04, 0x47,
	0x03, 0x6a, 0x97, 0x4b, 0x53, 0x21, 0x06, 0xb4, 0xaf, 0x29, 0x67, 0x5b, 0x7b, 0xea, 0x52, 0xb3,
	0x46, 0x00, 0x1a, 0x3f, 0xf9, 0xb3, 0x25, 0x75, 0x4c, 0x35, 0xfa, 0xe9, 0xdc, 0xf7, 0xbc, 0x05,
	0xe7, 0xf8, 0xa8, 0x91, 0x2e, 0x6a, 0xf0, 0x5d, 0x97, 0x3a, 0x53, 0x7b, 0xb6, 0x34, 0xeb, 0xa4,
	0x0d, 0x92, 0xdf, 0x6c, 0x5b, 0xdf, 0x03, 0x5c, 0xd0, 0xed, 0x99, 0xe3, 0x30, 0x1a, 0x86, 0x51,
	0xbe, 0x25, 0xdd, 0x0a, 0x75, 0x1d, 0x72, 0x00, 0x6d, 0x46, 0xe7, 0x0b, 0x7f, 0x35, 0x59, 0x38,
	0x42, 0x82, 0x46, 0x86, 0xa0, 0x05, 0x94, 0xb2, 0x81, 0x8a, 0x4f, 0xfa, 0xb8, 0x33, 0x8a, 0x4b,
	0xbd, 0xc2, 0x77, 0xd6, 0x19, 0x34, 0x2e, 0x36, 0x57, 0xf6, 0x82, 0x91, 0x37, 0xa0, 0x23, 0xcb,
	0xc4, 0x96, 0xa4, 0x82, 0x4d, 0x1f, 0xf7, 0xd3, 0x5a, 0x33, 0xf9, 0x0c, 0xa8, 0x6f, 0x6c, 0x77,
	0x4d, 0x05, 0x7d, 0xc7, 0xfa, 0x11, 0x8c, 0x73, 0xcf, 0xf9, 0x8e, 0xf2, 0x6b, 0x7a, 0xbf, 0xa6,
	0x21, 0xdf, 0x83, 0xa9, 0x07, 0xcd, 0x0d, 0x65, 0x21, 0xaa, 0x95, 0x52, 0x2d, 0x06, 0xdd, 0x84,
	0x2b, 0x0c, 0xfc, 0x55, 0x48, 0xc9, 0xa7, 0xd0, 0xc2, 0xd0, 0x49, 0xda, 0xff, 0x2c, 0x53, 0xc6,
	0xa2, 0xbc, 0x26, 0xf2, 0x02, 0x3a, 0x01, 0x5b, 0x78, 0x36, 0xdb, 0x4e, 0x5c, 0x6c, 0xaf, 0x28,
	0x5d, 0xbc, 0x8d, 0x9e, 0x26, 0x49, 0x4e, 0x4d, 0xe4, 0xfc, 0x45, 0xe4, 0xbc, 0x99, 0xd9, 0xab,
	0xfd, 0x0b, 0xc0, 0xb4, 0xee, 0x02, 0x1d, 0x13, 0x69, 0x8d, 0x6c, 0x3d, 0xaa, 0xe0, 0x5e, 0x83,
	0xf6, 0x03, 0xa7, 0xde, 0x53, 0xab, 0x88, 0x9d, 0x94, 0x35, 0xa4, 0x25, 0xa9, 0x95, 0x25, 0x69,
	0x95, 0x25, 0xd5, 0x45, 0xda, 0x6f, 0xa0, 0x97, 0x96, 0x14, 0xf7, 0x11, 0xa0, 0xe6, 0x2f, 0x45,
	0xee, 0x16, 0xf9, 0x08, 0x9a, 0x4c, 0x24, 0x0d, 0x31, 0x95, 0x8a, 0x62, 0x8c, 0x54, 0x4c, 0xa4,
	0xd6, 0xfa, 0x5b, 0x01, 0x82, 0xf8, 0x2b, 0x46, 0x7f, 0x63, 0x0b, 0x4e, 0x93, 0xb6, 0x7c, 0x88,
	0x73, 0xb4, 0xe6, 0x51, 0x3f, 0x22, 0x4c, 0x6f, 0xd7, 0x0f, 0x39, 0x40, 0xaf, 0x41, 0x73, 0xa8,
	0x9b, 0x50, 0x56, 0xb6, 0xcb, 0xc2, 0x76, 0xa1, 0xda, 0x10, 0x4b, 0x7a, 0x30, 0xa6, 0xba, 0xce,
	0xf7, 0xc0, 0x08, 0xb9, 0xcd, 0x78, 0xa1, 0xd0, 0x33, 0xe8, 0xe7, 0x84, 0x3e, 0xa3, 0x58, 0x17,
	0x4c, 0xa4, 0x90, 0x07, 0x2f, 0xa9, 0xb4, 0x94, 0x4d, 0x11, 0x07, 0xe9, 0x33, 0xe8, 0xa0, 0x3d,
	0x61, 0x3a, 0x18, 0x8f, 0x54, 0x7a, 0x08, 0xdd, 0x99, 0xa0, 0x9c, 0xe4, 0x07, 0xe2, 0x63, 0x38,
	0xc8, 0x64, 0x2b, 0xcb, 0xb5, 0x6e, 0x65, 0x80, 0x4b, 0xed, 0xd5, 0x3a, 0xd8, 0x7f, 0x20, 0x4b,
	0xca, 0xe5, 0xb9, 0xba, 0x11, 0x86, 0xa6, 0xac, 0xfb, 0x9d, 0xad, 0x72, 0x2d, 0x92, 0xf4, 0x57,
	0x78, 0x89, 0xa4, 0xd1, 0x8e, 0x8a, 0x36, 0xd4, 0xed, 0x1d, 0x5d, 0x3d, 0x6b, 0x09, 0x14, 0x47,
	0x58, 0x92, 0x1f, 0xc3, 0xb0, 0x8a, 0xbc, 0xc2, 0xe0, 0xc2, 0x3a, 0xfa, 0x43, 0x81, 0xa3, 0xb4,
	0xc5, 0xff, 0xaf, 0xa8, 0x87, 0x5c, 0x25, 0x7d, 0xd0, 0xe7, 0x94, 0x17, 0xf6, 0xca, 0x3b, 0x18,
	0x94, 0x75, 0xfc, 0xb7, 0xfe, 0xbf, 0x54, 0x68, 0x26, 0x7a, 0xad, 0xdc, 0x87, 0xe7, 0x45, 0x2a,
	0xf4, 0x67, 0x94, 0x68, 0xcf, 0xa9, 0x70, 0xe7, 0x0b, 0xd0, 0x67, 0x9e, 0x33, 0x89, 0xf2, 0x33,
	0x7a, 0x2f, 0x48, 0xf4, 0xf1, 0x61, 0x1a, 0x9a, 0x5f, 0xcd, 0x6f, 0xa1, 0x13, 0x05, 0x87, 0xb8,
	0x19, 0x44, 0xb4, 0xfc, 0x24, 0x1c, 0x65, 0xa3, 0xb3, 0x8b, 0xf0, 0x1d, 0x98, 0x51, 0x78, 0x10,
	0x9f, 0x2f, 0x01, 0xd1, 0x04, 0xe4, 0xfd, 0x2c, 0xa4, 0xb8, 0x28, 0xbe, 0xc2, 0x36, 0x21, 0x2c,
	0x6e, 0x55, 0x04, 0xaa, 0x0b, 0xd0, 0xcb, 0x2c, 0x28, 0x7f, 0xe2, 0xbe, 0x86, 0x9e, 0x80, 0xc8,
	0x09, 0x15, 0x98, 0x86, 0xc0, 0x0c, 0x73, 0x98, 0xfc, 0xb1, 0x38, 0x91, 0x79, 0xd8, 0x34, 0xad,
	0xbe, 0x29, 0x30, 0x56, 0x16, 0xf3, 0xc0, 0x7c, 0x9e, 0x02, 0xc9, 0x68, 0x4c, 0xf0, 0x2d, 0x81,
	0x7f, 0x55, 0xd6, 0x99, 0x47, 0x5b, 0xff, 0xa8, 0xd0, 0x4a, 0xcd, 0x7c, 0x8a, 0x4b, 0x71, 0xe3,
	0x65, 0x9e, 0x30, 0x88, 0x6d, 0x3a, 0x2a, 0xd9, 0x14, 0x53, 0x7e, 0x09, 0x46, 0xc6, 0x27, 0x8c,
	0x97, 0x46, 0x0d, 0xca, 0x46, 0xc5, 0x80, 0x63, 0x38, 0x28, 0x38, 0x85, 0x20, 0x69, 0xd5, 0x07,
	0xd5, 0x56, 0xc5, 0xc0, 0xa4, 0xf1, 0x89, 0x57, 0x08, 0xab, 0x57, 0x34, 0x3e, 0xbf, 0xb0, 0xe2,
	0xb9, 0xd8, 0xb9, 0x85, 0xa8, 0x46, 0x79, 0x2e, 0x8a, 0xfb, 0xe6, 0x54, 0xe6, 0x4a, 0xfd, 0x42,
	0x94, 0x34, 0xec, 0x93, 0x47, 0x0d, 0x8b, 0xd1, 0xdf, 0x42, 0xbf, 0xe4, 0x18, 0x32, 0x48, 0xcb,
	0x5e, 0x3f, 0x62, 0x99, 0xc4, 0x7f, 0xfe, 0xbb, 0x02, 0x7a, 0xd6, 0x12, 0xbc, 0x7a, 0xc9, 0xae,
	0xe3, 0xad, 0x4c, 0x87, 0x66, 0xdc, 0x51, 0xbc, 0x93, 0xf5, 0x40, 0xcf, 0x74, 0x2a, 0xbe, 0x98,
	0x25, 0xac, 0xf2, 0x62, 0xb6, 0x2b, 0x0e, 0x2f, 0x66, 0x87, 0x62, 0xb9, 0x16, 0x64, 0x9b, 0x0d,
	0xdc, 0x21, 0x66, 0x51, 0x8c, 0xd9, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x02, 0xac, 0x7a,
	0x82, 0x0a, 0x00, 0x00,
}