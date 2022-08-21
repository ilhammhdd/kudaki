package adapters

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ilhammhdd/kudaki-gateway-service/entities/aggregates/store"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases/events"

	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

// ----------------------------------------------

type AddStorefrontItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (asi *AddStorefrontItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.AddStorefrontItem)

	amount, err := strconv.ParseInt(r.MultipartForm.Value["amount"][0], 10, 32)
	errorkit.ErrorHandled(err)
	price, err := strconv.ParseInt(r.MultipartForm.Value["price"][0], 10, 32)
	errorkit.ErrorHandled(err)
	length, err := strconv.ParseInt(r.MultipartForm.Value["length"][0], 10, 32)
	errorkit.ErrorHandled(err)
	width, err := strconv.ParseInt(r.MultipartForm.Value["width"][0], 10, 32)
	errorkit.ErrorHandled(err)
	height, err := strconv.ParseInt(r.MultipartForm.Value["height"][0], 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.Amount = int32(amount)
	outEvent.Color = r.MultipartForm.Value["color"][0]
	outEvent.Description = r.MultipartForm.Value["description"][0]
	outEvent.Height = int32(height)
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Length = int32(length)
	outEvent.Name = r.MultipartForm.Value["name"][0]
	outEvent.Photo = r.MultipartForm.Value["photo"][0]
	outEvent.Price = int32(price)
	outEvent.PriceDuration = store.PriceDuration(store.PriceDuration_value[r.MultipartForm.Value["price_duration"][0]])
	outEvent.Uid = uuid.New().String()
	outEvent.Unit = r.MultipartForm.Value["unit"][0]
	outEvent.UnitOfMeasurement = store.UnitofMeasurement(store.UnitofMeasurement_value[r.MultipartForm.Value["price_duration"][0]])
	outEvent.Width = int32(width)

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (asi *AddStorefrontItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.StorefrontItemAdded)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (asi *AddStorefrontItem) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.StorefrontItemAdded

	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (asi *AddStorefrontItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       asi.Consumer,
		InTopic:        events.StorefrontServiceEventTopic_STOREFRONT_ITEM_ADDED.String(),
		OutTopic:       events.StorefrontServiceCommandTopic_ADD_STOREFRONT_ITEM.String(),
		Producer:       asi.Producer,
		InEventChecker: asi}
}

type DeleteStorefrontItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (dsi *DeleteStorefrontItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := /* new(events.DeleteStorefrontItemRequested) */ new(events.DeleteStorefrontItem)

	outEvent.ItemUuid = r.URL.Query().Get("item_uuid")
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (dsi *DeleteStorefrontItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.StorefrontItemDeleted)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (dsi *DeleteStorefrontItem) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.StorefrontItemDeleted

	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (dsi *DeleteStorefrontItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer: dsi.Consumer,
		InTopic:/* events.StoreTopic_STOREFRONT_ITEM_DELETED.String() */ events.StorefrontServiceEventTopic_STOREFRONT_ITEM_DELETED.String(),
		OutTopic:/* events.StoreTopic_DELETE_STOREFRONT_ITEM_REQUESTED.String() */ events.StorefrontServiceCommandTopic_DELETE_STOREFRONT_ITEM.String(),
		Producer:       dsi.Producer,
		InEventChecker: dsi}
}

type UpdateStorefrontItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (usi *UpdateStorefrontItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.UpdateStorefrontItem)

	amount, err := strconv.ParseInt(r.MultipartForm.Value["amount"][0], 10, 32)
	errorkit.ErrorHandled(err)
	price, err := strconv.ParseInt(r.MultipartForm.Value["price"][0], 10, 32)
	errorkit.ErrorHandled(err)
	length, err := strconv.ParseInt(r.MultipartForm.Value["length"][0], 10, 32)
	errorkit.ErrorHandled(err)
	width, err := strconv.ParseInt(r.MultipartForm.Value["width"][0], 10, 32)
	errorkit.ErrorHandled(err)
	height, err := strconv.ParseInt(r.MultipartForm.Value["height"][0], 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.Amount = int32(amount)
	outEvent.Description = r.MultipartForm.Value["description"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Name = r.MultipartForm.Value["name"][0]
	outEvent.Photo = r.MultipartForm.Value["photo"][0]
	outEvent.Price = int32(price)
	outEvent.Uid = uuid.New().String()
	outEvent.Unit = r.MultipartForm.Value["unit"][0]
	outEvent.ItemUuid = r.MultipartForm.Value["item_uuid"][0]

	outEvent.PriceDuration = store.PriceDuration(store.PriceDuration_value[r.MultipartForm.Value["price_duration"][0]])
	outEvent.Length = int32(length)
	outEvent.Width = int32(width)
	outEvent.Height = int32(height)
	outEvent.UnitOfMeasurement = store.UnitofMeasurement(store.UnitofMeasurement_value[r.MultipartForm.Value["unit_of_measurement"][0]])
	outEvent.Color = r.MultipartForm.Value["color"][0]

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (usi *UpdateStorefrontItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.StorefrontItemsUpdated)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (usi *UpdateStorefrontItem) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.StorefrontItemsUpdated

	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (usi *UpdateStorefrontItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       usi.Consumer,
		InTopic:        events.StorefrontServiceEventTopic_STOREFRONT_ITEMS_UPDATED.String(),
		OutTopic:       events.StorefrontServiceCommandTopic_UPDATE_STOREFRONT_ITEM.String(),
		Producer:       usi.Producer,
		InEventChecker: usi}
}

type RetrieveStorefrontItems struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rsfi *RetrieveStorefrontItems) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveStorefrontItems)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	if storefrontUUID := r.URL.Query().Get("storefront_uuid"); storefrontUUID != "" {
		outEvent.StorefrontUuid = storefrontUUID
	}
	outEvent.Uid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (rsfi *RetrieveStorefrontItems) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.StorefrontItemsRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}
	resBody.Data = json.RawMessage(inEvent.Result)

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (rsfi *RetrieveStorefrontItems) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.StorefrontItemsRetrieved
	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (rsfi *RetrieveStorefrontItems) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rsfi.Consumer,
		InTopic:        events.StorefrontServiceEventTopic_STOREFRONT_ITEMS_RETRIEVED.String(),
		InEventChecker: rsfi,
		OutTopic:/* events.StoreTopic_RETRIEVE_STOREFRONT_ITEMS_REQUESTED.String() */ events.StorefrontServiceCommandTopic_RETRIEVE_STOREFRONT_ITEMS.String(),
		Producer: rsfi.Producer}
}

// ----------------------------------------------

type RetrieveItems struct {
	Producer usecases.EventDrivenProducer
	Consumer usecases.EventDrivenConsumer
}

func (ri *RetrieveItems) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent := /* new(events.RetrieveItemsRequested) */ new(events.RetrieveItems)
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (ri *RetrieveItems) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.ItemsRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}
	resBody.Data = json.RawMessage(inEvent.Result)

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (ri *RetrieveItems) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.ItemsRetrieved
	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (ri *RetrieveItems) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       ri.Consumer,
		InEventChecker: ri,
		InTopic:        events.ItemServiceEventTopic_ITEMS_RETRIEVED.String(),
		OutTopic:/* events.StoreTopic_RETRIEVE_ITEMS_REQUESTED.String() */ events.ItemServiceCommandTopic_RETRIEVE_ITEMS.String(),
		Producer: ri.Producer}
}

type SearchItems struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (si *SearchItems) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := /* new(events.SearchItemsRequested) */ new(events.SearchItems)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.Keyword = r.URL.Query().Get("keyword")
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uid = uuid.New().String()

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (si *SearchItems) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.ItemsSearched
	if err := proto.Unmarshal(inVal, &inEvent); err == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

func (si *SearchItems) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.ItemsSearched)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}
	resBody.Data = json.RawMessage(inEvent.Result)

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (si *SearchItems) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       si.Consumer,
		InEventChecker: si,
		InTopic:        events.ItemServiceEventTopic_ITEMS_SEARCHED.String(),
		OutTopic:/* events.StoreTopic_SEARCH_ITEMS_REQUESTED.String() */ events.ItemServiceCommandTopic_SEARCH_ITEMS.String(),
		Producer: si.Producer}
}

type RetrieveItemReviews struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rir *RetrieveItemReviews) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveItemReviews)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.ItemUuid = r.URL.Query().Get("item_uuid")
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rir *RetrieveItemReviews) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.ItemReviewsRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	} else {
		resBody.Data = json.RawMessage(inEvent.Result)
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (rir *RetrieveItemReviews) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rir.Consumer,
		InEventChecker: rir,
		InTopic:        events.ItemReviewServiceEventTopic_ITEM_REVIEWS_RETRIEVED.String(),
		OutTopic:       events.ItemReviewServiceCommandTopic_RETRIEVE_ITEM_REVIEWS.String(),
		Producer:       rir.Producer}
}

func (rir *RetrieveItemReviews) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.ItemReviewsRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type CommentItemReview struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (cir *CommentItemReview) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.CommentItemReview)

	outEvent.Comment = r.MultipartForm.Value["comment"][0]
	outEvent.ItemReviewUuid = r.MultipartForm.Value["item_review_uuid"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (cir *CommentItemReview) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.ItemReviewCommented)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (cir *CommentItemReview) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       cir.Consumer,
		InEventChecker: cir,
		InTopic:        events.ItemReviewServiceEventTopic_ITEM_REVIEW_COMMENTED.String(),
		OutTopic:       events.ItemReviewServiceCommandTopic_COMMENT_ITEM_REVIEW.String(),
		Producer:       cir.Producer}
}

func (cir *CommentItemReview) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.ItemReviewCommented
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type RetrieveItemReviewComments struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rirc *RetrieveItemReviewComments) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveItemReviewComments)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.ItemReviewUuid = r.URL.Query().Get("item_review_uuid")
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rirc *RetrieveItemReviewComments) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.ItemReviewCommentsRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody.Errs = &inEvent.EventStatus.Errors
	} else {
		resBody.Data = json.RawMessage(inEvent.Result)
	}

	return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
}

func (rirc *RetrieveItemReviewComments) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rirc.Consumer,
		InEventChecker: rirc,
		InTopic:        events.ItemReviewServiceEventTopic_ITEM_REVIEW_COMMENTS_RETRIEVED.String(),
		OutTopic:       events.ItemReviewServiceCommandTopic_RETRIEVE_ITEM_REVIEW_COMMENTS.String(),
		Producer:       rirc.Producer}
}

func (rirc *RetrieveItemReviewComments) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.ItemReviewCommentsRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}
