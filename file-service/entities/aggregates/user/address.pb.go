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
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0xcf, 0x93, 0x30,
	0x18, 0xc7, 0x03, 0xef, 0xeb, 0xd8, 0xba, 0x4d, 0x93, 0x1a, 0x93, 0x4a, 0x34, 0xa2, 0x5e, 0xb8,
	0xac, 0x4d, 0xe6, 0xc9, 0xe3, 0xdc, 0x7d, 0x59, 0x88, 0x27, 0x2f, 0xa4, 0xd0, 0x67, 0xd0, 0x8c,
	0x52, 0x52, 0xda, 0x1d, 0xf6, 0x9d, 0xfd, 0x0e, 0x86, 0x32, 0xb6, 0x64, 0xde, 0x78, 0x9e, 0xdf,
	0x2f, 0xfc, 0xdb, 0x7f, 0xd1, 0x67, 0x5e, 0x55, 0x06, 0x2a, 0x6e, 0xa1, 0x67, 0xae, 0x07, 0xc3,
	0xb8, 0x10, 0x06, 0xfa, 0x9e, 0x76, 0x46, 0x5b, 0x8d, 0xdf, 0x3d, 0x30, 0x1d, 0x70, 0xfc, 0x9f,
	0xdf, 0x19, 0x7d, 0x92, 0x0d, 0x8c, 0x7e, 0xfc, 0xa5, 0xd2, 0xba, 0x6a, 0x80, 0xf9, 0xa9, 0x70,
	0x27, 0x66, 0xa5, 0x82, 0xde, 0x72, 0xd5, 0x8d, 0xc2, 0xb7, 0xbf, 0x21, 0x8a, 0x76, 0x63, 0x04,
	0x7e, 0x8b, 0x42, 0x29, 0x48, 0x90, 0x04, 0xe9, 0x4b, 0x16, 0x4a, 0x81, 0x31, 0x7a, 0x75, 0x4e,
	0x0a, 0x12, 0x26, 0x41, 0xba, 0xc8, 0xfc, 0x37, 0xde, 0xa2, 0xe8, 0x96, 0x40, 0x5e, 0x92, 0x20,
	0x5d, 0x6e, 0x09, 0x7d, 0x3a, 0x12, 0x3d, 0x8e, 0x3c, 0x9b, 0x44, 0xfc, 0x15, 0xad, 0x4e, 0xae,
	0x69, 0xf2, 0xdb, 0x55, 0xc8, 0xab, 0xff, 0xdf, 0x72, 0xd8, 0x4d, 0xd1, 0xdf, 0xd1, 0xda, 0x40,
	0x09, 0xf2, 0x02, 0x26, 0x6f, 0xb9, 0x02, 0xf2, 0xc6, 0x3b, 0xab, 0x69, 0x79, 0xe0, 0x0a, 0xf0,
	0x16, 0x7d, 0xb8, 0x4b, 0x5d, 0xad, 0x5b, 0xc8, 0x5b, 0xa7, 0x0a, 0x30, 0x64, 0xe6, 0xe5, 0xf7,
	0x13, 0x3c, 0x0e, 0xec, 0xe0, 0x11, 0xfe, 0x88, 0xe6, 0x57, 0xd9, 0xe5, 0xa5, 0x16, 0x40, 0x22,
	0xaf, 0x45, 0x57, 0xd9, 0xed, 0xb5, 0x00, 0x1c, 0xa3, 0x79, 0xc3, 0xad, 0xb4, 0x4e, 0x00, 0x59,
	0x26, 0x41, 0x1a, 0x66, 0xf7, 0x19, 0x7f, 0x42, 0x8b, 0x46, 0xb7, 0xd5, 0x08, 0x57, 0x1e, 0x3e,
	0x16, 0xf8, 0x27, 0x42, 0xa5, 0x01, 0x6e, 0x41, 0xe4, 0xdc, 0x92, 0xb5, 0xef, 0x21, 0xa6, 0x63,
	0xd5, 0x74, 0xaa, 0x9a, 0xfe, 0x9e, 0xaa, 0xce, 0x16, 0x37, 0x7b, 0x67, 0x7f, 0xed, 0xff, 0xec,
	0x2a, 0x69, 0x6b, 0x57, 0xd0, 0x52, 0x2b, 0x26, 0x9b, 0x9a, 0x2b, 0x55, 0x0b, 0xc1, 0xce, 0x4e,
	0xf0, 0xb3, 0xdc, 0x0c, 0x7d, 0x6d, 0x7a, 0x30, 0x17, 0x59, 0x02, 0x83, 0xd6, 0x4a, 0x2b, 0xa1,
	0x67, 0x4f, 0x6f, 0x5c, 0xcc, 0x7c, 0xc6, 0x8f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x10,
	0x64, 0xe7, 0x2d, 0x02, 0x00, 0x00,
}
