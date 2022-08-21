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
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0x49, 0x67, 0x14, 0x26, 0xa2, 0x42, 0x50, 0x28, 0xdd, 0x58, 0xc4, 0x45, 0x5d, 0x4c,
	0x02, 0xba, 0x72, 0xa9, 0xe2, 0x5e, 0x82, 0x6e, 0xdc, 0x0c, 0x69, 0xf3, 0x4c, 0x43, 0x9b, 0xa6,
	0x24, 0xaf, 0xfa, 0x31, 0xfe, 0xac, 0xb4, 0x9d, 0x61, 0xa0, 0x9b, 0xc0, 0x85, 0x93, 0x7b, 0x2e,
	0x8f, 0xde, 0x29, 0x63, 0x02, 0x18, 0x85, 0x10, 0xc5, 0x10, 0x21, 0x88, 0x00, 0x11, 0x70, 0xd7,
	0xab, 0x18, 0x7f, 0x7d, 0xd0, 0xbc, 0x0f, 0x1e, 0x3d, 0xbb, 0x3c, 0x52, 0x7c, 0xa4, 0xb2, 0x6c,
	0xf9, 0x6d, 0x7c, 0x66, 0x38, 0xbb, 0x31, 0xde, 0x9b, 0x16, 0xc4, 0x94, 0xca, 0xe1, 0x5b, 0xa0,
	0x75, 0x10, 0x51, 0xb9, 0x7e, 0x06, 0x6e, 0xff, 0x08, 0x3d, 0x97, 0xa3, 0xe6, 0x7d, 0x6f, 0x61,
	0x17, 0x34, 0xb1, 0x3a, 0x25, 0x39, 0x29, 0x56, 0x32, 0xb1, 0x9a, 0xdd, 0xd3, 0xf5, 0x58, 0x98,
	0x26, 0x39, 0x29, 0xce, 0x1e, 0xae, 0xf9, 0x42, 0xcf, 0x3f, 0x23, 0x04, 0x39, 0x21, 0xec, 0x8a,
	0x9e, 0xa0, 0x6f, 0xa0, 0x4b, 0x57, 0x39, 0x29, 0x36, 0x72, 0x0e, 0xec, 0x89, 0xd2, 0x2a, 0x80,
	0x42, 0xd0, 0x3b, 0x85, 0xe9, 0x7a, 0xaa, 0xc9, 0xf8, 0x3c, 0x8c, 0x1f, 0x86, 0xf1, 0x8f, 0xc3,
	0x30, 0xb9, 0xd9, 0xd3, 0xcf, 0xf8, 0xf2, 0xf6, 0xf5, 0x6a, 0x2c, 0xd6, 0x43, 0xc9, 0x2b, 0xef,
	0x84, 0x6d, 0x6b, 0xe5, 0x5c, 0xad, 0xb5, 0x68, 0x06, 0xad, 0x1a, 0xbb, 0x0d, 0xd0, 0xa1, 0x6a,
	0xb7, 0x11, 0xc2, 0x8f, 0xad, 0x40, 0x40, 0x87, 0x16, 0x2d, 0x44, 0xb1, 0x38, 0x48, 0x79, 0x3a,
	0x59, 0x1e, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xe9, 0x98, 0xd1, 0x61, 0x01, 0x00, 0x00,
}
