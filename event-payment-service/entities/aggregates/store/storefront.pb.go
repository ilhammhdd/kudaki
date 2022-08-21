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

func init() {
	proto.RegisterType((*Storefront)(nil), "aggregates.store.Storefront")
}

func init() { proto.RegisterFile("aggregates/store/storefront.proto", fileDescriptor_11e209133ba8f19a) }

var fileDescriptor_11e209133ba8f19a = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x49, 0xff, 0xe1, 0x46, 0x10, 0xc9, 0x41, 0x96, 0x8a, 0xb8, 0x7a, 0xda, 0x4b, 0x13,
	0xd0, 0x93, 0x47, 0xbd, 0x89, 0x78, 0x59, 0xf5, 0xe2, 0xa5, 0xa4, 0xcd, 0x34, 0x1d, 0xda, 0x6c,
	0x4a, 0x32, 0x29, 0xf8, 0xe9, 0xfc, 0x6a, 0xb2, 0xd9, 0x56, 0xc1, 0x4b, 0xfe, 0xbc, 0xdf, 0x9b,
	0x99, 0xc7, 0xf0, 0x1b, 0x6d, 0x6d, 0x00, 0xab, 0x09, 0xa2, 0x8a, 0xe4, 0x03, 0xf4, 0xe7, 0x2a,
	0xf8, 0x96, 0xe4, 0x2e, 0x78, 0xf2, 0xe2, 0xfc, 0xcf, 0x22, 0x33, 0x9c, 0x5e, 0x5b, 0xef, 0xed,
	0x16, 0x54, 0xe6, 0x8b, 0xb4, 0x52, 0x84, 0x0e, 0x22, 0x69, 0xb7, 0xeb, 0x4b, 0x6e, 0xbf, 0x19,
	0xe7, 0x6f, 0xbf, 0x7d, 0xc4, 0x19, 0x1f, 0xa0, 0x29, 0x59, 0xc5, 0xea, 0x61, 0x33, 0x40, 0x23,
	0x04, 0x1f, 0xa5, 0x84, 0xa6, 0x1c, 0x54, 0xac, 0x2e, 0x9a, 0xfc, 0x16, 0x97, 0xbc, 0x48, 0x11,
	0xc2, 0x3c, 0x83, 0x61, 0x06, 0x27, 0x9d, 0xf0, 0xd1, 0xc1, 0x2b, 0xce, 0xc9, 0x93, 0xde, 0xce,
	0x91, 0xc0, 0x95, 0xa3, 0x8a, 0xd5, 0xe3, 0xa6, 0xc8, 0xca, 0x33, 0x81, 0x13, 0x17, 0x7c, 0x12,
	0x34, 0x61, 0x6b, 0xcb, 0x71, 0xc5, 0x6a, 0xd6, 0x1c, 0x7e, 0xe2, 0x81, 0xf3, 0x65, 0x00, 0x4d,
	0x60, 0xe6, 0x9a, 0xca, 0x49, 0xc5, 0xea, 0xd3, 0xbb, 0xa9, 0xec, 0xc3, 0xcb, 0x63, 0x78, 0xf9,
	0x7e, 0x0c, 0xdf, 0x14, 0x07, 0xf7, 0x23, 0x3d, 0xbd, 0x7e, 0xbe, 0x58, 0xa4, 0x75, 0x5a, 0xc8,
	0xa5, 0x77, 0x0a, 0xb7, 0x6b, 0xed, 0xdc, 0xda, 0x18, 0xb5, 0x49, 0x46, 0x6f, 0x70, 0x06, 0x7b,
	0x68, 0x69, 0xb6, 0xd3, 0x5f, 0xae, 0xbb, 0x23, 0x84, 0x3d, 0x2e, 0x41, 0x41, 0x4b, 0x48, 0x08,
	0x51, 0xfd, 0x5f, 0xea, 0x62, 0x92, 0xa7, 0xdd, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xe6,
	0x01, 0x86, 0x6f, 0x01, 0x00, 0x00,
}
