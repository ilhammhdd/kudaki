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
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x69, 0x77, 0x15, 0x36, 0xa2, 0x42, 0x51, 0x28, 0xbd, 0x58, 0xc4, 0x43, 0x3d, 0x34,
	0x01, 0x3d, 0x79, 0xd4, 0x83, 0x78, 0x94, 0xa2, 0x17, 0x2f, 0x4b, 0xda, 0xcc, 0xa6, 0x43, 0x37,
	0x4d, 0x49, 0xa6, 0xfa, 0x63, 0xfc, 0xb3, 0x92, 0x76, 0x17, 0xa1, 0x97, 0xc0, 0x83, 0x2f, 0xef,
	0x7b, 0x0c, 0xbb, 0x93, 0x5a, 0x3b, 0xd0, 0x92, 0xc0, 0x8b, 0xd1, 0x83, 0x13, 0x0e, 0x3c, 0xd0,
	0x76, 0x90, 0xde, 0xff, 0x58, 0xa7, 0xf8, 0xe0, 0x2c, 0xd9, 0xe4, 0xf2, 0x9f, 0xe2, 0x81, 0xca,
	0xb2, 0xe5, 0xb7, 0xf0, 0xcc, 0x70, 0x76, 0xa3, 0xad, 0xd5, 0x7b, 0x10, 0x53, 0xaa, 0xc7, 0x9d,
	0x20, 0x34, 0xe0, 0x49, 0x9a, 0x61, 0x06, 0x6e, 0x7f, 0x23, 0x76, 0x5e, 0x05, 0xcd, 0xfb, 0xc1,
	0x92, 0x5c, 0xb0, 0x18, 0x55, 0x1a, 0xe5, 0x51, 0xb1, 0xaa, 0x62, 0x54, 0xc9, 0x3d, 0x5b, 0x87,
	0xc2, 0x34, 0xce, 0xa3, 0xe2, 0xec, 0xe1, 0x9a, 0x2f, 0xf4, 0xfc, 0xd3, 0x83, 0xab, 0x26, 0x24,
	0xb9, 0x62, 0x27, 0x64, 0x3b, 0xe8, 0xd3, 0x55, 0x1e, 0x15, 0x9b, 0x6a, 0x0e, 0xc9, 0x13, 0x63,
	0x8d, 0x03, 0x49, 0xa0, 0xb6, 0x92, 0xd2, 0xf5, 0x54, 0x93, 0xf1, 0x79, 0x18, 0x3f, 0x0e, 0xe3,
	0x1f, 0xc7, 0x61, 0xd5, 0xe6, 0x40, 0x3f, 0xd3, 0xcb, 0xdb, 0xd7, 0xab, 0x46, 0x6a, 0xc7, 0x9a,
	0x37, 0xd6, 0x08, 0xdc, 0xb7, 0xd2, 0x98, 0x56, 0x29, 0xd1, 0x8d, 0x4a, 0x76, 0x58, 0x06, 0x69,
	0x89, 0xfd, 0xce, 0x96, 0x1e, 0xdc, 0x37, 0x36, 0x20, 0xa0, 0x27, 0x24, 0x04, 0x2f, 0x16, 0x37,
	0xa9, 0x4f, 0x27, 0xd1, 0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0x4b, 0x54, 0x35, 0x64,
	0x01, 0x00, 0x00,
}