// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/rental/cart.proto

package rental

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

type Cart struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TotalPrice           int32                `protobuf:"varint,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	TotalItems           int32                `protobuf:"varint,5,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	Open                 bool                 `protobuf:"varint,6,opt,name=open,proto3" json:"open,omitempty"`
	Delivered            bool                 `protobuf:"varint,7,opt,name=delivered,proto3" json:"delivered,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef36b31966c981a, []int{0}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cart) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Cart) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Cart) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *Cart) GetTotalItems() int32 {
	if m != nil {
		return m.TotalItems
	}
	return 0
}

func (m *Cart) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

func (m *Cart) GetDelivered() bool {
	if m != nil {
		return m.Delivered
	}
	return false
}

func (m *Cart) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Cart)(nil), "aggregates.rental.Cart")
}

func init() { proto.RegisterFile("aggregates/rental/cart.proto", fileDescriptor_7ef36b31966c981a) }

var fileDescriptor_7ef36b31966c981a = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x49, 0x7f, 0xd9, 0x4d, 0x41, 0x30, 0xa7, 0x50, 0x0b, 0x5d, 0x3c, 0xed, 0xa5, 0x1b,
	0xd0, 0x93, 0x47, 0x2d, 0x1e, 0xbc, 0xc9, 0xa2, 0x17, 0x2f, 0x25, 0xdd, 0x8c, 0xdb, 0xa1, 0xbb,
	0xcd, 0x92, 0x4c, 0xfa, 0xa7, 0x7b, 0x96, 0xa4, 0x96, 0x0a, 0xde, 0x86, 0xef, 0x7b, 0x61, 0xf2,
	0x86, 0x2f, 0x74, 0xd3, 0x38, 0x68, 0x34, 0x81, 0x57, 0x0e, 0x0e, 0xa4, 0x5b, 0x55, 0x6b, 0x47,
	0x65, 0xef, 0x2c, 0x59, 0x71, 0x73, 0xb1, 0xe5, 0xc9, 0xce, 0x97, 0x8d, 0xb5, 0x4d, 0x0b, 0x2a,
	0x05, 0xb6, 0xe1, 0x4b, 0x11, 0x76, 0xe0, 0x49, 0x77, 0xfd, 0xe9, 0xcd, 0xdd, 0x37, 0xe3, 0xa3,
	0xb5, 0x76, 0x24, 0xae, 0xf9, 0x00, 0x8d, 0x64, 0x39, 0x2b, 0x86, 0xd5, 0x00, 0x8d, 0x10, 0x7c,
	0x14, 0x02, 0x1a, 0x39, 0xc8, 0x59, 0x91, 0x55, 0x69, 0x16, 0xb7, 0x3c, 0x0b, 0x1e, 0xdc, 0x26,
	0x89, 0x61, 0x12, 0xd3, 0x08, 0x3e, 0xa2, 0x5c, 0xf2, 0x19, 0x59, 0xd2, 0xed, 0xa6, 0x77, 0x58,
	0x83, 0x1c, 0xe5, 0xac, 0x18, 0x57, 0x3c, 0xa1, 0xb7, 0x48, 0x2e, 0x01, 0x24, 0xe8, 0xbc, 0x1c,
	0xff, 0x09, 0xbc, 0x46, 0x12, 0x57, 0xda, 0x1e, 0x0e, 0x72, 0x92, 0xb3, 0x62, 0x5a, 0xa5, 0x59,
	0x2c, 0x78, 0x66, 0xa0, 0xc5, 0x23, 0x38, 0x30, 0xf2, 0x2a, 0x89, 0x0b, 0x10, 0x8f, 0x9c, 0xd7,
	0x0e, 0x34, 0x81, 0xd9, 0x68, 0x92, 0xd3, 0x9c, 0x15, 0xb3, 0xfb, 0x79, 0x79, 0xea, 0x5c, 0x9e,
	0x3b, 0x97, 0xef, 0xe7, 0xce, 0x55, 0xf6, 0x9b, 0x7e, 0xa2, 0xe7, 0x97, 0xcf, 0x75, 0x83, 0xb4,
	0x0b, 0xdb, 0xb2, 0xb6, 0x9d, 0xc2, 0x76, 0xa7, 0xbb, 0x6e, 0x67, 0x8c, 0xda, 0x07, 0xa3, 0xf7,
	0xb8, 0x8a, 0x7f, 0x5c, 0x79, 0x70, 0x47, 0xac, 0x41, 0xc1, 0x81, 0x90, 0x10, 0xbc, 0xfa, 0x77,
	0xfe, 0xed, 0x24, 0x6d, 0x79, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0x98, 0x6c, 0xc1, 0xec, 0x9a,
	0x01, 0x00, 0x00,
}