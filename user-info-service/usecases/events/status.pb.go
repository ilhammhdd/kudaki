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
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0x29, 0xf3, 0xb7, 0x4c, 0xf3, 0xef, 0x22, 0x48, 0xa9, 0x0b, 0x8b, 0xab, 0x6e, 0x9a,
	0x80, 0x6e, 0xdc, 0x88, 0xa0, 0x6f, 0x50, 0x5d, 0xb9, 0x91, 0xb4, 0xb9, 0xd3, 0x86, 0x99, 0xcc,
	0x2d, 0xb9, 0x37, 0xf3, 0x22, 0xbe, 0xb0, 0x98, 0x3a, 0xe3, 0xf2, 0x7c, 0x7c, 0xe7, 0x24, 0x5c,
	0x71, 0x05, 0x27, 0x38, 0x32, 0x69, 0x62, 0xc3, 0x91, 0xd4, 0x12, 0x90, 0x51, 0xe6, 0x09, 0xd6,
	0xb7, 0x13, 0xe2, 0x74, 0x00, 0x9d, 0xe0, 0x10, 0x77, 0x9a, 0x9d, 0x07, 0x62, 0xe3, 0x97, 0xd5,
	0xbb, 0xfb, 0xca, 0x44, 0xf1, 0x96, 0x8a, 0xb2, 0x16, 0x5b, 0x0f, 0x44, 0x66, 0x02, 0xaa, 0x36,
	0xcd, 0xa6, 0x2d, 0xfb, 0x4b, 0x96, 0x8f, 0xa2, 0xbc, 0x34, 0xab, 0x7f, 0x4d, 0xd6, 0xfe, 0xbf,
	0xaf, 0xd5, 0xba, 0xad, 0xce, 0xdb, 0xea, 0xfd, 0x6c, 0xf4, 0x7f, 0xb2, 0xbc, 0x11, 0xe5, 0xcc,
	0xbc, 0x7c, 0x8e, 0x68, 0xa1, 0xca, 0x9b, 0xac, 0xcd, 0xfb, 0xed, 0x0f, 0x78, 0x45, 0x0b, 0xf2,
	0x5a, 0x14, 0x10, 0x02, 0x06, 0xaa, 0x8a, 0xf4, 0xe0, 0x6f, 0x7a, 0x79, 0xfe, 0x78, 0x9a, 0x1c,
	0xcf, 0x71, 0x50, 0x23, 0x7a, 0xed, 0x0e, 0xb3, 0xf1, 0x7e, 0xb6, 0x56, 0xef, 0xa3, 0x35, 0x7b,
	0xd7, 0x45, 0x82, 0xd0, 0xb9, 0xe3, 0x0e, 0x3b, 0x82, 0x70, 0x72, 0x23, 0xe8, 0x48, 0x30, 0x1a,
	0x02, 0xd2, 0xeb, 0x2d, 0x86, 0x22, 0x7d, 0xea, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x89, 0x13,
	0xa5, 0xf4, 0x1c, 0x01, 0x00, 0x00,
}
