package adapters

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases/events"
)

type AddKudakiEvent struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ake *AddKudakiEvent) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.AddKudakiEvent)

	durationFrom, err := strconv.ParseInt(r.MultipartForm.Value["duration_from"][0], 10, 64)
	errorkit.ErrorHandled(err)
	durationTo, err := strconv.ParseInt(r.MultipartForm.Value["duration_to"][0], 10, 64)
	errorkit.ErrorHandled(err)
	adDurationFrom, err := strconv.ParseInt(r.MultipartForm.Value["ad_duration_from"][0], 10, 64)
	errorkit.ErrorHandled(err)
	adDurationTo, err := strconv.ParseInt(r.MultipartForm.Value["ad_duration_to"][0], 10, 64)
	errorkit.ErrorHandled(err)

	outEvent.AdDurationFrom = adDurationFrom
	outEvent.AdDurationTo = adDurationTo
	outEvent.Description = r.MultipartForm.Value["description"][0]
	outEvent.DurationFrom = durationFrom
	outEvent.DurationTo = durationTo
	outEvent.FilePath = r.MultipartForm.Value["file_path"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Name = r.MultipartForm.Value["name"][0]
	outEvent.Uid = uuid.New().String()
	outEvent.Venue = r.MultipartForm.Value["venue"][0]

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ake *AddKudakiEvent) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.KudakiEventDokuInvoiceIssued)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	var responseData struct {
		Amount           float32 `json:"AMOUNT"`
		PurchaseAmount   float32 `json:"PURCHASEAMOUNT"`
		TransIDMerchant  string  `json:"TRANSIDMERCHANT"`
		Words            string  `json:"WORDS"`
		RequestDateTime  int64   `json:"REQUESTDATETIME"`
		Currency         int32   `json:"CURRENCY"`
		PurchaseCurrency int32   `json:"PURCHASECURRENCY"`
		SessionID        string  `json:"SESSIONID"`
		Name             string  `json:"NAME"`
		Email            string  `json:"EMAIL"`
		Basket           string  `json:"BASKET"`
	}

	responseData.Amount = inEvent.DokuInvoice.Amount
	responseData.Basket = inEvent.DokuInvoice.Basket
	responseData.Currency = inEvent.DokuInvoice.Currency
	responseData.Email = inEvent.DokuInvoice.Email
	responseData.Name = inEvent.DokuInvoice.Name
	responseData.PurchaseAmount = inEvent.DokuInvoice.PurchaseAmount
	responseData.PurchaseCurrency = inEvent.DokuInvoice.PurchaseCurrency
	responseData.RequestDateTime = inEvent.DokuInvoice.RequestDateTime
	responseData.SessionID = inEvent.DokuInvoice.SessionId
	responseData.TransIDMerchant = inEvent.DokuInvoice.TransactionIdMerchant
	responseData.Words = inEvent.DokuInvoice.Words

	resBody.Data = responseData

	return NewResponse(http.StatusOK, &resBody)
}

func (ake *AddKudakiEvent) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ake.Consumer,
		InEventChecker: ake,
		InTopic:        events.EventPaymentServiceEventTopic_EVENT_DOKU_INVOICE_ISSUED.String(),
		OutTopic:       events.EventServiceCommandTopic_ADD_KUDAKI_EVENT.String(),
		Producer:       ake.Producer}
}

func (ake *AddKudakiEvent) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.KudakiEventDokuInvoiceIssued
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// -------------------------------------------------------------------------------------------

type DeleteKudakiEvent struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ake *DeleteKudakiEvent) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.DeleteKudakiEvent)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.EventUuid = r.MultipartForm.Value["event_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ake *DeleteKudakiEvent) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.KudakiEventDeleted)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ake *DeleteKudakiEvent) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ake.Consumer,
		InEventChecker: ake,
		InTopic:        events.EventServiceEventTopic_KUDAKI_EVENT_DELETED.String(),
		OutTopic:       events.EventServiceCommandTopic_DELETE_KUDAKI_EVENT.String(),
		Producer:       ake.Producer}
}

func (ake *DeleteKudakiEvent) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.KudakiEventDeleted
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// -------------------------------------------------------------------------------------------

type RetrieveOrganizerInvoices struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ake *RetrieveOrganizerInvoices) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveOrganizerInvoices)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ake *RetrieveOrganizerInvoices) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.OrganizerInvoicesRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	} else {
		resBody.Data = json.RawMessage(inEvent.Result)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ake *RetrieveOrganizerInvoices) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ake.Consumer,
		InEventChecker: ake,
		InTopic:        events.EventServiceEventTopic_ORGANIZER_INVOICES_RETRIEVED.String(),
		OutTopic:       events.EventServiceCommandTopic_RETRIEVE_ORGANIZER_INVOICES.String(),
		Producer:       ake.Producer}
}

func (ake *RetrieveOrganizerInvoices) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.OrganizerInvoicesRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// -------------------------------------------------------------------------------------------

type PaymentRequest struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ake *PaymentRequest) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.PaymentRequestDoku)

	outEvent.HashedWords = r.MultipartForm.Value["hashed_words"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.SessionId = r.MultipartForm.Value["session_id"][0]
	outEvent.TransactionIdMerchant = r.MultipartForm.Value["transaction_id_merchant"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ake *PaymentRequest) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.PaymentRequestedDoku)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (ake *PaymentRequest) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ake.Consumer,
		InEventChecker: ake,
		InTopic:        events.EventPaymentServiceEventTopic_PAYMENT_REQUESTED_DOKU.String(),
		OutTopic:       events.EventPaymentServiceCommandTopic_PAYMENT_REQUEST_DOKU.String(),
		Producer:       ake.Producer}
}

func (ake *PaymentRequest) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.PaymentRequestedDoku
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// -------------------------------------------------------------------------------------------

type RetrieveKudakiEvent struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (ake *RetrieveKudakiEvent) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveKudakiEvent)

	outEvent.KudakiEventUuid = r.URL.Query().Get("kudaki_event_uuid")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (ake *RetrieveKudakiEvent) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.KudakiEventRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	resBody = ResponseBody{
		Data: map[string]interface{}{
			"uuid":             inEvent.KudakiEvent.Uuid,
			"name":             inEvent.KudakiEvent.Name,
			"venue":            inEvent.KudakiEvent.Venue,
			"description":      inEvent.KudakiEvent.Description,
			"duration_from":    inEvent.KudakiEvent.DurationFrom,
			"duration_to":      inEvent.KudakiEvent.DurationTo,
			"ad_duration_from": inEvent.KudakiEvent.AdDurationFrom,
			"ad_duration_to":   inEvent.KudakiEvent.AdDurationTo,
			"status":           inEvent.KudakiEvent.Status.String(),
			"poster":           inEvent.KudakiEvent.FilePath,
			"payment_status":   inEvent.PaymentStatus}}

	return NewResponse(http.StatusOK, &resBody)
}

func (ake *RetrieveKudakiEvent) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ake.Consumer,
		InEventChecker: ake,
		InTopic:        events.EventServiceEventTopic_KUDAKI_EVENT_RETRIEVED.String(),
		OutTopic:       events.EventServiceCommandTopic_RETRIEVE_KUDAKI_EVENT.String(),
		Producer:       ake.Producer}
}

func (ake *RetrieveKudakiEvent) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.KudakiEventRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// -------------------------------------------------------------------------------------------

type PublishKudakiEvent struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (pke *PublishKudakiEvent) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.PublishKudakiEvent)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.KudakiEventUuid = r.MultipartForm.Value["kudaki_event_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (pke *PublishKudakiEvent) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.KudakiEventPublished)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (pke *PublishKudakiEvent) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       pke.Consumer,
		InEventChecker: pke,
		InTopic:        events.EventServiceEventTopic_KUDAKI_EVENT_PUBLISHED.String(),
		OutTopic:       events.EventServiceCommandTopic_PUBLISH_KUDAKI_EVENT.String(),
		Producer:       pke.Producer}
}

func (pke *PublishKudakiEvent) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.KudakiEventPublished
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}
