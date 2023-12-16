// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/plugins/presto.proto

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

// This message works with the 'presto' task type in the SDK and is the object that will be in the 'custom' field
// of a Presto task's TaskTemplate
type PrestoQuery struct {
	RoutingGroup         string   `protobuf:"bytes,1,opt,name=routing_group,json=routingGroup,proto3" json:"routing_group,omitempty"`
	Catalog              string   `protobuf:"bytes,2,opt,name=catalog,proto3" json:"catalog,omitempty"`
	Schema               string   `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Statement            string   `protobuf:"bytes,4,opt,name=statement,proto3" json:"statement,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrestoQuery) Reset()         { *m = PrestoQuery{} }
func (m *PrestoQuery) String() string { return proto.CompactTextString(m) }
func (*PrestoQuery) ProtoMessage()    {}
func (*PrestoQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_881edc23a44f4737, []int{0}
}

func (m *PrestoQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrestoQuery.Unmarshal(m, b)
}
func (m *PrestoQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrestoQuery.Marshal(b, m, deterministic)
}
func (m *PrestoQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrestoQuery.Merge(m, src)
}
func (m *PrestoQuery) XXX_Size() int {
	return xxx_messageInfo_PrestoQuery.Size(m)
}
func (m *PrestoQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PrestoQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PrestoQuery proto.InternalMessageInfo

func (m *PrestoQuery) GetRoutingGroup() string {
	if m != nil {
		return m.RoutingGroup
	}
	return ""
}

func (m *PrestoQuery) GetCatalog() string {
	if m != nil {
		return m.Catalog
	}
	return ""
}

func (m *PrestoQuery) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *PrestoQuery) GetStatement() string {
	if m != nil {
		return m.Statement
	}
	return ""
}

func init() {
	proto.RegisterType((*PrestoQuery)(nil), "nebulaidl.plugins.PrestoQuery")
}

func init() { proto.RegisterFile("nebulaidl/plugins/presto.proto", fileDescriptor_881edc23a44f4737) }

var fileDescriptor_881edc23a44f4737 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcb, 0xa9, 0x2c,
	0x49, 0xcd, 0x4c, 0xc9, 0xd1, 0x2f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x2b, 0xd6, 0x2f, 0x28, 0x4a,
	0x2d, 0x2e, 0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x49, 0xeb, 0x41, 0xa5,
	0x95, 0x9a, 0x18, 0xb9, 0xb8, 0x03, 0xc0, 0x4a, 0x02, 0x4b, 0x53, 0x8b, 0x2a, 0x85, 0x94, 0xb9,
	0x78, 0x8b, 0xf2, 0x4b, 0x4b, 0x32, 0xf3, 0xd2, 0xe3, 0xd3, 0x8b, 0xf2, 0x4b, 0x0b, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x78, 0xa0, 0x82, 0xee, 0x20, 0x31, 0x21, 0x09, 0x2e, 0xf6, 0xe4,
	0xc4, 0x92, 0xc4, 0x9c, 0xfc, 0x74, 0x09, 0x26, 0xb0, 0x34, 0x8c, 0x2b, 0x24, 0xc6, 0xc5, 0x56,
	0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0x28, 0xc1, 0x0c, 0x96, 0x80, 0xf2, 0x84, 0x64, 0xb8, 0x38, 0x8b,
	0x4b, 0x12, 0x4b, 0x52, 0x73, 0x53, 0xf3, 0x4a, 0x24, 0x58, 0xc0, 0x52, 0x08, 0x01, 0x27, 0xcb,
	0x28, 0xf3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xb0, 0x1b, 0xf3,
	0x8b, 0xd2, 0xf5, 0xe1, 0x7e, 0x49, 0x4f, 0xcd, 0xd3, 0x2f, 0x48, 0xd2, 0x4d, 0xcf, 0xd7, 0x47,
	0xf7, 0x5e, 0x12, 0x1b, 0xd8, 0x63, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x67, 0xf9,
	0x0f, 0xf9, 0x00, 0x00, 0x00,
}
