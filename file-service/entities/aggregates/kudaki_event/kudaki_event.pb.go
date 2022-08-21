// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/kudaki_event/kudaki_event.proto

package kudaki_event

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

type KudakiEventStatus int32

const (
	KudakiEventStatus_UNPUBLISHED KudakiEventStatus = 0
	KudakiEventStatus_PUBLISHED   KudakiEventStatus = 1
	KudakiEventStatus_TAKEN_DOWN  KudakiEventStatus = 2
)

var KudakiEventStatus_name = map[int32]string{
	0: "UNPUBLISHED",
	1: "PUBLISHED",
	2: "TAKEN_DOWN",
}

var KudakiEventStatus_value = map[string]int32{
	"UNPUBLISHED": 0,
	"PUBLISHED":   1,
	"TAKEN_DOWN":  2,
}

func (x KudakiEventStatus) String() string {
	return proto.EnumName(KudakiEventStatus_name, int32(x))
}

func (KudakiEventStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c9e4a0310f5e824b, []int{0}
}

type KudakiEvent struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	OrganizerUserUuid    string               `protobuf:"bytes,3,opt,name=organizer_user_uuid,json=organizerUserUuid,proto3" json:"organizer_user_uuid,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Latitude             float64              `protobuf:"fixed64,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64              `protobuf:"fixed64,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Venue                string               `protobuf:"bytes,7,opt,name=venue,proto3" json:"venue,omitempty"`
	Description          string               `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	DurationFrom         *timestamp.Timestamp `protobuf:"bytes,9,opt,name=duration_from,json=durationFrom,proto3" json:"duration_from,omitempty"`
	DurationTo           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=duration_to,json=durationTo,proto3" json:"duration_to,omitempty"`
	AdDurationFrom       *timestamp.Timestamp `protobuf:"bytes,11,opt,name=ad_duration_from,json=adDurationFrom,proto3" json:"ad_duration_from,omitempty"`
	AdDurationTo         *timestamp.Timestamp `protobuf:"bytes,12,opt,name=ad_duration_to,json=adDurationTo,proto3" json:"ad_duration_to,omitempty"`
	Seen                 int32                `protobuf:"varint,13,opt,name=seen,proto3" json:"seen,omitempty"`
	Status               KudakiEventStatus    `protobuf:"varint,14,opt,name=status,proto3,enum=aggregates.kudaki_event.KudakiEventStatus" json:"status,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KudakiEvent) Reset()         { *m = KudakiEvent{} }
func (m *KudakiEvent) String() string { return proto.CompactTextString(m) }
func (*KudakiEvent) ProtoMessage()    {}
func (*KudakiEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e4a0310f5e824b, []int{0}
}

func (m *KudakiEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEvent.Unmarshal(m, b)
}
func (m *KudakiEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEvent.Marshal(b, m, deterministic)
}
func (m *KudakiEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEvent.Merge(m, src)
}
func (m *KudakiEvent) XXX_Size() int {
	return xxx_messageInfo_KudakiEvent.Size(m)
}
func (m *KudakiEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEvent.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEvent proto.InternalMessageInfo

func (m *KudakiEvent) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KudakiEvent) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *KudakiEvent) GetOrganizerUserUuid() string {
	if m != nil {
		return m.OrganizerUserUuid
	}
	return ""
}

func (m *KudakiEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KudakiEvent) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *KudakiEvent) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *KudakiEvent) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

func (m *KudakiEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *KudakiEvent) GetDurationFrom() *timestamp.Timestamp {
	if m != nil {
		return m.DurationFrom
	}
	return nil
}

func (m *KudakiEvent) GetDurationTo() *timestamp.Timestamp {
	if m != nil {
		return m.DurationTo
	}
	return nil
}

func (m *KudakiEvent) GetAdDurationFrom() *timestamp.Timestamp {
	if m != nil {
		return m.AdDurationFrom
	}
	return nil
}

func (m *KudakiEvent) GetAdDurationTo() *timestamp.Timestamp {
	if m != nil {
		return m.AdDurationTo
	}
	return nil
}

func (m *KudakiEvent) GetSeen() int32 {
	if m != nil {
		return m.Seen
	}
	return 0
}

func (m *KudakiEvent) GetStatus() KudakiEventStatus {
	if m != nil {
		return m.Status
	}
	return KudakiEventStatus_UNPUBLISHED
}

func (m *KudakiEvent) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("aggregates.kudaki_event.KudakiEventStatus", KudakiEventStatus_name, KudakiEventStatus_value)
	proto.RegisterType((*KudakiEvent)(nil), "aggregates.kudaki_event.KudakiEvent")
}

func init() {
	proto.RegisterFile("aggregates/kudaki_event/kudaki_event.proto", fileDescriptor_c9e4a0310f5e824b)
}

var fileDescriptor_c9e4a0310f5e824b = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcf, 0x6b, 0xdb, 0x3e,
	0x14, 0xff, 0x3a, 0x6d, 0xf2, 0x6d, 0x9e, 0x1b, 0x37, 0xd5, 0x06, 0x13, 0x61, 0x30, 0xb3, 0x93,
	0x09, 0xd4, 0x86, 0xee, 0x34, 0x76, 0xd8, 0x9a, 0x25, 0x63, 0x25, 0x23, 0x1b, 0x6e, 0xc2, 0x60,
	0x17, 0xa3, 0x44, 0x2f, 0x8e, 0x68, 0x6c, 0x05, 0x59, 0xca, 0x61, 0xa7, 0xfd, 0xe9, 0xc3, 0x72,
	0x7e, 0xb8, 0x8c, 0x92, 0x8b, 0x79, 0x4f, 0x9f, 0x1f, 0x7c, 0xfc, 0xa4, 0x07, 0x7d, 0x96, 0xa6,
	0x0a, 0x53, 0xa6, 0xb1, 0x88, 0x1e, 0x0d, 0x67, 0x8f, 0x22, 0xc1, 0x2d, 0xe6, 0xfa, 0x49, 0x13,
	0x6e, 0x94, 0xd4, 0x92, 0xbc, 0x3a, 0x72, 0xc3, 0x3a, 0xdc, 0x7b, 0x93, 0x4a, 0x99, 0xae, 0x31,
	0xb2, 0xb4, 0xb9, 0x59, 0x46, 0x5a, 0x64, 0x58, 0x68, 0x96, 0x6d, 0x2a, 0xe5, 0xdb, 0x3f, 0x4d,
	0x70, 0xc7, 0x56, 0x31, 0x2a, 0x05, 0xc4, 0x83, 0x86, 0xe0, 0xd4, 0xf1, 0x9d, 0xe0, 0x2c, 0x6e,
	0x08, 0x4e, 0x08, 0x9c, 0x1b, 0x23, 0x38, 0x6d, 0xf8, 0x4e, 0xd0, 0x8e, 0x6d, 0x4d, 0x42, 0x78,
	0x21, 0x55, 0xca, 0x72, 0xf1, 0x1b, 0x55, 0x62, 0x8a, 0xf2, 0x53, 0x52, 0xce, 0x2c, 0xe5, 0xfa,
	0x00, 0xcd, 0x0a, 0x54, 0x33, 0x53, 0x79, 0xe4, 0x2c, 0x43, 0x7a, 0x5e, 0x79, 0x94, 0x35, 0xe9,
	0xc1, 0xc5, 0x9a, 0x69, 0xa1, 0x0d, 0x47, 0xda, 0xf4, 0x9d, 0xc0, 0x89, 0x0f, 0x3d, 0x79, 0x0d,
	0xed, 0xb5, 0xcc, 0xd3, 0x0a, 0x6c, 0x59, 0xf0, 0x78, 0x40, 0x5e, 0x42, 0x73, 0x8b, 0xb9, 0x41,
	0xfa, 0xbf, 0xb5, 0xab, 0x1a, 0xe2, 0x83, 0xcb, 0xb1, 0x58, 0x28, 0xb1, 0xd1, 0x42, 0xe6, 0xf4,
	0xc2, 0x62, 0xf5, 0x23, 0xf2, 0x11, 0x3a, 0xdc, 0x28, 0x56, 0xd6, 0xc9, 0x52, 0xc9, 0x8c, 0xb6,
	0x7d, 0x27, 0x70, 0x6f, 0x7b, 0x61, 0x35, 0xa2, 0x70, 0x3f, 0xa2, 0x70, 0xba, 0x1f, 0x51, 0x7c,
	0xb9, 0x17, 0x7c, 0x51, 0x32, 0x23, 0x1f, 0xc0, 0x3d, 0x18, 0x68, 0x49, 0xe1, 0xa4, 0x1c, 0xf6,
	0xf4, 0xa9, 0x24, 0x43, 0xe8, 0x32, 0x9e, 0x3c, 0x0d, 0xe0, 0x9e, 0x74, 0xf0, 0x18, 0x1f, 0xd6,
	0x23, 0x7c, 0x02, 0xaf, 0xee, 0xa2, 0x25, 0xbd, 0x3c, 0xfd, 0x13, 0x47, 0x8f, 0xa9, 0x2c, 0xef,
	0xa2, 0x40, 0xcc, 0x69, 0xc7, 0x77, 0x82, 0x66, 0x6c, 0x6b, 0x32, 0x80, 0x56, 0xa1, 0x99, 0x36,
	0x05, 0xf5, 0x7c, 0x27, 0xf0, 0x6e, 0xfb, 0xe1, 0x33, 0xcf, 0x29, 0xac, 0xbd, 0x94, 0x07, 0xab,
	0x88, 0x77, 0x4a, 0xf2, 0x1e, 0x60, 0xa1, 0x90, 0x69, 0xe4, 0x09, 0xd3, 0xf4, 0xea, 0x64, 0xaa,
	0xf6, 0x8e, 0x7d, 0xa7, 0xfb, 0x9f, 0xe1, 0xfa, 0x1f, 0x5f, 0x72, 0x05, 0xee, 0x6c, 0xf2, 0x63,
	0x36, 0xf8, 0x76, 0xff, 0xf0, 0x75, 0x34, 0xec, 0xfe, 0x47, 0x3a, 0xd0, 0x3e, 0xb6, 0x0e, 0xf1,
	0x00, 0xa6, 0x77, 0xe3, 0xd1, 0x24, 0x19, 0x7e, 0xff, 0x39, 0xe9, 0x36, 0x06, 0xe3, 0x5f, 0xf7,
	0xa9, 0xd0, 0x2b, 0x33, 0x0f, 0x17, 0x32, 0x8b, 0xc4, 0x7a, 0xc5, 0xb2, 0x6c, 0xc5, 0xf9, 0x6e,
	0x59, 0x6e, 0x96, 0x62, 0x8d, 0x37, 0x05, 0xaa, 0xad, 0x58, 0x60, 0x84, 0xb9, 0x16, 0x5a, 0x60,
	0x11, 0x3d, 0xb3, 0x61, 0xf3, 0x96, 0x0d, 0xfc, 0xee, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5,
	0x7c, 0xd7, 0x11, 0x83, 0x03, 0x00, 0x00,
}
