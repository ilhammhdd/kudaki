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
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x49, 0xff, 0x68, 0x53, 0x10, 0x9c, 0xd5, 0x50, 0x85, 0x0e, 0xae, 0x66, 0xd3, 0x04,
	0x74, 0xe5, 0xd2, 0x1f, 0x04, 0x77, 0x32, 0xe8, 0xc6, 0x4d, 0x49, 0x27, 0xd7, 0xf4, 0xd2, 0xc9,
	0x64, 0x48, 0x6e, 0xfa, 0xac, 0x3e, 0x8e, 0x4c, 0xc6, 0x52, 0xc1, 0xdd, 0xe5, 0x7c, 0x5f, 0x38,
	0x9c, 0xf0, 0x6b, 0x65, 0x8c, 0x07, 0xa3, 0x08, 0x82, 0xf4, 0xd0, 0x92, 0x6a, 0x64, 0xad, 0x3c,
	0x89, 0xce, 0x3b, 0x72, 0xd9, 0xe5, 0x99, 0x8a, 0x81, 0xae, 0xd6, 0xc6, 0x39, 0xd3, 0x80, 0x4c,
	0xc2, 0x2e, 0x7e, 0x49, 0x42, 0x0b, 0x81, 0x94, 0xed, 0x86, 0x37, 0x37, 0xdf, 0x8c, 0x4f, 0x9e,
	0x94, 0xa7, 0xec, 0x82, 0x8f, 0x50, 0xe7, 0xac, 0x60, 0xe5, 0xb8, 0x1a, 0xa1, 0xce, 0x32, 0x3e,
	0x89, 0x11, 0x75, 0x3e, 0x2a, 0x58, 0xb9, 0xa8, 0xd2, 0x9d, 0x5d, 0xf1, 0x45, 0x0c, 0xe0, 0xb7,
	0x09, 0x8c, 0x13, 0x98, 0xf7, 0xc1, 0x47, 0x0f, 0xd7, 0x7c, 0x49, 0x8e, 0x54, 0xb3, 0xed, 0x3c,
	0xd6, 0x90, 0x4f, 0x0a, 0x56, 0x4e, 0x2b, 0x9e, 0xa2, 0xb7, 0x3e, 0x39, 0x0b, 0x48, 0x60, 0x43,
	0x3e, 0xfd, 0x23, 0xbc, 0xf6, 0x49, 0x5f, 0xe9, 0x3a, 0x68, 0xf3, 0x59, 0xc1, 0xca, 0x79, 0x95,
	0xee, 0xec, 0x9e, 0xf3, 0xda, 0x83, 0x22, 0xd0, 0x5b, 0x45, 0xf9, 0xbc, 0x60, 0xe5, 0xf2, 0x76,
	0x25, 0x86, 0x55, 0xe2, 0xb4, 0x4a, 0xbc, 0x9f, 0x56, 0x55, 0x8b, 0x5f, 0xfb, 0x81, 0x1e, 0x5f,
	0x3e, 0x9f, 0x0d, 0xd2, 0x3e, 0xee, 0x44, 0xed, 0xac, 0xc4, 0x66, 0xaf, 0xac, 0xdd, 0x6b, 0x2d,
	0x0f, 0x51, 0xab, 0x03, 0x6e, 0xe0, 0x08, 0x2d, 0x6d, 0x02, 0xf8, 0x23, 0xd6, 0x20, 0xa1, 0x25,
	0x24, 0x84, 0x20, 0xff, 0xfd, 0xf0, 0x6e, 0x96, 0x6a, 0xee, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x38, 0xb4, 0x73, 0xb9, 0x7d, 0x01, 0x00, 0x00,
}
