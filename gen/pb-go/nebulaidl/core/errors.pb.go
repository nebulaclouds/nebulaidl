// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/core/errors.proto

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

// Defines a generic error type that dictates the behavior of the retry strategy.
type ContainerError_Kind int32

const (
	ContainerError_NON_RECOVERABLE ContainerError_Kind = 0
	ContainerError_RECOVERABLE     ContainerError_Kind = 1
)

var ContainerError_Kind_name = map[int32]string{
	0: "NON_RECOVERABLE",
	1: "RECOVERABLE",
}

var ContainerError_Kind_value = map[string]int32{
	"NON_RECOVERABLE": 0,
	"RECOVERABLE":     1,
}

func (x ContainerError_Kind) String() string {
	return proto.EnumName(ContainerError_Kind_name, int32(x))
}

func (ContainerError_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2a49b99f342b879, []int{0, 0}
}

// Error message to propagate detailed errors from container executions to the execution
// engine.
type ContainerError struct {
	// A simplified code for errors, so that we can provide a glossary of all possible errors.
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// A detailed error message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// An abstract error kind for this error. Defaults to Non_Recoverable if not specified.
	Kind ContainerError_Kind `protobuf:"varint,3,opt,name=kind,proto3,enum=nebulaidl.core.ContainerError_Kind" json:"kind,omitempty"`
	// Defines the origin of the error (system, user, unknown).
	Origin               ExecutionError_ErrorKind `protobuf:"varint,4,opt,name=origin,proto3,enum=nebulaidl.core.ExecutionError_ErrorKind" json:"origin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ContainerError) Reset()         { *m = ContainerError{} }
func (m *ContainerError) String() string { return proto.CompactTextString(m) }
func (*ContainerError) ProtoMessage()    {}
func (*ContainerError) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2a49b99f342b879, []int{0}
}

func (m *ContainerError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerError.Unmarshal(m, b)
}
func (m *ContainerError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerError.Marshal(b, m, deterministic)
}
func (m *ContainerError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerError.Merge(m, src)
}
func (m *ContainerError) XXX_Size() int {
	return xxx_messageInfo_ContainerError.Size(m)
}
func (m *ContainerError) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerError.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerError proto.InternalMessageInfo

func (m *ContainerError) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ContainerError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ContainerError) GetKind() ContainerError_Kind {
	if m != nil {
		return m.Kind
	}
	return ContainerError_NON_RECOVERABLE
}

func (m *ContainerError) GetOrigin() ExecutionError_ErrorKind {
	if m != nil {
		return m.Origin
	}
	return ExecutionError_UNKNOWN
}

// Defines the errors.pb file format the container can produce to communicate
// failure reasons to the execution engine.
type ErrorDocument struct {
	// The error raised during execution.
	Error                *ContainerError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ErrorDocument) Reset()         { *m = ErrorDocument{} }
func (m *ErrorDocument) String() string { return proto.CompactTextString(m) }
func (*ErrorDocument) ProtoMessage()    {}
func (*ErrorDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2a49b99f342b879, []int{1}
}

func (m *ErrorDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorDocument.Unmarshal(m, b)
}
func (m *ErrorDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorDocument.Marshal(b, m, deterministic)
}
func (m *ErrorDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDocument.Merge(m, src)
}
func (m *ErrorDocument) XXX_Size() int {
	return xxx_messageInfo_ErrorDocument.Size(m)
}
func (m *ErrorDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDocument.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDocument proto.InternalMessageInfo

func (m *ErrorDocument) GetError() *ContainerError {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterEnum("nebulaidl.core.ContainerError_Kind", ContainerError_Kind_name, ContainerError_Kind_value)
	proto.RegisterType((*ContainerError)(nil), "nebulaidl.core.ContainerError")
	proto.RegisterType((*ErrorDocument)(nil), "nebulaidl.core.ErrorDocument")
}

func init() { proto.RegisterFile("nebulaidl/core/errors.proto", fileDescriptor_c2a49b99f342b879) }

var fileDescriptor_c2a49b99f342b879 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xad, 0xd6, 0x89, 0x6f, 0x6c, 0x93, 0x78, 0x29, 0x83, 0xc1, 0xe8, 0xc5, 0x1d, 0x34,
	0x81, 0x4d, 0x76, 0x15, 0xb7, 0xf5, 0xa4, 0x6c, 0xd0, 0x83, 0x07, 0x2f, 0xb2, 0xb6, 0x31, 0x06,
	0xd7, 0xbc, 0x91, 0xa6, 0xa0, 0x7f, 0xb0, 0xff, 0x87, 0xf4, 0xd5, 0x8a, 0xdd, 0xc1, 0x4b, 0xc8,
	0x7b, 0xef, 0xfb, 0x7d, 0x79, 0xf9, 0x60, 0xf8, 0xba, 0xfb, 0x74, 0x52, 0x67, 0x3b, 0x91, 0xa2,
	0x95, 0x42, 0x5a, 0x8b, 0xb6, 0xe0, 0x7b, 0x8b, 0x0e, 0x59, 0xaf, 0x99, 0xf1, 0x6a, 0x36, 0x1c,
	0x1d, 0x48, 0x3f, 0x64, 0x5a, 0x3a, 0x8d, 0xa6, 0x56, 0x87, 0x5f, 0x1e, 0xf4, 0x97, 0x68, 0xdc,
	0x56, 0x1b, 0x69, 0xa3, 0xca, 0x87, 0x31, 0xf0, 0x53, 0xcc, 0x64, 0xe0, 0x8d, 0xbd, 0xc9, 0x79,
	0x4c, 0x77, 0x16, 0xc0, 0x59, 0x2e, 0x8b, 0x62, 0xab, 0x64, 0x70, 0x4c, 0xed, 0xa6, 0x64, 0x73,
	0xf0, 0xdf, 0xb5, 0xc9, 0x82, 0x93, 0xb1, 0x37, 0xe9, 0x4f, 0x43, 0xde, 0x7a, 0x9d, 0xb7, 0xad,
	0xf9, 0x83, 0x36, 0x59, 0x4c, 0x7a, 0x76, 0x07, 0x1d, 0xb4, 0x5a, 0x69, 0x13, 0xf8, 0x44, 0x5e,
	0x1d, 0x90, 0x51, 0xb3, 0x68, 0x4d, 0xd2, 0x49, 0xf8, 0x0f, 0x16, 0x5e, 0x83, 0x5f, 0xd5, 0xec,
	0x12, 0x06, 0xeb, 0xcd, 0xfa, 0x25, 0x8e, 0x96, 0x9b, 0xa7, 0x28, 0xbe, 0x5f, 0x3c, 0x46, 0x17,
	0x47, 0x6c, 0x00, 0xdd, 0xbf, 0x0d, 0x2f, 0x5c, 0x41, 0x8f, 0x2c, 0x56, 0x98, 0x96, 0xb9, 0x34,
	0x8e, 0xcd, 0xe0, 0x94, 0x62, 0xa3, 0x6f, 0x76, 0xa7, 0xa3, 0x7f, 0x17, 0x8f, 0x6b, 0xed, 0x62,
	0xfe, 0x7c, 0xab, 0xb4, 0x7b, 0x2b, 0x13, 0x9e, 0x62, 0x2e, 0x88, 0x40, 0xab, 0xc4, 0x6f, 0xc4,
	0x4a, 0x1a, 0xb1, 0x4f, 0x6e, 0x14, 0x8a, 0x56, 0xea, 0x49, 0x87, 0xc2, 0x9e, 0x7d, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x79, 0x84, 0xe2, 0xf9, 0xb8, 0x01, 0x00, 0x00,
}
