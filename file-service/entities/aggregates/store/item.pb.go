// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/store/item.proto

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

type PriceDuration int32

const (
	PriceDuration_DAY   PriceDuration = 0
	PriceDuration_WEEK  PriceDuration = 1
	PriceDuration_MONTH PriceDuration = 2
	PriceDuration_YEAR  PriceDuration = 3
)

var PriceDuration_name = map[int32]string{
	0: "DAY",
	1: "WEEK",
	2: "MONTH",
	3: "YEAR",
}

var PriceDuration_value = map[string]int32{
	"DAY":   0,
	"WEEK":  1,
	"MONTH": 2,
	"YEAR":  3,
}

func (x PriceDuration) String() string {
	return proto.EnumName(PriceDuration_name, int32(x))
}

func (PriceDuration) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f845c3cdd9c1555f, []int{0}
}

type UnitofMeasurement int32

const (
	UnitofMeasurement_MM  UnitofMeasurement = 0
	UnitofMeasurement_CM  UnitofMeasurement = 1
	UnitofMeasurement_DM  UnitofMeasurement = 2
	UnitofMeasurement_M   UnitofMeasurement = 3
	UnitofMeasurement_DAM UnitofMeasurement = 4
	UnitofMeasurement_HM  UnitofMeasurement = 5
	UnitofMeasurement_KM  UnitofMeasurement = 6
)

var UnitofMeasurement_name = map[int32]string{
	0: "MM",
	1: "CM",
	2: "DM",
	3: "M",
	4: "DAM",
	5: "HM",
	6: "KM",
}

var UnitofMeasurement_value = map[string]int32{
	"MM":  0,
	"CM":  1,
	"DM":  2,
	"M":   3,
	"DAM": 4,
	"HM":  5,
	"KM":  6,
}

func (x UnitofMeasurement) String() string {
	return proto.EnumName(UnitofMeasurement_name, int32(x))
}

func (UnitofMeasurement) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f845c3cdd9c1555f, []int{1}
}

type ItemDimension struct {
	Length               int32             `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Width                int32             `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32             `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	UnitOfMeasurement    UnitofMeasurement `protobuf:"varint,4,opt,name=unit_of_measurement,json=unitOfMeasurement,proto3,enum=aggregates.store.UnitofMeasurement" json:"unit_of_measurement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ItemDimension) Reset()         { *m = ItemDimension{} }
func (m *ItemDimension) String() string { return proto.CompactTextString(m) }
func (*ItemDimension) ProtoMessage()    {}
func (*ItemDimension) Descriptor() ([]byte, []int) {
	return fileDescriptor_f845c3cdd9c1555f, []int{0}
}

func (m *ItemDimension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemDimension.Unmarshal(m, b)
}
func (m *ItemDimension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemDimension.Marshal(b, m, deterministic)
}
func (m *ItemDimension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemDimension.Merge(m, src)
}
func (m *ItemDimension) XXX_Size() int {
	return xxx_messageInfo_ItemDimension.Size(m)
}
func (m *ItemDimension) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemDimension.DiscardUnknown(m)
}

var xxx_messageInfo_ItemDimension proto.InternalMessageInfo

func (m *ItemDimension) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *ItemDimension) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ItemDimension) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ItemDimension) GetUnitOfMeasurement() UnitofMeasurement {
	if m != nil {
		return m.UnitOfMeasurement
	}
	return UnitofMeasurement_MM
}

type Item struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Storefront           *Storefront          `protobuf:"bytes,3,opt,name=storefront,proto3" json:"storefront,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Amount               int32                `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Unit                 string               `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Price                int32                `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	PriceDuration        PriceDuration        `protobuf:"varint,8,opt,name=price_duration,json=priceDuration,proto3,enum=aggregates.store.PriceDuration" json:"price_duration,omitempty"`
	Description          string               `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Photo                string               `protobuf:"bytes,10,opt,name=photo,proto3" json:"photo,omitempty"`
	Rating               float64              `protobuf:"fixed64,11,opt,name=rating,proto3" json:"rating,omitempty"`
	ItemDimension        *ItemDimension       `protobuf:"bytes,12,opt,name=item_dimension,json=itemDimension,proto3" json:"item_dimension,omitempty"`
	Color                string               `protobuf:"bytes,13,opt,name=color,proto3" json:"color,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	TotalRawRating       float64              `protobuf:"fixed64,15,opt,name=total_raw_rating,json=totalRawRating,proto3" json:"total_raw_rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_f845c3cdd9c1555f, []int{1}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Item) GetStorefront() *Storefront {
	if m != nil {
		return m.Storefront
	}
	return nil
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Item) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Item) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetPriceDuration() PriceDuration {
	if m != nil {
		return m.PriceDuration
	}
	return PriceDuration_DAY
}

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Item) GetPhoto() string {
	if m != nil {
		return m.Photo
	}
	return ""
}

func (m *Item) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Item) GetItemDimension() *ItemDimension {
	if m != nil {
		return m.ItemDimension
	}
	return nil
}

func (m *Item) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Item) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Item) GetTotalRawRating() float64 {
	if m != nil {
		return m.TotalRawRating
	}
	return 0
}

func init() {
	proto.RegisterEnum("aggregates.store.PriceDuration", PriceDuration_name, PriceDuration_value)
	proto.RegisterEnum("aggregates.store.UnitofMeasurement", UnitofMeasurement_name, UnitofMeasurement_value)
	proto.RegisterType((*ItemDimension)(nil), "aggregates.store.ItemDimension")
	proto.RegisterType((*Item)(nil), "aggregates.store.Item")
}

func init() { proto.RegisterFile("aggregates/store/item.proto", fileDescriptor_f845c3cdd9c1555f) }

var fileDescriptor_f845c3cdd9c1555f = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x51, 0x4f, 0xdb, 0x30,
	0x10, 0x26, 0x6d, 0x53, 0xa8, 0x59, 0x3b, 0xe3, 0xa1, 0xc9, 0x62, 0x93, 0xe8, 0xd8, 0x4b, 0x85,
	0x44, 0x22, 0xb1, 0x87, 0x69, 0xd2, 0x5e, 0x60, 0x65, 0x62, 0x42, 0x19, 0x93, 0x61, 0x9a, 0xd8,
	0x4b, 0x64, 0x1a, 0x37, 0xb1, 0x68, 0xec, 0x2a, 0xb9, 0x8c, 0xdf, 0xb2, 0xbf, 0xb0, 0x5f, 0x39,
	0xf9, 0x5c, 0x46, 0x4b, 0xf7, 0x12, 0xfb, 0xbb, 0xfb, 0x7c, 0xf7, 0xdd, 0xe9, 0x0b, 0x79, 0x25,
	0xf3, 0xbc, 0x52, 0xb9, 0x04, 0x55, 0xc7, 0x35, 0xd8, 0x4a, 0xc5, 0x1a, 0x54, 0x19, 0xcd, 0x2b,
	0x0b, 0x96, 0xd1, 0xc7, 0x64, 0x84, 0xc9, 0xbd, 0xfd, 0xdc, 0xda, 0x7c, 0xa6, 0x62, 0xcc, 0xdf,
	0x36, 0xd3, 0x18, 0x74, 0xa9, 0x6a, 0x90, 0xe5, 0xdc, 0x3f, 0xd9, 0x7b, 0xb3, 0x56, 0x0f, 0xbf,
	0xd3, 0xca, 0x1a, 0xf0, 0x94, 0x83, 0x3f, 0x01, 0xe9, 0x7f, 0x01, 0x55, 0x8e, 0x75, 0xa9, 0x4c,
	0xad, 0xad, 0x61, 0x2f, 0x49, 0x77, 0xa6, 0x4c, 0x0e, 0x05, 0x0f, 0x86, 0xc1, 0x28, 0x14, 0x0b,
	0xc4, 0x76, 0x49, 0x78, 0xaf, 0x33, 0x28, 0x78, 0x0b, 0xc3, 0x1e, 0x38, 0x76, 0xa1, 0x74, 0x5e,
	0x00, 0x6f, 0x7b, 0xb6, 0x47, 0xec, 0x8a, 0xbc, 0x68, 0x8c, 0x86, 0xd4, 0x4e, 0xd3, 0x52, 0xc9,
	0xba, 0xa9, 0x54, 0xa9, 0x0c, 0xf0, 0xce, 0x30, 0x18, 0x0d, 0x8e, 0xdf, 0x46, 0x4f, 0x67, 0x89,
	0xbe, 0x1b, 0x0d, 0x76, 0x9a, 0x3c, 0x52, 0xc5, 0x8e, 0x7b, 0x7f, 0xb9, 0x1c, 0x3a, 0xf8, 0xdd,
	0x21, 0x1d, 0x27, 0x96, 0x0d, 0x48, 0x4b, 0x67, 0xa8, 0xaf, 0x2d, 0x5a, 0x3a, 0x63, 0x8c, 0x74,
	0x9a, 0x46, 0x67, 0x28, 0xad, 0x27, 0xf0, 0xce, 0x3e, 0x12, 0xf2, 0x38, 0x2d, 0xaa, 0xdb, 0x3e,
	0x7e, 0xbd, 0xde, 0xf8, 0xea, 0x1f, 0x47, 0x2c, 0xf1, 0x5d, 0x45, 0x23, 0x4b, 0x85, 0x82, 0x7b,
	0x02, 0xef, 0x6e, 0x56, 0x59, 0xda, 0xc6, 0x00, 0x0f, 0xfd, 0xac, 0x1e, 0x61, 0x77, 0xa3, 0x81,
	0x77, 0x17, 0xdd, 0x8d, 0x06, 0xb7, 0xad, 0x79, 0xa5, 0x27, 0x8a, 0x6f, 0xfa, 0x6d, 0x21, 0x60,
	0x9f, 0xc9, 0x00, 0x2f, 0x69, 0xd6, 0x54, 0x12, 0xb4, 0x35, 0x7c, 0x0b, 0x17, 0xb2, 0xbf, 0xae,
	0xeb, 0x9b, 0xe3, 0x8d, 0x17, 0x34, 0xd1, 0x9f, 0x2f, 0x43, 0x36, 0x24, 0xdb, 0x99, 0xaa, 0x27,
	0x95, 0x9e, 0x63, 0x91, 0x1e, 0x36, 0x5e, 0x0e, 0x61, 0xff, 0xc2, 0x82, 0xe5, 0x04, 0x73, 0x1e,
	0xb8, 0x09, 0x5c, 0x05, 0x93, 0xf3, 0xed, 0x61, 0x30, 0x0a, 0xc4, 0x02, 0x39, 0x5d, 0xce, 0x69,
	0x69, 0xf6, 0xe0, 0x02, 0xfe, 0x0c, 0xf7, 0xf5, 0x1f, 0x5d, 0x2b, 0x66, 0x11, 0x7d, 0xbd, 0xe2,
	0x9d, 0x5d, 0x12, 0x4e, 0xec, 0xcc, 0x56, 0xbc, 0xef, 0xbb, 0x22, 0x60, 0x1f, 0x08, 0x99, 0x54,
	0x4a, 0x82, 0xca, 0x52, 0x09, 0x7c, 0x80, 0x95, 0xf7, 0x22, 0x6f, 0xde, 0xe8, 0xc1, 0xbc, 0xd1,
	0xf5, 0x83, 0x79, 0x45, 0x6f, 0xc1, 0x3e, 0x01, 0x36, 0x22, 0x14, 0x2c, 0xc8, 0x59, 0x5a, 0xc9,
	0xfb, 0x74, 0x21, 0xfd, 0x39, 0x4a, 0x1f, 0x60, 0x5c, 0xc8, 0x7b, 0x81, 0xd1, 0xc3, 0xf7, 0xa4,
	0xbf, 0xb2, 0x32, 0xb6, 0x49, 0xda, 0xe3, 0x93, 0x1b, 0xba, 0xc1, 0xb6, 0x48, 0xe7, 0xc7, 0xd9,
	0xd9, 0x05, 0x0d, 0x58, 0x8f, 0x84, 0xc9, 0xe5, 0xd7, 0xeb, 0x73, 0xda, 0x72, 0xc1, 0x9b, 0xb3,
	0x13, 0x41, 0xdb, 0x87, 0x17, 0x64, 0x67, 0xcd, 0x7c, 0xac, 0x4b, 0x5a, 0x49, 0x42, 0x37, 0xdc,
	0xf9, 0x29, 0xa1, 0x81, 0x3b, 0xc7, 0x09, 0x6d, 0xb1, 0x90, 0x04, 0x09, 0x6d, 0xfb, 0xda, 0x09,
	0xed, 0xb8, 0xf8, 0x79, 0x42, 0x43, 0x77, 0x5e, 0x24, 0xb4, 0x7b, 0x3a, 0xfe, 0x79, 0x9a, 0x6b,
	0x28, 0x9a, 0xdb, 0x68, 0x62, 0xcb, 0x58, 0xcf, 0x0a, 0x59, 0x96, 0x45, 0x96, 0xc5, 0x77, 0x4d,
	0x26, 0xef, 0xf4, 0xd1, 0x54, 0xcf, 0xd4, 0x51, 0xad, 0xaa, 0x5f, 0x7a, 0xa2, 0x62, 0x65, 0x40,
	0x83, 0x56, 0x75, 0xfc, 0xf4, 0x2f, 0xbd, 0xed, 0xe2, 0x52, 0xde, 0xfd, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xde, 0x0f, 0xa5, 0xe1, 0x10, 0x04, 0x00, 0x00,
}
