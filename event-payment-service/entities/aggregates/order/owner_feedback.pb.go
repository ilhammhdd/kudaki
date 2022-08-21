// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/order/owner_feedback.proto

package order

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

type OwnerFeedback struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Order                *Order               `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
	Rating               float64              `protobuf:"fixed64,5,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OwnerFeedback) Reset()         { *m = OwnerFeedback{} }
func (m *OwnerFeedback) String() string { return proto.CompactTextString(m) }
func (*OwnerFeedback) ProtoMessage()    {}
func (*OwnerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac2c0bc013486200, []int{0}
}

func (m *OwnerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerFeedback.Unmarshal(m, b)
}
func (m *OwnerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerFeedback.Marshal(b, m, deterministic)
}
func (m *OwnerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerFeedback.Merge(m, src)
}
func (m *OwnerFeedback) XXX_Size() int {
	return xxx_messageInfo_OwnerFeedback.Size(m)
}
func (m *OwnerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerFeedback proto.InternalMessageInfo

func (m *OwnerFeedback) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OwnerFeedback) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *OwnerFeedback) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *OwnerFeedback) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OwnerFeedback) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *OwnerFeedback) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*OwnerFeedback)(nil), "aggregates.order.OwnerFeedback")
}

func init() {
	proto.RegisterFile("aggregates/order/owner_feedback.proto", fileDescriptor_ac2c0bc013486200)
}

var fileDescriptor_ac2c0bc013486200 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4b, 0x33, 0x31,
	0x14, 0x85, 0x49, 0xbf, 0x78, 0x9b, 0x17, 0x45, 0xb2, 0xd0, 0xa1, 0x0a, 0x0e, 0x82, 0x30, 0x9b,
	0x26, 0xa0, 0x2b, 0x97, 0xba, 0x70, 0x23, 0x52, 0x18, 0x74, 0xe3, 0xa6, 0x64, 0x26, 0xb7, 0x69,
	0x68, 0x33, 0x29, 0x99, 0x9b, 0x8a, 0x7f, 0xd3, 0x5f, 0x24, 0xc9, 0x4c, 0x11, 0xea, 0xe6, 0xe6,
	0xe3, 0x39, 0xb9, 0x27, 0xe7, 0xd2, 0x5b, 0xa9, 0xb5, 0x07, 0x2d, 0x11, 0x5a, 0xe1, 0xbc, 0x02,
	0x2f, 0xdc, 0x67, 0x03, 0x7e, 0xb9, 0x02, 0x50, 0x95, 0xac, 0x37, 0x7c, 0xe7, 0x1d, 0x3a, 0x76,
	0xf6, 0x2b, 0xe3, 0x49, 0x36, 0xbb, 0xd6, 0xce, 0xe9, 0x2d, 0x88, 0xc4, 0xab, 0xb0, 0x12, 0x68,
	0x2c, 0xb4, 0x28, 0xed, 0xae, 0x7b, 0x32, 0xbb, 0xfa, 0xdb, 0x39, 0xd6, 0x8e, 0xde, 0x7c, 0x13,
	0x7a, 0xb2, 0x88, 0x4e, 0xcf, 0xbd, 0x11, 0x3b, 0xa5, 0x03, 0xa3, 0x32, 0x92, 0x93, 0x62, 0x58,
	0x0e, 0x8c, 0x62, 0x8c, 0x8e, 0x42, 0x30, 0x2a, 0x1b, 0xe4, 0xa4, 0x98, 0x96, 0x69, 0xcf, 0x2e,
	0xe9, 0x34, 0xb4, 0xe0, 0x97, 0x09, 0x0c, 0x13, 0xf8, 0x17, 0x2f, 0xde, 0x23, 0x9c, 0xd3, 0x71,
	0x72, 0xc8, 0x46, 0x39, 0x29, 0xfe, 0xdf, 0x5d, 0xf0, 0xe3, 0x3f, 0xf3, 0x45, 0xac, 0x65, 0xa7,
	0x62, 0xe7, 0x74, 0xe2, 0x25, 0x9a, 0x46, 0x67, 0xe3, 0x9c, 0x14, 0xa4, 0xec, 0x4f, 0xec, 0x81,
	0xd2, 0xda, 0x83, 0x44, 0x50, 0x4b, 0x89, 0xd9, 0x24, 0xf5, 0x9a, 0xf1, 0x2e, 0x2d, 0x3f, 0xa4,
	0xe5, 0x6f, 0x87, 0xb4, 0xe5, 0xb4, 0x57, 0x3f, 0xe2, 0xd3, 0xeb, 0xc7, 0x8b, 0x36, 0xb8, 0x0e,
	0x15, 0xaf, 0x9d, 0x15, 0x66, 0xbb, 0x96, 0xd6, 0xae, 0x95, 0x12, 0x9b, 0xa0, 0xe4, 0xc6, 0xcc,
	0x61, 0x0f, 0x0d, 0xce, 0x77, 0xf2, 0xcb, 0xc6, 0xb5, 0x05, 0xbf, 0x37, 0x35, 0x08, 0x68, 0xd0,
	0xa0, 0x81, 0x56, 0x1c, 0xcf, 0xab, 0x9a, 0x24, 0xb7, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc3, 0x55, 0xea, 0x99, 0xa4, 0x01, 0x00, 0x00,
}