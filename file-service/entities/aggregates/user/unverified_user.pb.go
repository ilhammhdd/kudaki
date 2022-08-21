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
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x69, 0x27, 0x82, 0x11, 0x26, 0x14, 0x84, 0xd2, 0x1b, 0x87, 0x20, 0xcc, 0x8b, 0xe5,
	0x80, 0x5e, 0x79, 0x39, 0x7d, 0x83, 0xe2, 0x6e, 0xbc, 0x19, 0x69, 0x73, 0x9a, 0x1e, 0xd6, 0x2c,
	0x23, 0x39, 0xd9, 0x23, 0xf8, 0xdc, 0xd2, 0xcc, 0x4d, 0xa8, 0x97, 0x7f, 0xf8, 0xfe, 0xfc, 0xdf,
	0x11, 0x4f, 0xca, 0x18, 0x8f, 0x46, 0x31, 0x06, 0x88, 0x01, 0x3d, 0xc4, 0xfd, 0x11, 0x3d, 0x75,
	0x84, 0x7a, 0x3b, 0x66, 0x79, 0xf0, 0x8e, 0x5d, 0x71, 0xf7, 0x87, 0xc9, 0xf1, 0xb9, 0xaa, 0xfe,
	0xf5, 0x2e, 0x70, 0xf5, 0x60, 0x9c, 0x33, 0x03, 0x42, 0x4a, 0x4d, 0xec, 0x80, 0xc9, 0x62, 0x60,
	0x65, 0x0f, 0x27, 0xe0, 0xf1, 0x3b, 0x13, 0xf3, 0xcd, 0x65, 0x67, 0x13, 0xd0, 0x17, 0x73, 0x91,
	0x93, 0x2e, 0xb3, 0x45, 0xb6, 0x9c, 0xd5, 0x39, 0xe9, 0xe2, 0x59, 0x5c, 0x8d, 0x3f, 0x96, 0xf9,
	0x22, 0x5b, 0xde, 0xbe, 0xdc, 0xcb, 0xc9, 0xbe, 0x1c, 0x4b, 0x75, 0x42, 0x8a, 0x37, 0x21, 0x5a,
	0x8f, 0x8a, 0x51, 0x6f, 0x15, 0x97, 0xb3, 0x54, 0xa8, 0xe4, 0xc9, 0x41, 0x9e, 0x1d, 0xe4, 0xe7,
	0xd9, 0xa1, 0xbe, 0xf9, 0xa5, 0xd7, 0xfc, 0xfe, 0xf1, 0xb5, 0x36, 0xc4, 0x7d, 0x6c, 0x64, 0xeb,
	0x2c, 0xd0, 0xd0, 0x2b, 0x6b, 0x7b, 0xad, 0x61, 0x17, 0xb5, 0xda, 0xd1, 0xaa, 0xa3, 0x01, 0x57,
	0x01, 0xfd, 0x91, 0x5a, 0x04, 0xdc, 0x33, 0x31, 0x61, 0x80, 0xc9, 0xe5, 0xcd, 0x75, 0xda, 0x78,
	0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x75, 0xf1, 0xf5, 0x4b, 0x01, 0x00, 0x00,
}