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
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcd, 0x6a, 0xf3, 0x30,
	0x10, 0xc4, 0xce, 0xcf, 0xe7, 0xe8, 0x83, 0x16, 0x54, 0x28, 0xc6, 0xa5, 0xd4, 0xe4, 0x64, 0x0a,
	0x91, 0xa0, 0x3d, 0xe5, 0xd8, 0x5e, 0x7a, 0x2a, 0xa5, 0xa6, 0xbd, 0xf4, 0x12, 0x14, 0x6b, 0xab,
	0x88, 0x44, 0x52, 0x90, 0x57, 0xe9, 0x6b, 0xf4, 0x91, 0x4b, 0x94, 0x38, 0xb9, 0xf8, 0x36, 0xb3,
	0x33, 0xf6, 0x8e, 0x66, 0xc9, 0xbd, 0x50, 0xca, 0x83, 0x12, 0x08, 0x2d, 0x37, 0x2e, 0x58, 0x14,
	0xda, 0x72, 0x0f, 0x8d, 0x33, 0x06, 0xac, 0x04, 0xb9, 0x50, 0x20, 0x3c, 0xdb, 0x7a, 0x87, 0x8e,
	0x5e, 0x9d, 0xbd, 0xac, 0xf3, 0x16, 0xd3, 0xbe, 0x1f, 0x74, 0xe0, 0xf0, 0x61, 0x71, 0xa7, 0x9c,
	0x53, 0x1b, 0xe0, 0x91, 0x2d, 0xc3, 0x37, 0x47, 0x6d, 0xa0, 0x45, 0x61, 0xb6, 0x07, 0xc3, 0xf4,
	0x37, 0x25, 0x97, 0xf5, 0x79, 0xe9, 0x0b, 0x08, 0x4f, 0x2f, 0x48, 0xaa, 0x65, 0x9e, 0x94, 0x49,
	0x35, 0xa8, 0x53, 0x2d, 0x29, 0x25, 0xc3, 0x10, 0xb4, 0xcc, 0xd3, 0x32, 0xa9, 0x26, 0x75, 0xc4,
	0xf4, 0x86, 0x4c, 0x42, 0x0b, 0x7e, 0x11, 0x85, 0x41, 0x14, 0xb2, 0xfd, 0xe0, 0x73, 0x2f, 0xce,
	0x49, 0xd6, 0xe5, 0xc8, 0x87, 0x65, 0x52, 0xfd, 0x7f, 0xb8, 0x65, 0x3d, 0x2f, 0x60, 0xaf, 0x47,
	0x50, 0x9f, 0xec, 0xf4, 0x9a, 0x8c, 0xc3, 0x76, 0xe7, 0x10, 0xf2, 0x51, 0x99, 0x54, 0xa3, 0xfa,
	0xc8, 0x68, 0x41, 0x32, 0xe9, 0x7e, 0x6c, 0x54, 0xc6, 0x51, 0x39, 0xf1, 0x7d, 0xbe, 0x16, 0xc0,
	0xe6, 0xff, 0xe2, 0x3c, 0x62, 0x3a, 0x27, 0xa4, 0xf1, 0x20, 0x10, 0xe4, 0x42, 0x60, 0x9e, 0xc5,
	0x10, 0x05, 0x3b, 0xb4, 0xc1, 0xba, 0x36, 0xd8, 0x47, 0xd7, 0x46, 0x3d, 0x39, 0xba, 0x9f, 0xf0,
	0xf9, 0xfd, 0xeb, 0x4d, 0x69, 0x5c, 0x85, 0x25, 0x6b, 0x9c, 0xe1, 0x7a, 0xb3, 0x12, 0xc6, 0xac,
	0xa4, 0xe4, 0xeb, 0x20, 0xc5, 0x5a, 0xcf, 0x4e, 0x27, 0x12, 0xa8, 0x9d, 0x9d, 0xb5, 0xe0, 0x77,
	0xba, 0x01, 0x0e, 0x16, 0x35, 0x6a, 0x68, 0x79, 0xcf, 0x55, 0x96, 0xe3, 0xb8, 0xf1, 0xf1, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0x0c, 0xce, 0x90, 0x98, 0xf4, 0x01, 0x00, 0x00,
}
