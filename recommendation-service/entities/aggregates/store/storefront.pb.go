// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/store/storefront.proto

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

type Storefront struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TotalItem            int32                `protobuf:"varint,4,opt,name=total_item,json=totalItem,proto3" json:"total_item,omitempty"`
	Rating               float64              `protobuf:"fixed64,5,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	TotalRawRating       float64              `protobuf:"fixed64,7,opt,name=total_raw_rating,json=totalRawRating,proto3" json:"total_raw_rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Storefront) Reset()         { *m = Storefront{} }
func (m *Storefront) String() string { return proto.CompactTextString(m) }
func (*Storefront) ProtoMessage()    {}
func (*Storefront) Descriptor() ([]byte, []int) {
	return fileDescriptor_11e209133ba8f19a, []int{0}
}

func (m *Storefront) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storefront.Unmarshal(m, b)
}
func (m *Storefront) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storefront.Marshal(b, m, deterministic)
}
func (m *Storefront) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storefront.Merge(m, src)
}
func (m *Storefront) XXX_Size() int {
	return xxx_messageInfo_Storefront.Size(m)
}
func (m *Storefront) XXX_DiscardUnknown() {
	xxx_messageInfo_Storefront.DiscardUnknown(m)
}

var xxx_messageInfo_Storefront proto.InternalMessageInfo

func (m *Storefront) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Storefront) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Storefront) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Storefront) GetTotalItem() int32 {
	if m != nil {
		return m.TotalItem
	}
	return 0
}

func (m *Storefront) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Storefront) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Storefront) GetTotalRawRating() float64 {
	if m != nil {
		return m.TotalRawRating
	}
	return 0
}

func init() {
	proto.RegisterType((*Storefront)(nil), "aggregates.store.Storefront")
}

func init() { proto.RegisterFile("aggregates/store/storefront.proto", fileDescriptor_11e209133ba8f19a) }

var fileDescriptor_11e209133ba8f19a = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0x49, 0xff, 0xbe, 0x6f, 0x22, 0x94, 0x92, 0x85, 0x0c, 0x15, 0x71, 0x74, 0x35, 0x9b,
	0x4e, 0x40, 0x57, 0x2e, 0x75, 0x27, 0x88, 0x8b, 0xa8, 0x1b, 0x37, 0x43, 0x3a, 0x39, 0x4d, 0x0f,
	0x6d, 0x26, 0x25, 0x73, 0x62, 0xaf, 0xd9, 0xbb, 0x90, 0x66, 0x5a, 0x05, 0x37, 0x21, 0x79, 0x9f,
	0x93, 0x37, 0x0f, 0xe1, 0xd7, 0xda, 0xda, 0x00, 0x56, 0x13, 0x74, 0xb2, 0x23, 0x1f, 0xa0, 0x5f,
	0x57, 0xc1, 0xb7, 0x54, 0xed, 0x82, 0x27, 0x2f, 0x66, 0xbf, 0x23, 0x55, 0x82, 0xf3, 0x2b, 0xeb,
	0xbd, 0xdd, 0x82, 0x4c, 0x7c, 0x19, 0x57, 0x92, 0xd0, 0x41, 0x47, 0xda, 0xed, 0xfa, 0x2b, 0x37,
	0x5f, 0x8c, 0xf3, 0xd7, 0x9f, 0x1e, 0x31, 0xe5, 0x03, 0x34, 0x39, 0x2b, 0x58, 0x39, 0x54, 0x03,
	0x34, 0x42, 0xf0, 0x51, 0x8c, 0x68, 0xf2, 0x41, 0xc1, 0xca, 0x4c, 0xa5, 0xbd, 0xb8, 0xe0, 0x59,
	0xec, 0x20, 0xd4, 0x09, 0x0c, 0x13, 0xf8, 0x7f, 0x08, 0xde, 0x0f, 0xf0, 0x92, 0x73, 0xf2, 0xa4,
	0xb7, 0x35, 0x12, 0xb8, 0x7c, 0x54, 0xb0, 0x72, 0xac, 0xb2, 0x94, 0x3c, 0x11, 0x38, 0x71, 0xce,
	0x27, 0x41, 0x13, 0xb6, 0x36, 0x1f, 0x17, 0xac, 0x64, 0xea, 0x78, 0x12, 0xf7, 0x9c, 0x37, 0x01,
	0x34, 0x81, 0xa9, 0x35, 0xe5, 0x93, 0x82, 0x95, 0x67, 0xb7, 0xf3, 0xaa, 0x97, 0xaf, 0x4e, 0xf2,
	0xd5, 0xdb, 0x49, 0x5e, 0x65, 0xc7, 0xe9, 0x07, 0x12, 0x25, 0x9f, 0xf5, 0x2f, 0x06, 0xbd, 0xaf,
	0x8f, 0xe5, 0xff, 0x52, 0xf9, 0x34, 0xe5, 0x4a, 0xef, 0x55, 0x4a, 0x1f, 0x5f, 0x3e, 0x9e, 0x2d,
	0xd2, 0x3a, 0x2e, 0xab, 0xc6, 0x3b, 0x89, 0xdb, 0xb5, 0x76, 0x6e, 0x6d, 0x8c, 0xdc, 0x44, 0xa3,
	0x37, 0xb8, 0x08, 0xd0, 0x78, 0xe7, 0xa0, 0x35, 0x9a, 0xd0, 0xb7, 0x8b, 0x0e, 0xc2, 0x27, 0x36,
	0x20, 0xa1, 0x25, 0x24, 0x84, 0x4e, 0xfe, 0xfd, 0xff, 0xe5, 0x24, 0x89, 0xdd, 0x7d, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x34, 0xa0, 0x87, 0xfe, 0x9a, 0x01, 0x00, 0x00,
}