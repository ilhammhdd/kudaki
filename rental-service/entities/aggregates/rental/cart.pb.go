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
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0x7e, 0xb9, 0x66, 0x20, 0x98, 0x53, 0x98, 0x83, 0x15, 0x4f, 0xbd, 0xac, 0x01,
	0x3d, 0x79, 0x54, 0x11, 0xf1, 0x26, 0x45, 0x2f, 0x5e, 0x46, 0xd6, 0x3c, 0xbb, 0xc7, 0xda, 0xa5,
	0x24, 0x2f, 0xfb, 0xd3, 0x3d, 0x4b, 0x53, 0x47, 0x05, 0x6f, 0x8f, 0xcf, 0xe7, 0x1b, 0x1e, 0xdf,
	0x3c, 0xbe, 0xd2, 0x55, 0xe5, 0xa0, 0xd2, 0x04, 0x5e, 0x39, 0x38, 0x92, 0xae, 0x55, 0xa9, 0x1d,
	0xe5, 0xad, 0xb3, 0x64, 0xc5, 0xd5, 0x60, 0xf3, 0xde, 0x2e, 0xd7, 0x95, 0xb5, 0x55, 0x0d, 0x2a,
	0x06, 0x76, 0xe1, 0x4b, 0x11, 0x36, 0xe0, 0x49, 0x37, 0x6d, 0xff, 0xe6, 0xe6, 0x9b, 0xf1, 0xc9,
	0x93, 0x76, 0x24, 0x2e, 0xf9, 0x08, 0x8d, 0x64, 0x29, 0xcb, 0xc6, 0xc5, 0x08, 0x8d, 0x10, 0x7c,
	0x12, 0x02, 0x1a, 0x39, 0x4a, 0x59, 0x96, 0x14, 0x71, 0x16, 0xd7, 0x3c, 0x09, 0x1e, 0xdc, 0x36,
	0x8a, 0x71, 0x14, 0xf3, 0x0e, 0x7c, 0x74, 0x72, 0xcd, 0x17, 0x64, 0x49, 0xd7, 0xdb, 0xd6, 0x61,
	0x09, 0x72, 0x92, 0xb2, 0x6c, 0x5a, 0xf0, 0x88, 0xde, 0x3a, 0x32, 0x04, 0x90, 0xa0, 0xf1, 0x72,
	0xfa, 0x27, 0xf0, 0xda, 0x91, 0x6e, 0xa5, 0x6d, 0xe1, 0x28, 0x67, 0x29, 0xcb, 0xe6, 0x45, 0x9c,
	0xc5, 0x8a, 0x27, 0x06, 0x6a, 0x3c, 0x81, 0x03, 0x23, 0x2f, 0xa2, 0x18, 0x80, 0xb8, 0xe7, 0xbc,
	0x74, 0xa0, 0x09, 0xcc, 0x56, 0x93, 0x9c, 0xa7, 0x2c, 0x5b, 0xdc, 0x2e, 0xf3, 0xbe, 0x73, 0x7e,
	0xee, 0x9c, 0xbf, 0x9f, 0x3b, 0x17, 0xc9, 0x6f, 0xfa, 0x81, 0x1e, 0x5f, 0x3e, 0x9f, 0x2b, 0xa4,
	0x7d, 0xd8, 0xe5, 0xa5, 0x6d, 0x14, 0xd6, 0x7b, 0xdd, 0x34, 0x7b, 0x63, 0xd4, 0x21, 0x18, 0x7d,
	0xc0, 0x4d, 0xff, 0x7f, 0x1b, 0x0f, 0xee, 0x84, 0x25, 0x28, 0x38, 0x12, 0x12, 0x82, 0x57, 0xff,
	0x0e, 0xb0, 0x9b, 0xc5, 0x3d, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xa3, 0x2b, 0xb0,
	0x9c, 0x01, 0x00, 0x00,
}
