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
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x49, 0xff, 0xf1, 0x35, 0x1f, 0x8a, 0x64, 0xa1, 0x43, 0x15, 0x1c, 0x04, 0x61, 0x36,
	0x93, 0x80, 0xae, 0x5c, 0xea, 0x42, 0x10, 0x17, 0x85, 0x41, 0x37, 0x6e, 0x4a, 0x66, 0x72, 0x9b,
	0x5e, 0xda, 0x69, 0x4a, 0xe6, 0xa6, 0x7d, 0x4f, 0x9f, 0x48, 0x26, 0x33, 0x45, 0xa8, 0x9b, 0x4b,
	0x92, 0xdf, 0xc9, 0x39, 0x39, 0xe1, 0xf7, 0xda, 0x5a, 0x0f, 0x56, 0x13, 0x34, 0xca, 0x79, 0x03,
	0x5e, 0xb9, 0xc3, 0x16, 0xfc, 0x62, 0x09, 0x60, 0x4a, 0x5d, 0xad, 0xe5, 0xce, 0x3b, 0x72, 0xe2,
	0xe2, 0x57, 0x26, 0xa3, 0x6c, 0x76, 0x6b, 0x9d, 0xb3, 0x1b, 0x50, 0x91, 0x97, 0x61, 0xa9, 0x08,
	0x6b, 0x68, 0x48, 0xd7, 0xbb, 0xee, 0xca, 0xec, 0xe6, 0xaf, 0x73, 0x3b, 0x3b, 0x7a, 0xf7, 0xcd,
	0xf8, 0xd9, 0xbc, 0x4d, 0x7a, 0xed, 0x83, 0xc4, 0x39, 0x1f, 0xa0, 0x49, 0x58, 0xca, 0xb2, 0x61,
	0x31, 0x40, 0x23, 0x04, 0x1f, 0x85, 0x80, 0x26, 0x19, 0xa4, 0x2c, 0x9b, 0x16, 0x71, 0x2d, 0xae,
	0xf9, 0x34, 0x34, 0xe0, 0x17, 0x11, 0x0c, 0x23, 0xf8, 0xd7, 0x1e, 0x7c, 0xb6, 0x30, 0xe7, 0xe3,
	0x98, 0x90, 0x8c, 0x52, 0x96, 0xfd, 0x7f, 0xb8, 0x92, 0xa7, 0x6f, 0x96, 0xf3, 0x76, 0x16, 0x9d,
	0x4a, 0x5c, 0xf2, 0x89, 0xd7, 0x84, 0x5b, 0x9b, 0x8c, 0x53, 0x96, 0xb1, 0xa2, 0xdf, 0x89, 0x27,
	0xce, 0x2b, 0x0f, 0x9a, 0xc0, 0x2c, 0x34, 0x25, 0x93, 0xe8, 0x35, 0x93, 0x5d, 0x5b, 0x79, 0x6c,
	0x2b, 0x3f, 0x8e, 0x6d, 0x8b, 0x69, 0xaf, 0x7e, 0xa6, 0x97, 0xf7, 0xaf, 0x37, 0x8b, 0xb4, 0x0a,
	0xa5, 0xac, 0x5c, 0xad, 0x70, 0xb3, 0xd2, 0x75, 0xbd, 0x32, 0x46, 0xad, 0x83, 0xd1, 0x6b, 0xcc,
	0x91, 0xa0, 0xce, 0x3d, 0xec, 0x11, 0x0e, 0x79, 0x03, 0x7e, 0x8f, 0x15, 0x28, 0xd8, 0x12, 0x12,
	0x42, 0xa3, 0x4e, 0x7f, 0xab, 0x9c, 0xc4, 0xac, 0xc7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87,
	0x0b, 0x0a, 0x55, 0xa2, 0x01, 0x00, 0x00,
}