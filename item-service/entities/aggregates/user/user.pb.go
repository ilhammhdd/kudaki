// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/user/user.proto

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

type UserRole int32

const (
	UserRole_ADMIN       UserRole = 0
	UserRole_USER        UserRole = 1
	UserRole_KUDAKI_TEAM UserRole = 2
	UserRole_ORGANIZER   UserRole = 3
)

var UserRole_name = map[int32]string{
	0: "ADMIN",
	1: "USER",
	2: "KUDAKI_TEAM",
	3: "ORGANIZER",
}

var UserRole_value = map[string]int32{
	"ADMIN":       0,
	"USER":        1,
	"KUDAKI_TEAM": 2,
	"ORGANIZER":   3,
}

func (x UserRole) String() string {
	return proto.EnumName(UserRole_name, int32(x))
}

func (UserRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7100ea213fb19611, []int{0}
}

type AccountType int32

const (
	AccountType_NATIVE   AccountType = 0
	AccountType_GOOGLE   AccountType = 1
	AccountType_FACEBOOK AccountType = 2
)

var AccountType_name = map[int32]string{
	0: "NATIVE",
	1: "GOOGLE",
	2: "FACEBOOK",
}

var AccountType_value = map[string]int32{
	"NATIVE":   0,
	"GOOGLE":   1,
	"FACEBOOK": 2,
}

func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}

func (AccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7100ea213fb19611, []int{1}
}

type User struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string               `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Token                string               `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	Role                 UserRole             `protobuf:"varint,6,opt,name=role,proto3,enum=aggregates.user.UserRole" json:"role,omitempty"`
	PhoneNumber          string               `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	AccountType          AccountType          `protobuf:"varint,8,opt,name=account_type,json=accountType,proto3,enum=aggregates.user.AccountType" json:"account_type,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_7100ea213fb19611, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *User) GetRole() UserRole {
	if m != nil {
		return m.Role
	}
	return UserRole_ADMIN
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *User) GetAccountType() AccountType {
	if m != nil {
		return m.AccountType
	}
	return AccountType_NATIVE
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("aggregates.user.UserRole", UserRole_name, UserRole_value)
	proto.RegisterEnum("aggregates.user.AccountType", AccountType_name, AccountType_value)
	proto.RegisterType((*User)(nil), "aggregates.user.User")
}

func init() { proto.RegisterFile("aggregates/user/user.proto", fileDescriptor_7100ea213fb19611) }

var fileDescriptor_7100ea213fb19611 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x14, 0x8c, 0xe4, 0x8f, 0x4a, 0x4f, 0x6e, 0x22, 0x96, 0x1e, 0x54, 0x53, 0xa8, 0xdb, 0x93, 0x09,
	0x58, 0x82, 0xe4, 0xd4, 0x53, 0xd8, 0x24, 0xaa, 0x31, 0x6e, 0x2c, 0xd8, 0xda, 0x3d, 0xe4, 0x62,
	0xd6, 0xd2, 0xab, 0xbc, 0x58, 0xd2, 0x8a, 0xd5, 0xaa, 0x25, 0xff, 0xae, 0x3f, 0xad, 0x78, 0x55,
	0x37, 0x25, 0xbd, 0x88, 0x37, 0xf3, 0xde, 0x0c, 0xa3, 0x61, 0x61, 0xcc, 0xf3, 0x5c, 0x61, 0xce,
	0x35, 0x36, 0x51, 0xdb, 0xa0, 0x32, 0x9f, 0xb0, 0x56, 0x52, 0x4b, 0x72, 0xf1, 0xbc, 0x0b, 0x8f,
	0xf4, 0xf8, 0x7d, 0x2e, 0x65, 0x5e, 0x60, 0x64, 0xd6, 0xbb, 0xf6, 0x7b, 0xa4, 0x45, 0x89, 0x8d,
	0xe6, 0x65, 0xdd, 0x29, 0x3e, 0xfe, 0xb2, 0xa1, 0xbf, 0x69, 0x50, 0x91, 0x73, 0xb0, 0x45, 0x16,
	0x58, 0x13, 0x6b, 0xda, 0x63, 0xb6, 0xc8, 0x08, 0x81, 0x7e, 0xdb, 0x8a, 0x2c, 0xb0, 0x27, 0xd6,
	0xd4, 0x65, 0x66, 0x26, 0x6f, 0x60, 0x80, 0x25, 0x17, 0x45, 0xd0, 0x33, 0x64, 0x07, 0xc8, 0x18,
	0x9c, 0x9a, 0x37, 0xcd, 0x4f, 0xa9, 0xb2, 0xa0, 0x6f, 0x16, 0x7f, 0xf1, 0x51, 0xa1, 0xe5, 0x01,
	0xab, 0x60, 0xd0, 0x29, 0x0c, 0x20, 0x33, 0xe8, 0x2b, 0x59, 0x60, 0x30, 0x9c, 0x58, 0xd3, 0xf3,
	0xab, 0xb7, 0xe1, 0x8b, 0xd4, 0xe1, 0x31, 0x10, 0x93, 0x05, 0x32, 0x73, 0x46, 0x3e, 0xc0, 0xa8,
	0xde, 0xcb, 0x0a, 0xb7, 0x55, 0x5b, 0xee, 0x50, 0x05, 0xaf, 0x8c, 0x97, 0x67, 0xb8, 0x95, 0xa1,
	0xc8, 0x0d, 0x8c, 0x78, 0x9a, 0xca, 0xb6, 0xd2, 0x5b, 0xfd, 0x54, 0x63, 0xe0, 0x18, 0xe7, 0x77,
	0xff, 0x39, 0xd3, 0xee, 0x68, 0xfd, 0x54, 0x23, 0xf3, 0xf8, 0x33, 0x20, 0x9f, 0x00, 0x52, 0x85,
	0x5c, 0x63, 0xb6, 0xe5, 0x3a, 0x70, 0x27, 0xd6, 0xd4, 0xbb, 0x1a, 0x87, 0x5d, 0x7b, 0xe1, 0xa9,
	0xbd, 0x70, 0x7d, 0x6a, 0x8f, 0xb9, 0x7f, 0xae, 0xa9, 0xbe, 0xbc, 0x01, 0xe7, 0x14, 0x98, 0xb8,
	0x30, 0xa0, 0xf7, 0x0f, 0x8b, 0x95, 0x7f, 0x46, 0x1c, 0xe8, 0x6f, 0xbe, 0xc6, 0xcc, 0xb7, 0xc8,
	0x05, 0x78, 0xcb, 0xcd, 0x3d, 0x5d, 0x2e, 0xb6, 0xeb, 0x98, 0x3e, 0xf8, 0x36, 0x79, 0x0d, 0x6e,
	0xc2, 0xe6, 0x74, 0xb5, 0x78, 0x8c, 0x99, 0xdf, 0xbb, 0xbc, 0x06, 0xef, 0x9f, 0x5c, 0x04, 0x60,
	0xb8, 0xa2, 0xeb, 0xc5, 0xb7, 0xd8, 0x3f, 0x3b, 0xce, 0xf3, 0x24, 0x99, 0x7f, 0x89, 0x7d, 0x8b,
	0x8c, 0xc0, 0xf9, 0x4c, 0xef, 0xe2, 0xdb, 0x24, 0x59, 0xfa, 0xf6, 0xed, 0xdd, 0x23, 0xcd, 0x85,
	0xde, 0xb7, 0xbb, 0x30, 0x95, 0x65, 0x24, 0x8a, 0x3d, 0x2f, 0xcb, 0x7d, 0x96, 0x45, 0x87, 0x36,
	0xe3, 0x07, 0x31, 0x13, 0x1a, 0xcb, 0x59, 0x83, 0xea, 0x87, 0x48, 0x31, 0xc2, 0x4a, 0x0b, 0x2d,
	0xb0, 0x89, 0x5e, 0x3c, 0x9d, 0xdd, 0xd0, 0xfc, 0xd9, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5f, 0xd9, 0x74, 0x50, 0x54, 0x02, 0x00, 0x00,
}
