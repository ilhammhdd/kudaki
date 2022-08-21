// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/order/owner_order.proto

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

type OwnerOrder struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Order                *Order               `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	OwnerUuid            string               `protobuf:"bytes,4,opt,name=owner_uuid,json=ownerUuid,proto3" json:"owner_uuid,omitempty"`
	TotalPrice           int32                `protobuf:"varint,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	TotalQuantity        int64                `protobuf:"varint,6,opt,name=total_quantity,json=totalQuantity,proto3" json:"total_quantity,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	OrderStatus          OrderStatus          `protobuf:"varint,8,opt,name=order_status,json=orderStatus,proto3,enum=aggregates.order.OrderStatus" json:"order_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OwnerOrder) Reset()         { *m = OwnerOrder{} }
func (m *OwnerOrder) String() string { return proto.CompactTextString(m) }
func (*OwnerOrder) ProtoMessage()    {}
func (*OwnerOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d297f26665e56f, []int{0}
}

func (m *OwnerOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerOrder.Unmarshal(m, b)
}
func (m *OwnerOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerOrder.Marshal(b, m, deterministic)
}
func (m *OwnerOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerOrder.Merge(m, src)
}
func (m *OwnerOrder) XXX_Size() int {
	return xxx_messageInfo_OwnerOrder.Size(m)
}
func (m *OwnerOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerOrder.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerOrder proto.InternalMessageInfo

func (m *OwnerOrder) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OwnerOrder) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *OwnerOrder) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OwnerOrder) GetOwnerUuid() string {
	if m != nil {
		return m.OwnerUuid
	}
	return ""
}

func (m *OwnerOrder) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *OwnerOrder) GetTotalQuantity() int64 {
	if m != nil {
		return m.TotalQuantity
	}
	return 0
}

func (m *OwnerOrder) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *OwnerOrder) GetOrderStatus() OrderStatus {
	if m != nil {
		return m.OrderStatus
	}
	return OrderStatus_PENDING
}

func init() {
	proto.RegisterType((*OwnerOrder)(nil), "aggregates.order.OwnerOrder")
}

func init() { proto.RegisterFile("aggregates/order/owner_order.proto", fileDescriptor_45d297f26665e56f) }

var fileDescriptor_45d297f26665e56f = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0xf7, 0x43, 0x97, 0xe9, 0x90, 0x5c, 0x0c, 0xc3, 0xb1, 0x32, 0x10, 0x7a, 0x59,
	0x02, 0xf3, 0xe4, 0x4d, 0x3d, 0x88, 0xb7, 0x69, 0xd5, 0x8b, 0x97, 0x92, 0x35, 0xb1, 0x0b, 0x5b,
	0x9b, 0x99, 0xbe, 0x28, 0xfe, 0x7d, 0xfe, 0x63, 0xd2, 0xd7, 0x8d, 0xc1, 0xc4, 0x5b, 0xfa, 0xe9,
	0xb7, 0x9f, 0x7e, 0xdf, 0x0b, 0x99, 0xc8, 0x3c, 0x77, 0x3a, 0x97, 0xa0, 0x2b, 0x61, 0x9d, 0xd2,
	0x4e, 0xd8, 0xaf, 0x52, 0xbb, 0x14, 0xcf, 0x7c, 0xe3, 0x2c, 0x58, 0x7a, 0xb6, 0xcf, 0x70, 0xe4,
	0xc3, 0x71, 0x6e, 0x6d, 0xbe, 0xd6, 0x02, 0xdf, 0x2f, 0xfc, 0xbb, 0x00, 0x53, 0xe8, 0x0a, 0x64,
	0xb1, 0x69, 0x3e, 0x19, 0x5e, 0xfc, 0xd5, 0xee, 0x85, 0x93, 0x9f, 0x90, 0x90, 0x79, 0xfd, 0x9b,
	0x79, 0x0d, 0xe9, 0x80, 0x84, 0x46, 0xb1, 0x20, 0x0a, 0xe2, 0x56, 0x12, 0x1a, 0x45, 0x29, 0x69,
	0x7b, 0x6f, 0x14, 0x0b, 0xa3, 0x20, 0xee, 0x25, 0x78, 0xa6, 0x53, 0xd2, 0x41, 0x03, 0x6b, 0x45,
	0x41, 0xdc, 0x9f, 0x9d, 0xf3, 0xc3, 0x4e, 0x1c, 0x5d, 0x49, 0x93, 0xa2, 0x23, 0x42, 0x9a, 0x39,
	0x50, 0xd4, 0x46, 0x51, 0x0f, 0xc9, 0x6b, 0x6d, 0x1b, 0x93, 0x3e, 0x58, 0x90, 0xeb, 0x74, 0xe3,
	0x4c, 0xa6, 0x59, 0x27, 0x0a, 0xe2, 0x4e, 0x42, 0x10, 0x3d, 0xd6, 0x84, 0x5e, 0x92, 0x41, 0x13,
	0xf8, 0xf0, 0xb2, 0x04, 0x03, 0xdf, 0xac, 0x8b, 0xf5, 0x4e, 0x91, 0x3e, 0x6d, 0x21, 0xbd, 0x26,
	0x24, 0x73, 0x5a, 0x82, 0x56, 0xa9, 0x04, 0x76, 0x84, 0xd5, 0x86, 0xbc, 0x59, 0x0e, 0xdf, 0x2d,
	0x87, 0xbf, 0xec, 0x96, 0x93, 0xf4, 0xb6, 0xe9, 0x5b, 0xa0, 0x37, 0xe4, 0x04, 0xab, 0xa6, 0x15,
	0x48, 0xf0, 0x15, 0x3b, 0x8e, 0x82, 0x78, 0x30, 0x1b, 0xfd, 0x33, 0xd7, 0x33, 0x86, 0x92, 0xbe,
	0xdd, 0x3f, 0xdc, 0x3d, 0xbc, 0xdd, 0xe7, 0x06, 0x96, 0x7e, 0xc1, 0x33, 0x5b, 0x08, 0xb3, 0x5e,
	0xca, 0xa2, 0x58, 0x2a, 0x25, 0x56, 0x5e, 0xc9, 0x95, 0x99, 0x16, 0xd6, 0x97, 0x20, 0x4d, 0x39,
	0xad, 0xb4, 0xfb, 0x34, 0x99, 0x16, 0xba, 0xee, 0x6c, 0x74, 0x25, 0x0e, 0xef, 0x66, 0xd1, 0xc5,
	0xaa, 0x57, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x56, 0x80, 0x28, 0x0d, 0x02, 0x00, 0x00,
}