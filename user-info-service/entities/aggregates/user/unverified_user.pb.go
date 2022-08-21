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
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x2a, 0x82, 0x2b, 0x54, 0x08, 0x08, 0x21, 0x17, 0x8b, 0x20, 0xd4, 0x43, 0x76,
	0x40, 0x4f, 0x1e, 0xf5, 0x20, 0x9e, 0x83, 0xbd, 0x78, 0x29, 0x9b, 0xec, 0x64, 0x33, 0xb4, 0x9b,
	0x2d, 0xbb, 0xb3, 0x7d, 0x04, 0x9f, 0x5b, 0x36, 0xb5, 0x15, 0xea, 0xf1, 0x1b, 0x7e, 0xdf, 0x9f,
	0x11, 0x0f, 0xca, 0x18, 0x8f, 0x46, 0x31, 0x06, 0x88, 0x01, 0x3d, 0xc4, 0x71, 0x8f, 0x9e, 0x7a,
	0x42, 0xbd, 0x4e, 0x5a, 0xee, 0xbc, 0x63, 0x57, 0xdc, 0xfc, 0x61, 0x32, 0x9d, 0xab, 0xea, 0x9f,
	0xef, 0x04, 0x57, 0x77, 0xc6, 0x39, 0xb3, 0x45, 0x98, 0x54, 0x1b, 0x7b, 0x60, 0xb2, 0x18, 0x58,
	0xd9, 0xdd, 0x01, 0xb8, 0xff, 0xce, 0xc4, 0x7c, 0x75, 0xea, 0x59, 0x05, 0xf4, 0xc5, 0x5c, 0xe4,
	0xa4, 0xcb, 0x6c, 0x91, 0x2d, 0x67, 0x4d, 0x4e, 0xba, 0x78, 0x14, 0x17, 0x29, 0xb1, 0xcc, 0x17,
	0xd9, 0xf2, 0xfa, 0xe9, 0x56, 0x9e, 0xf5, 0xcb, 0x64, 0x6a, 0x26, 0xa4, 0x78, 0x11, 0xa2, 0xf3,
	0xa8, 0x18, 0xf5, 0x5a, 0x71, 0x39, 0x9b, 0x0c, 0x95, 0x3c, 0x6c, 0x90, 0xc7, 0x0d, 0xf2, 0xf3,
	0xb8, 0xa1, 0xb9, 0xfa, 0xa5, 0x5f, 0xf9, 0xed, 0xe3, 0xeb, 0xdd, 0x10, 0x0f, 0xb1, 0x95, 0x9d,
	0xb3, 0x40, 0xdb, 0x41, 0x59, 0x3b, 0x68, 0x0d, 0x9b, 0xa8, 0xd5, 0x86, 0xea, 0x14, 0x5f, 0xd3,
	0xd8, 0xbb, 0x3a, 0xa0, 0xdf, 0x53, 0x87, 0x80, 0x23, 0x13, 0x13, 0x06, 0x38, 0x7b, 0xbf, 0xbd,
	0x9c, 0x8a, 0x9e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0x66, 0xd3, 0x1c, 0x50, 0x01, 0x00,
	0x00,
}
