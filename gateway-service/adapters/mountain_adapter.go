package adapters

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases/events"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-gateway-service/usecases"
)

type DeleteRecommendedGearItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (drgi *DeleteRecommendedGearItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.DeleteRecommendedGearItem)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.RecommendedGearItemUuid = r.URL.Query().Get("recommended_gear_item_uuid")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (drgi *DeleteRecommendedGearItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.RecommendedGearUpdated)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (drgi *DeleteRecommendedGearItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       drgi.Consumer,
		InEventChecker: drgi,
		InTopic:        events.RecommendationServiceEventTopic_RECOMMENDED_GEAR_UPDATED.String(),
		OutTopic:       events.RecommendationServiceCommandTopic_DELETE_RECOMMENDED_GEAR_ITEM.String(),
		Producer:       drgi.Producer}
}

func (drgi *DeleteRecommendedGearItem) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.RecommendedGearUpdated
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type AddRecommendedGear struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

type AddRecommendedGearReqData struct {
	MountainUUID string                             `json:"mountain_uuid"`
	Items        []*events.DataRecommendedGearItems `json:"items"`
}

func (arg *AddRecommendedGear) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	var reqBody AddRecommendedGearReqData

	bodyJSONDecoder := json.NewDecoder(r.Body)
	err := bodyJSONDecoder.Decode(&reqBody)
	errorkit.ErrorHandled(err)

	outEvent := new(events.AddRecommendedGear)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.MountainUuid = reqBody.MountainUUID
	outEvent.RecommendedGearItems = reqBody.Items
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (arg *AddRecommendedGear) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.RecommendedGearAdded)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (arg *AddRecommendedGear) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       arg.Consumer,
		InEventChecker: arg,
		InTopic:        events.RecommendationServiceEventTopic_RECOMMENDED_GEAR_ADDED.String(),
		OutTopic:       events.RecommendationServiceCommandTopic_ADD_RECOMMENDED_GEAR.String(),
		Producer:       arg.Producer}
}

func (arg *AddRecommendedGear) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.RecommendedGearAdded
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type DeleteRecommendedGear struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (drg *DeleteRecommendedGear) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.DeleteRecommendedGear)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.RecommendedGearUuid = r.MultipartForm.Value["recommended_gear_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (drg *DeleteRecommendedGear) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.RecommendedGearDeleted)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (drg *DeleteRecommendedGear) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       drg.Consumer,
		InEventChecker: drg,
		InTopic:        events.RecommendationServiceEventTopic_RECOMMENDED_GEAR_DELETED.String(),
		OutTopic:       events.RecommendationServiceCommandTopic_DELETE_RECOMMENDED_GEAR.String(),
		Producer:       drg.Producer}
}

func (drg *DeleteRecommendedGear) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.RecommendedGearDeleted
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type AddRecommendedGearItem struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (argi *AddRecommendedGearItem) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.AddRecommendedGearItem)

	total, err := strconv.ParseInt(r.MultipartForm.Value["total"][0], 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.ItemType = r.MultipartForm.Value["item_type"][0]
	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.RecommendedGearUuid = r.MultipartForm.Value["recommended_gear_uuid"][0]
	outEvent.Total = int32(total)
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (argi *AddRecommendedGearItem) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.RecommendedGearUpdated)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (argi *AddRecommendedGearItem) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       argi.Consumer,
		InEventChecker: argi,
		InTopic:        events.RecommendationServiceEventTopic_RECOMMENDED_GEAR_UPDATED.String(),
		OutTopic:       events.RecommendationServiceCommandTopic_ADD_RECOMMENDED_GEAR_ITEM.String(),
		Producer:       argi.Producer}
}

func (argi *AddRecommendedGearItem) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.RecommendedGearUpdated
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type RetrieveRecommendedGears struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rrg *RetrieveRecommendedGears) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveRecommendedGears)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.MountainUuid = r.URL.Query().Get("mountain_uuid")
	outEvent.Offset = int32(offset)
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rrg *RetrieveRecommendedGears) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.RecommendedGearsRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	resBody = ResponseBody{Data: json.RawMessage(inEvent.Result)}
	return NewResponse(http.StatusOK, &resBody)
}

func (rrg *RetrieveRecommendedGears) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rrg.Consumer,
		InEventChecker: rrg,
		InTopic:        events.RecommendationServiceEventTopic_RECOMMENDED_GEARS_RETRIEVED.String(),
		OutTopic:       events.RecommendationServiceCommandTopic_RETRIEVE_RECOMMENDED_GEARS.String(),
		Producer:       rrg.Producer}
}

func (rrg *RetrieveRecommendedGears) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.RecommendedGearsRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type RetrieveRecommendedGearItems struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (rrgi *RetrieveRecommendedGearItems) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveRecommendedGearItems)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.RecommendedGearUuid = r.URL.Query().Get("recommended_gear_uuid")
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (rrgi *RetrieveRecommendedGearItems) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.RecommendedGearItemsRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (rrgi *RetrieveRecommendedGearItems) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       rrgi.Consumer,
		InEventChecker: rrgi,
		InTopic:        events.RecommendationServiceEventTopic_RECOMMENDED_GEAR_ITEMS_RETRIEVED.String(),
		OutTopic:       events.RecommendationServiceCommandTopic_RETRIEVE_RECOMMENDED_GEAR_ITEMS.String(),
		Producer:       rrgi.Producer}
}

func (rrgi *RetrieveRecommendedGearItems) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.RecommendedGearItemsRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type UpVoteRecommendedGear struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (uvrg *UpVoteRecommendedGear) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.UpvoteRecommendedGear)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.RecommendedGearUuid = r.MultipartForm.Value["recommended_gear_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (uvrg *UpVoteRecommendedGear) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.RecommendedGearUpdated)
	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (uvrg *UpVoteRecommendedGear) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       uvrg.Consumer,
		InEventChecker: uvrg,
		InTopic:        events.RecommendationServiceEventTopic_RECOMMENDED_GEAR_UPDATED.String(),
		OutTopic:       events.RecommendationServiceCommandTopic_UPVOTE_RECOMMENDED_GEAR.String(),
		Producer:       uvrg.Producer}
}

func (uvrg *UpVoteRecommendedGear) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.RecommendedGearUpdated
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

type DownVoteRecommendedGear struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (dvrg *DownVoteRecommendedGear) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.DownvoteRecommendedGear)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.RecommendedGearUuid = r.MultipartForm.Value["recommended_gear_uuid"][0]
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (dvrg *DownVoteRecommendedGear) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.RecommendedGearUpdated)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (dvrg *DownVoteRecommendedGear) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       dvrg.Consumer,
		InEventChecker: dvrg,
		InTopic:        events.RecommendationServiceEventTopic_RECOMMENDED_GEAR_UPDATED.String(),
		OutTopic:       events.RecommendationServiceCommandTopic_DOWNVOTE_RECOMMENDED_GEAR.String(),
		Producer:       dvrg.Producer}
}

func (dvrg *DownVoteRecommendedGear) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.RecommendedGearUpdated
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// ----------------------------------------------------------------------------------------------

type RetrieveMountains struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (dvrg *RetrieveMountains) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.RetrieveMountains)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (dvrg *RetrieveMountains) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.MountainsRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	} else {
		resBody.Data = json.RawMessage(inEvent.Result)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (dvrg *RetrieveMountains) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       dvrg.Consumer,
		InEventChecker: dvrg,
		InTopic:        events.MountainServiceEventTopic_MOUNTAINS_RETRIEVED.String(),
		OutTopic:       events.MountainServiceCommandTopic_RETRIEVE_MOUNTAINS.String(),
		Producer:       dvrg.Producer}
}

func (dvrg *RetrieveMountains) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.MountainsRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}

// ----------------------------------------------------------------------------------------------

type SearchMountains struct {
	Consumer usecases.EventDrivenConsumer
	Producer usecases.EventDrivenProducer
}

func (dvrg *SearchMountains) ParseRequestToKafkaMessage(r *http.Request) (key string, message []byte) {
	outEvent := new(events.SearchMountains)

	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
	errorkit.ErrorHandled(err)
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 32)
	errorkit.ErrorHandled(err)

	outEvent.KudakiToken = r.Header.Get("Kudaki-Token")
	outEvent.Limit = int32(limit)
	outEvent.Offset = int32(offset)
	outEvent.Uid = uuid.New().String()

	out, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, out
}

func (dvrg *SearchMountains) ParseEventToResponse(in proto.Message) *Response {
	inEvent := in.(*events.MountainsRetrieved)

	var resBody ResponseBody
	if inEvent.EventStatus.HttpCode != http.StatusOK {
		resBody = ResponseBody{Errs: &inEvent.EventStatus.Errors}
		return NewResponse(int(inEvent.EventStatus.HttpCode), &resBody)
	} else {
		resBody.Data = json.RawMessage(inEvent.Result)
	}

	return NewResponse(http.StatusOK, &resBody)
}

func (dvrg *SearchMountains) initUsecaseHandler(outKey string) usecases.EventDrivenHandler {
	return &usecases.EventDrivenUsecase{
		Consumer:       dvrg.Consumer,
		InEventChecker: dvrg,
		InTopic:        events.MountainServiceEventTopic_MOUNTAINS_RETRIEVED.String(),
		OutTopic:       events.MountainServiceCommandTopic_RETRIEVE_MOUNTAINS.String(),
		Producer:       dvrg.Producer}
}

func (dvrg *SearchMountains) CheckInEvent(outKey string, inKey, inVal []byte) (proto.Message, bool) {
	var inEvent events.MountainsRetrieved
	if proto.Unmarshal(inVal, &inEvent) == nil {
		if outKey == string(inKey) {
			return &inEvent, true
		}
	}
	return nil, false
}
