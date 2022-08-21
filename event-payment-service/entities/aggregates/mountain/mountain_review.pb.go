// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/mountain/mountain_review.proto

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

type MountainReview struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Mountain             *Mountain            `protobuf:"bytes,4,opt,name=mountain,proto3" json:"mountain,omitempty"`
	Difficulty           float64              `protobuf:"fixed64,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Comment              string               `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MountainReview) Reset()         { *m = MountainReview{} }
func (m *MountainReview) String() string { return proto.CompactTextString(m) }
func (*MountainReview) ProtoMessage()    {}
func (*MountainReview) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a245687c67fd166, []int{0}
}

func (m *MountainReview) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountainReview.Unmarshal(m, b)
}
func (m *MountainReview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountainReview.Marshal(b, m, deterministic)
}
func (m *MountainReview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountainReview.Merge(m, src)
}
func (m *MountainReview) XXX_Size() int {
	return xxx_messageInfo_MountainReview.Size(m)
}
func (m *MountainReview) XXX_DiscardUnknown() {
	xxx_messageInfo_MountainReview.DiscardUnknown(m)
}

var xxx_messageInfo_MountainReview proto.InternalMessageInfo

func (m *MountainReview) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MountainReview) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *MountainReview) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *MountainReview) GetMountain() *Mountain {
	if m != nil {
		return m.Mountain
	}
	return nil
}

func (m *MountainReview) GetDifficulty() float64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *MountainReview) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *MountainReview) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*MountainReview)(nil), "aggregates.mountain.MountainReview")
}

func init() {
	proto.RegisterFile("aggregates/mountain/mountain_review.proto", fileDescriptor_7a245687c67fd166)
}

var fileDescriptor_7a245687c67fd166 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x95, 0xb4, 0xf4, 0x8f, 0x91, 0x3a, 0x98, 0xc5, 0x2a, 0x02, 0xa2, 0x4e, 0x61, 0xa8,
	0x2d, 0xc1, 0xd4, 0x11, 0x76, 0x10, 0x8a, 0x60, 0x61, 0xa9, 0xdc, 0xf8, 0x9a, 0x9e, 0x5a, 0xc7,
	0x55, 0x72, 0x2e, 0xea, 0xc8, 0x37, 0x47, 0x75, 0x9b, 0xc0, 0x50, 0xb1, 0xdd, 0xdd, 0x7b, 0x4f,
	0x3f, 0x3f, 0x99, 0xdd, 0xeb, 0xa2, 0xa8, 0xa0, 0xd0, 0x04, 0xb5, 0xb2, 0xce, 0x97, 0xa4, 0xb1,
	0x6c, 0x87, 0x79, 0x05, 0x3b, 0x84, 0x2f, 0xb9, 0xad, 0x1c, 0x39, 0x7e, 0xf5, 0x6b, 0x95, 0x8d,
	0x63, 0x3c, 0xf9, 0x2f, 0x7f, 0x0c, 0x8e, 0xef, 0x0a, 0xe7, 0x8a, 0x0d, 0xa8, 0xb0, 0x2d, 0xfc,
	0x52, 0x11, 0x5a, 0xa8, 0x49, 0xdb, 0xed, 0xd1, 0x30, 0xf9, 0x8e, 0xd9, 0xe8, 0xe5, 0x94, 0xc9,
	0x02, 0x92, 0x8f, 0x58, 0x8c, 0x46, 0x44, 0x49, 0x94, 0x76, 0xb2, 0x18, 0x0d, 0xe7, 0xac, 0xeb,
	0x3d, 0x1a, 0x11, 0x27, 0x51, 0x3a, 0xcc, 0xc2, 0xcc, 0xaf, 0xd9, 0xd0, 0xd7, 0x50, 0xcd, 0x83,
	0xd0, 0x09, 0xc2, 0xe0, 0x70, 0xf8, 0x38, 0x88, 0x33, 0x36, 0x68, 0x9e, 0x21, 0xba, 0x49, 0x94,
	0x5e, 0x3e, 0xdc, 0xc8, 0x33, 0x05, 0x64, 0xcb, 0x6d, 0xed, 0xfc, 0x96, 0x31, 0x83, 0xcb, 0x25,
	0xe6, 0x7e, 0x43, 0x7b, 0x71, 0x91, 0x44, 0x69, 0x94, 0xfd, 0xb9, 0x70, 0xc1, 0xfa, 0xb9, 0xb3,
	0x16, 0x4a, 0x12, 0xbd, 0x40, 0x6d, 0x56, 0x3e, 0x63, 0x2c, 0xaf, 0x40, 0x13, 0x98, 0xb9, 0x26,
	0xd1, 0x0f, 0xd8, 0xb1, 0x3c, 0xd6, 0x97, 0x4d, 0x7d, 0xf9, 0xde, 0xd4, 0xcf, 0x86, 0x27, 0xf7,
	0x13, 0x3d, 0xbf, 0x7d, 0xbe, 0x16, 0x48, 0x2b, 0xbf, 0x90, 0xb9, 0xb3, 0x0a, 0x37, 0x2b, 0x6d,
	0xed, 0xca, 0x18, 0xb5, 0xf6, 0x46, 0xaf, 0x71, 0x0a, 0x3b, 0x28, 0x69, 0xba, 0xd5, 0xfb, 0x03,
	0x67, 0x5a, 0x43, 0xb5, 0xc3, 0x1c, 0x14, 0x94, 0x84, 0x84, 0x50, 0xab, 0x33, 0xbf, 0xb0, 0xe8,
	0x05, 0xe0, 0xe3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x53, 0x39, 0x62, 0xe3, 0x01, 0x00,
	0x00,
}