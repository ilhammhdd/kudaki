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
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x6b, 0x3a, 0x31,
	0x10, 0xc5, 0x89, 0xfa, 0xf7, 0x5f, 0x53, 0x28, 0x25, 0x07, 0x59, 0xf4, 0xd0, 0xc5, 0xd3, 0x52,
	0x30, 0x81, 0xf6, 0xd4, 0x63, 0xa5, 0x97, 0x5e, 0x97, 0xf6, 0xd2, 0x8b, 0x44, 0x33, 0x8d, 0x83,
	0xc6, 0x95, 0x64, 0x62, 0xbf, 0x70, 0x3f, 0x48, 0xd9, 0x59, 0xa5, 0x20, 0xbd, 0xed, 0xcc, 0xef,
	0xed, 0x7b, 0xf3, 0x22, 0x67, 0xd6, 0xfb, 0x08, 0xde, 0x12, 0x24, 0x93, 0xa8, 0x89, 0x60, 0x90,
	0x20, 0x2c, 0x23, 0x1c, 0x11, 0xbe, 0xf4, 0x21, 0x36, 0xd4, 0xa8, 0xdb, 0x5f, 0x8d, 0x66, 0xcd,
	0xe4, 0xce, 0x37, 0x8d, 0xdf, 0x81, 0x61, 0xbe, 0xca, 0x9f, 0x86, 0x30, 0x40, 0x22, 0x1b, 0x0e,
	0xdd, 0x2f, 0x93, 0xe9, 0x9f, 0xb6, 0x1d, 0x9c, 0x7d, 0x0b, 0x29, 0x5f, 0x09, 0x42, 0xcd, 0x21,
	0xea, 0x46, 0xf6, 0xd0, 0x15, 0xa2, 0x14, 0x55, 0xbf, 0xee, 0xa1, 0x53, 0x4a, 0x0e, 0x72, 0x46,
	0x57, 0xf4, 0x4a, 0x51, 0x8d, 0x6a, 0xfe, 0x56, 0x53, 0x39, 0xca, 0x09, 0xe2, 0x92, 0x41, 0x9f,
	0xc1, 0x55, 0xbb, 0x78, 0x6f, 0xe1, 0xbd, 0x1c, 0xb4, 0xee, 0xc5, 0xa0, 0x14, 0xd5, 0xf5, 0xc3,
	0x58, 0x5f, 0x9e, 0xab, 0x39, 0x8c, 0x35, 0x6a, 0x2c, 0x87, 0x5d, 0xb7, 0xe2, 0x1f, 0xbb, 0x9c,
	0x26, 0xde, 0x5b, 0xc2, 0xbd, 0x2f, 0x86, 0xa5, 0xa8, 0x44, 0x7d, 0x9a, 0xd4, 0x93, 0x94, 0xeb,
	0x08, 0x96, 0xc0, 0x2d, 0x2d, 0x15, 0xff, 0x39, 0x61, 0xa2, 0xbb, 0xfa, 0xfa, 0x5c, 0x5f, 0xbf,
	0x9d, 0xeb, 0xd7, 0xa3, 0x93, 0xfa, 0x99, 0x16, 0x2f, 0x1f, 0x0b, 0x8f, 0xb4, 0xc9, 0x2b, 0xbd,
	0x6e, 0x82, 0xc1, 0xdd, 0xc6, 0x86, 0xb0, 0x71, 0xce, 0x6c, 0xb3, 0xb3, 0x5b, 0x9c, 0xb7, 0xe7,
	0xcc, 0x13, 0xc4, 0x23, 0xae, 0xc1, 0xc0, 0x9e, 0x90, 0x10, 0x92, 0xb9, 0x7c, 0xb7, 0xd5, 0x90,
	0x43, 0x1e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xf7, 0x71, 0x85, 0xa9, 0x01, 0x00, 0x00,
}