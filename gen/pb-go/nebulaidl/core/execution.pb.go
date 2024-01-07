// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/core/execution.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WorkflowExecution_Phase int32

const (
	WorkflowExecution_UNDEFINED  WorkflowExecution_Phase = 0
	WorkflowExecution_QUEUED     WorkflowExecution_Phase = 1
	WorkflowExecution_RUNNING    WorkflowExecution_Phase = 2
	WorkflowExecution_SUCCEEDING WorkflowExecution_Phase = 3
	WorkflowExecution_SUCCEEDED  WorkflowExecution_Phase = 4
	WorkflowExecution_FAILING    WorkflowExecution_Phase = 5
	WorkflowExecution_FAILED     WorkflowExecution_Phase = 6
	WorkflowExecution_ABORTED    WorkflowExecution_Phase = 7
	WorkflowExecution_TIMED_OUT  WorkflowExecution_Phase = 8
	WorkflowExecution_ABORTING   WorkflowExecution_Phase = 9
)

var WorkflowExecution_Phase_name = map[int32]string{
	0: "UNDEFINED",
	1: "QUEUED",
	2: "RUNNING",
	3: "SUCCEEDING",
	4: "SUCCEEDED",
	5: "FAILING",
	6: "FAILED",
	7: "ABORTED",
	8: "TIMED_OUT",
	9: "ABORTING",
}

var WorkflowExecution_Phase_value = map[string]int32{
	"UNDEFINED":  0,
	"QUEUED":     1,
	"RUNNING":    2,
	"SUCCEEDING": 3,
	"SUCCEEDED":  4,
	"FAILING":    5,
	"FAILED":     6,
	"ABORTED":    7,
	"TIMED_OUT":  8,
	"ABORTING":   9,
}

func (x WorkflowExecution_Phase) String() string {
	return proto.EnumName(WorkflowExecution_Phase_name, int32(x))
}

func (WorkflowExecution_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{0, 0}
}

type NodeExecution_Phase int32

const (
	NodeExecution_UNDEFINED       NodeExecution_Phase = 0
	NodeExecution_QUEUED          NodeExecution_Phase = 1
	NodeExecution_RUNNING         NodeExecution_Phase = 2
	NodeExecution_SUCCEEDED       NodeExecution_Phase = 3
	NodeExecution_FAILING         NodeExecution_Phase = 4
	NodeExecution_FAILED          NodeExecution_Phase = 5
	NodeExecution_ABORTED         NodeExecution_Phase = 6
	NodeExecution_SKIPPED         NodeExecution_Phase = 7
	NodeExecution_TIMED_OUT       NodeExecution_Phase = 8
	NodeExecution_DYNAMIC_RUNNING NodeExecution_Phase = 9
	NodeExecution_RECOVERED       NodeExecution_Phase = 10
)

var NodeExecution_Phase_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "QUEUED",
	2:  "RUNNING",
	3:  "SUCCEEDED",
	4:  "FAILING",
	5:  "FAILED",
	6:  "ABORTED",
	7:  "SKIPPED",
	8:  "TIMED_OUT",
	9:  "DYNAMIC_RUNNING",
	10: "RECOVERED",
}

var NodeExecution_Phase_value = map[string]int32{
	"UNDEFINED":       0,
	"QUEUED":          1,
	"RUNNING":         2,
	"SUCCEEDED":       3,
	"FAILING":         4,
	"FAILED":          5,
	"ABORTED":         6,
	"SKIPPED":         7,
	"TIMED_OUT":       8,
	"DYNAMIC_RUNNING": 9,
	"RECOVERED":       10,
}

func (x NodeExecution_Phase) String() string {
	return proto.EnumName(NodeExecution_Phase_name, int32(x))
}

func (NodeExecution_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{1, 0}
}

type TaskExecution_Phase int32

const (
	TaskExecution_UNDEFINED TaskExecution_Phase = 0
	TaskExecution_QUEUED    TaskExecution_Phase = 1
	TaskExecution_RUNNING   TaskExecution_Phase = 2
	TaskExecution_SUCCEEDED TaskExecution_Phase = 3
	TaskExecution_ABORTED   TaskExecution_Phase = 4
	TaskExecution_FAILED    TaskExecution_Phase = 5
	// To indicate cases where task is initializing, like: ErrImagePull, ContainerCreating, PodInitializing
	TaskExecution_INITIALIZING TaskExecution_Phase = 6
	// To address cases, where underlying resource is not available: Backoff error, Resource quota exceeded
	TaskExecution_WAITING_FOR_RESOURCES TaskExecution_Phase = 7
)

var TaskExecution_Phase_name = map[int32]string{
	0: "UNDEFINED",
	1: "QUEUED",
	2: "RUNNING",
	3: "SUCCEEDED",
	4: "ABORTED",
	5: "FAILED",
	6: "INITIALIZING",
	7: "WAITING_FOR_RESOURCES",
}

var TaskExecution_Phase_value = map[string]int32{
	"UNDEFINED":             0,
	"QUEUED":                1,
	"RUNNING":               2,
	"SUCCEEDED":             3,
	"ABORTED":               4,
	"FAILED":                5,
	"INITIALIZING":          6,
	"WAITING_FOR_RESOURCES": 7,
}

func (x TaskExecution_Phase) String() string {
	return proto.EnumName(TaskExecution_Phase_name, int32(x))
}

func (TaskExecution_Phase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{2, 0}
}

// Error type: System or User
type ExecutionError_ErrorKind int32

const (
	ExecutionError_UNKNOWN ExecutionError_ErrorKind = 0
	ExecutionError_USER    ExecutionError_ErrorKind = 1
	ExecutionError_SYSTEM  ExecutionError_ErrorKind = 2
)

var ExecutionError_ErrorKind_name = map[int32]string{
	0: "UNKNOWN",
	1: "USER",
	2: "SYSTEM",
}

var ExecutionError_ErrorKind_value = map[string]int32{
	"UNKNOWN": 0,
	"USER":    1,
	"SYSTEM":  2,
}

func (x ExecutionError_ErrorKind) String() string {
	return proto.EnumName(ExecutionError_ErrorKind_name, int32(x))
}

func (ExecutionError_ErrorKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{3, 0}
}

type TaskLog_MessageFormat int32

const (
	TaskLog_UNKNOWN TaskLog_MessageFormat = 0
	TaskLog_CSV     TaskLog_MessageFormat = 1
	TaskLog_JSON    TaskLog_MessageFormat = 2
)

var TaskLog_MessageFormat_name = map[int32]string{
	0: "UNKNOWN",
	1: "CSV",
	2: "JSON",
}

var TaskLog_MessageFormat_value = map[string]int32{
	"UNKNOWN": 0,
	"CSV":     1,
	"JSON":    2,
}

func (x TaskLog_MessageFormat) String() string {
	return proto.EnumName(TaskLog_MessageFormat_name, int32(x))
}

func (TaskLog_MessageFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{4, 0}
}

type QualityOfService_Tier int32

const (
	// Default: no quality of service specified.
	QualityOfService_UNDEFINED QualityOfService_Tier = 0
	QualityOfService_HIGH      QualityOfService_Tier = 1
	QualityOfService_MEDIUM    QualityOfService_Tier = 2
	QualityOfService_LOW       QualityOfService_Tier = 3
)

var QualityOfService_Tier_name = map[int32]string{
	0: "UNDEFINED",
	1: "HIGH",
	2: "MEDIUM",
	3: "LOW",
}

var QualityOfService_Tier_value = map[string]int32{
	"UNDEFINED": 0,
	"HIGH":      1,
	"MEDIUM":    2,
	"LOW":       3,
}

func (x QualityOfService_Tier) String() string {
	return proto.EnumName(QualityOfService_Tier_name, int32(x))
}

func (QualityOfService_Tier) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{6, 0}
}

// Indicates various phases of Workflow Execution
type WorkflowExecution struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecution) Reset()         { *m = WorkflowExecution{} }
func (m *WorkflowExecution) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecution) ProtoMessage()    {}
func (*WorkflowExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{0}
}

func (m *WorkflowExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecution.Unmarshal(m, b)
}
func (m *WorkflowExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecution.Marshal(b, m, deterministic)
}
func (m *WorkflowExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecution.Merge(m, src)
}
func (m *WorkflowExecution) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecution.Size(m)
}
func (m *WorkflowExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecution.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecution proto.InternalMessageInfo

// Indicates various phases of Node Execution that only include the time spent to run the nodes/workflows
type NodeExecution struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecution) Reset()         { *m = NodeExecution{} }
func (m *NodeExecution) String() string { return proto.CompactTextString(m) }
func (*NodeExecution) ProtoMessage()    {}
func (*NodeExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{1}
}

func (m *NodeExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecution.Unmarshal(m, b)
}
func (m *NodeExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecution.Marshal(b, m, deterministic)
}
func (m *NodeExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecution.Merge(m, src)
}
func (m *NodeExecution) XXX_Size() int {
	return xxx_messageInfo_NodeExecution.Size(m)
}
func (m *NodeExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecution.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecution proto.InternalMessageInfo

// Phases that task plugins can go through. Not all phases may be applicable to a specific plugin task,
// but this is the cumulative list that customers may want to know about for their task.
type TaskExecution struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecution) Reset()         { *m = TaskExecution{} }
func (m *TaskExecution) String() string { return proto.CompactTextString(m) }
func (*TaskExecution) ProtoMessage()    {}
func (*TaskExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{2}
}

func (m *TaskExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecution.Unmarshal(m, b)
}
func (m *TaskExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecution.Marshal(b, m, deterministic)
}
func (m *TaskExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecution.Merge(m, src)
}
func (m *TaskExecution) XXX_Size() int {
	return xxx_messageInfo_TaskExecution.Size(m)
}
func (m *TaskExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecution.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecution proto.InternalMessageInfo

// Represents the error message from the execution.
type ExecutionError struct {
	// Error code indicates a grouping of a type of error.
	// More Info: <Link>
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// Detailed description of the error - including stack trace.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Full error contents accessible via a URI
	ErrorUri             string                   `protobuf:"bytes,3,opt,name=error_uri,json=errorUri,proto3" json:"error_uri,omitempty"`
	Kind                 ExecutionError_ErrorKind `protobuf:"varint,4,opt,name=kind,proto3,enum=nebulaidl.core.ExecutionError_ErrorKind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExecutionError) Reset()         { *m = ExecutionError{} }
func (m *ExecutionError) String() string { return proto.CompactTextString(m) }
func (*ExecutionError) ProtoMessage()    {}
func (*ExecutionError) Descriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{3}
}

func (m *ExecutionError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionError.Unmarshal(m, b)
}
func (m *ExecutionError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionError.Marshal(b, m, deterministic)
}
func (m *ExecutionError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionError.Merge(m, src)
}
func (m *ExecutionError) XXX_Size() int {
	return xxx_messageInfo_ExecutionError.Size(m)
}
func (m *ExecutionError) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionError.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionError proto.InternalMessageInfo

func (m *ExecutionError) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ExecutionError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ExecutionError) GetErrorUri() string {
	if m != nil {
		return m.ErrorUri
	}
	return ""
}

func (m *ExecutionError) GetKind() ExecutionError_ErrorKind {
	if m != nil {
		return m.Kind
	}
	return ExecutionError_UNKNOWN
}

// Log information for the task that is specific to a log sink
// When our log story is flushed out, we may have more metadata here like log link expiry
type TaskLog struct {
	Uri                  string                `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MessageFormat        TaskLog_MessageFormat `protobuf:"varint,3,opt,name=message_format,json=messageFormat,proto3,enum=nebulaidl.core.TaskLog_MessageFormat" json:"message_format,omitempty"`
	Ttl                  *duration.Duration    `protobuf:"bytes,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TaskLog) Reset()         { *m = TaskLog{} }
func (m *TaskLog) String() string { return proto.CompactTextString(m) }
func (*TaskLog) ProtoMessage()    {}
func (*TaskLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{4}
}

func (m *TaskLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLog.Unmarshal(m, b)
}
func (m *TaskLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLog.Marshal(b, m, deterministic)
}
func (m *TaskLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLog.Merge(m, src)
}
func (m *TaskLog) XXX_Size() int {
	return xxx_messageInfo_TaskLog.Size(m)
}
func (m *TaskLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskLog.DiscardUnknown(m)
}

var xxx_messageInfo_TaskLog proto.InternalMessageInfo

func (m *TaskLog) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *TaskLog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskLog) GetMessageFormat() TaskLog_MessageFormat {
	if m != nil {
		return m.MessageFormat
	}
	return TaskLog_UNKNOWN
}

func (m *TaskLog) GetTtl() *duration.Duration {
	if m != nil {
		return m.Ttl
	}
	return nil
}

// Represents customized execution run-time attributes.
type QualityOfServiceSpec struct {
	// Indicates how much queueing delay an execution can tolerate.
	QueueingBudget       *duration.Duration `protobuf:"bytes,1,opt,name=queueing_budget,json=queueingBudget,proto3" json:"queueing_budget,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *QualityOfServiceSpec) Reset()         { *m = QualityOfServiceSpec{} }
func (m *QualityOfServiceSpec) String() string { return proto.CompactTextString(m) }
func (*QualityOfServiceSpec) ProtoMessage()    {}
func (*QualityOfServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{5}
}

func (m *QualityOfServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QualityOfServiceSpec.Unmarshal(m, b)
}
func (m *QualityOfServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QualityOfServiceSpec.Marshal(b, m, deterministic)
}
func (m *QualityOfServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QualityOfServiceSpec.Merge(m, src)
}
func (m *QualityOfServiceSpec) XXX_Size() int {
	return xxx_messageInfo_QualityOfServiceSpec.Size(m)
}
func (m *QualityOfServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_QualityOfServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_QualityOfServiceSpec proto.InternalMessageInfo

func (m *QualityOfServiceSpec) GetQueueingBudget() *duration.Duration {
	if m != nil {
		return m.QueueingBudget
	}
	return nil
}

// Indicates the priority of an execution.
type QualityOfService struct {
	// Types that are valid to be assigned to Designation:
	//	*QualityOfService_Tier_
	//	*QualityOfService_Spec
	Designation          isQualityOfService_Designation `protobuf_oneof:"designation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *QualityOfService) Reset()         { *m = QualityOfService{} }
func (m *QualityOfService) String() string { return proto.CompactTextString(m) }
func (*QualityOfService) ProtoMessage()    {}
func (*QualityOfService) Descriptor() ([]byte, []int) {
	return fileDescriptor_9397d6d150f54adb, []int{6}
}

func (m *QualityOfService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QualityOfService.Unmarshal(m, b)
}
func (m *QualityOfService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QualityOfService.Marshal(b, m, deterministic)
}
func (m *QualityOfService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QualityOfService.Merge(m, src)
}
func (m *QualityOfService) XXX_Size() int {
	return xxx_messageInfo_QualityOfService.Size(m)
}
func (m *QualityOfService) XXX_DiscardUnknown() {
	xxx_messageInfo_QualityOfService.DiscardUnknown(m)
}

var xxx_messageInfo_QualityOfService proto.InternalMessageInfo

type isQualityOfService_Designation interface {
	isQualityOfService_Designation()
}

type QualityOfService_Tier_ struct {
	Tier QualityOfService_Tier `protobuf:"varint,1,opt,name=tier,proto3,enum=nebulaidl.core.QualityOfService_Tier,oneof"`
}

type QualityOfService_Spec struct {
	Spec *QualityOfServiceSpec `protobuf:"bytes,2,opt,name=spec,proto3,oneof"`
}

func (*QualityOfService_Tier_) isQualityOfService_Designation() {}

func (*QualityOfService_Spec) isQualityOfService_Designation() {}

func (m *QualityOfService) GetDesignation() isQualityOfService_Designation {
	if m != nil {
		return m.Designation
	}
	return nil
}

func (m *QualityOfService) GetTier() QualityOfService_Tier {
	if x, ok := m.GetDesignation().(*QualityOfService_Tier_); ok {
		return x.Tier
	}
	return QualityOfService_UNDEFINED
}

func (m *QualityOfService) GetSpec() *QualityOfServiceSpec {
	if x, ok := m.GetDesignation().(*QualityOfService_Spec); ok {
		return x.Spec
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QualityOfService) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QualityOfService_Tier_)(nil),
		(*QualityOfService_Spec)(nil),
	}
}

func init() {
	proto.RegisterEnum("nebulaidl.core.WorkflowExecution_Phase", WorkflowExecution_Phase_name, WorkflowExecution_Phase_value)
	proto.RegisterEnum("nebulaidl.core.NodeExecution_Phase", NodeExecution_Phase_name, NodeExecution_Phase_value)
	proto.RegisterEnum("nebulaidl.core.TaskExecution_Phase", TaskExecution_Phase_name, TaskExecution_Phase_value)
	proto.RegisterEnum("nebulaidl.core.ExecutionError_ErrorKind", ExecutionError_ErrorKind_name, ExecutionError_ErrorKind_value)
	proto.RegisterEnum("nebulaidl.core.TaskLog_MessageFormat", TaskLog_MessageFormat_name, TaskLog_MessageFormat_value)
	proto.RegisterEnum("nebulaidl.core.QualityOfService_Tier", QualityOfService_Tier_name, QualityOfService_Tier_value)
	proto.RegisterType((*WorkflowExecution)(nil), "nebulaidl.core.WorkflowExecution")
	proto.RegisterType((*NodeExecution)(nil), "nebulaidl.core.NodeExecution")
	proto.RegisterType((*TaskExecution)(nil), "nebulaidl.core.TaskExecution")
	proto.RegisterType((*ExecutionError)(nil), "nebulaidl.core.ExecutionError")
	proto.RegisterType((*TaskLog)(nil), "nebulaidl.core.TaskLog")
	proto.RegisterType((*QualityOfServiceSpec)(nil), "nebulaidl.core.QualityOfServiceSpec")
	proto.RegisterType((*QualityOfService)(nil), "nebulaidl.core.QualityOfService")
}

func init() { proto.RegisterFile("nebulaidl/core/execution.proto", fileDescriptor_9397d6d150f54adb) }

var fileDescriptor_9397d6d150f54adb = []byte{
	// 733 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xc7, 0x63, 0x70, 0xf8, 0x38, 0x04, 0x32, 0x9d, 0xb6, 0x12, 0x69, 0xa5, 0x28, 0xb2, 0x5a,
	0x29, 0x52, 0x55, 0x5b, 0xa2, 0xbd, 0x4a, 0x73, 0x03, 0x78, 0x48, 0xdc, 0x80, 0x9d, 0xd8, 0x38,
	0x28, 0xb9, 0x41, 0xc6, 0x1e, 0x1c, 0x2b, 0xe0, 0xa1, 0xc6, 0x6e, 0xbb, 0xf7, 0xfb, 0x02, 0x7b,
	0xb3, 0x4f, 0xb0, 0xd2, 0xbe, 0xc1, 0xbe, 0xc3, 0xde, 0xed, 0xfd, 0xbe, 0xcc, 0x6a, 0x06, 0x08,
	0x1f, 0x8a, 0x76, 0xb5, 0xd2, 0xde, 0xa0, 0x39, 0x73, 0xce, 0xff, 0xcc, 0xef, 0x7f, 0xf0, 0x0c,
	0x1c, 0xc7, 0x74, 0x94, 0x4d, 0xbc, 0x28, 0x98, 0x68, 0x3e, 0x4b, 0xa8, 0x46, 0xff, 0xa7, 0x7e,
	0x96, 0x46, 0x2c, 0x56, 0x67, 0x09, 0x4b, 0x19, 0xae, 0x3d, 0xe5, 0x55, 0x9e, 0xff, 0xe9, 0x38,
	0x64, 0x2c, 0x9c, 0x50, 0x4d, 0x64, 0x47, 0xd9, 0x58, 0x0b, 0xb2, 0xc4, 0x5b, 0xd7, 0x2b, 0x6f,
	0x25, 0xf8, 0x6e, 0xc0, 0x92, 0xc7, 0xf1, 0x84, 0xfd, 0x47, 0x56, 0xbd, 0x94, 0x57, 0x12, 0xec,
	0x5f, 0x3f, 0x78, 0x73, 0x8a, 0xab, 0x50, 0x76, 0x4d, 0x9d, 0x74, 0x0c, 0x93, 0xe8, 0x68, 0x0f,
	0x03, 0x14, 0x6e, 0x5c, 0xe2, 0x12, 0x1d, 0x49, 0xb8, 0x02, 0x45, 0xdb, 0x35, 0x4d, 0xc3, 0xbc,
	0x40, 0x39, 0x5c, 0x03, 0x70, 0xdc, 0x76, 0x9b, 0x10, 0x9d, 0xc7, 0x79, 0xae, 0x5b, 0xc6, 0x44,
	0x47, 0x32, 0xaf, 0xed, 0x34, 0x8d, 0x2e, 0xcf, 0xed, 0xf3, 0x26, 0x3c, 0x20, 0x3a, 0x2a, 0xf0,
	0x44, 0xb3, 0x65, 0xd9, 0x7d, 0xa2, 0xa3, 0x22, 0x17, 0xf5, 0x8d, 0x1e, 0xd1, 0x87, 0x96, 0xdb,
	0x47, 0x25, 0x7c, 0x00, 0x25, 0x91, 0xe3, 0xaa, 0xb2, 0xf2, 0x4e, 0x82, 0xaa, 0xc9, 0x02, 0xba,
	0xa6, 0x7c, 0xf3, 0xd5, 0x94, 0x5b, 0x54, 0xf9, 0x4d, 0x2a, 0x79, 0x83, 0x6a, 0x7f, 0x93, 0x4a,
	0x20, 0x3a, 0x57, 0xc6, 0xf5, 0xf5, 0x73, 0x88, 0xdf, 0xc3, 0xa1, 0x7e, 0x67, 0x36, 0x7b, 0x46,
	0x7b, 0xb8, 0x3a, 0xa5, 0xcc, 0x6b, 0x6c, 0xd2, 0xb6, 0x6e, 0x89, 0x4d, 0x74, 0x04, 0xca, 0x6b,
	0x09, 0xaa, 0x7d, 0x6f, 0xfe, 0xb8, 0x06, 0x7f, 0xf9, 0x0d, 0xc0, 0x57, 0x7c, 0xdb, 0xe0, 0x08,
	0x0e, 0x0c, 0xd3, 0xe8, 0x1b, 0xcd, 0xae, 0x71, 0xcf, 0x95, 0x05, 0x7c, 0x04, 0x3f, 0x0e, 0x9a,
	0x06, 0x9f, 0xe1, 0xb0, 0x63, 0xd9, 0x43, 0x9b, 0x38, 0x96, 0x6b, 0xb7, 0x89, 0x83, 0x8a, 0xca,
	0x7b, 0x09, 0x6a, 0x4f, 0x50, 0x24, 0x49, 0x58, 0x82, 0x31, 0xc8, 0x3e, 0x0b, 0x68, 0x5d, 0x3a,
	0x91, 0x4e, 0xcb, 0xb6, 0x58, 0xe3, 0x3a, 0x14, 0xa7, 0x74, 0x3e, 0xf7, 0x42, 0x5a, 0xcf, 0x89,
	0xed, 0x55, 0x88, 0x7f, 0x86, 0x32, 0xe5, 0xb2, 0x61, 0x96, 0x44, 0xf5, 0xbc, 0xc8, 0x95, 0xc4,
	0x86, 0x9b, 0x44, 0xf8, 0x1c, 0xe4, 0xc7, 0x28, 0x0e, 0xea, 0xf2, 0x89, 0x74, 0x5a, 0x6b, 0x9c,
	0xaa, 0xdb, 0x1f, 0xa6, 0xba, 0x7d, 0xb0, 0x2a, 0x7e, 0xaf, 0xa2, 0x38, 0xb0, 0x85, 0x4a, 0x51,
	0xa1, 0xfc, 0xb4, 0xc5, 0xed, 0xba, 0xe6, 0x95, 0x69, 0x0d, 0x4c, 0xb4, 0x87, 0x4b, 0x20, 0xbb,
	0x0e, 0xb1, 0x91, 0xc4, 0x8d, 0x3b, 0x77, 0x4e, 0x9f, 0xf4, 0x50, 0x4e, 0xf9, 0x28, 0x41, 0x91,
	0x0f, 0xb9, 0xcb, 0x42, 0x8c, 0x20, 0xcf, 0x81, 0x16, 0x1e, 0xf8, 0x92, 0xdb, 0x8a, 0xbd, 0xe9,
	0x8a, 0x5f, 0xac, 0x71, 0x17, 0x6a, 0x4b, 0x1f, 0xc3, 0x31, 0x4b, 0xa6, 0x5e, 0x2a, 0x1c, 0xd4,
	0x1a, 0xbf, 0xee, 0x92, 0x2e, 0xdb, 0xaa, 0xbd, 0x45, 0x75, 0x47, 0x14, 0xdb, 0xd5, 0xe9, 0x66,
	0x88, 0x7f, 0x83, 0x7c, 0x9a, 0x4e, 0x84, 0xd9, 0x4a, 0xe3, 0x48, 0x5d, 0xdc, 0x3a, 0x75, 0x75,
	0xeb, 0x54, 0x7d, 0x79, 0xeb, 0x6c, 0x5e, 0xa5, 0x68, 0x50, 0xdd, 0x6a, 0xb6, 0x6d, 0xb0, 0x08,
	0xf9, 0xb6, 0x73, 0x8b, 0x24, 0xee, 0xf4, 0x6f, 0xc7, 0x32, 0x51, 0x4e, 0xb9, 0x87, 0x1f, 0x6e,
	0x32, 0x6f, 0x12, 0xa5, 0x2f, 0xac, 0xb1, 0x43, 0x93, 0x7f, 0x23, 0x9f, 0x3a, 0x33, 0xea, 0xe3,
	0x16, 0x1c, 0xfe, 0x93, 0xd1, 0x8c, 0x46, 0x71, 0x38, 0x1c, 0x65, 0x41, 0x48, 0x53, 0xe1, 0xfa,
	0xb3, 0x04, 0xb5, 0x95, 0xa2, 0x25, 0x04, 0xca, 0x07, 0x09, 0xd0, 0x6e, 0x73, 0xfc, 0x17, 0xc8,
	0x69, 0x44, 0x13, 0xd1, 0xed, 0x99, 0x91, 0xec, 0xd6, 0xab, 0xfd, 0x88, 0x26, 0x97, 0x7b, 0xb6,
	0x10, 0xe1, 0x33, 0x90, 0xe7, 0x33, 0xea, 0x8b, 0x69, 0x57, 0x1a, 0xbf, 0x7c, 0x49, 0xcc, 0x9d,
	0x70, 0x2d, 0xd7, 0x28, 0x7f, 0x82, 0xcc, 0x7b, 0xed, 0x5e, 0x8c, 0x12, 0xc8, 0x97, 0xc6, 0xc5,
	0xe5, 0xe2, 0x4f, 0xef, 0x11, 0xdd, 0x70, 0x7b, 0x28, 0xc7, 0x27, 0xd5, 0xb5, 0x06, 0x28, 0xdf,
	0xaa, 0x42, 0x25, 0xa0, 0xf3, 0x28, 0x8c, 0x85, 0xc5, 0xd6, 0xf9, 0xfd, 0x59, 0x18, 0xa5, 0x0f,
	0xd9, 0x48, 0xf5, 0xd9, 0x54, 0x5b, 0x1c, 0xef, 0x4f, 0x58, 0x16, 0xcc, 0xb5, 0xf5, 0xf3, 0x19,
	0xd2, 0x58, 0x9b, 0x8d, 0x7e, 0x0f, 0x99, 0xb6, 0xfd, 0xa4, 0x8e, 0x0a, 0x62, 0x66, 0x7f, 0x7c,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x4c, 0x41, 0x92, 0x6b, 0x05, 0x00, 0x00,
}
