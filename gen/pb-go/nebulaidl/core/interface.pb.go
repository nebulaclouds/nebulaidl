// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/core/interface.proto

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

// Defines a strongly typed variable.
type Variable struct {
	// Variable literal type.
	Type *LiteralType `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	//+optional string describing input variable
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Variable) Reset()         { *m = Variable{} }
func (m *Variable) String() string { return proto.CompactTextString(m) }
func (*Variable) ProtoMessage()    {}
func (*Variable) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef1f35079e48714, []int{0}
}

func (m *Variable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Variable.Unmarshal(m, b)
}
func (m *Variable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Variable.Marshal(b, m, deterministic)
}
func (m *Variable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Variable.Merge(m, src)
}
func (m *Variable) XXX_Size() int {
	return xxx_messageInfo_Variable.Size(m)
}
func (m *Variable) XXX_DiscardUnknown() {
	xxx_messageInfo_Variable.DiscardUnknown(m)
}

var xxx_messageInfo_Variable proto.InternalMessageInfo

func (m *Variable) GetType() *LiteralType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Variable) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// A map of Variables
type VariableMap struct {
	// Defines a map of variable names to variables.
	Variables            map[string]*Variable `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VariableMap) Reset()         { *m = VariableMap{} }
func (m *VariableMap) String() string { return proto.CompactTextString(m) }
func (*VariableMap) ProtoMessage()    {}
func (*VariableMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef1f35079e48714, []int{1}
}

func (m *VariableMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VariableMap.Unmarshal(m, b)
}
func (m *VariableMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VariableMap.Marshal(b, m, deterministic)
}
func (m *VariableMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VariableMap.Merge(m, src)
}
func (m *VariableMap) XXX_Size() int {
	return xxx_messageInfo_VariableMap.Size(m)
}
func (m *VariableMap) XXX_DiscardUnknown() {
	xxx_messageInfo_VariableMap.DiscardUnknown(m)
}

var xxx_messageInfo_VariableMap proto.InternalMessageInfo

func (m *VariableMap) GetVariables() map[string]*Variable {
	if m != nil {
		return m.Variables
	}
	return nil
}

// Defines strongly typed inputs and outputs.
type TypedInterface struct {
	Inputs               *VariableMap `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              *VariableMap `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TypedInterface) Reset()         { *m = TypedInterface{} }
func (m *TypedInterface) String() string { return proto.CompactTextString(m) }
func (*TypedInterface) ProtoMessage()    {}
func (*TypedInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef1f35079e48714, []int{2}
}

func (m *TypedInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedInterface.Unmarshal(m, b)
}
func (m *TypedInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedInterface.Marshal(b, m, deterministic)
}
func (m *TypedInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedInterface.Merge(m, src)
}
func (m *TypedInterface) XXX_Size() int {
	return xxx_messageInfo_TypedInterface.Size(m)
}
func (m *TypedInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedInterface.DiscardUnknown(m)
}

var xxx_messageInfo_TypedInterface proto.InternalMessageInfo

func (m *TypedInterface) GetInputs() *VariableMap {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TypedInterface) GetOutputs() *VariableMap {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// A parameter is used as input to a launch plan and has
// the special ability to have a default value or mark itself as required.
type Parameter struct {
	//+required Variable. Defines the type of the variable backing this parameter.
	Var *Variable `protobuf:"bytes,1,opt,name=var,proto3" json:"var,omitempty"`
	//+optional
	//
	// Types that are valid to be assigned to Behavior:
	//	*Parameter_Default
	//	*Parameter_Required
	Behavior             isParameter_Behavior `protobuf_oneof:"behavior"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef1f35079e48714, []int{3}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetVar() *Variable {
	if m != nil {
		return m.Var
	}
	return nil
}

type isParameter_Behavior interface {
	isParameter_Behavior()
}

type Parameter_Default struct {
	Default *Literal `protobuf:"bytes,2,opt,name=default,proto3,oneof"`
}

type Parameter_Required struct {
	Required bool `protobuf:"varint,3,opt,name=required,proto3,oneof"`
}

func (*Parameter_Default) isParameter_Behavior() {}

func (*Parameter_Required) isParameter_Behavior() {}

func (m *Parameter) GetBehavior() isParameter_Behavior {
	if m != nil {
		return m.Behavior
	}
	return nil
}

func (m *Parameter) GetDefault() *Literal {
	if x, ok := m.GetBehavior().(*Parameter_Default); ok {
		return x.Default
	}
	return nil
}

func (m *Parameter) GetRequired() bool {
	if x, ok := m.GetBehavior().(*Parameter_Required); ok {
		return x.Required
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Parameter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Parameter_Default)(nil),
		(*Parameter_Required)(nil),
	}
}

// A map of Parameters.
type ParameterMap struct {
	// Defines a map of parameter names to parameters.
	Parameters           map[string]*Parameter `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ParameterMap) Reset()         { *m = ParameterMap{} }
func (m *ParameterMap) String() string { return proto.CompactTextString(m) }
func (*ParameterMap) ProtoMessage()    {}
func (*ParameterMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_bef1f35079e48714, []int{4}
}

func (m *ParameterMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParameterMap.Unmarshal(m, b)
}
func (m *ParameterMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParameterMap.Marshal(b, m, deterministic)
}
func (m *ParameterMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterMap.Merge(m, src)
}
func (m *ParameterMap) XXX_Size() int {
	return xxx_messageInfo_ParameterMap.Size(m)
}
func (m *ParameterMap) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterMap.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterMap proto.InternalMessageInfo

func (m *ParameterMap) GetParameters() map[string]*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func init() {
	proto.RegisterType((*Variable)(nil), "nebulaidl.core.Variable")
	proto.RegisterType((*VariableMap)(nil), "nebulaidl.core.VariableMap")
	proto.RegisterMapType((map[string]*Variable)(nil), "nebulaidl.core.VariableMap.VariablesEntry")
	proto.RegisterType((*TypedInterface)(nil), "nebulaidl.core.TypedInterface")
	proto.RegisterType((*Parameter)(nil), "nebulaidl.core.Parameter")
	proto.RegisterType((*ParameterMap)(nil), "nebulaidl.core.ParameterMap")
	proto.RegisterMapType((map[string]*Parameter)(nil), "nebulaidl.core.ParameterMap.ParametersEntry")
}

func init() { proto.RegisterFile("nebulaidl/core/interface.proto", fileDescriptor_bef1f35079e48714) }

var fileDescriptor_bef1f35079e48714 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0x69, 0x75, 0xb7, 0xbd, 0x91, 0x2a, 0xf3, 0x62, 0x8c, 0x1f, 0x84, 0x3c, 0x95, 0x45,
	0x13, 0x68, 0x11, 0x64, 0xf1, 0x69, 0x41, 0xa8, 0xb0, 0x82, 0x0c, 0xb2, 0x88, 0xe0, 0xc3, 0x24,
	0xb9, 0xdb, 0x1d, 0x9c, 0xcd, 0x8c, 0x93, 0x99, 0x42, 0xc1, 0xdf, 0xe1, 0xdf, 0xf0, 0xcd, 0xdf,
	0x27, 0x49, 0x93, 0x36, 0x0d, 0x86, 0x7d, 0x9b, 0x3b, 0xe7, 0x9c, 0x9c, 0x93, 0x39, 0x5c, 0x78,
	0x55, 0x60, 0xea, 0x24, 0x17, 0xb9, 0x4c, 0x32, 0x65, 0x30, 0x11, 0x85, 0x45, 0x73, 0xc3, 0x33,
	0x8c, 0xb5, 0x51, 0x56, 0xd1, 0xd9, 0x1e, 0x8f, 0x2b, 0x3c, 0x08, 0x7a, 0x7c, 0xbb, 0xd5, 0x58,
	0xee, 0xb8, 0xc1, 0xcb, 0x1e, 0x26, 0x85, 0x45, 0xc3, 0x65, 0x03, 0x47, 0xdf, 0x61, 0x72, 0xcd,
	0x8d, 0xe0, 0xa9, 0x44, 0x9a, 0xc0, 0x83, 0x4a, 0xe9, 0x93, 0x90, 0xcc, 0xbd, 0xc5, 0xf3, 0xf8,
	0xd8, 0x25, 0xbe, 0xda, 0x29, 0xbf, 0x6c, 0x35, 0xb2, 0x9a, 0x48, 0x43, 0xf0, 0x72, 0x2c, 0x33,
	0x23, 0xb4, 0x15, 0xaa, 0xf0, 0x47, 0x21, 0x99, 0x4f, 0x59, 0xf7, 0x2a, 0xfa, 0x43, 0xc0, 0x6b,
	0xbf, 0xff, 0x89, 0x6b, 0xba, 0x82, 0xe9, 0xa6, 0x19, 0x4b, 0x9f, 0x84, 0xe3, 0xb9, 0xb7, 0x38,
	0xef, 0xfb, 0x74, 0xf8, 0xfb, 0x73, 0xf9, 0xa1, 0xb0, 0x66, 0xcb, 0x0e, 0xe2, 0xe0, 0x1a, 0x66,
	0xc7, 0x20, 0x7d, 0x02, 0xe3, 0x1f, 0xb8, 0xad, 0xd3, 0x4f, 0x59, 0x75, 0xa4, 0x31, 0x3c, 0xdc,
	0x70, 0xe9, 0xb0, 0x4e, 0xe6, 0x2d, 0xfc, 0x21, 0x27, 0xb6, 0xa3, 0x5d, 0x8c, 0xde, 0x91, 0xe8,
	0x17, 0xcc, 0xaa, 0x3f, 0xcc, 0x3f, 0xb6, 0x6f, 0x4e, 0x97, 0x70, 0x2a, 0x0a, 0xed, 0x6c, 0x39,
	0xf4, 0x30, 0x9d, 0xc0, 0xac, 0xa1, 0xd2, 0xb7, 0x70, 0xa6, 0x9c, 0xad, 0x55, 0xa3, 0xfb, 0x55,
	0x2d, 0x37, 0xfa, 0x4d, 0x60, 0xfa, 0x99, 0x1b, 0x7e, 0x87, 0x16, 0x0d, 0x3d, 0x87, 0xf1, 0x86,
	0x9b, 0xc6, 0x76, 0x38, 0x7d, 0x45, 0xa2, 0x4b, 0x38, 0xcb, 0xf1, 0x86, 0x3b, 0x69, 0x1b, 0xc3,
	0xa7, 0x03, 0xfd, 0xad, 0x4e, 0x58, 0xcb, 0xa4, 0x2f, 0x60, 0x62, 0xf0, 0xa7, 0x13, 0x06, 0x73,
	0x7f, 0x1c, 0x92, 0xf9, 0x64, 0x75, 0xc2, 0xf6, 0x37, 0x97, 0x00, 0x93, 0x14, 0x6f, 0xf9, 0x46,
	0x28, 0x13, 0xfd, 0x25, 0xf0, 0x68, 0x1f, 0xac, 0x6a, 0xf2, 0x0a, 0x40, 0xb7, 0x73, 0x5b, 0xe5,
	0xeb, 0xbe, 0x65, 0x57, 0x71, 0x18, 0x9a, 0x32, 0x3b, 0xfa, 0xe0, 0x2b, 0x3c, 0xee, 0xc1, 0xff,
	0xa9, 0x33, 0x39, 0xae, 0xf3, 0xd9, 0xa0, 0x5b, 0xa7, 0xcf, 0xcb, 0xf7, 0xdf, 0x2e, 0xd6, 0xc2,
	0xde, 0xba, 0x34, 0xce, 0xd4, 0x5d, 0xb2, 0x53, 0x64, 0x52, 0xb9, 0xbc, 0x4c, 0x0e, 0x9b, 0xb1,
	0xc6, 0x22, 0xd1, 0xe9, 0x9b, 0xb5, 0x4a, 0x8e, 0xb7, 0x25, 0x3d, 0xad, 0xb7, 0x64, 0xf9, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0xc6, 0xcc, 0xa2, 0x53, 0x92, 0x03, 0x00, 0x00,
}
