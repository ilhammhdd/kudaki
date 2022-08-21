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
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0xcb, 0xd3, 0x30,
	0x18, 0xc7, 0x69, 0xdf, 0xd7, 0x75, 0xcb, 0x36, 0x85, 0x88, 0x10, 0x8b, 0x62, 0xd5, 0x4b, 0x2f,
	0x4b, 0x60, 0x9e, 0x3c, 0x4e, 0x41, 0x3c, 0x8d, 0x51, 0x3c, 0x79, 0x29, 0x69, 0xf3, 0xac, 0x0d,
	0x6b, 0x9a, 0x92, 0x26, 0x3b, 0xec, 0x3b, 0xfb, 0x1d, 0xa4, 0xe9, 0xba, 0xc1, 0xde, 0x5b, 0x9f,
	0xe7, 0xf7, 0xa3, 0xff, 0xe4, 0x1f, 0xf4, 0x91, 0x57, 0x95, 0x81, 0x8a, 0x5b, 0xe8, 0x99, 0xeb,
	0xc1, 0x30, 0x2e, 0x84, 0x81, 0xbe, 0xa7, 0x9d, 0xd1, 0x56, 0xe3, 0x37, 0x77, 0x4c, 0x07, 0x1c,
	0xbf, 0xf0, 0x3b, 0xa3, 0x8f, 0xb2, 0x81, 0xd1, 0x8f, 0x3f, 0x55, 0x5a, 0x57, 0x0d, 0x30, 0x3f,
	0x15, 0xee, 0xc8, 0xac, 0x54, 0xd0, 0x5b, 0xae, 0xba, 0x51, 0xf8, 0xf2, 0x2f, 0x44, 0xd1, 0x6e,
	0x8c, 0xc0, 0xaf, 0x51, 0x28, 0x05, 0x09, 0x92, 0x20, 0x7d, 0xca, 0x42, 0x29, 0x30, 0x46, 0xcf,
	0xce, 0x49, 0x41, 0xc2, 0x24, 0x48, 0x17, 0x99, 0xff, 0xc6, 0x5b, 0x14, 0x5d, 0x13, 0xc8, 0x53,
	0x12, 0xa4, 0xcb, 0x2d, 0xa1, 0x0f, 0x47, 0xa2, 0x87, 0x91, 0x67, 0x93, 0x88, 0x3f, 0xa3, 0xd5,
	0xd1, 0x35, 0x4d, 0x7e, 0xbd, 0x0a, 0x79, 0xf6, 0xff, 0x5b, 0x0e, 0xbb, 0x29, 0xfa, 0x2b, 0x5a,
	0x1b, 0x28, 0x41, 0x9e, 0xc1, 0xe4, 0x2d, 0x57, 0x40, 0x5e, 0x79, 0x67, 0x35, 0x2d, 0xf7, 0x5c,
	0x01, 0xde, 0xa2, 0x77, 0x37, 0xa9, 0xab, 0x75, 0x0b, 0x79, 0xeb, 0x54, 0x01, 0x86, 0xcc, 0xbc,
	0xfc, 0x76, 0x82, 0x87, 0x81, 0xed, 0x3d, 0xc2, 0xef, 0xd1, 0xfc, 0x22, 0xbb, 0xbc, 0xd4, 0x02,
	0x48, 0xe4, 0xb5, 0xe8, 0x22, 0xbb, 0x9f, 0x5a, 0x00, 0x8e, 0xd1, 0xbc, 0xe1, 0x56, 0x5a, 0x27,
	0x80, 0x2c, 0x93, 0x20, 0x0d, 0xb3, 0xdb, 0x8c, 0x3f, 0xa0, 0x45, 0xa3, 0xdb, 0x6a, 0x84, 0x2b,
	0x0f, 0xef, 0x0b, 0xfc, 0x1d, 0xa1, 0xd2, 0x00, 0xb7, 0x20, 0x72, 0x6e, 0xc9, 0xda, 0xf7, 0x10,
	0xd3, 0xb1, 0x6a, 0x3a, 0x55, 0x4d, 0xff, 0x4c, 0x55, 0x67, 0x8b, 0xab, 0xbd, 0xb3, 0x3f, 0x7e,
	0xff, 0xfd, 0x55, 0x49, 0x5b, 0xbb, 0x82, 0x96, 0x5a, 0x31, 0xd9, 0xd4, 0x5c, 0xa9, 0x5a, 0x08,
	0x76, 0x72, 0x82, 0x9f, 0xe4, 0x66, 0x28, 0x70, 0xc3, 0x9d, 0xad, 0x37, 0x3d, 0x98, 0xb3, 0x2c,
	0x81, 0x41, 0x6b, 0xa5, 0x95, 0xd0, 0xb3, 0x87, 0x87, 0x2e, 0x66, 0x3e, 0xe8, 0xdb, 0xff, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1e, 0xe4, 0x48, 0x0f, 0x32, 0x02, 0x00, 0x00,
}