// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/user/reset_password.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ResetPassword struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *User                `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Token                string               `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ResetPassword) Reset()         { *m = ResetPassword{} }
func (m *ResetPassword) String() string { return proto.CompactTextString(m) }
func (*ResetPassword) ProtoMessage()    {}
func (*ResetPassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_174109f3c45dea44, []int{0}
}

func (m *ResetPassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPassword.Unmarshal(m, b)
}
func (m *ResetPassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPassword.Marshal(b, m, deterministic)
}
func (m *ResetPassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPassword.Merge(m, src)
}
func (m *ResetPassword) XXX_Size() int {
	return xxx_messageInfo_ResetPassword.Size(m)
}
func (m *ResetPassword) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPassword.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPassword proto.InternalMessageInfo

func (m *ResetPassword) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ResetPassword) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ResetPassword) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ResetPassword) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ResetPassword)(nil), "aggregates.user.ResetPassword")
}

func init() {
	proto.RegisterFile("aggregates/user/reset_password.proto", fileDescriptor_174109f3c45dea44)
}

var fileDescriptor_174109f3c45dea44 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0x69, 0x67, 0x14, 0x26, 0xa2, 0x42, 0x50, 0x28, 0xdd, 0x58, 0xc4, 0x45, 0x5d, 0x4c,
	0x02, 0xba, 0x72, 0x39, 0xfa, 0x03, 0x12, 0x74, 0xe3, 0x66, 0x48, 0x9b, 0x67, 0xfa, 0xe8, 0x64,
	0x52, 0x92, 0x57, 0xfd, 0x18, 0x7f, 0x56, 0xda, 0xce, 0x30, 0xd0, 0x4d, 0xe0, 0xc2, 0xc9, 0x3d,
	0x97, 0xc7, 0x1e, 0xb4, 0xb5, 0x01, 0xac, 0x26, 0x88, 0xb2, 0x8f, 0x10, 0x64, 0x80, 0x08, 0xb4,
	0xed, 0x74, 0x8c, 0xbf, 0x3e, 0x18, 0xd1, 0x05, 0x4f, 0x9e, 0x5f, 0x9f, 0x28, 0x31, 0x50, 0x79,
	0x3e, 0xff, 0x36, 0x3c, 0x13, 0x9c, 0xdf, 0x59, 0xef, 0xed, 0x0e, 0xe4, 0x98, 0xaa, 0xfe, 0x5b,
	0x12, 0x3a, 0x88, 0xa4, 0x5d, 0x37, 0x01, 0xf7, 0x7f, 0x09, 0xbb, 0x54, 0x83, 0xe6, 0xfd, 0x60,
	0xe1, 0x57, 0x2c, 0x45, 0x93, 0x25, 0x45, 0x52, 0x2e, 0x54, 0x8a, 0x86, 0x3f, 0xb2, 0xe5, 0x50,
	0x98, 0xa5, 0x45, 0x52, 0x5e, 0x3c, 0xdd, 0x8a, 0x99, 0x5e, 0x7c, 0x46, 0x08, 0x6a, 0x44, 0xf8,
	0x0d, 0x3b, 0x23, 0xdf, 0xc2, 0x3e, 0x5b, 0x14, 0x49, 0xb9, 0x52, 0x53, 0xe0, 0x2f, 0x8c, 0xd5,
	0x01, 0x34, 0x81, 0xd9, 0x6a, 0xca, 0x96, 0x63, 0x4d, 0x2e, 0xa6, 0x61, 0xe2, 0x38, 0x4c, 0x7c,
	0x1c, 0x87, 0xa9, 0xd5, 0x81, 0xde, 0xd0, 0xeb, 0xdb, 0xd7, 0xc6, 0x22, 0x35, 0x7d, 0x25, 0x6a,
	0xef, 0x24, 0xee, 0x1a, 0xed, 0x5c, 0x63, 0x8c, 0x6c, 0x7b, 0xa3, 0x5b, 0x5c, 0x23, 0x81, 0x5b,
	0x47, 0x08, 0x3f, 0x58, 0x83, 0x84, 0x3d, 0x21, 0x21, 0x44, 0x39, 0x3b, 0x47, 0x75, 0x3e, 0x3a,
	0x9e, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x18, 0x9b, 0xe3, 0xed, 0x5f, 0x01, 0x00, 0x00,
}