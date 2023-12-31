// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/admin/notification.proto

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

// Represents the Email object that is sent to a publisher/subscriber
// to forward the notification.
// Note: This is internal to Admin and doesn't need to be exposed to other components.
type EmailMessage struct {
	// The list of email addresses to receive an email with the content populated in the other fields.
	// Currently, each email recipient will receive its own email.
	// This populates the TO field.
	RecipientsEmail []string `protobuf:"bytes,1,rep,name=recipients_email,json=recipientsEmail,proto3" json:"recipients_email,omitempty"`
	// The email of the sender.
	// This populates the FROM field.
	SenderEmail string `protobuf:"bytes,2,opt,name=sender_email,json=senderEmail,proto3" json:"sender_email,omitempty"`
	// The content of the subject line.
	// This populates the SUBJECT field.
	SubjectLine string `protobuf:"bytes,3,opt,name=subject_line,json=subjectLine,proto3" json:"subject_line,omitempty"`
	// The content of the email body.
	// This populates the BODY field.
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailMessage) Reset()         { *m = EmailMessage{} }
func (m *EmailMessage) String() string { return proto.CompactTextString(m) }
func (*EmailMessage) ProtoMessage()    {}
func (*EmailMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ae0b0f9890c99ab, []int{0}
}

func (m *EmailMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailMessage.Unmarshal(m, b)
}
func (m *EmailMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailMessage.Marshal(b, m, deterministic)
}
func (m *EmailMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailMessage.Merge(m, src)
}
func (m *EmailMessage) XXX_Size() int {
	return xxx_messageInfo_EmailMessage.Size(m)
}
func (m *EmailMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EmailMessage proto.InternalMessageInfo

func (m *EmailMessage) GetRecipientsEmail() []string {
	if m != nil {
		return m.RecipientsEmail
	}
	return nil
}

func (m *EmailMessage) GetSenderEmail() string {
	if m != nil {
		return m.SenderEmail
	}
	return ""
}

func (m *EmailMessage) GetSubjectLine() string {
	if m != nil {
		return m.SubjectLine
	}
	return ""
}

func (m *EmailMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*EmailMessage)(nil), "nebulaidl.admin.EmailMessage")
}

func init() { proto.RegisterFile("nebulaidl/admin/notification.proto", fileDescriptor_7ae0b0f9890c99ab) }

var fileDescriptor_7ae0b0f9890c99ab = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x3f, 0x4b, 0xc5, 0x30,
	0x14, 0x47, 0xa9, 0xef, 0x21, 0xbc, 0xf8, 0xe0, 0x49, 0xa6, 0x8e, 0xcf, 0x4e, 0x75, 0xb0, 0x19,
	0x1c, 0xc5, 0x45, 0x70, 0xd3, 0xa5, 0xa3, 0x4b, 0xc9, 0x9f, 0x6b, 0xbc, 0x92, 0xde, 0x5b, 0x9a,
	0x64, 0xf0, 0x73, 0xf8, 0x85, 0xc5, 0xb4, 0x58, 0x70, 0x0b, 0xe7, 0x77, 0x02, 0xf7, 0x88, 0x86,
	0xc0, 0xe4, 0xa0, 0xd1, 0x05, 0xa5, 0xdd, 0x88, 0xa4, 0x88, 0x13, 0xbe, 0xa3, 0xd5, 0x09, 0x99,
	0xba, 0x69, 0xe6, 0xc4, 0xf2, 0xf4, 0xe7, 0x74, 0xc5, 0x69, 0xbe, 0x2b, 0x71, 0x7c, 0x1e, 0x35,
	0x86, 0x57, 0x88, 0x51, 0x7b, 0x90, 0xb7, 0xe2, 0x7a, 0x06, 0x8b, 0x13, 0x02, 0xa5, 0x38, 0xc0,
	0xef, 0x54, 0x57, 0xe7, 0x5d, 0x7b, 0xe8, 0x4f, 0x1b, 0x2f, 0x3f, 0xe4, 0x8d, 0x38, 0x46, 0x20,
	0x07, 0xf3, 0xaa, 0x5d, 0x9c, 0xab, 0xf6, 0xd0, 0x5f, 0x2d, 0x6c, 0x53, 0xb2, 0xf9, 0x04, 0x9b,
	0x86, 0x80, 0x04, 0xf5, 0x6e, 0x55, 0x16, 0xf6, 0x82, 0x04, 0x52, 0x8a, 0xbd, 0x61, 0xf7, 0x55,
	0xef, 0xcb, 0x54, 0xde, 0x4f, 0x8f, 0x6f, 0x0f, 0x1e, 0xd3, 0x47, 0x36, 0x9d, 0xe5, 0x51, 0x2d,
	0x37, 0xdb, 0xc0, 0xd9, 0x45, 0xb5, 0x45, 0x7a, 0x20, 0x35, 0x99, 0x3b, 0xcf, 0xea, 0x5f, 0xb8,
	0xb9, 0x2c, 0xb1, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x2f, 0xfc, 0xb7, 0x12, 0x01,
	0x00, 0x00,
}
