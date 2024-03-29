// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/kudaki_event.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	kudaki_event "github.com/ilhammhdd/kudaki-event-payment-service/entities/aggregates/kudaki_event"
	user "github.com/ilhammhdd/kudaki-event-payment-service/entities/aggregates/user"
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

type EventServiceCommandTopic int32

const (
	EventServiceCommandTopic_ADD_KUDAKI_EVENT            EventServiceCommandTopic = 0
	EventServiceCommandTopic_UPDATE_KUDAKI_EVENT         EventServiceCommandTopic = 1
	EventServiceCommandTopic_DELETE_KUDAKI_EVENT         EventServiceCommandTopic = 2
	EventServiceCommandTopic_RETRIEVE_KUDAKI_EVENT       EventServiceCommandTopic = 3
	EventServiceCommandTopic_RETRIEVE_KUDAKI_EVENTS      EventServiceCommandTopic = 4
	EventServiceCommandTopic_RETRIEVE_ORGANIZER_INVOICES EventServiceCommandTopic = 5
)

var EventServiceCommandTopic_name = map[int32]string{
	0: "ADD_KUDAKI_EVENT",
	1: "UPDATE_KUDAKI_EVENT",
	2: "DELETE_KUDAKI_EVENT",
	3: "RETRIEVE_KUDAKI_EVENT",
	4: "RETRIEVE_KUDAKI_EVENTS",
	5: "RETRIEVE_ORGANIZER_INVOICES",
}

var EventServiceCommandTopic_value = map[string]int32{
	"ADD_KUDAKI_EVENT":            0,
	"UPDATE_KUDAKI_EVENT":         1,
	"DELETE_KUDAKI_EVENT":         2,
	"RETRIEVE_KUDAKI_EVENT":       3,
	"RETRIEVE_KUDAKI_EVENTS":      4,
	"RETRIEVE_ORGANIZER_INVOICES": 5,
}

func (x EventServiceCommandTopic) String() string {
	return proto.EnumName(EventServiceCommandTopic_name, int32(x))
}

func (EventServiceCommandTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{0}
}

type EventServiceEventTopic int32

const (
	EventServiceEventTopic_KUDAKI_EVENT_ADDED           EventServiceEventTopic = 0
	EventServiceEventTopic_KUDAKI_EVENT_UPDATED         EventServiceEventTopic = 1
	EventServiceEventTopic_KUDAKI_EVENT_DELETED         EventServiceEventTopic = 2
	EventServiceEventTopic_KUDAKI_EVENT_RETRIEVED       EventServiceEventTopic = 3
	EventServiceEventTopic_KUDAKI_EVENTS_RETRIEVED      EventServiceEventTopic = 4
	EventServiceEventTopic_ORGANIZER_INVOICES_RETRIEVED EventServiceEventTopic = 5
)

var EventServiceEventTopic_name = map[int32]string{
	0: "KUDAKI_EVENT_ADDED",
	1: "KUDAKI_EVENT_UPDATED",
	2: "KUDAKI_EVENT_DELETED",
	3: "KUDAKI_EVENT_RETRIEVED",
	4: "KUDAKI_EVENTS_RETRIEVED",
	5: "ORGANIZER_INVOICES_RETRIEVED",
}

var EventServiceEventTopic_value = map[string]int32{
	"KUDAKI_EVENT_ADDED":           0,
	"KUDAKI_EVENT_UPDATED":         1,
	"KUDAKI_EVENT_DELETED":         2,
	"KUDAKI_EVENT_RETRIEVED":       3,
	"KUDAKI_EVENTS_RETRIEVED":      4,
	"ORGANIZER_INVOICES_RETRIEVED": 5,
}

func (x EventServiceEventTopic) String() string {
	return proto.EnumName(EventServiceEventTopic_name, int32(x))
}

func (EventServiceEventTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{1}
}

type EventPaymentServiceCommandTopic int32

const (
	EventPaymentServiceCommandTopic_PAYMENT_REQUEST_DOKU EventPaymentServiceCommandTopic = 0
)

var EventPaymentServiceCommandTopic_name = map[int32]string{
	0: "PAYMENT_REQUEST_DOKU",
}

var EventPaymentServiceCommandTopic_value = map[string]int32{
	"PAYMENT_REQUEST_DOKU": 0,
}

func (x EventPaymentServiceCommandTopic) String() string {
	return proto.EnumName(EventPaymentServiceCommandTopic_name, int32(x))
}

func (EventPaymentServiceCommandTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{2}
}

type EventPaymentServiceEventTopic int32

const (
	EventPaymentServiceEventTopic_EVENT_DOKU_INVOICE_ISSUED   EventPaymentServiceEventTopic = 0
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_IDENTIFY EventPaymentServiceEventTopic = 1
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_REDIRECT EventPaymentServiceEventTopic = 2
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_NOTIFY   EventPaymentServiceEventTopic = 3
	EventPaymentServiceEventTopic_PAYMENT_REQUESTED_DOKU      EventPaymentServiceEventTopic = 4
)

var EventPaymentServiceEventTopic_name = map[int32]string{
	0: "EVENT_DOKU_INVOICE_ISSUED",
	1: "EVENT_DOKU_PAYMENT_IDENTIFY",
	2: "EVENT_DOKU_PAYMENT_REDIRECT",
	3: "EVENT_DOKU_PAYMENT_NOTIFY",
	4: "PAYMENT_REQUESTED_DOKU",
}

var EventPaymentServiceEventTopic_value = map[string]int32{
	"EVENT_DOKU_INVOICE_ISSUED":   0,
	"EVENT_DOKU_PAYMENT_IDENTIFY": 1,
	"EVENT_DOKU_PAYMENT_REDIRECT": 2,
	"EVENT_DOKU_PAYMENT_NOTIFY":   3,
	"PAYMENT_REQUESTED_DOKU":      4,
}

func (x EventPaymentServiceEventTopic) String() string {
	return proto.EnumName(EventPaymentServiceEventTopic_name, int32(x))
}

func (EventPaymentServiceEventTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{3}
}

type AddKudakiEvent struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Venue                string   `protobuf:"bytes,2,opt,name=venue,proto3" json:"venue,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DurationFrom         int64    `protobuf:"varint,4,opt,name=duration_from,json=durationFrom,proto3" json:"duration_from,omitempty"`
	DurationTo           int64    `protobuf:"varint,5,opt,name=duration_to,json=durationTo,proto3" json:"duration_to,omitempty"`
	AdDurationFrom       int64    `protobuf:"varint,7,opt,name=ad_duration_from,json=adDurationFrom,proto3" json:"ad_duration_from,omitempty"`
	AdDurationTo         int64    `protobuf:"varint,8,opt,name=ad_duration_to,json=adDurationTo,proto3" json:"ad_duration_to,omitempty"`
	KudakiToken          string   `protobuf:"bytes,9,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddKudakiEvent) Reset()         { *m = AddKudakiEvent{} }
func (m *AddKudakiEvent) String() string { return proto.CompactTextString(m) }
func (*AddKudakiEvent) ProtoMessage()    {}
func (*AddKudakiEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{0}
}

func (m *AddKudakiEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKudakiEvent.Unmarshal(m, b)
}
func (m *AddKudakiEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKudakiEvent.Marshal(b, m, deterministic)
}
func (m *AddKudakiEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKudakiEvent.Merge(m, src)
}
func (m *AddKudakiEvent) XXX_Size() int {
	return xxx_messageInfo_AddKudakiEvent.Size(m)
}
func (m *AddKudakiEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKudakiEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AddKudakiEvent proto.InternalMessageInfo

func (m *AddKudakiEvent) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AddKudakiEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddKudakiEvent) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

func (m *AddKudakiEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddKudakiEvent) GetDurationFrom() int64 {
	if m != nil {
		return m.DurationFrom
	}
	return 0
}

func (m *AddKudakiEvent) GetDurationTo() int64 {
	if m != nil {
		return m.DurationTo
	}
	return 0
}

func (m *AddKudakiEvent) GetAdDurationFrom() int64 {
	if m != nil {
		return m.AdDurationFrom
	}
	return 0
}

func (m *AddKudakiEvent) GetAdDurationTo() int64 {
	if m != nil {
		return m.AdDurationTo
	}
	return 0
}

func (m *AddKudakiEvent) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

type DeleteKudakiEvent struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	EventUuid            string   `protobuf:"bytes,2,opt,name=event_uuid,json=eventUuid,proto3" json:"event_uuid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,3,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteKudakiEvent) Reset()         { *m = DeleteKudakiEvent{} }
func (m *DeleteKudakiEvent) String() string { return proto.CompactTextString(m) }
func (*DeleteKudakiEvent) ProtoMessage()    {}
func (*DeleteKudakiEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{1}
}

func (m *DeleteKudakiEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteKudakiEvent.Unmarshal(m, b)
}
func (m *DeleteKudakiEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteKudakiEvent.Marshal(b, m, deterministic)
}
func (m *DeleteKudakiEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteKudakiEvent.Merge(m, src)
}
func (m *DeleteKudakiEvent) XXX_Size() int {
	return xxx_messageInfo_DeleteKudakiEvent.Size(m)
}
func (m *DeleteKudakiEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteKudakiEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteKudakiEvent proto.InternalMessageInfo

func (m *DeleteKudakiEvent) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeleteKudakiEvent) GetEventUuid() string {
	if m != nil {
		return m.EventUuid
	}
	return ""
}

func (m *DeleteKudakiEvent) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

type RetrieveOrganizerInvoices struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	ResultSchema         []byte   `protobuf:"bytes,3,opt,name=result_schema,json=resultSchema,proto3" json:"result_schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveOrganizerInvoices) Reset()         { *m = RetrieveOrganizerInvoices{} }
func (m *RetrieveOrganizerInvoices) String() string { return proto.CompactTextString(m) }
func (*RetrieveOrganizerInvoices) ProtoMessage()    {}
func (*RetrieveOrganizerInvoices) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{2}
}

func (m *RetrieveOrganizerInvoices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveOrganizerInvoices.Unmarshal(m, b)
}
func (m *RetrieveOrganizerInvoices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveOrganizerInvoices.Marshal(b, m, deterministic)
}
func (m *RetrieveOrganizerInvoices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveOrganizerInvoices.Merge(m, src)
}
func (m *RetrieveOrganizerInvoices) XXX_Size() int {
	return xxx_messageInfo_RetrieveOrganizerInvoices.Size(m)
}
func (m *RetrieveOrganizerInvoices) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveOrganizerInvoices.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveOrganizerInvoices proto.InternalMessageInfo

func (m *RetrieveOrganizerInvoices) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RetrieveOrganizerInvoices) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *RetrieveOrganizerInvoices) GetResultSchema() []byte {
	if m != nil {
		return m.ResultSchema
	}
	return nil
}

type KudakiEventAdded struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.User                `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	KudakiEvent          *kudaki_event.KudakiEvent `protobuf:"bytes,4,opt,name=kudaki_event,json=kudakiEvent,proto3" json:"kudaki_event,omitempty"`
	AddKudakiEvent       *AddKudakiEvent           `protobuf:"bytes,5,opt,name=add_kudaki_event,json=addKudakiEvent,proto3" json:"add_kudaki_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KudakiEventAdded) Reset()         { *m = KudakiEventAdded{} }
func (m *KudakiEventAdded) String() string { return proto.CompactTextString(m) }
func (*KudakiEventAdded) ProtoMessage()    {}
func (*KudakiEventAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{3}
}

func (m *KudakiEventAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEventAdded.Unmarshal(m, b)
}
func (m *KudakiEventAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEventAdded.Marshal(b, m, deterministic)
}
func (m *KudakiEventAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEventAdded.Merge(m, src)
}
func (m *KudakiEventAdded) XXX_Size() int {
	return xxx_messageInfo_KudakiEventAdded.Size(m)
}
func (m *KudakiEventAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEventAdded.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEventAdded proto.InternalMessageInfo

func (m *KudakiEventAdded) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KudakiEventAdded) GetOrganizer() *user.User {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *KudakiEventAdded) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *KudakiEventAdded) GetKudakiEvent() *kudaki_event.KudakiEvent {
	if m != nil {
		return m.KudakiEvent
	}
	return nil
}

func (m *KudakiEventAdded) GetAddKudakiEvent() *AddKudakiEvent {
	if m != nil {
		return m.AddKudakiEvent
	}
	return nil
}

type KudakiEventDeleted struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.User                `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	KudakiEvent          *kudaki_event.KudakiEvent `protobuf:"bytes,4,opt,name=kudaki_event,json=kudakiEvent,proto3" json:"kudaki_event,omitempty"`
	DeleteKudakiEvent    *DeleteKudakiEvent        `protobuf:"bytes,5,opt,name=delete_kudaki_event,json=deleteKudakiEvent,proto3" json:"delete_kudaki_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KudakiEventDeleted) Reset()         { *m = KudakiEventDeleted{} }
func (m *KudakiEventDeleted) String() string { return proto.CompactTextString(m) }
func (*KudakiEventDeleted) ProtoMessage()    {}
func (*KudakiEventDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{4}
}

func (m *KudakiEventDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEventDeleted.Unmarshal(m, b)
}
func (m *KudakiEventDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEventDeleted.Marshal(b, m, deterministic)
}
func (m *KudakiEventDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEventDeleted.Merge(m, src)
}
func (m *KudakiEventDeleted) XXX_Size() int {
	return xxx_messageInfo_KudakiEventDeleted.Size(m)
}
func (m *KudakiEventDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEventDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEventDeleted proto.InternalMessageInfo

func (m *KudakiEventDeleted) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KudakiEventDeleted) GetOrganizer() *user.User {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *KudakiEventDeleted) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *KudakiEventDeleted) GetKudakiEvent() *kudaki_event.KudakiEvent {
	if m != nil {
		return m.KudakiEvent
	}
	return nil
}

func (m *KudakiEventDeleted) GetDeleteKudakiEvent() *DeleteKudakiEvent {
	if m != nil {
		return m.DeleteKudakiEvent
	}
	return nil
}

type OrganizerInvoicesRetrieved struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.User `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status    `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	Result               []byte     `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OrganizerInvoicesRetrieved) Reset()         { *m = OrganizerInvoicesRetrieved{} }
func (m *OrganizerInvoicesRetrieved) String() string { return proto.CompactTextString(m) }
func (*OrganizerInvoicesRetrieved) ProtoMessage()    {}
func (*OrganizerInvoicesRetrieved) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{5}
}

func (m *OrganizerInvoicesRetrieved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizerInvoicesRetrieved.Unmarshal(m, b)
}
func (m *OrganizerInvoicesRetrieved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizerInvoicesRetrieved.Marshal(b, m, deterministic)
}
func (m *OrganizerInvoicesRetrieved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizerInvoicesRetrieved.Merge(m, src)
}
func (m *OrganizerInvoicesRetrieved) XXX_Size() int {
	return xxx_messageInfo_OrganizerInvoicesRetrieved.Size(m)
}
func (m *OrganizerInvoicesRetrieved) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizerInvoicesRetrieved.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizerInvoicesRetrieved proto.InternalMessageInfo

func (m *OrganizerInvoicesRetrieved) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *OrganizerInvoicesRetrieved) GetOrganizer() *user.User {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *OrganizerInvoicesRetrieved) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *OrganizerInvoicesRetrieved) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

type PaymentRequestDoku struct {
	Uid                   string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken           string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	TransactionIdMerchant string   `protobuf:"bytes,3,opt,name=transaction_id_merchant,json=transactionIdMerchant,proto3" json:"transaction_id_merchant,omitempty"`
	SessionId             string   `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	HashedWords           string   `protobuf:"bytes,5,opt,name=hashed_words,json=hashedWords,proto3" json:"hashed_words,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *PaymentRequestDoku) Reset()         { *m = PaymentRequestDoku{} }
func (m *PaymentRequestDoku) String() string { return proto.CompactTextString(m) }
func (*PaymentRequestDoku) ProtoMessage()    {}
func (*PaymentRequestDoku) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{6}
}

func (m *PaymentRequestDoku) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentRequestDoku.Unmarshal(m, b)
}
func (m *PaymentRequestDoku) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentRequestDoku.Marshal(b, m, deterministic)
}
func (m *PaymentRequestDoku) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentRequestDoku.Merge(m, src)
}
func (m *PaymentRequestDoku) XXX_Size() int {
	return xxx_messageInfo_PaymentRequestDoku.Size(m)
}
func (m *PaymentRequestDoku) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentRequestDoku.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentRequestDoku proto.InternalMessageInfo

func (m *PaymentRequestDoku) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *PaymentRequestDoku) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *PaymentRequestDoku) GetTransactionIdMerchant() string {
	if m != nil {
		return m.TransactionIdMerchant
	}
	return ""
}

func (m *PaymentRequestDoku) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *PaymentRequestDoku) GetHashedWords() string {
	if m != nil {
		return m.HashedWords
	}
	return ""
}

type KudakiEventDokuInvoiceIssued struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.User                `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	DokuInvoice          *kudaki_event.DokuInvoice `protobuf:"bytes,4,opt,name=doku_invoice,json=dokuInvoice,proto3" json:"doku_invoice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KudakiEventDokuInvoiceIssued) Reset()         { *m = KudakiEventDokuInvoiceIssued{} }
func (m *KudakiEventDokuInvoiceIssued) String() string { return proto.CompactTextString(m) }
func (*KudakiEventDokuInvoiceIssued) ProtoMessage()    {}
func (*KudakiEventDokuInvoiceIssued) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{7}
}

func (m *KudakiEventDokuInvoiceIssued) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEventDokuInvoiceIssued.Unmarshal(m, b)
}
func (m *KudakiEventDokuInvoiceIssued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEventDokuInvoiceIssued.Marshal(b, m, deterministic)
}
func (m *KudakiEventDokuInvoiceIssued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEventDokuInvoiceIssued.Merge(m, src)
}
func (m *KudakiEventDokuInvoiceIssued) XXX_Size() int {
	return xxx_messageInfo_KudakiEventDokuInvoiceIssued.Size(m)
}
func (m *KudakiEventDokuInvoiceIssued) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEventDokuInvoiceIssued.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEventDokuInvoiceIssued proto.InternalMessageInfo

func (m *KudakiEventDokuInvoiceIssued) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KudakiEventDokuInvoiceIssued) GetOrganizer() *user.User {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *KudakiEventDokuInvoiceIssued) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *KudakiEventDokuInvoiceIssued) GetDokuInvoice() *kudaki_event.DokuInvoice {
	if m != nil {
		return m.DokuInvoice
	}
	return nil
}

type PaymentRequestedDoku struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.User                `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	DokuInvoice          *kudaki_event.DokuInvoice `protobuf:"bytes,4,opt,name=doku_invoice,json=dokuInvoice,proto3" json:"doku_invoice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *PaymentRequestedDoku) Reset()         { *m = PaymentRequestedDoku{} }
func (m *PaymentRequestedDoku) String() string { return proto.CompactTextString(m) }
func (*PaymentRequestedDoku) ProtoMessage()    {}
func (*PaymentRequestedDoku) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{8}
}

func (m *PaymentRequestedDoku) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentRequestedDoku.Unmarshal(m, b)
}
func (m *PaymentRequestedDoku) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentRequestedDoku.Marshal(b, m, deterministic)
}
func (m *PaymentRequestedDoku) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentRequestedDoku.Merge(m, src)
}
func (m *PaymentRequestedDoku) XXX_Size() int {
	return xxx_messageInfo_PaymentRequestedDoku.Size(m)
}
func (m *PaymentRequestedDoku) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentRequestedDoku.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentRequestedDoku proto.InternalMessageInfo

func (m *PaymentRequestedDoku) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *PaymentRequestedDoku) GetOrganizer() *user.User {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *PaymentRequestedDoku) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *PaymentRequestedDoku) GetDokuInvoice() *kudaki_event.DokuInvoice {
	if m != nil {
		return m.DokuInvoice
	}
	return nil
}

func init() {
	proto.RegisterEnum("aggregates.event.EventServiceCommandTopic", EventServiceCommandTopic_name, EventServiceCommandTopic_value)
	proto.RegisterEnum("aggregates.event.EventServiceEventTopic", EventServiceEventTopic_name, EventServiceEventTopic_value)
	proto.RegisterEnum("aggregates.event.EventPaymentServiceCommandTopic", EventPaymentServiceCommandTopic_name, EventPaymentServiceCommandTopic_value)
	proto.RegisterEnum("aggregates.event.EventPaymentServiceEventTopic", EventPaymentServiceEventTopic_name, EventPaymentServiceEventTopic_value)
	proto.RegisterType((*AddKudakiEvent)(nil), "aggregates.event.AddKudakiEvent")
	proto.RegisterType((*DeleteKudakiEvent)(nil), "aggregates.event.DeleteKudakiEvent")
	proto.RegisterType((*RetrieveOrganizerInvoices)(nil), "aggregates.event.RetrieveOrganizerInvoices")
	proto.RegisterType((*KudakiEventAdded)(nil), "aggregates.event.KudakiEventAdded")
	proto.RegisterType((*KudakiEventDeleted)(nil), "aggregates.event.KudakiEventDeleted")
	proto.RegisterType((*OrganizerInvoicesRetrieved)(nil), "aggregates.event.OrganizerInvoicesRetrieved")
	proto.RegisterType((*PaymentRequestDoku)(nil), "aggregates.event.PaymentRequestDoku")
	proto.RegisterType((*KudakiEventDokuInvoiceIssued)(nil), "aggregates.event.KudakiEventDokuInvoiceIssued")
	proto.RegisterType((*PaymentRequestedDoku)(nil), "aggregates.event.PaymentRequestedDoku")
}

func init() { proto.RegisterFile("events/kudaki_event.proto", fileDescriptor_fa36ed941739bd5b) }

var fileDescriptor_fa36ed941739bd5b = []byte{
	// 960 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0x36, 0xf5, 0xc7, 0xbf, 0x9f, 0x46, 0x8a, 0xc1, 0xac, 0x65, 0x9b, 0x96, 0x63, 0x58, 0x55,
	0x72, 0x30, 0x04, 0x58, 0x2a, 0x1c, 0xa0, 0x97, 0x9e, 0x54, 0x93, 0x09, 0x58, 0x35, 0x92, 0x4b,
	0x49, 0x2e, 0x92, 0x0b, 0xc1, 0x68, 0xb7, 0x12, 0x21, 0x93, 0xeb, 0x72, 0x49, 0x15, 0xed, 0xd3,
	0xf4, 0xd6, 0x43, 0xd1, 0x17, 0x68, 0x0f, 0x3d, 0xf4, 0x25, 0x7a, 0xec, 0xa3, 0x14, 0x9c, 0xa5,
	0xec, 0xa5, 0xa4, 0x34, 0x40, 0x2f, 0x41, 0x7b, 0x31, 0xb8, 0xdf, 0x37, 0x9c, 0xf9, 0x66, 0xe6,
	0x5b, 0x99, 0x70, 0xcc, 0x96, 0x2c, 0x8c, 0x45, 0x77, 0x91, 0x50, 0x6f, 0xe1, 0xbb, 0x78, 0xea,
	0xdc, 0x45, 0x3c, 0xe6, 0x44, 0xf7, 0x66, 0xb3, 0x88, 0xcd, 0xbc, 0x98, 0x89, 0x0e, 0xe2, 0x8d,
	0xc6, 0x03, 0xd2, 0x4d, 0x04, 0x8b, 0xf0, 0x8f, 0x8c, 0x6e, 0xb4, 0x15, 0x4e, 0x4d, 0xd6, 0xa5,
	0x7c, 0x91, 0xb8, 0x7e, 0xb8, 0xe4, 0xfe, 0x94, 0xbd, 0x2f, 0x76, 0x53, 0x45, 0x63, 0x3f, 0x13,
	0x28, 0x62, 0x2f, 0x4e, 0x84, 0x04, 0x5b, 0x3f, 0x16, 0x60, 0xaf, 0x47, 0x69, 0x1f, 0xc3, 0xad,
	0x34, 0x80, 0xe8, 0x50, 0x4c, 0x7c, 0x6a, 0x68, 0x4d, 0xed, 0xbc, 0xe2, 0xa4, 0x8f, 0x84, 0x40,
	0x29, 0xf4, 0x02, 0x66, 0xec, 0x22, 0x84, 0xcf, 0xa4, 0x0e, 0xe5, 0x25, 0x0b, 0x13, 0x66, 0x14,
	0x10, 0x94, 0x07, 0xd2, 0x84, 0x2a, 0x65, 0x62, 0x1a, 0xf9, 0x77, 0xb1, 0xcf, 0x43, 0xa3, 0x88,
	0x9c, 0x0a, 0x91, 0xa7, 0xf0, 0x88, 0x26, 0x91, 0x97, 0x3e, 0xbb, 0x5f, 0x47, 0x3c, 0x30, 0x4a,
	0x4d, 0xed, 0xbc, 0xe8, 0xd4, 0x56, 0xe0, 0x8b, 0x88, 0x07, 0xe4, 0x0c, 0xaa, 0xf7, 0x41, 0x31,
	0x37, 0xca, 0x18, 0x02, 0x2b, 0x68, 0xcc, 0xc9, 0x39, 0xe8, 0x1e, 0x75, 0xf3, 0x89, 0xfe, 0x87,
	0x51, 0x7b, 0x1e, 0x35, 0xd5, 0x54, 0xcf, 0x60, 0x4f, 0x8d, 0x8c, 0xb9, 0xf1, 0x7f, 0x59, 0xf0,
	0x21, 0x6e, 0xcc, 0xc9, 0x47, 0x50, 0xcb, 0x26, 0x16, 0xf3, 0x05, 0x0b, 0x8d, 0x8a, 0x14, 0x2e,
	0xb1, 0x71, 0x0a, 0xb5, 0x66, 0xf0, 0xd8, 0x64, 0xb7, 0x2c, 0x66, 0x7f, 0x3f, 0xab, 0x53, 0x00,
	0x9c, 0xb3, 0x9b, 0xa4, 0x84, 0x1c, 0x4e, 0x05, 0x91, 0x49, 0x4a, 0xaf, 0x17, 0x2a, 0x6e, 0x16,
	0x4a, 0xe0, 0xd8, 0x61, 0x71, 0xe4, 0xb3, 0x25, 0x1b, 0x46, 0x33, 0x2f, 0xf4, 0xbf, 0x67, 0x91,
	0x2d, 0xb7, 0x2e, 0xb6, 0x14, 0x5c, 0xcf, 0x58, 0xd8, 0xc8, 0x98, 0xce, 0x3c, 0x62, 0x22, 0xb9,
	0x8d, 0x5d, 0x31, 0x9d, 0xb3, 0xc0, 0xc3, 0xaa, 0x35, 0xa7, 0x26, 0xc1, 0x11, 0x62, 0xad, 0x1f,
	0x0a, 0xa0, 0x2b, 0xad, 0xf5, 0x28, 0x65, 0x74, 0x4b, 0xb9, 0xe7, 0x50, 0xe1, 0x2b, 0x55, 0x58,
	0xab, 0x7a, 0x79, 0xd0, 0x51, 0xfc, 0x8d, 0x46, 0x9e, 0x08, 0x16, 0x39, 0x0f, 0x71, 0xe4, 0x63,
	0xa8, 0xc9, 0xa1, 0x48, 0xef, 0x61, 0xfd, 0xea, 0xe5, 0x23, 0x79, 0x19, 0x3a, 0x23, 0x04, 0x9d,
	0x2a, 0x9e, 0xe4, 0x81, 0xbc, 0xbc, 0xef, 0x0a, 0x51, 0x74, 0x49, 0xf5, 0xf2, 0x99, 0x5a, 0x29,
	0x67, 0x71, 0x45, 0xf9, 0xaa, 0x77, 0xb9, 0xa1, 0xcf, 0x53, 0xa7, 0x50, 0x37, 0x97, 0xac, 0x8c,
	0xc9, 0x9a, 0x9d, 0xf5, 0x6b, 0xd9, 0xc9, 0xdf, 0x84, 0xd4, 0x4b, 0xea, 0xb9, 0xf5, 0x53, 0x01,
	0x88, 0x72, 0x96, 0x76, 0xf8, 0x2f, 0x0c, 0x69, 0x04, 0xfb, 0x14, 0x9b, 0xd9, 0x36, 0xa7, 0xa7,
	0x9b, 0x73, 0xda, 0xb8, 0x08, 0xce, 0x63, 0xba, 0x0e, 0xb5, 0x7e, 0xd6, 0xa0, 0xb1, 0x61, 0xe0,
	0x95, 0xb3, 0x3f, 0xe0, 0xd4, 0x0e, 0x61, 0x57, 0x1a, 0x1f, 0xe7, 0x55, 0x73, 0xb2, 0x53, 0xeb,
	0x77, 0x0d, 0xc8, 0xb5, 0xf7, 0x5d, 0x90, 0xb6, 0xc3, 0xbe, 0x49, 0x98, 0x88, 0x4d, 0xbe, 0x48,
	0xfe, 0xd9, 0x8d, 0xfb, 0x04, 0x8e, 0xe2, 0xc8, 0x0b, 0x85, 0x37, 0xc5, 0x5f, 0x1d, 0x9f, 0xba,
	0x01, 0x8b, 0xa6, 0x73, 0x2f, 0x8c, 0xb3, 0x1b, 0x7f, 0xa0, 0xd0, 0x36, 0x7d, 0x95, 0x91, 0xe9,
	0xaf, 0x87, 0x60, 0x42, 0xc8, 0x77, 0x50, 0x5f, 0xc5, 0xa9, 0x64, 0x88, 0x8d, 0x95, 0xe7, 0x9e,
	0x98, 0x33, 0xea, 0x7e, 0xcb, 0x23, 0x2a, 0x70, 0x41, 0x15, 0xa7, 0x2a, 0xb1, 0xaf, 0x52, 0xa8,
	0xf5, 0xa7, 0x06, 0x4f, 0x54, 0x8f, 0xf2, 0x45, 0x92, 0xcd, 0xde, 0x16, 0x22, 0xf9, 0xc0, 0x6e,
	0x55, 0xff, 0x83, 0xbd, 0xd7, 0xad, 0x8a, 0x74, 0xa7, 0x4a, 0x1f, 0x0e, 0xad, 0x3f, 0x34, 0xa8,
	0xe7, 0x17, 0xc5, 0xe8, 0x3b, 0x56, 0xf5, 0x2f, 0x6b, 0xad, 0xfd, 0x8b, 0x06, 0x06, 0xee, 0x6d,
	0xc4, 0xa2, 0xa5, 0x3f, 0x65, 0x57, 0x3c, 0x08, 0xbc, 0x90, 0x8e, 0xf9, 0x9d, 0x3f, 0x25, 0x75,
	0xd0, 0x7b, 0xa6, 0xe9, 0xf6, 0x27, 0x66, 0xaf, 0x6f, 0xbb, 0xd6, 0x8d, 0x35, 0x18, 0xeb, 0x3b,
	0xe4, 0x08, 0xf6, 0x27, 0xd7, 0x66, 0x6f, 0x6c, 0xe5, 0x09, 0x2d, 0x25, 0x4c, 0xeb, 0x0b, 0x6b,
	0x9d, 0x28, 0x90, 0x63, 0x38, 0x70, 0xac, 0xb1, 0x63, 0x5b, 0x37, 0x6b, 0x54, 0x91, 0x34, 0xe0,
	0x70, 0x2b, 0x35, 0xd2, 0x4b, 0xe4, 0x0c, 0x4e, 0xee, 0xb9, 0xa1, 0xf3, 0xb2, 0x37, 0xb0, 0xdf,
	0x58, 0x8e, 0x6b, 0x0f, 0x6e, 0x86, 0xf6, 0x95, 0x35, 0xd2, 0xcb, 0xed, 0xdf, 0x34, 0x38, 0x54,
	0xc5, 0xe3, 0xb3, 0x94, 0x7e, 0x08, 0x44, 0x4d, 0xe7, 0xf6, 0x4c, 0xd3, 0x32, 0xf5, 0x1d, 0x62,
	0x40, 0x3d, 0x87, 0xcb, 0x4e, 0x4c, 0x5d, 0xdb, 0x60, 0x64, 0x2b, 0xa6, 0x5e, 0x48, 0x35, 0xe6,
	0x98, 0x95, 0x28, 0x53, 0x2f, 0x92, 0x13, 0x38, 0xca, 0xc9, 0x56, 0xc8, 0x12, 0x69, 0xc2, 0x93,
	0x4d, 0xdd, 0x4a, 0x44, 0xb9, 0xfd, 0x29, 0x9c, 0xa1, 0xe8, 0xcc, 0x5d, 0xdb, 0x96, 0x60, 0x40,
	0xfd, 0xba, 0xf7, 0xfa, 0x95, 0x2c, 0xfc, 0xe5, 0xc4, 0x1a, 0x8d, 0x5d, 0x73, 0xd8, 0x9f, 0xe8,
	0x3b, 0xed, 0x5f, 0x35, 0x38, 0xdd, 0xf2, 0xb6, 0x32, 0x85, 0x53, 0x38, 0xce, 0x9a, 0x19, 0xf6,
	0x27, 0x2b, 0x05, 0xae, 0x3d, 0x1a, 0x4d, 0x70, 0x18, 0x67, 0x70, 0xa2, 0xd0, 0xab, 0x2a, 0xb6,
	0x69, 0x0d, 0xc6, 0xf6, 0x8b, 0xd7, 0xba, 0xf6, 0x8e, 0x00, 0xc7, 0x32, 0x6d, 0xc7, 0xba, 0x4a,
	0x37, 0x9b, 0x2f, 0xb0, 0x0a, 0x18, 0x0c, 0xf1, 0x7d, 0xdc, 0xee, 0x9a, 0x76, 0xcb, 0x94, 0xea,
	0x4b, 0x9f, 0x5d, 0xbd, 0xe9, 0xcd, 0xfc, 0x78, 0x9e, 0xbc, 0xed, 0x4c, 0x79, 0xd0, 0xf5, 0x6f,
	0xe7, 0x5e, 0x10, 0xcc, 0x29, 0xcd, 0x3e, 0x24, 0x2f, 0xd0, 0xb7, 0x17, 0x77, 0xb2, 0xb3, 0x0b,
	0x21, 0x5b, 0x4b, 0xbf, 0x5a, 0xa7, 0x9e, 0x60, 0xa2, 0x2b, 0x3f, 0x2d, 0xdf, 0xee, 0xe2, 0x47,
	0xe5, 0xf3, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xea, 0xf4, 0x82, 0x00, 0x0c, 0x0b, 0x00, 0x00,
}
