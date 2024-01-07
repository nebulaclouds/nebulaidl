// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/service/identity.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoRequest) Reset()         { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()    {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_79a484abbd97cb78, []int{0}
}

func (m *UserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRequest.Unmarshal(m, b)
}
func (m *UserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRequest.Marshal(b, m, deterministic)
}
func (m *UserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRequest.Merge(m, src)
}
func (m *UserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoRequest.Size(m)
}
func (m *UserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRequest proto.InternalMessageInfo

// See the OpenID Connect spec at https://openid.net/specs/openid-connect-core-1_0.html#UserInfoResponse for more information.
type UserInfoResponse struct {
	// Locally unique and never reassigned identifier within the Issuer for the End-User, which is intended to be consumed
	// by the Client.
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// Full name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Shorthand name by which the End-User wishes to be referred to
	PreferredUsername string `protobuf:"bytes,3,opt,name=preferred_username,json=preferredUsername,proto3" json:"preferred_username,omitempty"`
	// Given name(s) or first name(s)
	GivenName string `protobuf:"bytes,4,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	// Surname(s) or last name(s)
	FamilyName string `protobuf:"bytes,5,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	// Preferred e-mail address
	Email string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	// Profile picture URL
	Picture string `protobuf:"bytes,7,opt,name=picture,proto3" json:"picture,omitempty"`
	// Additional claims
	AdditionalClaims     *_struct.Struct `protobuf:"bytes,8,opt,name=additional_claims,json=additionalClaims,proto3" json:"additional_claims,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_79a484abbd97cb78, []int{1}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *UserInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfoResponse) GetPreferredUsername() string {
	if m != nil {
		return m.PreferredUsername
	}
	return ""
}

func (m *UserInfoResponse) GetGivenName() string {
	if m != nil {
		return m.GivenName
	}
	return ""
}

func (m *UserInfoResponse) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *UserInfoResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfoResponse) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *UserInfoResponse) GetAdditionalClaims() *_struct.Struct {
	if m != nil {
		return m.AdditionalClaims
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfoRequest)(nil), "nebulaidl.service.UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "nebulaidl.service.UserInfoResponse")
}

func init() { proto.RegisterFile("nebulaidl/service/identity.proto", fileDescriptor_79a484abbd97cb78) }

var fileDescriptor_79a484abbd97cb78 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x8a, 0xdb, 0x30,
	0x10, 0xc7, 0xc9, 0x77, 0xa2, 0x1c, 0x12, 0x8b, 0x42, 0x4d, 0x68, 0x69, 0x70, 0x2f, 0xb9, 0x44,
	0x82, 0xf4, 0x5c, 0x0a, 0x6d, 0x2f, 0xb9, 0xf4, 0x90, 0x90, 0x4b, 0x2f, 0xa9, 0x6c, 0x8f, 0x5d,
	0x2d, 0xb2, 0xe4, 0xd5, 0x47, 0x20, 0xd7, 0x7d, 0x85, 0x7d, 0x86, 0x7d, 0xa2, 0x7d, 0x85, 0x7d,
	0x90, 0x25, 0x92, 0x93, 0xc0, 0x06, 0xf6, 0xe6, 0x99, 0xdf, 0x6f, 0x40, 0xfe, 0xcf, 0xa0, 0xb9,
	0x84, 0xd4, 0x09, 0xc6, 0x73, 0x41, 0x0d, 0xe8, 0x03, 0xcf, 0x80, 0xf2, 0x1c, 0xa4, 0xe5, 0xf6,
	0x48, 0x6a, 0xad, 0xac, 0xc2, 0xd1, 0xc5, 0x20, 0x8d, 0x31, 0xfb, 0x54, 0x2a, 0x55, 0x0a, 0xa0,
	0xac, 0xe6, 0x94, 0x49, 0xa9, 0x2c, 0xb3, 0x5c, 0x49, 0x13, 0x06, 0x2e, 0xd4, 0x57, 0xa9, 0x2b,
	0xa8, 0xb1, 0xda, 0x65, 0x36, 0xd0, 0x24, 0x42, 0x93, 0x9d, 0x01, 0xbd, 0x96, 0x85, 0xda, 0xc0,
	0xbd, 0x03, 0x63, 0x93, 0xa7, 0x36, 0x9a, 0x5e, 0x7b, 0xa6, 0x56, 0xd2, 0x00, 0x8e, 0xd1, 0xc0,
	0xb8, 0xf4, 0x0e, 0x32, 0x1b, 0xb7, 0xe6, 0xad, 0xc5, 0x68, 0x73, 0x2e, 0x31, 0x46, 0x5d, 0xc9,
	0x2a, 0x88, 0xdb, 0xbe, 0xed, 0xbf, 0xf1, 0x12, 0xe1, 0x5a, 0x43, 0x01, 0x5a, 0x43, 0xbe, 0x77,
	0x06, 0xb4, 0x37, 0x3a, 0xde, 0x88, 0x2e, 0x64, 0xd7, 0x00, 0xfc, 0x19, 0xa1, 0x92, 0x1f, 0x40,
	0xee, 0xbd, 0xd6, 0xf5, 0xda, 0xc8, 0x77, 0xfe, 0x9c, 0xf0, 0x17, 0x34, 0x2e, 0x58, 0xc5, 0xc5,
	0x31, 0xf0, 0x9e, 0xe7, 0x28, 0xb4, 0xbc, 0xf0, 0x01, 0xf5, 0xa0, 0x62, 0x5c, 0xc4, 0x7d, 0x8f,
	0x42, 0x71, 0x7a, 0x72, 0xcd, 0x33, 0xeb, 0x34, 0xc4, 0x83, 0xf0, 0xe4, 0xa6, 0xc4, 0xbf, 0x51,
	0xc4, 0xf2, 0x9c, 0x9f, 0x52, 0x62, 0x62, 0x9f, 0x09, 0xc6, 0x2b, 0x13, 0x0f, 0xe7, 0xad, 0xc5,
	0x78, 0xf5, 0x91, 0x84, 0xb8, 0xc8, 0x39, 0x2e, 0xb2, 0xf5, 0x71, 0x6d, 0xa6, 0xd7, 0x89, 0x5f,
	0x7e, 0x60, 0x65, 0xd0, 0x64, 0xdd, 0xec, 0x66, 0x1b, 0x36, 0x81, 0xff, 0xa1, 0xe1, 0x39, 0x39,
	0x9c, 0x90, 0x9b, 0x4d, 0x91, 0x37, 0x51, 0xcf, 0xbe, 0xbe, 0xeb, 0x84, 0xe8, 0x93, 0xf1, 0xc3,
	0xf3, 0xcb, 0x63, 0xbb, 0x87, 0x3b, 0xb4, 0x82, 0x9f, 0x3f, 0xfe, 0x7e, 0x2f, 0xb9, 0xfd, 0xef,
	0x52, 0x92, 0xa9, 0x8a, 0x86, 0xe9, 0x4c, 0x28, 0x97, 0x1b, 0x7a, 0x3d, 0x9d, 0x12, 0x24, 0xad,
	0xd3, 0x65, 0xa9, 0xe8, 0xcd, 0x39, 0xa5, 0x7d, 0xff, 0x63, 0xdf, 0x5e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x22, 0x02, 0xb1, 0xb2, 0x6a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityServiceClient interface {
	// Retrieves user information about the currently logged in user.
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
}

type identityServiceClient struct {
	cc *grpc.ClientConn
}

func NewIdentityServiceClient(cc *grpc.ClientConn) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/nebulaidl.service.IdentityService/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
type IdentityServiceServer interface {
	// Retrieves user information about the currently logged in user.
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
}

// UnimplementedIdentityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (*UnimplementedIdentityServiceServer) UserInfo(ctx context.Context, req *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}

func RegisterIdentityServiceServer(s *grpc.Server, srv IdentityServiceServer) {
	s.RegisterService(&_IdentityService_serviceDesc, srv)
}

func _IdentityService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebulaidl.service.IdentityService/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nebulaidl.service.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserInfo",
			Handler:    _IdentityService_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebulaidl/service/identity.proto",
}
