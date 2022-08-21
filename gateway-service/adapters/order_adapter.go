package adapters

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ilhammhdd/kudaki-gateway-service/entities/aggregates/order"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases/events"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

type RetrieveOwnerOrderHistories struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rooh *RetrieveOwnerOrderHistories) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveOwnersOrderHistories)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.OrderStatus = order.OrderStatus(order.OrderStatus_value[r.URL.Query().Get("order_status")])
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rooh *RetrieveOwnerOrderHistories) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OwnersOrderHistoriesRetrieved)
	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}
	resBody = ResponseBody{Data: json.RawMessage(inEvent.Result)}

	return NewResponse(http.StatusOK, &resBody)
}

func (rooh *RetrieveOwnerOrderHistories) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rooh.Consumer,
		InEventChecker: rooh,
		InTopic:        events.OrderServiceEventTopic_OWNERS_ORDER_HISTORIES_RETRIEVED.String(),
		OutTopic:       events.OrderServiceCommandTopic_RETRIEVE_OWNERS_ORDER_HISTORIES.String(),
		Producer:       rooh.Producer}
}

func (rooh *RetrieveOwnerOrderHistories) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OwnersOrderHistoriesRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type RetrieveTenantOrderHistories struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rtoh *RetrieveTenantOrderHistories) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveTenantsOrderHistories)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.OrderStatus = order.OrderStatus(order.OrderStatus_value[r.URL.Query().Get("order_status")])
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rtoh *RetrieveTenantOrderHistories) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.TenantOrderHistoriesRetrieved)
	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}
	resBody.Data = json.RawMessage(inEvent.Result)

	return NewResponse(http.StatusOK, &resBody)
}

func (rtoh *RetrieveTenantOrderHistories) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rtoh.Consumer,
		InEventChecker: rtoh,
		InTopic:        events.OrderServiceEventTopic_TENANTS_ORDER_HISTORIES_RETRIEVED.String(),
		OutTopic:       events.OrderServiceCommandTopic_RETRIEVE_TENANTS_ORDER_HISTORIES.String(),
		Producer:       rtoh.Producer}
}

func (rtoh *RetrieveTenantOrderHistories) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.TenantOrderHistoriesRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type TenantReviewOwnerOrder struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (tro *TenantReviewOwnerOrder) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.TenantReviewOwnerOrder)

	rating, err := strconv.ParseFloat(r.MultipartForm.Value["rating"][0], 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OwnerOrderUuid = r.MultipartForm.Value["owner_order_uuid"][0]
	outEvent.Rating = rating
	outEvent.Review = r.MultipartForm.Value["review"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (tro *TenantReviewOwnerOrder) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.TenantReviewedOwnerOrder)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (tro *TenantReviewOwnerOrder) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       tro.Consumer,
		InEventChecker: tro,
		InTopic:        events.OrderServiceEventTopic_TENANT_REVIEWED_OWNER_ORDER.String(),
		OutTopic:       events.OrderServiceCommandTopic_TENANT_REVIEW_OWNER_ORDER.String(),
		Producer:       tro.Producer}
}

func (tro *TenantReviewOwnerOrder) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.TenantReviewedOwnerOrder
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type ApproveOwnerOrder struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ao *ApproveOwnerOrder) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.ApproveOwnerOrder)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OwnerOrderUuid = r.MultipartForm.Value["owner_order_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ao *ApproveOwnerOrder) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OwnerOrderApproved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ao *ApproveOwnerOrder) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ao.Consumer,
		InEventChecker: ao,
		InTopic:        events.OrderServiceEventTopic_OWNER_ORDER_APPROVED.String(),
		OutTopic:       events.OrderServiceCommandTopic_APPROVE_OWNER_ORDER.String(),
		Producer:       ao.Producer}
}

func (ao *ApproveOwnerOrder) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OwnerOrderApproved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type DisapproveOwnerOrder struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (do *DisapproveOwnerOrder) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.DisapproveOwnerOrder)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OwnerOrderUuid = r.MultipartForm.Value["owner_order_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (do *DisapproveOwnerOrder) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OrderDisapproved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (do *DisapproveOwnerOrder) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       do.Consumer,
		InEventChecker: do,
		InTopic:        events.OrderServiceEventTopic_OWNER_ORDER_DISAPPROVED.String(),
		OutTopic:       events.OrderServiceCommandTopic_DISAPPROVE_OWNER_ORDER.String(),
		Producer:       do.Producer}
}

func (do *DisapproveOwnerOrder) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OrderDisapproved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type CheckOut struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (co *CheckOut) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.CheckOut)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.CartUuid = r.MultipartForm.Value["cart_uuid"][0]
	// outEvent.AddressUuid = r.MultipartForm.Value["address_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (co *CheckOut) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.CheckedOut)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (co *CheckOut) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       co.Consumer,
		InEventChecker: co,
		InTopic:        events.OrderServiceEventTopic_CHECKED_OUT.String(),
		OutTopic:       events.OrderServiceCommandTopic_CHECK_OUT.String(),
		Producer:       co.Producer}
}

func (co *CheckOut) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.CheckedOut
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type OwnerConfirmReturnment struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ocr *OwnerConfirmReturnment) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.OwnerConfirmReturnment)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OwnerOrderUuid = r.MultipartForm.Value["owner_order_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ocr *OwnerConfirmReturnment) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OwnerConfirmedReturnment)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ocr *OwnerConfirmReturnment) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ocr.Consumer,
		InEventChecker: ocr,
		InTopic:        events.OrderServiceEventTopic_OWNER_CONFIRMED_RETURNMENT.String(),
		OutTopic:       events.OrderServiceCommandTopic_OWNER_CONFIRM_RETURNMENT.String(),
		Producer:       ocr.Producer}
}

func (ocr *OwnerConfirmReturnment) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OwnerConfirmedReturnment
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// -------------------------------------------------------------------------

type OwnerOrderRented struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (oor *OwnerOrderRented) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.OwnerOrderRented)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.OwnerOrderUuid = r.MultipartForm.Value["owner_order_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (oor *OwnerOrderRented) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OwnerOrderRentedOut)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (oor *OwnerOrderRented) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       oor.Consumer,
		InEventChecker: oor,
		InTopic:        events.OrderServiceEventTopic_OWNER_ORDER_RENTED_OUT.String(),
		OutTopic:       events.OrderServiceCommandTopic_OWNER_ORDER_RENTED.String(),
		Producer:       oor.Producer}
}

func (oor *OwnerOrderRented) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OwnerOrderRentedOut
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}
