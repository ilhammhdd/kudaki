// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/user/profile.proto

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

type Profile struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	User                 *User                `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	FullName             string               `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Photo                string               `protobuf:"bytes,5,opt,name=photo,proto3" json:"photo,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_11073363221d01be, []int{0}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Profile) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Profile) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Profile) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Profile) GetPhoto() string {
	if m != nil {
		return m.Photo
	}
	return ""
}

func (m *Profile) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "aggregates.user.Profile")
}

func init() { proto.RegisterFile("aggregates/user/profile.proto", fileDescriptor_11073363221d01be) }

var fileDescriptor_11073363221d01be = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xe9, 0xfc, 0xf9, 0xe6, 0x9b, 0x08, 0x0a, 0x41, 0x21, 0x54, 0xc4, 0xc1, 0xd5, 0xb8,
	0x98, 0x04, 0x74, 0xe5, 0x52, 0x97, 0x22, 0x22, 0x45, 0x37, 0x6e, 0x86, 0x4c, 0x73, 0x9b, 0x86,
	0x69, 0x9a, 0x92, 0xdc, 0x0c, 0xf8, 0x82, 0x3e, 0x97, 0x34, 0x75, 0x10, 0xba, 0x49, 0x72, 0x73,
	0x7e, 0x9c, 0x93, 0x13, 0x72, 0x25, 0xb5, 0xf6, 0xa0, 0x25, 0x42, 0x10, 0x31, 0x80, 0x17, 0x9d,
	0x77, 0x95, 0x69, 0x80, 0x77, 0xde, 0xa1, 0xa3, 0x67, 0x7f, 0x32, 0xef, 0xe5, 0x3c, 0x1f, 0xf3,
	0xfd, 0x32, 0xc0, 0xf9, 0xb5, 0x76, 0x4e, 0x37, 0x20, 0xd2, 0xb4, 0x8b, 0x95, 0x40, 0x63, 0x21,
	0xa0, 0xb4, 0xdd, 0x00, 0xdc, 0x7c, 0x67, 0x64, 0xf1, 0x36, 0xf8, 0xd3, 0x53, 0x32, 0x31, 0x8a,
	0x65, 0xab, 0x6c, 0x3d, 0x2d, 0x26, 0x46, 0x51, 0x4a, 0x66, 0x31, 0x1a, 0xc5, 0x26, 0xab, 0x6c,
	0xbd, 0x2c, 0xd2, 0x99, 0xde, 0x92, 0x59, 0x6f, 0xcf, 0xa6, 0xab, 0x6c, 0x7d, 0x72, 0x77, 0xc1,
	0x47, 0x8f, 0xe1, 0x1f, 0x01, 0x7c, 0x91, 0x10, 0x7a, 0x49, 0x96, 0x55, 0x6c, 0x9a, 0x6d, 0x2b,
	0x2d, 0xb0, 0x59, 0xf2, 0xf8, 0xdf, 0x5f, 0xbc, 0x4a, 0x0b, 0xf4, 0x9c, 0xcc, 0xbb, 0xda, 0xa1,
	0x63, 0xf3, 0x24, 0x0c, 0x03, 0x7d, 0x20, 0xa4, 0xf4, 0x20, 0x11, 0xd4, 0x56, 0x22, 0x5b, 0xa4,
	0x8c, 0x9c, 0x0f, 0x1d, 0xf8, 0xb1, 0x03, 0x7f, 0x3f, 0x76, 0x28, 0x96, 0xbf, 0xf4, 0x23, 0x3e,
	0xbd, 0x7c, 0x3e, 0x6b, 0x83, 0x75, 0xdc, 0xf1, 0xd2, 0x59, 0x61, 0x9a, 0x5a, 0x5a, 0x5b, 0x2b,
	0x25, 0xf6, 0x51, 0xc9, 0xbd, 0xd9, 0xc0, 0x01, 0x5a, 0xdc, 0x74, 0xf2, 0xcb, 0xf6, 0x7b, 0x00,
	0x7f, 0x30, 0x25, 0x08, 0x68, 0xd1, 0xa0, 0x81, 0x20, 0x46, 0x5f, 0xb8, 0xfb, 0x97, 0xc2, 0xee,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x51, 0xc1, 0xca, 0x24, 0x8c, 0x01, 0x00, 0x00,
}
