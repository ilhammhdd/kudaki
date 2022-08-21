// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/mountain/mountain.proto

package mountain

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

type Mountain struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Height               float32              `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty"`
	Latitude             float64              `protobuf:"fixed64,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64              `protobuf:"fixed64,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Difficulty           float64              `protobuf:"fixed64,7,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Description          string               `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Mountain) Reset()         { *m = Mountain{} }
func (m *Mountain) String() string { return proto.CompactTextString(m) }
func (*Mountain) ProtoMessage()    {}
func (*Mountain) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d670c60b152a810, []int{0}
}

func (m *Mountain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mountain.Unmarshal(m, b)
}
func (m *Mountain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mountain.Marshal(b, m, deterministic)
}
func (m *Mountain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mountain.Merge(m, src)
}
func (m *Mountain) XXX_Size() int {
	return xxx_messageInfo_Mountain.Size(m)
}
func (m *Mountain) XXX_DiscardUnknown() {
	xxx_messageInfo_Mountain.DiscardUnknown(m)
}

var xxx_messageInfo_Mountain proto.InternalMessageInfo

func (m *Mountain) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Mountain) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Mountain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mountain) GetHeight() float32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Mountain) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Mountain) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Mountain) GetDifficulty() float64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Mountain) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Mountain) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Mountain)(nil), "aggregates.mountain.Mountain")
}

func init() { proto.RegisterFile("aggregates/mountain/mountain.proto", fileDescriptor_0d670c60b152a810) }

var fileDescriptor_0d670c60b152a810 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x14, 0xc4, 0xe5, 0xd0, 0x96, 0xc6, 0x95, 0x18, 0x8c, 0x84, 0xac, 0x0a, 0x41, 0xd4, 0x29, 0x4b,
	0x63, 0x09, 0x26, 0x46, 0x18, 0x11, 0x2c, 0x11, 0x13, 0x0b, 0x72, 0xe3, 0x57, 0xe7, 0xa9, 0x71,
	0x5c, 0x25, 0xcf, 0x48, 0x7c, 0x0c, 0xbe, 0x31, 0xaa, 0xd3, 0x7f, 0x03, 0xdb, 0xdd, 0xef, 0xce,
	0xb2, 0x75, 0xe6, 0x0b, 0x6d, 0x6d, 0x07, 0x56, 0x13, 0xf4, 0xca, 0xf9, 0xd0, 0x92, 0xc6, 0xf6,
	0x28, 0x8a, 0x6d, 0xe7, 0xc9, 0x8b, 0xeb, 0x53, 0xa7, 0x38, 0x44, 0xf3, 0x7b, 0xeb, 0xbd, 0x6d,
	0x40, 0xc5, 0xca, 0x2a, 0xac, 0x15, 0xa1, 0x83, 0x9e, 0xb4, 0xdb, 0x0e, 0xa7, 0x16, 0xbf, 0x09,
	0x9f, 0xbe, 0xef, 0xdb, 0xe2, 0x8a, 0x27, 0x68, 0x24, 0xcb, 0x58, 0x3e, 0x2e, 0x13, 0x34, 0x42,
	0xf0, 0x51, 0x08, 0x68, 0x64, 0x92, 0xb1, 0x3c, 0x2d, 0xa3, 0xde, 0xb1, 0x56, 0x3b, 0x90, 0x17,
	0x03, 0xdb, 0x69, 0x71, 0xc3, 0x27, 0x35, 0xa0, 0xad, 0x49, 0x8e, 0x32, 0x96, 0x27, 0xe5, 0xde,
	0x89, 0x39, 0x9f, 0x36, 0x9a, 0x90, 0x82, 0x01, 0x39, 0xce, 0x58, 0xce, 0xca, 0xa3, 0x17, 0xb7,
	0x3c, 0x6d, 0x7c, 0x6b, 0x87, 0x70, 0x12, 0xc3, 0x13, 0x10, 0x77, 0x9c, 0x1b, 0x5c, 0xaf, 0xb1,
	0x0a, 0x0d, 0xfd, 0xc8, 0xcb, 0x18, 0x9f, 0x11, 0x91, 0xf1, 0x99, 0x81, 0xbe, 0xea, 0x70, 0x4b,
	0xe8, 0x5b, 0x39, 0x8d, 0x8f, 0x39, 0x47, 0xe2, 0x89, 0xf3, 0xaa, 0x03, 0x4d, 0x60, 0xbe, 0x34,
	0xc9, 0x34, 0x63, 0xf9, 0xec, 0x61, 0x5e, 0x0c, 0x73, 0x14, 0x87, 0x39, 0x8a, 0x8f, 0xc3, 0x1c,
	0x65, 0xba, 0x6f, 0x3f, 0xd3, 0xcb, 0xdb, 0xe7, 0xab, 0x45, 0xaa, 0xc3, 0xaa, 0xa8, 0xbc, 0x53,
	0xd8, 0xd4, 0xda, 0xb9, 0xda, 0x18, 0xb5, 0x09, 0x46, 0x6f, 0x70, 0x19, 0x7a, 0xe8, 0x96, 0x3a,
	0x50, 0xbd, 0xec, 0xa1, 0xfb, 0xc6, 0x0a, 0x14, 0xb4, 0x84, 0x84, 0xd0, 0xab, 0x7f, 0xbe, 0x69,
	0x35, 0x89, 0x97, 0x3d, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xac, 0xdd, 0x9f, 0x1e, 0xc4, 0x01,
	0x00, 0x00,
}
