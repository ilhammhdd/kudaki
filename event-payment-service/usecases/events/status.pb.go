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
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x95, 0x44, 0x8d, 0xd9, 0x8c, 0x84, 0xa2, 0x30, 0x10, 0x31, 0x65, 0xa9, 0x2d,
	0xc1, 0xc2, 0x0a, 0xfd, 0x07, 0x81, 0x89, 0x05, 0x39, 0xf1, 0x91, 0x58, 0xad, 0xb9, 0xc8, 0x77,
	0xae, 0xc4, 0xef, 0xe0, 0x0f, 0xa3, 0x3a, 0xb4, 0x9d, 0xac, 0xf7, 0xf4, 0xbd, 0xe7, 0xbb, 0x13,
	0x37, 0x70, 0x80, 0x6f, 0x26, 0x4d, 0x6c, 0x38, 0x92, 0x9a, 0x03, 0x32, 0xca, 0x3c, 0x99, 0xf5,
	0xfd, 0x88, 0x38, 0xee, 0x41, 0x27, 0xb3, 0x8f, 0x5f, 0x9a, 0x9d, 0x07, 0x62, 0xe3, 0xe7, 0x85,
	0x7b, 0xf8, 0xcd, 0x44, 0xf1, 0x96, 0x82, 0xb2, 0x16, 0x6b, 0x0f, 0x44, 0x66, 0x04, 0xaa, 0x56,
	0xcd, 0xaa, 0x2d, 0xbb, 0xb3, 0x96, 0xcf, 0xa2, 0x3c, 0x27, 0xab, 0xab, 0x26, 0x6b, 0xaf, 0x1f,
	0x6b, 0xb5, 0x74, 0xab, 0x53, 0xb7, 0x7a, 0x3f, 0x11, 0xdd, 0x05, 0x96, 0x77, 0xa2, 0x9c, 0x98,
	0xe7, 0xcf, 0x01, 0x2d, 0x54, 0x79, 0x93, 0xb5, 0x79, 0xb7, 0x3e, 0x1a, 0x5b, 0xb4, 0x20, 0x6f,
	0x45, 0x01, 0x21, 0x60, 0xa0, 0xaa, 0x48, 0x1f, 0xfe, 0xab, 0xd7, 0xed, 0xc7, 0xcb, 0xe8, 0x78,
	0x8a, 0xbd, 0x1a, 0xd0, 0x6b, 0xb7, 0x9f, 0x8c, 0xf7, 0x93, 0xb5, 0x7a, 0x17, 0xad, 0xd9, 0xb9,
	0x4d, 0xda, 0x6d, 0x33, 0x9b, 0x1f, 0x7f, 0x7c, 0x09, 0xc2, 0xc1, 0x0d, 0xa0, 0x23, 0xc1, 0x60,
	0x08, 0x48, 0x2f, 0xf7, 0xe8, 0x8b, 0x34, 0xd8, 0xd3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37,
	0xbb, 0xfb, 0x09, 0x20, 0x01, 0x00, 0x00,
}
