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
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x95, 0x44, 0x8d, 0xd9, 0x8c, 0x84, 0xa2, 0x30, 0x10, 0x31, 0x65, 0xa9, 0x2d,
	0xc1, 0xc2, 0x84, 0x10, 0xbc, 0x41, 0x60, 0x62, 0x41, 0x4e, 0x7c, 0x4d, 0xac, 0xd6, 0xbd, 0xc8,
	0x77, 0xee, 0x8b, 0xf0, 0xc2, 0x08, 0x87, 0xb6, 0xe3, 0xff, 0xe9, 0xfb, 0x7f, 0x5b, 0x27, 0x6e,
	0xe0, 0x08, 0x07, 0x26, 0x4d, 0x6c, 0x38, 0x92, 0x9a, 0x03, 0x32, 0xca, 0x3c, 0xc1, 0xfa, 0x7e,
	0x44, 0x1c, 0xf7, 0xa0, 0x13, 0xec, 0xe3, 0x56, 0xb3, 0xf3, 0x40, 0x6c, 0xfc, 0xbc, 0x78, 0x0f,
	0x3f, 0x99, 0x28, 0x3e, 0x52, 0x51, 0xd6, 0x62, 0xed, 0x81, 0xc8, 0x8c, 0x40, 0xd5, 0xaa, 0x59,
	0xb5, 0x65, 0x77, 0xce, 0xf2, 0x59, 0x94, 0xe7, 0x66, 0x75, 0xd5, 0x64, 0xed, 0xf5, 0x63, 0xad,
	0x96, 0x6d, 0x75, 0xda, 0x56, 0x9f, 0x27, 0xa3, 0xbb, 0xc8, 0xf2, 0x4e, 0x94, 0x13, 0xf3, 0xfc,
	0x3d, 0xa0, 0x85, 0x2a, 0x6f, 0xb2, 0x36, 0xef, 0xd6, 0x7f, 0xe0, 0x1d, 0x2d, 0xc8, 0x5b, 0x51,
	0x40, 0x08, 0x18, 0xa8, 0x2a, 0xd2, 0x83, 0xff, 0xe9, 0xed, 0xf5, 0xeb, 0x65, 0x74, 0x3c, 0xc5,
	0x5e, 0x0d, 0xe8, 0xb5, 0xdb, 0x4f, 0xc6, 0xfb, 0xc9, 0x5a, 0xbd, 0x8b, 0xd6, 0xec, 0xdc, 0x86,
	0x18, 0x03, 0x6c, 0x03, 0x1e, 0x78, 0x43, 0x10, 0x8e, 0x6e, 0x00, 0x1d, 0x09, 0x06, 0x43, 0x40,
	0x7a, 0x39, 0x46, 0x5f, 0xa4, 0x5f, 0x3d, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x78, 0x2e, 0xaf,
	0xe6, 0x1d, 0x01, 0x00, 0x00,
}
