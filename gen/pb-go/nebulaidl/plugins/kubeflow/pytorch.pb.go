// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/plugins/kubeflow/pytorch.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	core "github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core"
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

// Custom proto for torch elastic config for distributed training using
// https://github.com/kubeflow/training-operator/blob/master/pkg/apis/kubeflow.org/v1/pytorch_types.go
type ElasticConfig struct {
	RdzvBackend          string   `protobuf:"bytes,1,opt,name=rdzv_backend,json=rdzvBackend,proto3" json:"rdzv_backend,omitempty"`
	MinReplicas          int32    `protobuf:"varint,2,opt,name=min_replicas,json=minReplicas,proto3" json:"min_replicas,omitempty"`
	MaxReplicas          int32    `protobuf:"varint,3,opt,name=max_replicas,json=maxReplicas,proto3" json:"max_replicas,omitempty"`
	NprocPerNode         int32    `protobuf:"varint,4,opt,name=nproc_per_node,json=nprocPerNode,proto3" json:"nproc_per_node,omitempty"`
	MaxRestarts          int32    `protobuf:"varint,5,opt,name=max_restarts,json=maxRestarts,proto3" json:"max_restarts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElasticConfig) Reset()         { *m = ElasticConfig{} }
func (m *ElasticConfig) String() string { return proto.CompactTextString(m) }
func (*ElasticConfig) ProtoMessage()    {}
func (*ElasticConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_886ca6225d4c0b10, []int{0}
}

func (m *ElasticConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElasticConfig.Unmarshal(m, b)
}
func (m *ElasticConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElasticConfig.Marshal(b, m, deterministic)
}
func (m *ElasticConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElasticConfig.Merge(m, src)
}
func (m *ElasticConfig) XXX_Size() int {
	return xxx_messageInfo_ElasticConfig.Size(m)
}
func (m *ElasticConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ElasticConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ElasticConfig proto.InternalMessageInfo

func (m *ElasticConfig) GetRdzvBackend() string {
	if m != nil {
		return m.RdzvBackend
	}
	return ""
}

func (m *ElasticConfig) GetMinReplicas() int32 {
	if m != nil {
		return m.MinReplicas
	}
	return 0
}

func (m *ElasticConfig) GetMaxReplicas() int32 {
	if m != nil {
		return m.MaxReplicas
	}
	return 0
}

func (m *ElasticConfig) GetNprocPerNode() int32 {
	if m != nil {
		return m.NprocPerNode
	}
	return 0
}

func (m *ElasticConfig) GetMaxRestarts() int32 {
	if m != nil {
		return m.MaxRestarts
	}
	return 0
}

// Proto for plugin that enables distributed training using https://github.com/kubeflow/pytorch-operator
type DistributedPyTorchTrainingTask struct {
	// Worker replicas spec
	WorkerReplicas *DistributedPyTorchTrainingReplicaSpec `protobuf:"bytes,1,opt,name=worker_replicas,json=workerReplicas,proto3" json:"worker_replicas,omitempty"`
	// Master replicas spec, master replicas can only have 1 replica
	MasterReplicas *DistributedPyTorchTrainingReplicaSpec `protobuf:"bytes,2,opt,name=master_replicas,json=masterReplicas,proto3" json:"master_replicas,omitempty"`
	// RunPolicy encapsulates various runtime policies of the distributed training
	// job, for example how to clean up resources and how long the job can stay
	// active.
	RunPolicy *RunPolicy `protobuf:"bytes,3,opt,name=run_policy,json=runPolicy,proto3" json:"run_policy,omitempty"`
	// config for an elastic pytorch job
	ElasticConfig        *ElasticConfig `protobuf:"bytes,4,opt,name=elastic_config,json=elasticConfig,proto3" json:"elastic_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DistributedPyTorchTrainingTask) Reset()         { *m = DistributedPyTorchTrainingTask{} }
func (m *DistributedPyTorchTrainingTask) String() string { return proto.CompactTextString(m) }
func (*DistributedPyTorchTrainingTask) ProtoMessage()    {}
func (*DistributedPyTorchTrainingTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_886ca6225d4c0b10, []int{1}
}

func (m *DistributedPyTorchTrainingTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedPyTorchTrainingTask.Unmarshal(m, b)
}
func (m *DistributedPyTorchTrainingTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedPyTorchTrainingTask.Marshal(b, m, deterministic)
}
func (m *DistributedPyTorchTrainingTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedPyTorchTrainingTask.Merge(m, src)
}
func (m *DistributedPyTorchTrainingTask) XXX_Size() int {
	return xxx_messageInfo_DistributedPyTorchTrainingTask.Size(m)
}
func (m *DistributedPyTorchTrainingTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedPyTorchTrainingTask.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedPyTorchTrainingTask proto.InternalMessageInfo

func (m *DistributedPyTorchTrainingTask) GetWorkerReplicas() *DistributedPyTorchTrainingReplicaSpec {
	if m != nil {
		return m.WorkerReplicas
	}
	return nil
}

func (m *DistributedPyTorchTrainingTask) GetMasterReplicas() *DistributedPyTorchTrainingReplicaSpec {
	if m != nil {
		return m.MasterReplicas
	}
	return nil
}

func (m *DistributedPyTorchTrainingTask) GetRunPolicy() *RunPolicy {
	if m != nil {
		return m.RunPolicy
	}
	return nil
}

func (m *DistributedPyTorchTrainingTask) GetElasticConfig() *ElasticConfig {
	if m != nil {
		return m.ElasticConfig
	}
	return nil
}

type DistributedPyTorchTrainingReplicaSpec struct {
	// Number of replicas
	Replicas int32 `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// Image used for the replica group
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// Resources required for the replica group
	Resources *core.Resources `protobuf:"bytes,3,opt,name=resources,proto3" json:"resources,omitempty"`
	// RestartPolicy determines whether pods will be restarted when they exit
	RestartPolicy        RestartPolicy `protobuf:"varint,4,opt,name=restart_policy,json=restartPolicy,proto3,enum=nebulaidl.plugins.kubeflow.RestartPolicy" json:"restart_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DistributedPyTorchTrainingReplicaSpec) Reset()         { *m = DistributedPyTorchTrainingReplicaSpec{} }
func (m *DistributedPyTorchTrainingReplicaSpec) String() string { return proto.CompactTextString(m) }
func (*DistributedPyTorchTrainingReplicaSpec) ProtoMessage()    {}
func (*DistributedPyTorchTrainingReplicaSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_886ca6225d4c0b10, []int{2}
}

func (m *DistributedPyTorchTrainingReplicaSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedPyTorchTrainingReplicaSpec.Unmarshal(m, b)
}
func (m *DistributedPyTorchTrainingReplicaSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedPyTorchTrainingReplicaSpec.Marshal(b, m, deterministic)
}
func (m *DistributedPyTorchTrainingReplicaSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedPyTorchTrainingReplicaSpec.Merge(m, src)
}
func (m *DistributedPyTorchTrainingReplicaSpec) XXX_Size() int {
	return xxx_messageInfo_DistributedPyTorchTrainingReplicaSpec.Size(m)
}
func (m *DistributedPyTorchTrainingReplicaSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedPyTorchTrainingReplicaSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedPyTorchTrainingReplicaSpec proto.InternalMessageInfo

func (m *DistributedPyTorchTrainingReplicaSpec) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *DistributedPyTorchTrainingReplicaSpec) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DistributedPyTorchTrainingReplicaSpec) GetResources() *core.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DistributedPyTorchTrainingReplicaSpec) GetRestartPolicy() RestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return RestartPolicy_RESTART_POLICY_NEVER
}

func init() {
	proto.RegisterType((*ElasticConfig)(nil), "nebulaidl.plugins.kubeflow.ElasticConfig")
	proto.RegisterType((*DistributedPyTorchTrainingTask)(nil), "nebulaidl.plugins.kubeflow.DistributedPyTorchTrainingTask")
	proto.RegisterType((*DistributedPyTorchTrainingReplicaSpec)(nil), "nebulaidl.plugins.kubeflow.DistributedPyTorchTrainingReplicaSpec")
}

func init() {
	proto.RegisterFile("nebulaidl/plugins/kubeflow/pytorch.proto", fileDescriptor_886ca6225d4c0b10)
}

var fileDescriptor_886ca6225d4c0b10 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd5, 0x8d, 0x22, 0xea, 0xac, 0x9d, 0x14, 0x71, 0x28, 0x3d, 0xa0, 0x51, 0x31, 0x51,
	0x0e, 0x24, 0x52, 0x39, 0x70, 0x42, 0x88, 0x31, 0xae, 0xa8, 0x32, 0x3d, 0x71, 0x89, 0x1c, 0xe7,
	0x2d, 0xf3, 0x92, 0xd8, 0xd1, 0xb3, 0xcd, 0x56, 0x3e, 0x05, 0x5f, 0x89, 0xef, 0xc2, 0x07, 0x41,
	0xb1, 0xd3, 0x24, 0x13, 0x5a, 0xc5, 0x81, 0x5b, 0xde, 0xcb, 0xcf, 0xff, 0xe7, 0xff, 0xf3, 0x7b,
	0x64, 0x25, 0x21, 0xb5, 0x25, 0x13, 0x59, 0x19, 0xd7, 0xa5, 0xcd, 0x85, 0xd4, 0x71, 0x61, 0x53,
	0xb8, 0x2a, 0xd5, 0x6d, 0x5c, 0xef, 0x8c, 0x42, 0x7e, 0x1d, 0xd5, 0xa8, 0x8c, 0x0a, 0x17, 0x1d,
	0x19, 0xb5, 0x64, 0xb4, 0x27, 0x17, 0xfd, 0xbf, 0x98, 0x2b, 0x84, 0xd8, 0x30, 0x5d, 0x68, 0x7f,
	0x6e, 0xf1, 0xea, 0x40, 0x05, 0xae, 0xaa, 0x4a, 0x49, 0x0f, 0x2e, 0x7f, 0x8d, 0xc8, 0xf4, 0x73,
	0xc9, 0xb4, 0x11, 0xfc, 0x93, 0x92, 0x57, 0x22, 0x0f, 0x5f, 0x90, 0x13, 0xcc, 0x7e, 0x7c, 0x4f,
	0x52, 0xc6, 0x0b, 0x90, 0xd9, 0x7c, 0x74, 0x36, 0x5a, 0x4d, 0x68, 0xd0, 0xe4, 0x2e, 0x7c, 0xaa,
	0x41, 0x2a, 0x21, 0x13, 0x84, 0xba, 0x14, 0x9c, 0xe9, 0xf9, 0xd1, 0xd9, 0x68, 0x35, 0xa6, 0x41,
	0x25, 0x24, 0x6d, 0x53, 0x0e, 0x61, 0x77, 0x3d, 0x72, 0xdc, 0x22, 0xec, 0xae, 0x43, 0x5e, 0x92,
	0x99, 0xac, 0x51, 0xf1, 0xa4, 0x06, 0x4c, 0xa4, 0xca, 0x60, 0xfe, 0xc8, 0x41, 0x27, 0x2e, 0xbb,
	0x01, 0xfc, 0xa2, 0x32, 0xe8, 0x85, 0xb4, 0x61, 0x68, 0xf4, 0x7c, 0x3c, 0x10, 0xf2, 0xa9, 0xe5,
	0xcf, 0x63, 0xf2, 0xfc, 0x52, 0x68, 0x83, 0x22, 0xb5, 0x06, 0xb2, 0xcd, 0x6e, 0xdb, 0x74, 0x70,
	0x8b, 0x4c, 0x48, 0x21, 0xf3, 0x2d, 0xd3, 0x45, 0x78, 0x43, 0x4e, 0x6f, 0x15, 0x16, 0x80, 0xfd,
	0x8d, 0x1a, 0x5f, 0xc1, 0xfa, 0x63, 0xf4, 0x70, 0x87, 0xa3, 0x87, 0x45, 0x5b, 0x13, 0x5f, 0x6b,
	0xe0, 0x74, 0xe6, 0x95, 0x3b, 0x5f, 0x37, 0xe4, 0xb4, 0x62, 0xda, 0x0c, 0x6b, 0x1d, 0xfd, 0xb7,
	0x5a, 0x5e, 0xb9, 0xab, 0x75, 0x49, 0x08, 0x5a, 0x99, 0xd4, 0xaa, 0x14, 0x7c, 0xe7, 0x9a, 0x1c,
	0xac, 0xcf, 0x0f, 0x95, 0xa1, 0x56, 0x6e, 0x1c, 0x4c, 0x27, 0xb8, 0xff, 0x0c, 0x37, 0x64, 0x06,
	0x7e, 0x06, 0x12, 0xee, 0x86, 0xc0, 0xbd, 0x44, 0xb0, 0x7e, 0x7d, 0x48, 0xe9, 0xde, 0xd4, 0xd0,
	0x29, 0x0c, 0xc3, 0xe5, 0xef, 0x11, 0x39, 0xff, 0x27, 0x47, 0xe1, 0x82, 0x3c, 0xb9, 0xf7, 0x24,
	0x63, 0xda, 0xc5, 0xe1, 0x53, 0x32, 0x16, 0x15, 0xcb, 0xc1, 0xf5, 0x6f, 0x42, 0x7d, 0x10, 0xbe,
	0x23, 0x13, 0x04, 0xad, 0x2c, 0x72, 0xd0, 0xad, 0xe5, 0x67, 0x83, 0x8b, 0x36, 0xbb, 0x10, 0xd1,
	0x3d, 0x40, 0x7b, 0xb6, 0xb1, 0xd9, 0x8e, 0xd1, 0xbe, 0x61, 0x8d, 0xcd, 0xd9, 0x61, 0x9b, 0xed,
	0x94, 0xb5, 0x4d, 0x9b, 0xe2, 0x30, 0xbc, 0xf8, 0xf0, 0xed, 0x7d, 0x2e, 0xcc, 0xb5, 0x4d, 0x23,
	0xae, 0xaa, 0xd8, 0xab, 0xf0, 0x52, 0xd9, 0x4c, 0xc7, 0xfd, 0x02, 0xe6, 0x20, 0xe3, 0x3a, 0x7d,
	0x93, 0xab, 0xf8, 0xaf, 0xa5, 0x4c, 0x1f, 0xbb, 0x2d, 0x7c, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x95, 0x3e, 0x93, 0xf7, 0x12, 0x04, 0x00, 0x00,
}
