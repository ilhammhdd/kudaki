// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/mountain/recommended_gear_item.proto

package mountain

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

type RecommendedGearItem struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	RecommendedGear      *RecommendedGear     `protobuf:"bytes,3,opt,name=recommended_gear,json=recommendedGear,proto3" json:"recommended_gear,omitempty"`
	ItemType             string               `protobuf:"bytes,4,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	Total                int32                `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RecommendedGearItem) Reset()         { *m = RecommendedGearItem{} }
func (m *RecommendedGearItem) String() string { return proto.CompactTextString(m) }
func (*RecommendedGearItem) ProtoMessage()    {}
func (*RecommendedGearItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9694773ce2bb3b2f, []int{0}
}

func (m *RecommendedGearItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendedGearItem.Unmarshal(m, b)
}
func (m *RecommendedGearItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendedGearItem.Marshal(b, m, deterministic)
}
func (m *RecommendedGearItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendedGearItem.Merge(m, src)
}
func (m *RecommendedGearItem) XXX_Size() int {
	return xxx_messageInfo_RecommendedGearItem.Size(m)
}
func (m *RecommendedGearItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendedGearItem.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendedGearItem proto.InternalMessageInfo

func (m *RecommendedGearItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RecommendedGearItem) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *RecommendedGearItem) GetRecommendedGear() *RecommendedGear {
	if m != nil {
		return m.RecommendedGear
	}
	return nil
}

func (m *RecommendedGearItem) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *RecommendedGearItem) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RecommendedGearItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*RecommendedGearItem)(nil), "aggregates.mountain.RecommendedGearItem")
}

func init() {
	proto.RegisterFile("aggregates/mountain/recommended_gear_item.proto", fileDescriptor_9694773ce2bb3b2f)
}

var fileDescriptor_9694773ce2bb3b2f = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0x87, 0xd9, 0xfe, 0xe3, 0x6d, 0x5e, 0x50, 0x49, 0x3d, 0x2c, 0xf5, 0xe0, 0x22, 0x1e, 0x16,
	0xa1, 0x09, 0xe8, 0xc9, 0xa3, 0x82, 0x14, 0x4f, 0x42, 0xe8, 0xc9, 0xcb, 0x92, 0x6e, 0xc6, 0x74,
	0x68, 0xb3, 0x59, 0xb2, 0x13, 0xa1, 0xdf, 0xdc, 0xa3, 0x74, 0xd7, 0x52, 0x2c, 0x3d, 0x78, 0xcb,
	0x0c, 0xf3, 0x9b, 0x27, 0x0f, 0xc3, 0xa4, 0xb6, 0x36, 0x80, 0xd5, 0x04, 0x8d, 0x74, 0x3e, 0x56,
	0xa4, 0xb1, 0x92, 0x01, 0x4a, 0xef, 0x1c, 0x54, 0x06, 0x4c, 0x61, 0x41, 0x87, 0x02, 0x09, 0x9c,
	0xa8, 0x83, 0x27, 0xcf, 0x27, 0x87, 0x80, 0xd8, 0x07, 0xa6, 0x77, 0x7f, 0xd9, 0xd2, 0x2d, 0x98,
	0x5e, 0x5b, 0xef, 0xed, 0x06, 0x64, 0x5b, 0x2d, 0xe3, 0x87, 0x24, 0x74, 0xd0, 0x90, 0x76, 0x75,
	0x37, 0x70, 0xf3, 0x95, 0xb0, 0x89, 0x3a, 0x64, 0xe7, 0xa0, 0xc3, 0x2b, 0x81, 0xe3, 0x67, 0xac,
	0x87, 0x26, 0x4d, 0xb2, 0x24, 0xef, 0xab, 0x1e, 0x1a, 0xce, 0xd9, 0x20, 0x46, 0x34, 0x69, 0x2f,
	0x4b, 0xf2, 0xb1, 0x6a, 0xdf, 0xfc, 0x8d, 0x5d, 0x1c, 0x63, 0xd3, 0x7e, 0x96, 0xe4, 0xff, 0xef,
	0x6f, 0xc5, 0x89, 0x8f, 0x8b, 0x23, 0x8e, 0x3a, 0x0f, 0xbf, 0x1b, 0xfc, 0x8a, 0x8d, 0x77, 0xf2,
	0x05, 0x6d, 0x6b, 0x48, 0x07, 0x2d, 0xe9, 0xdf, 0xae, 0xb1, 0xd8, 0xd6, 0xc0, 0x2f, 0xd9, 0x90,
	0x3c, 0xe9, 0x4d, 0x3a, 0xcc, 0x92, 0x7c, 0xa8, 0xba, 0x82, 0x3f, 0x32, 0x56, 0x06, 0xd0, 0x04,
	0xa6, 0xd0, 0x94, 0x8e, 0x5a, 0xfa, 0x54, 0x74, 0xd6, 0x62, 0x6f, 0x2d, 0x16, 0x7b, 0x6b, 0x35,
	0xfe, 0x99, 0x7e, 0xa2, 0xe7, 0xf9, 0xfb, 0x8b, 0x45, 0x5a, 0xc5, 0xa5, 0x28, 0xbd, 0x93, 0xb8,
	0x59, 0x69, 0xe7, 0x56, 0xc6, 0xc8, 0x75, 0x34, 0x7a, 0x8d, 0xb3, 0x1d, 0x78, 0xd6, 0x40, 0xf8,
	0xc4, 0x12, 0x24, 0x54, 0x84, 0x84, 0xd0, 0x9c, 0xba, 0xe0, 0x72, 0xd4, 0x72, 0x1e, 0xbe, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x99, 0x85, 0x36, 0xb1, 0xdf, 0x01, 0x00, 0x00,
}
