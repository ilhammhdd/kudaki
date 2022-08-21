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
	0x14, 0x45, 0xe9, 0xcc, 0x28, 0x4c, 0x44, 0x85, 0xa0, 0x50, 0xba, 0xb1, 0x88, 0x8b, 0xba, 0x98,
	0x04, 0x74, 0xe5, 0x72, 0xf4, 0x07, 0x24, 0xe8, 0xc6, 0xcd, 0x90, 0x36, 0x6f, 0xd2, 0x47, 0x9b,
	0x49, 0x49, 0x5e, 0xf5, 0x63, 0xfc, 0x59, 0x69, 0x3b, 0x83, 0xd0, 0x4d, 0xe0, 0xc2, 0xc9, 0x3d,
	0x97, 0xc7, 0x1e, 0xb4, 0xb5, 0x01, 0xac, 0x26, 0x88, 0xb2, 0x8f, 0x10, 0x64, 0x80, 0x08, 0xb4,
	0xeb, 0x74, 0x8c, 0x3f, 0x3e, 0x18, 0xd1, 0x05, 0x4f, 0x9e, 0x5f, 0xff, 0x53, 0x62, 0xa0, 0xb2,
	0x6c, 0xfe, 0x6d, 0x78, 0x26, 0x38, 0xbb, 0xb3, 0xde, 0xdb, 0x16, 0xe4, 0x98, 0xca, 0x7e, 0x2f,
	0x09, 0x1d, 0x44, 0xd2, 0xae, 0x9b, 0x80, 0xfb, 0xdf, 0x84, 0x5d, 0xaa, 0x41, 0xf3, 0x7e, 0xb4,
	0xf0, 0x2b, 0xb6, 0x40, 0x93, 0x26, 0x79, 0x52, 0x2c, 0xd5, 0x02, 0x0d, 0x7f, 0x64, 0xab, 0xa1,
	0x30, 0x5d, 0xe4, 0x49, 0x71, 0xf1, 0x74, 0x2b, 0x66, 0x7a, 0xf1, 0x19, 0x21, 0xa8, 0x11, 0xe1,
	0x37, 0xec, 0x8c, 0x7c, 0x03, 0x87, 0x74, 0x99, 0x27, 0xc5, 0x5a, 0x4d, 0x81, 0xbf, 0x30, 0x56,
	0x05, 0xd0, 0x04, 0x66, 0xa7, 0x29, 0x5d, 0x8d, 0x35, 0x99, 0x98, 0x86, 0x89, 0xd3, 0x30, 0xf1,
	0x71, 0x1a, 0xa6, 0xd6, 0x47, 0x7a, 0x4b, 0xaf, 0x6f, 0x5f, 0x5b, 0x8b, 0x54, 0xf7, 0xa5, 0xa8,
	0xbc, 0x93, 0xd8, 0xd6, 0xda, 0xb9, 0xda, 0x18, 0xd9, 0xf4, 0x46, 0x37, 0xb8, 0xd9, 0x63, 0x0b,
	0x9b, 0x08, 0xe1, 0x1b, 0x2b, 0x90, 0x70, 0x20, 0x24, 0x84, 0x28, 0x67, 0xe7, 0x28, 0xcf, 0x47,
	0xc7, 0xf3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xf0, 0x29, 0xf1, 0x5f, 0x01, 0x00, 0x00,
}