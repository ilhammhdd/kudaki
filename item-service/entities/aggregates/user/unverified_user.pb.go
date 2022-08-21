// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/user/unverified_user.proto

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

type UnverifiedUser struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *User                `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UnverifiedUser) Reset()         { *m = UnverifiedUser{} }
func (m *UnverifiedUser) String() string { return proto.CompactTextString(m) }
func (*UnverifiedUser) ProtoMessage()    {}
func (*UnverifiedUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66d3b9a6d578134, []int{0}
}

func (m *UnverifiedUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnverifiedUser.Unmarshal(m, b)
}
func (m *UnverifiedUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnverifiedUser.Marshal(b, m, deterministic)
}
func (m *UnverifiedUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnverifiedUser.Merge(m, src)
}
func (m *UnverifiedUser) XXX_Size() int {
	return xxx_messageInfo_UnverifiedUser.Size(m)
}
func (m *UnverifiedUser) XXX_DiscardUnknown() {
	xxx_messageInfo_UnverifiedUser.DiscardUnknown(m)
}

var xxx_messageInfo_UnverifiedUser proto.InternalMessageInfo

func (m *UnverifiedUser) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UnverifiedUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UnverifiedUser) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*UnverifiedUser)(nil), "aggregates.user.UnverifiedUser")
}

func init() {
	proto.RegisterFile("aggregates/user/unverified_user.proto", fileDescriptor_d66d3b9a6d578134)
}

var fileDescriptor_d66d3b9a6d578134 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x2a, 0x82, 0x2b, 0x54, 0x08, 0x08, 0x21, 0x17, 0x8b, 0x20, 0xd4, 0x43, 0x67,
	0x40, 0x4f, 0x1e, 0xab, 0x6f, 0x10, 0xec, 0xc5, 0x4b, 0xd9, 0x64, 0xa7, 0x9b, 0xa1, 0xdd, 0x6e,
	0xd9, 0x9d, 0xed, 0x23, 0xf8, 0xdc, 0x92, 0xd4, 0x56, 0x88, 0xc7, 0x7f, 0xf9, 0xfe, 0xfd, 0xbf,
	0x51, 0x4f, 0xda, 0xda, 0x40, 0x56, 0x0b, 0x45, 0x4c, 0x91, 0x02, 0xa6, 0xfd, 0x91, 0x02, 0x6f,
	0x98, 0xcc, 0xba, 0xcf, 0x70, 0x08, 0x5e, 0x7c, 0x71, 0xf7, 0x87, 0x41, 0xff, 0x5c, 0x55, 0xff,
	0x7a, 0x17, 0xb8, 0x7a, 0xb0, 0xde, 0xdb, 0x1d, 0xe1, 0x90, 0x9a, 0xb4, 0x41, 0x61, 0x47, 0x51,
	0xb4, 0x3b, 0x9c, 0x80, 0xc7, 0xef, 0x4c, 0x4d, 0x57, 0x97, 0x9d, 0x55, 0xa4, 0x50, 0x4c, 0x55,
	0xce, 0xa6, 0xcc, 0x66, 0xd9, 0x7c, 0x52, 0xe7, 0x6c, 0x8a, 0x67, 0x75, 0xd5, 0xff, 0x58, 0xe6,
	0xb3, 0x6c, 0x7e, 0xfb, 0x72, 0x0f, 0xa3, 0x7d, 0xe8, 0x4b, 0xf5, 0x80, 0x14, 0x6f, 0x4a, 0xb5,
	0x81, 0xb4, 0x90, 0x59, 0x6b, 0x29, 0x27, 0x43, 0xa1, 0x82, 0x93, 0x03, 0x9c, 0x1d, 0xe0, 0xf3,
	0xec, 0x50, 0xdf, 0xfc, 0xd2, 0x4b, 0x79, 0xff, 0xf8, 0x5a, 0x5a, 0x96, 0x2e, 0x35, 0xd0, 0x7a,
	0x87, 0xbc, 0xeb, 0xb4, 0x73, 0x9d, 0x31, 0xb8, 0x4d, 0x46, 0x6f, 0x79, 0xc1, 0x42, 0x6e, 0x11,
	0x29, 0x1c, 0xb9, 0x25, 0xa4, 0xbd, 0xb0, 0x30, 0x45, 0x1c, 0x5d, 0xde, 0x5c, 0x0f, 0x1b, 0xaf,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0x1e, 0x3b, 0xe9, 0x4b, 0x01, 0x00, 0x00,
}
