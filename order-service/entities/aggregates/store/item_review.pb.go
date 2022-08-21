// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/store/item_review.proto

package store

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

type ItemReview struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Item                 *Item                `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
	Review               string               `protobuf:"bytes,5,opt,name=review,proto3" json:"review,omitempty"`
	Rating               float64              `protobuf:"fixed64,6,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ItemReview) Reset()         { *m = ItemReview{} }
func (m *ItemReview) String() string { return proto.CompactTextString(m) }
func (*ItemReview) ProtoMessage()    {}
func (*ItemReview) Descriptor() ([]byte, []int) {
	return fileDescriptor_3806212aa403f188, []int{0}
}

func (m *ItemReview) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemReview.Unmarshal(m, b)
}
func (m *ItemReview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemReview.Marshal(b, m, deterministic)
}
func (m *ItemReview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemReview.Merge(m, src)
}
func (m *ItemReview) XXX_Size() int {
	return xxx_messageInfo_ItemReview.Size(m)
}
func (m *ItemReview) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemReview.DiscardUnknown(m)
}

var xxx_messageInfo_ItemReview proto.InternalMessageInfo

func (m *ItemReview) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemReview) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ItemReview) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *ItemReview) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ItemReview) GetReview() string {
	if m != nil {
		return m.Review
	}
	return ""
}

func (m *ItemReview) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *ItemReview) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemReview)(nil), "aggregates.store.ItemReview")
}

func init() { proto.RegisterFile("aggregates/store/item_review.proto", fileDescriptor_3806212aa403f188) }

var fileDescriptor_3806212aa403f188 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0x49, 0xdb, 0x7f, 0xff, 0x36, 0x82, 0x48, 0x0e, 0x65, 0x69, 0x0f, 0x2e, 0x3d, 0x2d,
	0x42, 0x13, 0xd0, 0x93, 0x47, 0x15, 0x0f, 0x5e, 0x17, 0xbd, 0x78, 0x29, 0x69, 0x33, 0xa6, 0x43,
	0x9b, 0xa6, 0x64, 0x27, 0xf5, 0x0b, 0xfb, 0x41, 0xa4, 0xb3, 0x5b, 0x84, 0xe2, 0x2d, 0x33, 0xbf,
	0x97, 0xf7, 0xf2, 0x22, 0x67, 0xd6, 0xfb, 0x04, 0xde, 0x12, 0x34, 0xa6, 0xa1, 0x98, 0xc0, 0x20,
	0x41, 0x58, 0x24, 0x38, 0x20, 0x7c, 0xe9, 0x7d, 0x8a, 0x14, 0xd5, 0xf5, 0xaf, 0x46, 0xb3, 0x66,
	0x72, 0xe3, 0x63, 0xf4, 0x5b, 0x30, 0xcc, 0x97, 0xf9, 0xd3, 0x10, 0x06, 0x68, 0xc8, 0x86, 0x7d,
	0x7b, 0x65, 0x32, 0xfd, 0xd3, 0xb6, 0x85, 0xb3, 0x6f, 0x21, 0xe5, 0x2b, 0x41, 0xa8, 0x39, 0x44,
	0x5d, 0xc9, 0x1e, 0xba, 0x42, 0x94, 0xa2, 0xea, 0xd7, 0x3d, 0x74, 0x4a, 0xc9, 0x41, 0xce, 0xe8,
	0x8a, 0x5e, 0x29, 0xaa, 0x51, 0xcd, 0x67, 0x35, 0x95, 0xa3, 0xdc, 0x40, 0x5a, 0x30, 0xe8, 0x33,
	0xb8, 0x38, 0x2e, 0xde, 0x8f, 0xf0, 0x56, 0x0e, 0x8e, 0xee, 0xc5, 0xa0, 0x14, 0xd5, 0xe5, 0xdd,
	0x58, 0x9f, 0x3f, 0x57, 0x73, 0x18, 0x6b, 0xd4, 0x58, 0x0e, 0xdb, 0x6e, 0xc5, 0x3f, 0x76, 0xe9,
	0x26, 0xde, 0x5b, 0xc2, 0x9d, 0x2f, 0x86, 0xa5, 0xa8, 0x44, 0xdd, 0x4d, 0xea, 0x41, 0xca, 0x55,
	0x02, 0x4b, 0xe0, 0x16, 0x96, 0x8a, 0xff, 0x9c, 0x30, 0xd1, 0x6d, 0x7d, 0x7d, 0xaa, 0xaf, 0xdf,
	0x4e, 0xf5, 0xeb, 0x51, 0xa7, 0x7e, 0xa4, 0xa7, 0x97, 0x8f, 0x67, 0x8f, 0xb4, 0xce, 0x4b, 0xbd,
	0x8a, 0xc1, 0xe0, 0x76, 0x6d, 0x43, 0x58, 0x3b, 0x67, 0x36, 0xd9, 0xd9, 0x0d, 0xce, 0x63, 0x72,
	0x90, 0xe6, 0x0d, 0xa4, 0x03, 0xae, 0xc0, 0xc0, 0x8e, 0x90, 0x10, 0x1a, 0x73, 0xfe, 0x71, 0xcb,
	0x21, 0xa7, 0xdc, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x94, 0x09, 0x68, 0x81, 0xaa, 0x01, 0x00,
	0x00,
}
