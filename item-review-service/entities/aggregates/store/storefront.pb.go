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
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0x49, 0xbf, 0x70, 0x23, 0x88, 0xe4, 0x20, 0x4b, 0x45, 0x5c, 0x3d, 0xed, 0xa5, 0x09,
	0xe8, 0xc9, 0xa3, 0xde, 0x8a, 0xb7, 0x55, 0x2f, 0x5e, 0x4a, 0xda, 0x4c, 0xd3, 0xa1, 0xcd, 0xa6,
	0x24, 0x93, 0xfa, 0xf3, 0xfc, 0x6b, 0xb2, 0xd9, 0x56, 0xc1, 0x4b, 0x48, 0xde, 0xe7, 0x4d, 0xf2,
	0x30, 0xfc, 0x4e, 0x5b, 0x1b, 0xc0, 0x6a, 0x82, 0xa8, 0x22, 0xf9, 0x00, 0xfd, 0xba, 0x0e, 0xbe,
	0x25, 0xb9, 0x0f, 0x9e, 0xbc, 0xb8, 0xfc, 0xab, 0xc8, 0x0c, 0xa7, 0xb7, 0xd6, 0x7b, 0xbb, 0x03,
	0x95, 0xf9, 0x32, 0xad, 0x15, 0xa1, 0x83, 0x48, 0xda, 0xed, 0xfb, 0x2b, 0xf7, 0xdf, 0x8c, 0xf3,
	0xb7, 0xdf, 0x77, 0xc4, 0x05, 0x1f, 0xa0, 0x29, 0x59, 0xc5, 0xea, 0x61, 0x33, 0x40, 0x23, 0x04,
	0x1f, 0xa5, 0x84, 0xa6, 0x1c, 0x54, 0xac, 0x2e, 0x9a, 0xbc, 0x17, 0xd7, 0xbc, 0x48, 0x11, 0xc2,
	0x22, 0x83, 0x61, 0x06, 0x67, 0x5d, 0xf0, 0xd1, 0xc1, 0x1b, 0xce, 0xc9, 0x93, 0xde, 0x2d, 0x90,
	0xc0, 0x95, 0xa3, 0x8a, 0xd5, 0xe3, 0xa6, 0xc8, 0xc9, 0x9c, 0xc0, 0x89, 0x2b, 0x3e, 0x09, 0x9a,
	0xb0, 0xb5, 0xe5, 0xb8, 0x62, 0x35, 0x6b, 0x8e, 0x27, 0xf1, 0xc4, 0xf9, 0x2a, 0x80, 0x26, 0x30,
	0x0b, 0x4d, 0xe5, 0xa4, 0x62, 0xf5, 0xf9, 0xc3, 0x54, 0xf6, 0xf2, 0xf2, 0x24, 0x2f, 0xdf, 0x4f,
	0xf2, 0x4d, 0x71, 0x6c, 0x3f, 0xd3, 0xcb, 0xeb, 0xe7, 0xdc, 0x22, 0x6d, 0xd2, 0x52, 0xae, 0xbc,
	0x53, 0xb8, 0xdb, 0x68, 0xe7, 0x36, 0xc6, 0xa8, 0x6d, 0x32, 0x7a, 0x8b, 0xb3, 0xce, 0x63, 0x16,
	0xe0, 0x80, 0xf0, 0x35, 0x8b, 0x10, 0x0e, 0xb8, 0x02, 0x05, 0x2d, 0x21, 0x21, 0x44, 0xf5, 0x7f,
	0xa4, 0xcb, 0x49, 0xfe, 0xeb, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x61, 0x9d, 0x27, 0x1b, 0x6d,
	0x01, 0x00, 0x00,
}
