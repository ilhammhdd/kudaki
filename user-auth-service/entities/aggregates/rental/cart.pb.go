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
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xc9, 0x7e, 0xb9, 0x66, 0x20, 0x98, 0x53, 0x98, 0x83, 0x15, 0x4f, 0xbd, 0xac, 0x01,
	0x3d, 0x79, 0x54, 0x2f, 0xea, 0x49, 0x8a, 0x5e, 0xbc, 0x8c, 0xac, 0xf9, 0x6c, 0x3f, 0xd6, 0x2e,
	0x25, 0xfd, 0xb2, 0x3f, 0xdd, 0xb3, 0x24, 0x73, 0x54, 0xf0, 0xf6, 0xf1, 0x3c, 0x6f, 0x78, 0x79,
	0x09, 0x5f, 0xe9, 0xaa, 0x72, 0x50, 0x69, 0x82, 0x5e, 0x39, 0x38, 0x90, 0x6e, 0x54, 0xa9, 0x1d,
	0xe5, 0x9d, 0xb3, 0x64, 0xc5, 0xd5, 0x60, 0xf3, 0x93, 0x5d, 0xae, 0x2b, 0x6b, 0xab, 0x06, 0x54,
	0x0c, 0xec, 0xfc, 0x97, 0x22, 0x6c, 0xa1, 0x27, 0xdd, 0x76, 0xa7, 0x37, 0x37, 0xdf, 0x8c, 0x4f,
	0x9e, 0xb4, 0x23, 0x71, 0xc9, 0x47, 0x68, 0x24, 0x4b, 0x59, 0x36, 0x2e, 0x46, 0x68, 0x84, 0xe0,
	0x13, 0xef, 0xd1, 0xc8, 0x51, 0xca, 0xb2, 0xa4, 0x88, 0xb7, 0xb8, 0xe6, 0x89, 0xef, 0xc1, 0x6d,
	0xa3, 0x18, 0x47, 0x31, 0x0f, 0xe0, 0x23, 0xc8, 0x35, 0x5f, 0x90, 0x25, 0xdd, 0x6c, 0x3b, 0x87,
	0x25, 0xc8, 0x49, 0xca, 0xb2, 0x69, 0xc1, 0x23, 0x7a, 0x0b, 0x64, 0x08, 0x20, 0x41, 0xdb, 0xcb,
	0xe9, 0x9f, 0xc0, 0x4b, 0x20, 0xa1, 0xd2, 0x76, 0x70, 0x90, 0xb3, 0x94, 0x65, 0xf3, 0x22, 0xde,
	0x62, 0xc5, 0x13, 0x03, 0x0d, 0x1e, 0xc1, 0x81, 0x91, 0x17, 0x51, 0x0c, 0x40, 0xdc, 0x73, 0x5e,
	0x3a, 0xd0, 0x04, 0x66, 0xab, 0x49, 0xce, 0x53, 0x96, 0x2d, 0x6e, 0x97, 0xf9, 0x69, 0x73, 0x7e,
	0xde, 0x9c, 0xbf, 0x9f, 0x37, 0x17, 0xc9, 0x6f, 0xfa, 0x81, 0x1e, 0x5f, 0x3f, 0x9f, 0x2b, 0xa4,
	0xda, 0xef, 0xf2, 0xd2, 0xb6, 0x0a, 0x9b, 0x5a, 0xb7, 0x6d, 0x6d, 0x8c, 0xda, 0x7b, 0xa3, 0xf7,
	0xb8, 0x09, 0xb3, 0x36, 0xda, 0x53, 0xbd, 0xe9, 0xc1, 0x1d, 0xb1, 0x04, 0x05, 0x07, 0x42, 0x42,
	0xe8, 0xd5, 0xbf, 0x3f, 0xd8, 0xcd, 0x62, 0xd5, 0xdd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac,
	0xcd, 0x1c, 0x7c, 0x9f, 0x01, 0x00, 0x00,
}