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
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x69, 0x4f, 0xe4, 0x1a, 0xc5, 0xa1, 0x2e, 0xa5, 0x22, 0x96, 0x5b, 0xec, 0x72, 0x09,
	0xe8, 0x74, 0xa3, 0x0e, 0x82, 0x88, 0x20, 0xc5, 0xc9, 0xa5, 0xa4, 0xcd, 0xbb, 0xe4, 0x71, 0x4d,
	0x53, 0xda, 0x57, 0xff, 0x2f, 0xff, 0x43, 0xb9, 0xf6, 0x72, 0x2e, 0x87, 0x5b, 0x7e, 0x7c, 0xbe,
	0xdf, 0xf7, 0x49, 0xd8, 0xbd, 0xd4, 0xba, 0x07, 0x2d, 0x09, 0x06, 0x61, 0xdd, 0xd8, 0x92, 0xc4,
	0xf6, 0xb8, 0x28, 0xb7, 0xd8, 0x00, 0xef, 0x7a, 0x47, 0x2e, 0xbe, 0xfe, 0x03, 0xb9, 0xbf, 0x4f,
	0x57, 0xff, 0xa5, 0xe7, 0x60, 0x7a, 0xa7, 0x9d, 0xd3, 0x0d, 0x88, 0x69, 0x57, 0x8d, 0x5b, 0x41,
	0x68, 0x61, 0x20, 0x69, 0xbb, 0x19, 0x58, 0xfd, 0x04, 0xec, 0xf2, 0xfd, 0x90, 0x79, 0xc1, 0x06,
	0xe2, 0x2b, 0x16, 0xa2, 0x4a, 0x82, 0x2c, 0xc8, 0x17, 0x45, 0x88, 0x2a, 0xde, 0xb0, 0xa5, 0xef,
	0x4c, 0xc2, 0x2c, 0xc8, 0x2f, 0x1e, 0x6e, 0xf9, 0x09, 0x1b, 0xee, 0x4b, 0x8a, 0x23, 0x1e, 0xdf,
	0xb0, 0x68, 0xff, 0x86, 0xb2, 0x93, 0x64, 0x92, 0x45, 0x16, 0xe4, 0x51, 0xb1, 0xdc, 0x1f, 0x7c,
	0x48, 0x32, 0xf1, 0x86, 0xb1, 0xba, 0x07, 0x49, 0xa0, 0x4a, 0x49, 0xc9, 0xd9, 0xd4, 0x9c, 0xf2,
	0x59, 0x97, 0x7b, 0x5d, 0xfe, 0xe9, 0x75, 0x8b, 0xe8, 0x40, 0x3f, 0xd1, 0xf3, 0xdb, 0xd7, 0xab,
	0x46, 0x32, 0x63, 0xc5, 0x6b, 0x67, 0x05, 0x36, 0x46, 0x5a, 0x6b, 0x94, 0x12, 0xbb, 0x51, 0xc9,
	0x1d, 0xae, 0xfd, 0xfc, 0xf5, 0x00, 0xfd, 0x37, 0xd6, 0x20, 0xa0, 0x25, 0x24, 0x84, 0x41, 0x9c,
	0xf8, 0xb0, 0xea, 0x7c, 0x9a, 0xf5, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x33, 0x11, 0x86,
	0x8c, 0x01, 0x00, 0x00,
}