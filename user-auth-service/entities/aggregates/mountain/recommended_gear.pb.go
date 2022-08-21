// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/mountain/recommended_gear.proto

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

type RecommendedGear struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Mountain             *Mountain            `protobuf:"bytes,4,opt,name=mountain,proto3" json:"mountain,omitempty"`
	Upvote               int32                `protobuf:"varint,5,opt,name=upvote,proto3" json:"upvote,omitempty"`
	Downvote             int32                `protobuf:"varint,6,opt,name=downvote,proto3" json:"downvote,omitempty"`
	Seen                 int32                `protobuf:"varint,7,opt,name=seen,proto3" json:"seen,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RecommendedGear) Reset()         { *m = RecommendedGear{} }
func (m *RecommendedGear) String() string { return proto.CompactTextString(m) }
func (*RecommendedGear) ProtoMessage()    {}
func (*RecommendedGear) Descriptor() ([]byte, []int) {
	return fileDescriptor_50de5b97abeab0b3, []int{0}
}

func (m *RecommendedGear) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendedGear.Unmarshal(m, b)
}
func (m *RecommendedGear) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendedGear.Marshal(b, m, deterministic)
}
func (m *RecommendedGear) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendedGear.Merge(m, src)
}
func (m *RecommendedGear) XXX_Size() int {
	return xxx_messageInfo_RecommendedGear.Size(m)
}
func (m *RecommendedGear) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendedGear.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendedGear proto.InternalMessageInfo

func (m *RecommendedGear) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RecommendedGear) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RecommendedGear) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *RecommendedGear) GetMountain() *Mountain {
	if m != nil {
		return m.Mountain
	}
	return nil
}

func (m *RecommendedGear) GetUpvote() int32 {
	if m != nil {
		return m.Upvote
	}
	return 0
}

func (m *RecommendedGear) GetDownvote() int32 {
	if m != nil {
		return m.Downvote
	}
	return 0
}

func (m *RecommendedGear) GetSeen() int32 {
	if m != nil {
		return m.Seen
	}
	return 0
}

func (m *RecommendedGear) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*RecommendedGear)(nil), "aggregates.mountain.RecommendedGear")
}

func init() {
	proto.RegisterFile("aggregates/mountain/recommended_gear.proto", fileDescriptor_50de5b97abeab0b3)
}

var fileDescriptor_50de5b97abeab0b3 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x6a, 0xf3, 0x30,
	0x10, 0xc4, 0xce, 0xcf, 0xe7, 0xe8, 0x83, 0x16, 0x54, 0x28, 0xc6, 0xa5, 0xd4, 0xe4, 0x64, 0x0a,
	0x91, 0xa0, 0x3d, 0xe5, 0xd8, 0x5e, 0x0a, 0xa5, 0xbd, 0x98, 0xf6, 0xd2, 0x4b, 0x50, 0xac, 0xad,
	0x2c, 0x12, 0x59, 0x41, 0x5e, 0xa5, 0xaf, 0xd1, 0x47, 0x2e, 0x56, 0xe2, 0xe4, 0xe2, 0xdb, 0xcc,
	0xce, 0xd8, 0x3b, 0x3b, 0x22, 0xf7, 0x42, 0x29, 0x07, 0x4a, 0x20, 0xb4, 0xdc, 0x58, 0xdf, 0xa0,
	0xd0, 0x0d, 0x77, 0x50, 0x59, 0x63, 0xa0, 0x91, 0x20, 0x57, 0x0a, 0x84, 0x63, 0x3b, 0x67, 0xd1,
	0xd2, 0xab, 0xb3, 0x97, 0xf5, 0xde, 0x6c, 0x3e, 0xf4, 0x83, 0x1e, 0x1c, 0x3e, 0xcc, 0xee, 0x94,
	0xb5, 0x6a, 0x0b, 0x3c, 0xb0, 0xb5, 0xff, 0xe6, 0xa8, 0x0d, 0xb4, 0x28, 0xcc, 0xee, 0x60, 0x98,
	0xff, 0xc6, 0xe4, 0xb2, 0x3c, 0x2f, 0x7d, 0x01, 0xe1, 0xe8, 0x05, 0x89, 0xb5, 0x4c, 0xa3, 0x3c,
	0x2a, 0x46, 0x65, 0xac, 0x25, 0xa5, 0x64, 0xec, 0xbd, 0x96, 0x69, 0x9c, 0x47, 0xc5, 0xac, 0x0c,
	0x98, 0xde, 0x90, 0x99, 0x6f, 0xc1, 0xad, 0x82, 0x30, 0x0a, 0x42, 0xd2, 0x0d, 0x3e, 0x3b, 0x71,
	0x49, 0x92, 0x3e, 0x47, 0x3a, 0xce, 0xa3, 0xe2, 0xff, 0xc3, 0x2d, 0x1b, 0xb8, 0x80, 0xbd, 0x1f,
	0x41, 0x79, 0xb2, 0xd3, 0x6b, 0x32, 0xf5, 0xbb, 0xbd, 0x45, 0x48, 0x27, 0x79, 0x54, 0x4c, 0xca,
	0x23, 0xa3, 0x19, 0x49, 0xa4, 0xfd, 0x69, 0x82, 0x32, 0x0d, 0xca, 0x89, 0x77, 0xf9, 0x5a, 0x80,
	0x26, 0xfd, 0x17, 0xe6, 0x01, 0xd3, 0x25, 0x21, 0x95, 0x03, 0x81, 0x20, 0x57, 0x02, 0xd3, 0x24,
	0x84, 0xc8, 0xd8, 0xa1, 0x0d, 0xd6, 0xb7, 0xc1, 0x3e, 0xfa, 0x36, 0xca, 0xd9, 0xd1, 0xfd, 0x84,
	0xcf, 0x6f, 0x5f, 0xaf, 0x4a, 0x63, 0xed, 0xd7, 0xac, 0xb2, 0x86, 0xeb, 0x6d, 0x2d, 0x8c, 0xa9,
	0xa5, 0xe4, 0x1b, 0x2f, 0xc5, 0x46, 0x2f, 0xba, 0x2b, 0x17, 0xc2, 0x63, 0xbd, 0x68, 0xc1, 0xed,
	0x75, 0x05, 0x1c, 0x1a, 0xd4, 0xa8, 0xa1, 0xe5, 0x03, 0x0f, 0xb2, 0x9e, 0x86, 0x65, 0x8f, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x95, 0x75, 0xdf, 0xef, 0x01, 0x00, 0x00,
}
