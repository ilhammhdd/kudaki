// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/mountain/mountain_file.proto

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

type MountainFile struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mountain             *Mountain            `protobuf:"bytes,2,opt,name=mountain,proto3" json:"mountain,omitempty"`
	FilePath             string               `protobuf:"bytes,3,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MountainFile) Reset()         { *m = MountainFile{} }
func (m *MountainFile) String() string { return proto.CompactTextString(m) }
func (*MountainFile) ProtoMessage()    {}
func (*MountainFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb58f8d9bcb324f, []int{0}
}

func (m *MountainFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountainFile.Unmarshal(m, b)
}
func (m *MountainFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountainFile.Marshal(b, m, deterministic)
}
func (m *MountainFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountainFile.Merge(m, src)
}
func (m *MountainFile) XXX_Size() int {
	return xxx_messageInfo_MountainFile.Size(m)
}
func (m *MountainFile) XXX_DiscardUnknown() {
	xxx_messageInfo_MountainFile.DiscardUnknown(m)
}

var xxx_messageInfo_MountainFile proto.InternalMessageInfo

func (m *MountainFile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MountainFile) GetMountain() *Mountain {
	if m != nil {
		return m.Mountain
	}
	return nil
}

func (m *MountainFile) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *MountainFile) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*MountainFile)(nil), "aggregates.mountain.MountainFile")
}

func init() {
	proto.RegisterFile("aggregates/mountain/mountain_file.proto", fileDescriptor_2eb58f8d9bcb324f)
}

var fileDescriptor_2eb58f8d9bcb324f = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x16, 0xa1, 0xc6, 0x20, 0x86, 0xb0, 0x44, 0x41, 0x88, 0xa8, 0x0b, 0x59, 0x6a,
	0x4b, 0x30, 0x75, 0x84, 0x81, 0x05, 0x55, 0x42, 0x11, 0x13, 0x4b, 0xe4, 0xc4, 0x17, 0xe7, 0xd4,
	0x38, 0x8e, 0x9c, 0x0b, 0xff, 0x8b, 0x7f, 0x88, 0x9a, 0xd4, 0x65, 0xa9, 0xd8, 0xfc, 0xf1, 0xbc,
	0xaf, 0x1f, 0x1f, 0x7b, 0x94, 0x5a, 0x3b, 0xd0, 0x92, 0x60, 0x10, 0xc6, 0x8e, 0x1d, 0x49, 0xec,
	0x4e, 0x8b, 0xa2, 0xc6, 0x16, 0x78, 0xef, 0x2c, 0xd9, 0xe8, 0xf6, 0x0f, 0xe4, 0xfe, 0x3e, 0x59,
	0xff, 0x97, 0x9e, 0x83, 0xc9, 0x83, 0xb6, 0x56, 0xb7, 0x20, 0xa6, 0x5d, 0x39, 0xd6, 0x82, 0xd0,
	0xc0, 0x40, 0xd2, 0xf4, 0x33, 0xb0, 0xfe, 0x09, 0xd8, 0xf5, 0xee, 0x98, 0x79, 0xc3, 0x16, 0xa2,
	0x1b, 0xb6, 0x40, 0x15, 0x07, 0x69, 0x90, 0x2d, 0xf3, 0x05, 0xaa, 0x68, 0xcb, 0x56, 0xbe, 0x33,
	0x5e, 0xa4, 0x41, 0x76, 0xf5, 0x74, 0xcf, 0xcf, 0xd8, 0x70, 0x5f, 0x92, 0x9f, 0xf0, 0xe8, 0x8e,
	0x85, 0x87, 0x3f, 0x14, 0xbd, 0xa4, 0x26, 0x5e, 0xa6, 0x41, 0x16, 0xe6, 0xab, 0xc3, 0xc1, 0x87,
	0xa4, 0x26, 0xda, 0x32, 0x56, 0x39, 0x90, 0x04, 0xaa, 0x90, 0x14, 0x5f, 0x4c, 0xcd, 0x09, 0x9f,
	0x75, 0xb9, 0xd7, 0xe5, 0x9f, 0x5e, 0x37, 0x0f, 0x8f, 0xf4, 0x0b, 0xbd, 0xee, 0xbe, 0xde, 0x35,
	0x52, 0x33, 0x96, 0xbc, 0xb2, 0x46, 0x60, 0xdb, 0x48, 0x63, 0x1a, 0xa5, 0xc4, 0x7e, 0x54, 0x72,
	0x8f, 0x9b, 0x81, 0xac, 0x83, 0xda, 0xd9, 0x8e, 0x36, 0x03, 0xb8, 0x6f, 0xac, 0x40, 0x40, 0x47,
	0x48, 0x08, 0x83, 0x38, 0x33, 0xb2, 0xf2, 0x72, 0x7a, 0xed, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xc7, 0xbd, 0x5f, 0x72, 0x8e, 0x01, 0x00, 0x00,
}