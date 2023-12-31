// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/plugins/kubeflow/common.proto

package plugins

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

type RestartPolicy int32

const (
	RestartPolicy_RESTART_POLICY_NEVER      RestartPolicy = 0
	RestartPolicy_RESTART_POLICY_ON_FAILURE RestartPolicy = 1
	RestartPolicy_RESTART_POLICY_ALWAYS     RestartPolicy = 2
)

var RestartPolicy_name = map[int32]string{
	0: "RESTART_POLICY_NEVER",
	1: "RESTART_POLICY_ON_FAILURE",
	2: "RESTART_POLICY_ALWAYS",
}

var RestartPolicy_value = map[string]int32{
	"RESTART_POLICY_NEVER":      0,
	"RESTART_POLICY_ON_FAILURE": 1,
	"RESTART_POLICY_ALWAYS":     2,
}

func (x RestartPolicy) String() string {
	return proto.EnumName(RestartPolicy_name, int32(x))
}

func (RestartPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f5a1b1dc5fbe87b, []int{0}
}

type CleanPodPolicy int32

const (
	CleanPodPolicy_CLEANPOD_POLICY_NONE    CleanPodPolicy = 0
	CleanPodPolicy_CLEANPOD_POLICY_RUNNING CleanPodPolicy = 1
	CleanPodPolicy_CLEANPOD_POLICY_ALL     CleanPodPolicy = 2
)

var CleanPodPolicy_name = map[int32]string{
	0: "CLEANPOD_POLICY_NONE",
	1: "CLEANPOD_POLICY_RUNNING",
	2: "CLEANPOD_POLICY_ALL",
}

var CleanPodPolicy_value = map[string]int32{
	"CLEANPOD_POLICY_NONE":    0,
	"CLEANPOD_POLICY_RUNNING": 1,
	"CLEANPOD_POLICY_ALL":     2,
}

func (x CleanPodPolicy) String() string {
	return proto.EnumName(CleanPodPolicy_name, int32(x))
}

func (CleanPodPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f5a1b1dc5fbe87b, []int{1}
}

type RunPolicy struct {
	// Defines the policy to kill pods after the job completes. Default to None.
	CleanPodPolicy CleanPodPolicy `protobuf:"varint,1,opt,name=clean_pod_policy,json=cleanPodPolicy,proto3,enum=nebulaidl.plugins.kubeflow.CleanPodPolicy" json:"clean_pod_policy,omitempty"`
	// TTL to clean up jobs. Default to infinite.
	TtlSecondsAfterFinished int32 `protobuf:"varint,2,opt,name=ttl_seconds_after_finished,json=ttlSecondsAfterFinished,proto3" json:"ttl_seconds_after_finished,omitempty"`
	// Specifies the duration in seconds relative to the startTime that the job may be active
	// before the system tries to terminate it; value must be positive integer.
	ActiveDeadlineSeconds int32 `protobuf:"varint,3,opt,name=active_deadline_seconds,json=activeDeadlineSeconds,proto3" json:"active_deadline_seconds,omitempty"`
	// Number of retries before marking this job failed.
	BackoffLimit         int32    `protobuf:"varint,4,opt,name=backoff_limit,json=backoffLimit,proto3" json:"backoff_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunPolicy) Reset()         { *m = RunPolicy{} }
func (m *RunPolicy) String() string { return proto.CompactTextString(m) }
func (*RunPolicy) ProtoMessage()    {}
func (*RunPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5a1b1dc5fbe87b, []int{0}
}

func (m *RunPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunPolicy.Unmarshal(m, b)
}
func (m *RunPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunPolicy.Marshal(b, m, deterministic)
}
func (m *RunPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunPolicy.Merge(m, src)
}
func (m *RunPolicy) XXX_Size() int {
	return xxx_messageInfo_RunPolicy.Size(m)
}
func (m *RunPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RunPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RunPolicy proto.InternalMessageInfo

func (m *RunPolicy) GetCleanPodPolicy() CleanPodPolicy {
	if m != nil {
		return m.CleanPodPolicy
	}
	return CleanPodPolicy_CLEANPOD_POLICY_NONE
}

func (m *RunPolicy) GetTtlSecondsAfterFinished() int32 {
	if m != nil {
		return m.TtlSecondsAfterFinished
	}
	return 0
}

func (m *RunPolicy) GetActiveDeadlineSeconds() int32 {
	if m != nil {
		return m.ActiveDeadlineSeconds
	}
	return 0
}

func (m *RunPolicy) GetBackoffLimit() int32 {
	if m != nil {
		return m.BackoffLimit
	}
	return 0
}

func init() {
	proto.RegisterEnum("nebulaidl.plugins.kubeflow.RestartPolicy", RestartPolicy_name, RestartPolicy_value)
	proto.RegisterEnum("nebulaidl.plugins.kubeflow.CleanPodPolicy", CleanPodPolicy_name, CleanPodPolicy_value)
	proto.RegisterType((*RunPolicy)(nil), "nebulaidl.plugins.kubeflow.RunPolicy")
}

func init() {
	proto.RegisterFile("nebulaidl/plugins/kubeflow/common.proto", fileDescriptor_2f5a1b1dc5fbe87b)
}

var fileDescriptor_2f5a1b1dc5fbe87b = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdf, 0xab, 0xd3, 0x30,
	0x1c, 0xc5, 0x6f, 0xe7, 0x0f, 0x30, 0x78, 0x47, 0x89, 0x5e, 0xda, 0x7b, 0x45, 0x18, 0xfa, 0xe0,
	0x18, 0xd8, 0x82, 0x82, 0x2f, 0x22, 0x52, 0xb7, 0x4e, 0x06, 0xa5, 0x2d, 0xd9, 0xa6, 0xcc, 0x97,
	0xd8, 0x26, 0x69, 0x17, 0x96, 0x26, 0x65, 0x4d, 0x15, 0xff, 0x76, 0x5f, 0xa4, 0x3f, 0x36, 0x59,
	0xc5, 0xc7, 0x9c, 0xcf, 0xf9, 0x9e, 0x6f, 0x48, 0x0e, 0x78, 0x25, 0x59, 0x5a, 0x8b, 0x84, 0x53,
	0xe1, 0x96, 0xa2, 0xce, 0xb9, 0xac, 0xdc, 0x43, 0x9d, 0xb2, 0x4c, 0xa8, 0x9f, 0x2e, 0x51, 0x45,
	0xa1, 0xa4, 0x53, 0x1e, 0x95, 0x56, 0xf0, 0xee, 0x6c, 0x74, 0x7a, 0xa3, 0x73, 0x32, 0xbe, 0xf8,
	0x6d, 0x80, 0x47, 0xa8, 0x96, 0xb1, 0x12, 0x9c, 0xfc, 0x82, 0x1b, 0x60, 0x12, 0xc1, 0x12, 0x89,
	0x4b, 0x45, 0x71, 0xd9, 0x6a, 0xb6, 0x31, 0x31, 0xa6, 0xe3, 0x37, 0x33, 0xe7, 0xff, 0x21, 0xce,
	0xbc, 0x99, 0x89, 0x15, 0xed, 0x52, 0xd0, 0x98, 0x5c, 0x9c, 0xe1, 0x7b, 0x70, 0xa7, 0xb5, 0xc0,
	0x15, 0x23, 0x4a, 0xd2, 0x0a, 0x27, 0x99, 0x66, 0x47, 0x9c, 0x71, 0xc9, 0xab, 0x3d, 0xa3, 0xf6,
	0x68, 0x62, 0x4c, 0x1f, 0x20, 0x4b, 0x6b, 0xb1, 0xee, 0x0c, 0x5e, 0xc3, 0x97, 0x3d, 0x86, 0xef,
	0x80, 0x95, 0x10, 0xcd, 0x7f, 0x30, 0x4c, 0x59, 0x42, 0x05, 0x97, 0xec, 0x14, 0x64, 0xdf, 0x6b,
	0x27, 0x6f, 0x3a, 0xbc, 0xe8, 0x69, 0x1f, 0x02, 0x5f, 0x82, 0xeb, 0x34, 0x21, 0x07, 0x95, 0x65,
	0x58, 0xf0, 0x82, 0x6b, 0xfb, 0x7e, 0xeb, 0x7e, 0xdc, 0x8b, 0x41, 0xa3, 0xcd, 0x08, 0xb8, 0x46,
	0xac, 0xd2, 0xc9, 0x51, 0xf7, 0x57, 0xb5, 0xc1, 0x53, 0xe4, 0xaf, 0x37, 0x1e, 0xda, 0xe0, 0x38,
	0x0a, 0x56, 0xf3, 0x1d, 0x0e, 0xfd, 0x2f, 0x3e, 0x32, 0xaf, 0xe0, 0x73, 0x70, 0x3b, 0x20, 0x51,
	0x88, 0x97, 0xde, 0x2a, 0xd8, 0x22, 0xdf, 0x34, 0xe0, 0x2d, 0xb8, 0x19, 0x60, 0x2f, 0xf8, 0xea,
	0xed, 0xd6, 0xe6, 0x68, 0xf6, 0x1d, 0x8c, 0x2f, 0x1f, 0xa8, 0xd9, 0x32, 0x0f, 0x7c, 0x2f, 0x8c,
	0xa3, 0xc5, 0x79, 0x4d, 0x14, 0xfa, 0xe6, 0x15, 0x7c, 0x06, 0xac, 0x21, 0x41, 0xdb, 0x30, 0x5c,
	0x85, 0x9f, 0x4d, 0x03, 0x5a, 0xe0, 0xc9, 0x10, 0x7a, 0x41, 0x60, 0x8e, 0x3e, 0x7d, 0xfc, 0xf6,
	0x21, 0xe7, 0x7a, 0x5f, 0xa7, 0x0e, 0x51, 0x85, 0xdb, 0x7d, 0x14, 0x11, 0xaa, 0xa6, 0x95, 0xfb,
	0xb7, 0x23, 0x39, 0x93, 0x6e, 0x99, 0xbe, 0xce, 0x95, 0xfb, 0x4f, 0x6f, 0xd2, 0x87, 0x6d, 0x51,
	0xde, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x81, 0x76, 0x3e, 0x1d, 0x53, 0x02, 0x00, 0x00,
}
