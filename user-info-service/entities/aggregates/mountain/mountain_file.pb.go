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
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4f, 0x84, 0x30,
	0x14, 0xc7, 0x03, 0x67, 0xcc, 0x51, 0x8d, 0x03, 0x2e, 0x04, 0x63, 0x24, 0xb7, 0xc8, 0x42, 0x9b,
	0xe8, 0x74, 0xa3, 0x0e, 0x0e, 0x46, 0x13, 0x43, 0x9c, 0x5c, 0x48, 0xa1, 0x8f, 0xf2, 0x72, 0x94,
	0x12, 0x78, 0xf8, 0x7f, 0xf9, 0x1f, 0x9a, 0x83, 0xeb, 0xb9, 0x5c, 0xdc, 0xfa, 0xe3, 0xf3, 0xfd,
	0xf6, 0xd3, 0xc7, 0xee, 0xa5, 0xd6, 0x03, 0x68, 0x49, 0x30, 0x0a, 0x63, 0xa7, 0x8e, 0x24, 0x76,
	0xc7, 0x45, 0x51, 0x63, 0x0b, 0xbc, 0x1f, 0x2c, 0xd9, 0xf0, 0xfa, 0x0f, 0xe4, 0xee, 0x3e, 0xde,
	0xfc, 0x97, 0x5e, 0x82, 0xf1, 0x9d, 0xb6, 0x56, 0xb7, 0x20, 0xe6, 0x5d, 0x39, 0xd5, 0x82, 0xd0,
	0xc0, 0x48, 0xd2, 0xf4, 0x0b, 0xb0, 0xf9, 0xf1, 0xd8, 0xe5, 0xfb, 0x21, 0xf3, 0x82, 0x2d, 0x84,
	0x57, 0xcc, 0x47, 0x15, 0x79, 0x89, 0x97, 0xae, 0x72, 0x1f, 0x55, 0xb8, 0x65, 0x6b, 0xd7, 0x19,
	0xf9, 0x89, 0x97, 0x5e, 0x3c, 0xdc, 0xf2, 0x13, 0x36, 0xdc, 0x95, 0xe4, 0x47, 0x3c, 0xbc, 0x61,
	0xc1, 0xfe, 0x0f, 0x45, 0x2f, 0xa9, 0x89, 0x56, 0x89, 0x97, 0x06, 0xf9, 0x7a, 0x7f, 0xf0, 0x21,
	0xa9, 0x09, 0xb7, 0x8c, 0x55, 0x03, 0x48, 0x02, 0x55, 0x48, 0x8a, 0xce, 0xe6, 0xe6, 0x98, 0x2f,
	0xba, 0xdc, 0xe9, 0xf2, 0x4f, 0xa7, 0x9b, 0x07, 0x07, 0xfa, 0x89, 0x9e, 0xdf, 0xbe, 0x5e, 0x35,
	0x52, 0x33, 0x95, 0xbc, 0xb2, 0x46, 0x60, 0xdb, 0x48, 0x63, 0x1a, 0xa5, 0xc4, 0x6e, 0x52, 0x72,
	0x87, 0xd9, 0x34, 0xc2, 0x90, 0x61, 0x57, 0xdb, 0x6c, 0x84, 0xe1, 0x1b, 0x2b, 0x10, 0xd0, 0x11,
	0x12, 0xc2, 0x28, 0x4e, 0x4c, 0xac, 0x3c, 0x9f, 0x1f, 0x7b, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xf2, 0xdf, 0xf9, 0x0f, 0x8d, 0x01, 0x00, 0x00,
}
