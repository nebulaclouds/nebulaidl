// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/admin/project_attributes.proto

package admin

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

// Defines a set of custom matching attributes at the project level.
// For more info on matchable attributes, see :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration`
type ProjectAttributes struct {
	// Unique project id for which this set of attributes will be applied.
	Project              string              `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	MatchingAttributes   *MatchingAttributes `protobuf:"bytes,2,opt,name=matching_attributes,json=matchingAttributes,proto3" json:"matching_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProjectAttributes) Reset()         { *m = ProjectAttributes{} }
func (m *ProjectAttributes) String() string { return proto.CompactTextString(m) }
func (*ProjectAttributes) ProtoMessage()    {}
func (*ProjectAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8dc9faa9640256, []int{0}
}

func (m *ProjectAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectAttributes.Unmarshal(m, b)
}
func (m *ProjectAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectAttributes.Marshal(b, m, deterministic)
}
func (m *ProjectAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectAttributes.Merge(m, src)
}
func (m *ProjectAttributes) XXX_Size() int {
	return xxx_messageInfo_ProjectAttributes.Size(m)
}
func (m *ProjectAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectAttributes proto.InternalMessageInfo

func (m *ProjectAttributes) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectAttributes) GetMatchingAttributes() *MatchingAttributes {
	if m != nil {
		return m.MatchingAttributes
	}
	return nil
}

// Sets custom attributes for a project
// For more info on matchable attributes, see :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration`
type ProjectAttributesUpdateRequest struct {
	// +required
	Attributes           *ProjectAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ProjectAttributesUpdateRequest) Reset()         { *m = ProjectAttributesUpdateRequest{} }
func (m *ProjectAttributesUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectAttributesUpdateRequest) ProtoMessage()    {}
func (*ProjectAttributesUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8dc9faa9640256, []int{1}
}

func (m *ProjectAttributesUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectAttributesUpdateRequest.Unmarshal(m, b)
}
func (m *ProjectAttributesUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectAttributesUpdateRequest.Marshal(b, m, deterministic)
}
func (m *ProjectAttributesUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectAttributesUpdateRequest.Merge(m, src)
}
func (m *ProjectAttributesUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectAttributesUpdateRequest.Size(m)
}
func (m *ProjectAttributesUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectAttributesUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectAttributesUpdateRequest proto.InternalMessageInfo

func (m *ProjectAttributesUpdateRequest) GetAttributes() *ProjectAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Purposefully empty, may be populated in the future.
type ProjectAttributesUpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectAttributesUpdateResponse) Reset()         { *m = ProjectAttributesUpdateResponse{} }
func (m *ProjectAttributesUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectAttributesUpdateResponse) ProtoMessage()    {}
func (*ProjectAttributesUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8dc9faa9640256, []int{2}
}

func (m *ProjectAttributesUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectAttributesUpdateResponse.Unmarshal(m, b)
}
func (m *ProjectAttributesUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectAttributesUpdateResponse.Marshal(b, m, deterministic)
}
func (m *ProjectAttributesUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectAttributesUpdateResponse.Merge(m, src)
}
func (m *ProjectAttributesUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectAttributesUpdateResponse.Size(m)
}
func (m *ProjectAttributesUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectAttributesUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectAttributesUpdateResponse proto.InternalMessageInfo

// Request to get an individual project level attribute override.
// For more info on matchable attributes, see :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration`
type ProjectAttributesGetRequest struct {
	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Which type of matchable attributes to return.
	// +required
	ResourceType         MatchableResource `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=nebulaidl.admin.MatchableResource" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProjectAttributesGetRequest) Reset()         { *m = ProjectAttributesGetRequest{} }
func (m *ProjectAttributesGetRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectAttributesGetRequest) ProtoMessage()    {}
func (*ProjectAttributesGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8dc9faa9640256, []int{3}
}

func (m *ProjectAttributesGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectAttributesGetRequest.Unmarshal(m, b)
}
func (m *ProjectAttributesGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectAttributesGetRequest.Marshal(b, m, deterministic)
}
func (m *ProjectAttributesGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectAttributesGetRequest.Merge(m, src)
}
func (m *ProjectAttributesGetRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectAttributesGetRequest.Size(m)
}
func (m *ProjectAttributesGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectAttributesGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectAttributesGetRequest proto.InternalMessageInfo

func (m *ProjectAttributesGetRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectAttributesGetRequest) GetResourceType() MatchableResource {
	if m != nil {
		return m.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

// Response to get an individual project level attribute override.
// For more info on matchable attributes, see :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration`
type ProjectAttributesGetResponse struct {
	Attributes           *ProjectAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ProjectAttributesGetResponse) Reset()         { *m = ProjectAttributesGetResponse{} }
func (m *ProjectAttributesGetResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectAttributesGetResponse) ProtoMessage()    {}
func (*ProjectAttributesGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8dc9faa9640256, []int{4}
}

func (m *ProjectAttributesGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectAttributesGetResponse.Unmarshal(m, b)
}
func (m *ProjectAttributesGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectAttributesGetResponse.Marshal(b, m, deterministic)
}
func (m *ProjectAttributesGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectAttributesGetResponse.Merge(m, src)
}
func (m *ProjectAttributesGetResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectAttributesGetResponse.Size(m)
}
func (m *ProjectAttributesGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectAttributesGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectAttributesGetResponse proto.InternalMessageInfo

func (m *ProjectAttributesGetResponse) GetAttributes() *ProjectAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Request to delete a set matchable project level attribute override.
// For more info on matchable attributes, see :ref:`ref_nebulaidl.admin.MatchableAttributesConfiguration`
type ProjectAttributesDeleteRequest struct {
	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Which type of matchable attributes to delete.
	// +required
	ResourceType         MatchableResource `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=nebulaidl.admin.MatchableResource" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProjectAttributesDeleteRequest) Reset()         { *m = ProjectAttributesDeleteRequest{} }
func (m *ProjectAttributesDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectAttributesDeleteRequest) ProtoMessage()    {}
func (*ProjectAttributesDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8dc9faa9640256, []int{5}
}

func (m *ProjectAttributesDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectAttributesDeleteRequest.Unmarshal(m, b)
}
func (m *ProjectAttributesDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectAttributesDeleteRequest.Marshal(b, m, deterministic)
}
func (m *ProjectAttributesDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectAttributesDeleteRequest.Merge(m, src)
}
func (m *ProjectAttributesDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectAttributesDeleteRequest.Size(m)
}
func (m *ProjectAttributesDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectAttributesDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectAttributesDeleteRequest proto.InternalMessageInfo

func (m *ProjectAttributesDeleteRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectAttributesDeleteRequest) GetResourceType() MatchableResource {
	if m != nil {
		return m.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

// Purposefully empty, may be populated in the future.
type ProjectAttributesDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectAttributesDeleteResponse) Reset()         { *m = ProjectAttributesDeleteResponse{} }
func (m *ProjectAttributesDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectAttributesDeleteResponse) ProtoMessage()    {}
func (*ProjectAttributesDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb8dc9faa9640256, []int{6}
}

func (m *ProjectAttributesDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectAttributesDeleteResponse.Unmarshal(m, b)
}
func (m *ProjectAttributesDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectAttributesDeleteResponse.Marshal(b, m, deterministic)
}
func (m *ProjectAttributesDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectAttributesDeleteResponse.Merge(m, src)
}
func (m *ProjectAttributesDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectAttributesDeleteResponse.Size(m)
}
func (m *ProjectAttributesDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectAttributesDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectAttributesDeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProjectAttributes)(nil), "nebulaidl.admin.ProjectAttributes")
	proto.RegisterType((*ProjectAttributesUpdateRequest)(nil), "nebulaidl.admin.ProjectAttributesUpdateRequest")
	proto.RegisterType((*ProjectAttributesUpdateResponse)(nil), "nebulaidl.admin.ProjectAttributesUpdateResponse")
	proto.RegisterType((*ProjectAttributesGetRequest)(nil), "nebulaidl.admin.ProjectAttributesGetRequest")
	proto.RegisterType((*ProjectAttributesGetResponse)(nil), "nebulaidl.admin.ProjectAttributesGetResponse")
	proto.RegisterType((*ProjectAttributesDeleteRequest)(nil), "nebulaidl.admin.ProjectAttributesDeleteRequest")
	proto.RegisterType((*ProjectAttributesDeleteResponse)(nil), "nebulaidl.admin.ProjectAttributesDeleteResponse")
}

func init() {
	proto.RegisterFile("nebulaidl/admin/project_attributes.proto", fileDescriptor_cb8dc9faa9640256)
}

var fileDescriptor_cb8dc9faa9640256 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x65, 0x06, 0x10, 0x06, 0x2a, 0x11, 0x96, 0x0a, 0x10, 0xb4, 0x5e, 0xe8, 0x42, 0x2c,
	0x81, 0x10, 0x73, 0x11, 0x82, 0xa9, 0x12, 0x32, 0xb0, 0xb0, 0x54, 0x4e, 0x7a, 0xa4, 0x41, 0x49,
	0x6c, 0x9c, 0xcb, 0x90, 0x09, 0xa9, 0xbf, 0x1c, 0xc9, 0x89, 0xdb, 0xb4, 0x49, 0x37, 0x18, 0x13,
	0xbd, 0x7b, 0xef, 0xf3, 0xbb, 0xa3, 0x57, 0x9f, 0x49, 0x89, 0x10, 0xcf, 0x12, 0x2e, 0x67, 0x69,
	0x9c, 0x71, 0x6d, 0xd4, 0x17, 0x84, 0x38, 0x95, 0x88, 0x26, 0x0e, 0x0a, 0x84, 0xdc, 0xd7, 0x46,
	0xa1, 0xf2, 0x7a, 0x4e, 0xe8, 0x5b, 0xe1, 0xe9, 0xe6, 0x60, 0x2a, 0x31, 0x9c, 0xcb, 0x20, 0x81,
	0xa9, 0x81, 0x5c, 0x15, 0x26, 0x84, 0x6a, 0x90, 0x2d, 0x08, 0x3d, 0x7e, 0xa9, 0x5c, 0xc7, 0x4b,
	0x53, 0xaf, 0x4f, 0xf7, 0xea, 0xa8, 0x3e, 0x19, 0x90, 0xd1, 0xbe, 0x70, 0x9f, 0xde, 0x2b, 0x3d,
	0xb1, 0x5e, 0x71, 0x16, 0x35, 0x28, 0xfa, 0x3b, 0x03, 0x32, 0x3a, 0xb8, 0x61, 0xfe, 0x3a, 0x86,
	0x3f, 0xa9, 0xa5, 0x2b, 0x6b, 0xe1, 0xa5, 0xad, 0x7f, 0x2c, 0xa4, 0x17, 0x2d, 0x86, 0x77, 0x3d,
	0x93, 0x08, 0x02, 0xbe, 0x0b, 0xc8, 0xd1, 0x1b, 0x53, 0xda, 0x48, 0x23, 0x36, 0x6d, 0xb8, 0x99,
	0xd6, 0xf2, 0x10, 0x8d, 0x21, 0x36, 0xa4, 0x97, 0x5b, 0x43, 0x72, 0xad, 0xb2, 0x1c, 0xd8, 0x0f,
	0x3d, 0x6b, 0x49, 0x9e, 0x01, 0x1d, 0xc4, 0xf6, 0x56, 0x9e, 0xe8, 0x91, 0xeb, 0x75, 0x8a, 0xa5,
	0x06, 0xdb, 0x47, 0xaf, 0x4d, 0x38, 0x71, 0x6b, 0x10, 0xb5, 0x5a, 0x1c, 0xba, 0xb9, 0xb7, 0x52,
	0x03, 0x93, 0xf4, 0xbc, 0x1b, 0xa0, 0x02, 0xfc, 0x8b, 0x1a, 0x16, 0xa4, 0xa3, 0xec, 0x47, 0x48,
	0x60, 0x55, 0xf6, 0xff, 0xbf, 0xb3, 0x6b, 0x17, 0x8e, 0xa1, 0x7a, 0xea, 0xc3, 0xfd, 0xc7, 0x5d,
	0x14, 0xe3, 0xbc, 0x08, 0xfc, 0x50, 0xa5, 0xdc, 0xfa, 0x2b, 0x13, 0xf1, 0xe5, 0x5d, 0x47, 0x90,
	0x71, 0x1d, 0x5c, 0x47, 0x8a, 0xaf, 0x9f, 0x7a, 0xb0, 0x6b, 0x0f, 0xfb, 0xf6, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x34, 0x6f, 0x88, 0x40, 0x3c, 0x03, 0x00, 0x00,
}
