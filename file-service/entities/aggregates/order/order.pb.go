// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/order/order.proto

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

type OrderStatus int32

const (
	OrderStatus_PENDING     OrderStatus = 0
	OrderStatus_APPROVED    OrderStatus = 1
	OrderStatus_DISAPPROVED OrderStatus = 2
	OrderStatus_RENTED      OrderStatus = 4
	OrderStatus_DONE        OrderStatus = 5
)

var OrderStatus_name = map[int32]string{
	0: "PENDING",
	1: "APPROVED",
	2: "DISAPPROVED",
	4: "RENTED",
	5: "DONE",
}

var OrderStatus_value = map[string]int32{
	"PENDING":     0,
	"APPROVED":    1,
	"DISAPPROVED": 2,
	"RENTED":      4,
	"DONE":        5,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2a6b118f4f9db3fd, []int{0}
}

type Order struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	CartUuid             string               `protobuf:"bytes,3,opt,name=cart_uuid,json=cartUuid,proto3" json:"cart_uuid,omitempty"`
	OrderNum             string               `protobuf:"bytes,4,opt,name=order_num,json=orderNum,proto3" json:"order_num,omitempty"`
	AddressUuid          string               `protobuf:"bytes,5,opt,name=address_uuid,json=addressUuid,proto3" json:"address_uuid,omitempty"`
	Status               OrderStatus          `protobuf:"varint,6,opt,name=status,proto3,enum=aggregates.order.OrderStatus" json:"status,omitempty"`
	ShipmentFee          int32                `protobuf:"varint,7,opt,name=shipment_fee,json=shipmentFee,proto3" json:"shipment_fee,omitempty"`
	Delivered            bool                 `protobuf:"varint,8,opt,name=delivered,proto3" json:"delivered,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	TenantUuid           string               `protobuf:"bytes,11,opt,name=tenant_uuid,json=tenantUuid,proto3" json:"tenant_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6b118f4f9db3fd, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Order) GetCartUuid() string {
	if m != nil {
		return m.CartUuid
	}
	return ""
}

func (m *Order) GetOrderNum() string {
	if m != nil {
		return m.OrderNum
	}
	return ""
}

func (m *Order) GetAddressUuid() string {
	if m != nil {
		return m.AddressUuid
	}
	return ""
}

func (m *Order) GetStatus() OrderStatus {
	if m != nil {
		return m.Status
	}
	return OrderStatus_PENDING
}

func (m *Order) GetShipmentFee() int32 {
	if m != nil {
		return m.ShipmentFee
	}
	return 0
}

func (m *Order) GetDelivered() bool {
	if m != nil {
		return m.Delivered
	}
	return false
}

func (m *Order) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Order) GetTenantUuid() string {
	if m != nil {
		return m.TenantUuid
	}
	return ""
}

func init() {
	proto.RegisterEnum("aggregates.order.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterType((*Order)(nil), "aggregates.order.Order")
}

func init() { proto.RegisterFile("aggregates/order/order.proto", fileDescriptor_2a6b118f4f9db3fd) }

var fileDescriptor_2a6b118f4f9db3fd = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0x67, 0x37, 0x49, 0xed, 0xd7, 0xa5, 0x33, 0x3a, 0x99, 0xae, 0xa3, 0xde, 0x4e, 0x66,
	0x50, 0x19, 0x3a, 0x76, 0xd8, 0xb1, 0xc5, 0xde, 0xe8, 0xc5, 0x09, 0x6a, 0xb7, 0xc3, 0x2e, 0x41,
	0x89, 0xde, 0x38, 0xa2, 0x91, 0x1d, 0xf4, 0xa7, 0x9f, 0x76, 0x1f, 0x66, 0x44, 0x76, 0xd6, 0xd2,
	0x8b, 0x31, 0xbf, 0xe7, 0x79, 0xa5, 0xf7, 0x79, 0x10, 0x5c, 0xf2, 0xb6, 0xd5, 0xd8, 0x72, 0x8b,
	0xa6, 0xec, 0xb5, 0x40, 0x3d, 0x7c, 0xe9, 0x5e, 0xf7, 0xb6, 0x27, 0xe9, 0x8b, 0x4a, 0x3d, 0xbf,
	0xb8, 0x6a, 0xfb, 0xbe, 0xdd, 0x61, 0xe9, 0xf5, 0x95, 0xdb, 0x94, 0x56, 0x2a, 0x34, 0x96, 0xab,
	0xfd, 0x30, 0xf2, 0xf9, 0x6f, 0x08, 0xd3, 0xf9, 0xc1, 0x4a, 0xce, 0x21, 0x94, 0x22, 0x0b, 0xf2,
	0xa0, 0x38, 0x61, 0xa1, 0x14, 0x84, 0xc0, 0xc4, 0x39, 0x29, 0xb2, 0x30, 0x0f, 0x8a, 0x98, 0xf9,
	0x7f, 0xf2, 0x01, 0xe2, 0x35, 0xd7, 0x76, 0xe9, 0x85, 0x13, 0x2f, 0x44, 0x07, 0xf0, 0x6b, 0x14,
	0xfd, 0xa5, 0xcb, 0xce, 0xa9, 0x6c, 0x32, 0x88, 0x1e, 0x34, 0x4e, 0x91, 0x4f, 0x70, 0xc6, 0x85,
	0xd0, 0x68, 0xcc, 0x30, 0x3c, 0xf5, 0x7a, 0x32, 0x32, 0x3f, 0xff, 0x0d, 0x66, 0xc6, 0x72, 0xeb,
	0x4c, 0x36, 0xcb, 0x83, 0xe2, 0xfc, 0xe6, 0x23, 0x7d, 0x1b, 0x87, 0xfa, 0x4d, 0x1f, 0xbc, 0x89,
	0x8d, 0xe6, 0xc3, 0xc9, 0x66, 0x2b, 0xf7, 0x0a, 0x3b, 0xbb, 0xdc, 0x20, 0x66, 0xa7, 0x79, 0x50,
	0x4c, 0x59, 0x72, 0x64, 0x3f, 0x10, 0xc9, 0x25, 0xc4, 0x02, 0x77, 0xf2, 0x19, 0x35, 0x8a, 0x2c,
	0xca, 0x83, 0x22, 0x62, 0x2f, 0x80, 0x7c, 0x07, 0x58, 0x6b, 0xe4, 0x16, 0xc5, 0x92, 0xdb, 0x2c,
	0xce, 0x83, 0x22, 0xb9, 0xb9, 0xa0, 0x43, 0x71, 0xf4, 0x58, 0x1c, 0x7d, 0x3c, 0x16, 0xc7, 0xe2,
	0xd1, 0x7d, 0x6b, 0xc9, 0x15, 0x24, 0x16, 0x3b, 0xde, 0x8d, 0x8d, 0x24, 0x3e, 0x14, 0x0c, 0xe8,
	0x90, 0xe9, 0xcb, 0x1c, 0x92, 0x57, 0x3b, 0x93, 0x04, 0x4e, 0x17, 0x75, 0x53, 0xdd, 0x37, 0x3f,
	0xd3, 0x77, 0xe4, 0x0c, 0xa2, 0xdb, 0xc5, 0x82, 0xcd, 0x7f, 0xd7, 0x55, 0x1a, 0x90, 0xf7, 0x90,
	0x54, 0xf7, 0x0f, 0xff, 0x41, 0x48, 0x00, 0x66, 0xac, 0x6e, 0x1e, 0xeb, 0x2a, 0x9d, 0x90, 0x08,
	0x26, 0xd5, 0xbc, 0xa9, 0xd3, 0xe9, 0x5d, 0xf5, 0xe7, 0xae, 0x95, 0x76, 0xeb, 0x56, 0x74, 0xdd,
	0xab, 0x52, 0xee, 0xb6, 0x5c, 0xa9, 0xad, 0x10, 0xe5, 0x93, 0x13, 0xfc, 0x49, 0x5e, 0x6f, 0xe4,
	0x0e, 0xaf, 0x0d, 0xea, 0x67, 0xb9, 0xc6, 0x12, 0x3b, 0x2b, 0xad, 0x44, 0x53, 0xbe, 0x7d, 0x34,
	0xab, 0x99, 0x8f, 0xf5, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x8d, 0xe0, 0x8a, 0x4f,
	0x02, 0x00, 0x00,
}
