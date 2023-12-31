// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/plugins/tensorflow.proto

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

// Custom proto for plugin that enables distributed training using https://github.com/kubeflow/tf-operator
type DistributedTensorflowTrainingTask struct {
	// number of worker, ps, chief replicas spawned in the cluster for this job
	Workers int32 `protobuf:"varint,1,opt,name=workers,proto3" json:"workers,omitempty"`
	// PS -> Parameter server
	PsReplicas           int32    `protobuf:"varint,2,opt,name=ps_replicas,json=psReplicas,proto3" json:"ps_replicas,omitempty"`
	ChiefReplicas        int32    `protobuf:"varint,3,opt,name=chief_replicas,json=chiefReplicas,proto3" json:"chief_replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributedTensorflowTrainingTask) Reset()         { *m = DistributedTensorflowTrainingTask{} }
func (m *DistributedTensorflowTrainingTask) String() string { return proto.CompactTextString(m) }
func (*DistributedTensorflowTrainingTask) ProtoMessage()    {}
func (*DistributedTensorflowTrainingTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e751b9753142802, []int{0}
}

func (m *DistributedTensorflowTrainingTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedTensorflowTrainingTask.Unmarshal(m, b)
}
func (m *DistributedTensorflowTrainingTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedTensorflowTrainingTask.Marshal(b, m, deterministic)
}
func (m *DistributedTensorflowTrainingTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedTensorflowTrainingTask.Merge(m, src)
}
func (m *DistributedTensorflowTrainingTask) XXX_Size() int {
	return xxx_messageInfo_DistributedTensorflowTrainingTask.Size(m)
}
func (m *DistributedTensorflowTrainingTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedTensorflowTrainingTask.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedTensorflowTrainingTask proto.InternalMessageInfo

func (m *DistributedTensorflowTrainingTask) GetWorkers() int32 {
	if m != nil {
		return m.Workers
	}
	return 0
}

func (m *DistributedTensorflowTrainingTask) GetPsReplicas() int32 {
	if m != nil {
		return m.PsReplicas
	}
	return 0
}

func (m *DistributedTensorflowTrainingTask) GetChiefReplicas() int32 {
	if m != nil {
		return m.ChiefReplicas
	}
	return 0
}

func init() {
	proto.RegisterType((*DistributedTensorflowTrainingTask)(nil), "nebulaidl.plugins.DistributedTensorflowTrainingTask")
}

func init() { proto.RegisterFile("nebulaidl/plugins/tensorflow.proto", fileDescriptor_7e751b9753142802) }

var fileDescriptor_7e751b9753142802 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xcf, 0xbf, 0x4a, 0xc7, 0x30,
	0x10, 0x07, 0x70, 0xaa, 0xa8, 0x10, 0x51, 0x30, 0x53, 0x37, 0xb5, 0x20, 0xb8, 0xd8, 0x0c, 0xce,
	0x22, 0x88, 0x4f, 0x50, 0x3a, 0xb9, 0x48, 0x92, 0xa6, 0xe9, 0xd1, 0x98, 0x0b, 0xb9, 0x84, 0x3e,
	0x81, 0xef, 0x2d, 0xc4, 0xfe, 0x19, 0x7e, 0xe3, 0x7d, 0xef, 0x73, 0x70, 0x5f, 0xd6, 0x78, 0xa3,
	0xb2, 0x93, 0x30, 0x38, 0x11, 0x5c, 0xb6, 0xe0, 0x49, 0x24, 0xe3, 0x09, 0xe3, 0xe8, 0x70, 0x69,
	0x43, 0xc4, 0x84, 0xfc, 0x6e, 0x37, 0xed, 0x6a, 0x9a, 0xdf, 0x8a, 0x3d, 0x7e, 0x02, 0xa5, 0x08,
	0x2a, 0x27, 0x33, 0xf4, 0xfb, 0x49, 0x1f, 0x25, 0x78, 0xf0, 0xb6, 0x97, 0x34, 0xf3, 0x9a, 0x5d,
	0x2d, 0x18, 0x67, 0x13, 0xa9, 0xae, 0x1e, 0xaa, 0xe7, 0x8b, 0x6e, 0x1b, 0xf9, 0x3d, 0xbb, 0x0e,
	0xf4, 0x1d, 0x4d, 0x70, 0xa0, 0x25, 0xd5, 0x67, 0x65, 0xcb, 0x02, 0x75, 0x6b, 0xc2, 0x9f, 0xd8,
	0xad, 0x9e, 0xc0, 0x8c, 0x87, 0x39, 0x2f, 0xe6, 0xa6, 0xa4, 0x1b, 0xfb, 0x78, 0xff, 0x7a, 0xb3,
	0x90, 0xa6, 0xac, 0x5a, 0x8d, 0x3f, 0xe2, 0xff, 0x4f, 0xed, 0x30, 0x0f, 0x24, 0x8e, 0x62, 0xd6,
	0x78, 0x11, 0xd4, 0x8b, 0x45, 0x71, 0x52, 0x56, 0x5d, 0x96, 0x8a, 0xaf, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x6b, 0xaa, 0x2b, 0x6b, 0x08, 0x01, 0x00, 0x00,
}
