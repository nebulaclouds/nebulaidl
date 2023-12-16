// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/core/identifier.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Indicates a resource type within Nebula.
type ResourceType int32

const (
	ResourceType_UNSPECIFIED ResourceType = 0
	ResourceType_TASK        ResourceType = 1
	ResourceType_WORKFLOW    ResourceType = 2
	ResourceType_LAUNCH_PLAN ResourceType = 3
	// A dataset represents an entity modeled in Nebula DataCatalog. A Dataset is also a versioned entity and can be a compilation of multiple individual objects.
	// Eventually all Catalog objects should be modeled similar to Nebula Objects. The Dataset entities makes it possible for the UI  and CLI to act on the objects
	// in a similar manner to other Nebula objects
	ResourceType_DATASET ResourceType = 4
)

var ResourceType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "TASK",
	2: "WORKFLOW",
	3: "LAUNCH_PLAN",
	4: "DATASET",
}

var ResourceType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"TASK":        1,
	"WORKFLOW":    2,
	"LAUNCH_PLAN": 3,
	"DATASET":     4,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{0}
}

// Encapsulation of fields that uniquely identifies a Nebula resource.
type Identifier struct {
	// Identifies the specific type of resource that this identifier corresponds to.
	ResourceType ResourceType `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3,enum=nebulaidl.core.ResourceType" json:"resource_type,omitempty"`
	// Name of the project the resource belongs to.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the domain the resource belongs to.
	// A domain can be considered as a subset within a specific project.
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	// User provided value for the resource.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Specific version of the resource.
	Version              string   `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identifier) Reset()         { *m = Identifier{} }
func (m *Identifier) String() string { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()    {}
func (*Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{0}
}

func (m *Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identifier.Unmarshal(m, b)
}
func (m *Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identifier.Marshal(b, m, deterministic)
}
func (m *Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identifier.Merge(m, src)
}
func (m *Identifier) XXX_Size() int {
	return xxx_messageInfo_Identifier.Size(m)
}
func (m *Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_Identifier proto.InternalMessageInfo

func (m *Identifier) GetResourceType() ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return ResourceType_UNSPECIFIED
}

func (m *Identifier) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Identifier) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Identifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Identifier) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Encapsulation of fields that uniquely identifies a Nebula workflow execution
type WorkflowExecutionIdentifier struct {
	// Name of the project the resource belongs to.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the domain the resource belongs to.
	// A domain can be considered as a subset within a specific project.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// User or system provided value for the resource.
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecutionIdentifier) Reset()         { *m = WorkflowExecutionIdentifier{} }
func (m *WorkflowExecutionIdentifier) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionIdentifier) ProtoMessage()    {}
func (*WorkflowExecutionIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{1}
}

func (m *WorkflowExecutionIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionIdentifier.Unmarshal(m, b)
}
func (m *WorkflowExecutionIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionIdentifier.Marshal(b, m, deterministic)
}
func (m *WorkflowExecutionIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionIdentifier.Merge(m, src)
}
func (m *WorkflowExecutionIdentifier) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionIdentifier.Size(m)
}
func (m *WorkflowExecutionIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionIdentifier proto.InternalMessageInfo

func (m *WorkflowExecutionIdentifier) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *WorkflowExecutionIdentifier) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *WorkflowExecutionIdentifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Encapsulation of fields that identify a Nebula node execution entity.
type NodeExecutionIdentifier struct {
	NodeId               string                       `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ExecutionId          *WorkflowExecutionIdentifier `protobuf:"bytes,2,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *NodeExecutionIdentifier) Reset()         { *m = NodeExecutionIdentifier{} }
func (m *NodeExecutionIdentifier) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionIdentifier) ProtoMessage()    {}
func (*NodeExecutionIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{2}
}

func (m *NodeExecutionIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionIdentifier.Unmarshal(m, b)
}
func (m *NodeExecutionIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionIdentifier.Marshal(b, m, deterministic)
}
func (m *NodeExecutionIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionIdentifier.Merge(m, src)
}
func (m *NodeExecutionIdentifier) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionIdentifier.Size(m)
}
func (m *NodeExecutionIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionIdentifier proto.InternalMessageInfo

func (m *NodeExecutionIdentifier) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *NodeExecutionIdentifier) GetExecutionId() *WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

// Encapsulation of fields that identify a Nebula task execution entity.
type TaskExecutionIdentifier struct {
	TaskId               *Identifier              `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	NodeExecutionId      *NodeExecutionIdentifier `protobuf:"bytes,2,opt,name=node_execution_id,json=nodeExecutionId,proto3" json:"node_execution_id,omitempty"`
	RetryAttempt         uint32                   `protobuf:"varint,3,opt,name=retry_attempt,json=retryAttempt,proto3" json:"retry_attempt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TaskExecutionIdentifier) Reset()         { *m = TaskExecutionIdentifier{} }
func (m *TaskExecutionIdentifier) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionIdentifier) ProtoMessage()    {}
func (*TaskExecutionIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{3}
}

func (m *TaskExecutionIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionIdentifier.Unmarshal(m, b)
}
func (m *TaskExecutionIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionIdentifier.Marshal(b, m, deterministic)
}
func (m *TaskExecutionIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionIdentifier.Merge(m, src)
}
func (m *TaskExecutionIdentifier) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionIdentifier.Size(m)
}
func (m *TaskExecutionIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionIdentifier proto.InternalMessageInfo

func (m *TaskExecutionIdentifier) GetTaskId() *Identifier {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *TaskExecutionIdentifier) GetNodeExecutionId() *NodeExecutionIdentifier {
	if m != nil {
		return m.NodeExecutionId
	}
	return nil
}

func (m *TaskExecutionIdentifier) GetRetryAttempt() uint32 {
	if m != nil {
		return m.RetryAttempt
	}
	return 0
}

// Encapsulation of fields the uniquely identify a signal.
type SignalIdentifier struct {
	// Unique identifier for a signal.
	SignalId string `protobuf:"bytes,1,opt,name=signal_id,json=signalId,proto3" json:"signal_id,omitempty"`
	// Identifies the Nebula workflow execution this signal belongs to.
	ExecutionId          *WorkflowExecutionIdentifier `protobuf:"bytes,2,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SignalIdentifier) Reset()         { *m = SignalIdentifier{} }
func (m *SignalIdentifier) String() string { return proto.CompactTextString(m) }
func (*SignalIdentifier) ProtoMessage()    {}
func (*SignalIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_adfa846a86e1fa0c, []int{4}
}

func (m *SignalIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalIdentifier.Unmarshal(m, b)
}
func (m *SignalIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalIdentifier.Marshal(b, m, deterministic)
}
func (m *SignalIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalIdentifier.Merge(m, src)
}
func (m *SignalIdentifier) XXX_Size() int {
	return xxx_messageInfo_SignalIdentifier.Size(m)
}
func (m *SignalIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_SignalIdentifier proto.InternalMessageInfo

func (m *SignalIdentifier) GetSignalId() string {
	if m != nil {
		return m.SignalId
	}
	return ""
}

func (m *SignalIdentifier) GetExecutionId() *WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

func init() {
	proto.RegisterEnum("nebulaidl.core.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterType((*Identifier)(nil), "nebulaidl.core.Identifier")
	proto.RegisterType((*WorkflowExecutionIdentifier)(nil), "nebulaidl.core.WorkflowExecutionIdentifier")
	proto.RegisterType((*NodeExecutionIdentifier)(nil), "nebulaidl.core.NodeExecutionIdentifier")
	proto.RegisterType((*TaskExecutionIdentifier)(nil), "nebulaidl.core.TaskExecutionIdentifier")
	proto.RegisterType((*SignalIdentifier)(nil), "nebulaidl.core.SignalIdentifier")
}

func init() { proto.RegisterFile("nebulaidl/core/identifier.proto", fileDescriptor_adfa846a86e1fa0c) }

var fileDescriptor_adfa846a86e1fa0c = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x69, 0x48, 0xd2, 0x49, 0x42, 0xcd, 0x1e, 0x88, 0x51, 0x24, 0x54, 0x19, 0x09, 0x55,
	0x95, 0xb0, 0xa5, 0x80, 0x38, 0x63, 0xda, 0x54, 0x58, 0x0d, 0x69, 0xe5, 0x38, 0x8a, 0xc4, 0x25,
	0x72, 0xec, 0x89, 0x59, 0x92, 0xec, 0x5a, 0xeb, 0x0d, 0xe0, 0x0b, 0x12, 0x3f, 0xc4, 0x67, 0xf0,
	0x5d, 0x68, 0x5d, 0xa7, 0x75, 0xa2, 0x14, 0x2e, 0xbd, 0x79, 0xe6, 0xbd, 0x9d, 0xf7, 0x66, 0x3c,
	0x03, 0x2f, 0xe6, 0xcb, 0x4c, 0x22, 0x8d, 0x96, 0x76, 0xc8, 0x05, 0xda, 0x34, 0x42, 0x26, 0xe9,
	0x9c, 0xa2, 0xb0, 0x12, 0xc1, 0x25, 0x27, 0xed, 0x0d, 0x6e, 0x29, 0xdc, 0xfc, 0xad, 0x01, 0xb8,
	0xb7, 0x1c, 0xf2, 0x1e, 0xda, 0x02, 0x53, 0xbe, 0x16, 0x21, 0x4e, 0x65, 0x96, 0xa0, 0xa1, 0x1d,
	0x6b, 0x27, 0x4f, 0x7a, 0x5d, 0x6b, 0xeb, 0x95, 0xe5, 0x15, 0x1c, 0x3f, 0x4b, 0xd0, 0x6b, 0x89,
	0x52, 0x44, 0x0c, 0xa8, 0x27, 0x82, 0x7f, 0xc5, 0x50, 0x1a, 0x95, 0x63, 0xed, 0xe4, 0xd0, 0xdb,
	0x84, 0xe4, 0x19, 0xd4, 0x22, 0xbe, 0x0a, 0x28, 0x33, 0x0e, 0x72, 0xa0, 0x88, 0x08, 0x81, 0x2a,
	0x0b, 0x56, 0x68, 0x54, 0xf3, 0x6c, 0xfe, 0xad, 0xaa, 0x7c, 0x43, 0x91, 0x52, 0xce, 0x8c, 0xc7,
	0x37, 0x55, 0x8a, 0xd0, 0x0c, 0xa1, 0x3b, 0xe1, 0x62, 0x31, 0x5f, 0xf2, 0xef, 0xfd, 0x1f, 0x18,
	0xae, 0x25, 0xe5, 0xac, 0xd4, 0x40, 0x49, 0x5e, 0xbb, 0x4f, 0xbe, 0xf2, 0x3f, 0x79, 0xf3, 0x97,
	0x06, 0x9d, 0x21, 0x8f, 0x70, 0x9f, 0x42, 0x07, 0xea, 0x8c, 0x47, 0x38, 0xa5, 0x51, 0xa1, 0x50,
	0x53, 0xa1, 0x1b, 0x91, 0x4f, 0xd0, 0xc2, 0x0d, 0x5f, 0xa1, 0x4a, 0xa6, 0xd9, 0x3b, 0xdd, 0x19,
	0xdd, 0x3f, 0xcc, 0x7b, 0x4d, 0xbc, 0x4b, 0x9a, 0x7f, 0x34, 0xe8, 0xf8, 0x41, 0xba, 0xd8, 0xe7,
	0xa1, 0x07, 0x75, 0x19, 0xa4, 0x8b, 0x8d, 0x87, 0x66, 0xef, 0xf9, 0x8e, 0x4a, 0xa9, 0x68, 0x4d,
	0x31, 0xdd, 0x88, 0x78, 0xf0, 0x34, 0xf7, 0xbd, 0xc7, 0xe3, 0xab, 0x9d, 0xd7, 0xf7, 0xb4, 0xee,
	0x1d, 0xb1, 0x6d, 0x80, 0xbc, 0x54, 0xeb, 0x22, 0x45, 0x36, 0x0d, 0xa4, 0xc4, 0x55, 0x22, 0xf3,
	0x3f, 0xdb, 0x56, 0x1b, 0x21, 0x45, 0xe6, 0xdc, 0xe4, 0xcc, 0x9f, 0xa0, 0x8f, 0x68, 0xcc, 0x82,
	0x65, 0xa9, 0x81, 0x2e, 0x1c, 0xa6, 0x79, 0xee, 0x6e, 0x8c, 0x8d, 0xb4, 0x20, 0x3d, 0xf0, 0x20,
	0x4f, 0xc7, 0xd0, 0x2a, 0xef, 0x2b, 0x39, 0x82, 0xe6, 0x78, 0x38, 0xba, 0xee, 0x9f, 0xb9, 0x17,
	0x6e, 0xff, 0x5c, 0x7f, 0x44, 0x1a, 0x50, 0xf5, 0x9d, 0xd1, 0xa5, 0xae, 0x91, 0x16, 0x34, 0x26,
	0x57, 0xde, 0xe5, 0xc5, 0xe0, 0x6a, 0xa2, 0x57, 0x14, 0x71, 0xe0, 0x8c, 0x87, 0x67, 0x1f, 0xa7,
	0xd7, 0x03, 0x67, 0xa8, 0x1f, 0x90, 0x26, 0xd4, 0xcf, 0x1d, 0xdf, 0x19, 0xf5, 0x7d, 0xbd, 0xfa,
	0xe1, 0xdd, 0xe7, 0xb7, 0x31, 0x95, 0x5f, 0xd6, 0x33, 0x2b, 0xe4, 0x2b, 0x3b, 0xf7, 0xc6, 0x45,
	0x6c, 0xdf, 0x9e, 0x5f, 0x8c, 0xcc, 0x4e, 0x66, 0xaf, 0x63, 0x6e, 0x6f, 0x5d, 0xe4, 0xac, 0x96,
	0xdf, 0xe1, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x5c, 0x4c, 0x03, 0xa9, 0x03, 0x00,
	0x00,
}
