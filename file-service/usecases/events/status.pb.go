// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/status.proto

package events

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

type Status struct {
	Messages             []string             `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	HttpCode             int32                `protobuf:"varint,5,opt,name=http_code,json=httpCode,proto3" json:"http_code,omitempty"`
	Errors               []string             `protobuf:"bytes,6,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_24a8aa99d70e3e0b, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *Status) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Status) GetHttpCode() int32 {
	if m != nil {
		return m.HttpCode
	}
	return 0
}

func (m *Status) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "event.Status")
}

func init() { proto.RegisterFile("events/status.proto", fileDescriptor_24a8aa99d70e3e0b) }

var fileDescriptor_24a8aa99d70e3e0b = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x95, 0x44, 0x8d, 0xd9, 0x8c, 0x84, 0xa2, 0x30, 0x10, 0x31, 0x65, 0xa9, 0x2d,
	0xc1, 0x52, 0x31, 0xc2, 0x1b, 0x04, 0x26, 0x16, 0xe4, 0xc4, 0xd7, 0xc4, 0x6a, 0xcc, 0x45, 0xbe,
	0x73, 0x5f, 0x84, 0x17, 0x46, 0x38, 0xb4, 0x8c, 0xff, 0xa7, 0xef, 0xff, 0x6d, 0x9d, 0xb8, 0x81,
	0x13, 0x7c, 0x31, 0x69, 0x62, 0xc3, 0x91, 0xd4, 0x12, 0x90, 0x51, 0xe6, 0x09, 0xd6, 0xf7, 0x23,
	0xe2, 0x38, 0x83, 0x4e, 0xb0, 0x8f, 0x07, 0xcd, 0xce, 0x03, 0xb1, 0xf1, 0xcb, 0xea, 0x3d, 0x7c,
	0x67, 0xa2, 0x78, 0x4b, 0x45, 0x59, 0x8b, 0xad, 0x07, 0x22, 0x33, 0x02, 0x55, 0x9b, 0x66, 0xd3,
	0x96, 0xdd, 0x25, 0xcb, 0xbd, 0x28, 0x2f, 0xcd, 0xea, 0xaa, 0xc9, 0xda, 0xeb, 0xc7, 0x5a, 0xad,
	0xdb, 0xea, 0xbc, 0xad, 0xde, 0xcf, 0x46, 0xf7, 0x2f, 0xcb, 0x3b, 0x51, 0x4e, 0xcc, 0xcb, 0xe7,
	0x80, 0x16, 0xaa, 0xbc, 0xc9, 0xda, 0xbc, 0xdb, 0xfe, 0x82, 0x57, 0xb4, 0x20, 0x6f, 0x45, 0x01,
	0x21, 0x60, 0xa0, 0xaa, 0x48, 0x0f, 0xfe, 0xa5, 0x97, 0xe7, 0x8f, 0xfd, 0xe8, 0x78, 0x8a, 0xbd,
	0x1a, 0xd0, 0x6b, 0x37, 0x4f, 0xc6, 0xfb, 0xc9, 0x5a, 0x7d, 0x8c, 0xd6, 0x1c, 0xdd, 0xee, 0xe0,
	0x66, 0xd8, 0x11, 0x84, 0x93, 0x1b, 0x40, 0x47, 0x82, 0xc1, 0x10, 0x90, 0x5e, 0xcf, 0xd0, 0x17,
	0xe9, 0x3f, 0x4f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x58, 0x0e, 0x74, 0x0f, 0x17, 0x01, 0x00,
	0x00,
}