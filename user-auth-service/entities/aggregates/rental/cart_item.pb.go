// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/rental/cart_item.proto

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

type CartItem struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Cart                 *Cart                `protobuf:"bytes,3,opt,name=cart,proto3" json:"cart,omitempty"`
	ItemUuid             string               `protobuf:"bytes,4,opt,name=item_uuid,json=itemUuid,proto3" json:"item_uuid,omitempty"`
	TotalItem            int32                `protobuf:"varint,5,opt,name=total_item,json=totalItem,proto3" json:"total_item,omitempty"`
	TotalPrice           int32                `protobuf:"varint,6,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	UnitPrice            int32                `protobuf:"varint,7,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	DurationFrom         *timestamp.Timestamp `protobuf:"bytes,8,opt,name=duration_from,json=durationFrom,proto3" json:"duration_from,omitempty"`
	DurationTo           *timestamp.Timestamp `protobuf:"bytes,9,opt,name=duration_to,json=durationTo,proto3" json:"duration_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CartItem) Reset()         { *m = CartItem{} }
func (m *CartItem) String() string { return proto.CompactTextString(m) }
func (*CartItem) ProtoMessage()    {}
func (*CartItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_8026613f915ecff3, []int{0}
}

func (m *CartItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartItem.Unmarshal(m, b)
}
func (m *CartItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartItem.Marshal(b, m, deterministic)
}
func (m *CartItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartItem.Merge(m, src)
}
func (m *CartItem) XXX_Size() int {
	return xxx_messageInfo_CartItem.Size(m)
}
func (m *CartItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CartItem.DiscardUnknown(m)
}

var xxx_messageInfo_CartItem proto.InternalMessageInfo

func (m *CartItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CartItem) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CartItem) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

func (m *CartItem) GetItemUuid() string {
	if m != nil {
		return m.ItemUuid
	}
	return ""
}

func (m *CartItem) GetTotalItem() int32 {
	if m != nil {
		return m.TotalItem
	}
	return 0
}

func (m *CartItem) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *CartItem) GetUnitPrice() int32 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *CartItem) GetDurationFrom() *timestamp.Timestamp {
	if m != nil {
		return m.DurationFrom
	}
	return nil
}

func (m *CartItem) GetDurationTo() *timestamp.Timestamp {
	if m != nil {
		return m.DurationTo
	}
	return nil
}

func init() {
	proto.RegisterType((*CartItem)(nil), "aggregates.rental.CartItem")
}

func init() { proto.RegisterFile("aggregates/rental/cart_item.proto", fileDescriptor_8026613f915ecff3) }

var fileDescriptor_8026613f915ecff3 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x18, 0xc4, 0x49, 0xfa, 0xe7, 0x6d, 0xb6, 0xaf, 0x82, 0xb9, 0x18, 0xaa, 0xd2, 0xe8, 0x29, 0x20,
	0xdd, 0x05, 0x3d, 0x7a, 0x10, 0x14, 0x44, 0x3d, 0x49, 0xa8, 0x17, 0x2f, 0x65, 0x9b, 0x6c, 0x93,
	0x87, 0x66, 0xb3, 0x65, 0xf3, 0xac, 0xdf, 0xd4, 0xef, 0x23, 0xbb, 0x6b, 0xf4, 0x50, 0xc1, 0x5b,
	0x98, 0x99, 0xdf, 0x64, 0xf2, 0x84, 0x9c, 0xf3, 0xaa, 0xd2, 0xa2, 0xe2, 0x28, 0x3a, 0xa6, 0x45,
	0x8b, 0xbc, 0x61, 0x05, 0xd7, 0xb8, 0x02, 0x14, 0x92, 0xee, 0xb4, 0x42, 0x15, 0x1f, 0xfd, 0x44,
	0xa8, 0x8f, 0xcc, 0xe6, 0x95, 0x52, 0x55, 0x23, 0x98, 0x0b, 0xac, 0xcd, 0x86, 0x21, 0x48, 0xd1,
	0x21, 0x97, 0x3b, 0xcf, 0xcc, 0x4e, 0x7f, 0xaf, 0xf5, 0xee, 0xc5, 0x47, 0x48, 0x26, 0xf7, 0x5c,
	0xe3, 0x13, 0x0a, 0x19, 0x1f, 0x92, 0x10, 0xca, 0x24, 0x48, 0x83, 0x6c, 0x90, 0x87, 0x50, 0xc6,
	0x31, 0x19, 0x1a, 0x03, 0x65, 0x12, 0xa6, 0x41, 0x16, 0xe5, 0xee, 0x39, 0xbe, 0x24, 0x43, 0x8b,
	0x27, 0x83, 0x34, 0xc8, 0xa6, 0x57, 0xc7, 0x74, 0x6f, 0x11, 0xb5, 0x75, 0xb9, 0x0b, 0xc5, 0x27,
	0x24, 0xb2, 0xeb, 0x57, 0xae, 0x65, 0xe8, 0x5a, 0x26, 0x56, 0x78, 0xb5, 0x4d, 0x67, 0x84, 0xa0,
	0x42, 0xde, 0xb8, 0x0f, 0x4c, 0x46, 0x69, 0x90, 0x8d, 0xf2, 0xc8, 0x29, 0x6e, 0xcc, 0x9c, 0x4c,
	0xbd, 0xbd, 0xd3, 0x50, 0x88, 0x64, 0xec, 0x7c, 0x4f, 0xbc, 0x58, 0xc5, 0xf2, 0xa6, 0x05, 0xfc,
	0xf2, 0xff, 0x79, 0xde, 0x2a, 0xde, 0xbe, 0x25, 0x07, 0xa5, 0xd1, 0x1c, 0x41, 0xb5, 0xab, 0x8d,
	0x56, 0x32, 0x99, 0xb8, 0xc5, 0x33, 0xea, 0x0f, 0x46, 0xfb, 0x83, 0xd1, 0x65, 0x7f, 0xb0, 0xfc,
	0x7f, 0x0f, 0x3c, 0x68, 0x25, 0xe3, 0x1b, 0x32, 0xfd, 0x2e, 0x40, 0x95, 0x44, 0x7f, 0xe2, 0xa4,
	0x8f, 0x2f, 0xd5, 0xdd, 0xf3, 0xdb, 0x63, 0x05, 0x58, 0x9b, 0x35, 0x2d, 0x94, 0x64, 0xd0, 0xd4,
	0x5c, 0xca, 0xba, 0x2c, 0xd9, 0xd6, 0x94, 0x7c, 0x0b, 0x0b, 0xd3, 0x09, 0xbd, 0xe0, 0x06, 0xeb,
	0x45, 0x27, 0xf4, 0x3b, 0x14, 0x82, 0x89, 0x16, 0x01, 0x41, 0x74, 0x6c, 0xef, 0x77, 0xad, 0xc7,
	0xee, 0x5d, 0xd7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x7c, 0x8d, 0x66, 0x21, 0x02, 0x00,
	0x00,
}
