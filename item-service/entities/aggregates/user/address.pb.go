// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/user/address.proto

package user

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

type Address struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Profile              *Profile             `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	FullAddress          string               `protobuf:"bytes,4,opt,name=full_address,json=fullAddress,proto3" json:"full_address,omitempty"`
	ReceiverName         string               `protobuf:"bytes,5,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name,omitempty"`
	ReceiverPhoneNumber  string               `protobuf:"bytes,6,opt,name=receiver_phone_number,json=receiverPhoneNumber,proto3" json:"receiver_phone_number,omitempty"`
	ZipCode              string               `protobuf:"bytes,7,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Latitude             float32              `protobuf:"fixed32,11,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32              `protobuf:"fixed32,12,opt,name=longitude,proto3" json:"longitude,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_10ae9d3b22a24a4f, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Address) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Address) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *Address) GetFullAddress() string {
	if m != nil {
		return m.FullAddress
	}
	return ""
}

func (m *Address) GetReceiverName() string {
	if m != nil {
		return m.ReceiverName
	}
	return ""
}

func (m *Address) GetReceiverPhoneNumber() string {
	if m != nil {
		return m.ReceiverPhoneNumber
	}
	return ""
}

func (m *Address) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *Address) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Address) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Address) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Address)(nil), "aggregates.user.Address")
}

func init() { proto.RegisterFile("aggregates/user/address.proto", fileDescriptor_10ae9d3b22a24a4f) }

var fileDescriptor_10ae9d3b22a24a4f = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0xec, 0xb2, 0xd9, 0xba, 0x5d, 0x90, 0x8c, 0x90, 0x4c, 0x04, 0x22, 0xc0, 0x25,
	0x97, 0xb5, 0xa5, 0x72, 0xe2, 0x58, 0xf6, 0xbe, 0xaa, 0x22, 0x4e, 0x5c, 0x22, 0x27, 0x9e, 0x26,
	0xa3, 0xc6, 0x71, 0xe4, 0xd8, 0x3d, 0xf4, 0x9d, 0x79, 0x07, 0x54, 0xa7, 0x69, 0xa5, 0x72, 0xcb,
	0xcc, 0xf7, 0x29, 0xbf, 0xfd, 0x9b, 0x7c, 0x96, 0x4d, 0x63, 0xa1, 0x91, 0x0e, 0x46, 0xe1, 0x47,
	0xb0, 0x42, 0x2a, 0x65, 0x61, 0x1c, 0xf9, 0x60, 0x8d, 0x33, 0xf4, 0xdd, 0x15, 0xf3, 0x13, 0x4e,
	0xff, 0xf3, 0x07, 0x6b, 0x76, 0xd8, 0xc1, 0xe4, 0xa7, 0x5f, 0x1a, 0x63, 0x9a, 0x0e, 0x44, 0x98,
	0x2a, 0xbf, 0x13, 0x0e, 0x35, 0x8c, 0x4e, 0xea, 0x61, 0x12, 0xbe, 0xfd, 0x8d, 0x49, 0xb2, 0x99,
	0x22, 0xe8, 0x5b, 0x12, 0xa3, 0x62, 0x51, 0x16, 0xe5, 0x77, 0x45, 0x8c, 0x8a, 0x52, 0x72, 0xef,
	0x3d, 0x2a, 0x16, 0x67, 0x51, 0xbe, 0x28, 0xc2, 0x37, 0x5d, 0x93, 0xe4, 0x9c, 0xc0, 0xee, 0xb2,
	0x28, 0x5f, 0xae, 0x19, 0xbf, 0x39, 0x12, 0xdf, 0x4e, 0xbc, 0x98, 0x45, 0xfa, 0x95, 0xac, 0x76,
	0xbe, 0xeb, 0xca, 0xf3, 0x55, 0xd8, 0x7d, 0xf8, 0xdf, 0xf2, 0xb4, 0x9b, 0xa3, 0xbf, 0x93, 0x27,
	0x0b, 0x35, 0xe0, 0x01, 0x6c, 0xd9, 0x4b, 0x0d, 0xec, 0x4d, 0x70, 0x56, 0xf3, 0xf2, 0x55, 0x6a,
	0xa0, 0x6b, 0xf2, 0xe1, 0x22, 0x0d, 0xad, 0xe9, 0xa1, 0xec, 0xbd, 0xae, 0xc0, 0xb2, 0x87, 0x20,
	0xbf, 0x9f, 0xe1, 0xf6, 0xc4, 0x5e, 0x03, 0xa2, 0x1f, 0xc9, 0xe3, 0x11, 0x87, 0xb2, 0x36, 0x0a,
	0x58, 0x12, 0xb4, 0xe4, 0x88, 0xc3, 0x8b, 0x51, 0x40, 0x53, 0xf2, 0xd8, 0x49, 0x87, 0xce, 0x2b,
	0x60, 0xcb, 0x2c, 0xca, 0xe3, 0xe2, 0x32, 0xd3, 0x4f, 0x64, 0xd1, 0x99, 0xbe, 0x99, 0xe0, 0x2a,
	0xc0, 0xeb, 0x82, 0xfe, 0x24, 0xa4, 0xb6, 0x20, 0x1d, 0xa8, 0x52, 0x3a, 0xf6, 0x14, 0x7a, 0x48,
	0xf9, 0x54, 0x35, 0x9f, 0xab, 0xe6, 0xbf, 0xe7, 0xaa, 0x8b, 0xc5, 0xd9, 0xde, 0xb8, 0x5f, 0x2f,
	0x7f, 0x36, 0x0d, 0xba, 0xd6, 0x57, 0xbc, 0x36, 0x5a, 0x60, 0xd7, 0x4a, 0xad, 0x5b, 0xa5, 0xc4,
	0xde, 0x2b, 0xb9, 0xc7, 0x67, 0x74, 0xa0, 0x9f, 0x47, 0xb0, 0x07, 0xac, 0x41, 0x40, 0xef, 0xd0,
	0x21, 0x8c, 0xe2, 0xe6, 0x8d, 0xab, 0x87, 0x90, 0xf1, 0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa2, 0x7b, 0xae, 0xfb, 0x2d, 0x02, 0x00, 0x00,
}
