package usecases

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/golang/protobuf/ptypes"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-recommendation-service/entities/aggregates/mountain"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type AddRecommendationGear struct {
	DBO DBOperator
}

func (arg *AddRecommendationGear) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := arg.initInOutEvent(in)
	log.Println(inEvent, outEvent)

	mt := arg.retrieveMountain(inEvent.MountainUuid)
	if mt == nil {
		outEvent.EventStatus.Errors = []string{"mountain with the given uuid not found"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	recommendationGear := arg.initRecommendedGear(mt, inEvent, outEvent)
	recommendationGearItems := arg.initRecommendedGearItems(inEvent, recommendationGear)

	outEvent.RecommendedGear = recommendationGear
	outEvent.RecommendedGearItems = recommendationGearItems

	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (arg *AddRecommendationGear) initInOutEvent(in proto.Message) (inEvent *events.AddRecommendedGear, outEvent *events.RecommendedGearAdded) {
	inEvent = in.(*events.AddRecommendedGear)

	outEvent = new(events.RecommendedGearAdded)
	outEvent.AddRecommendedGear = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (arg *AddRecommendationGear) retrieveMountain(mountainUuid string) *mountain.Mountain {
	row, err := arg.DBO.QueryRow("SELECT id,uuid,name,height,latitude,longitude,difficulty,description,created_at FROM kudaki_mountain.mountains WHERE uuid=?;",
		mountainUuid)
	errorkit.ErrorHandled(err)

	var mountain mountain.Mountain
	var createdAt int64
	if row.Scan(
		&mountain.Id,
		&mountain.Uuid,
		&mountain.Name,
		&mountain.Height,
		&mountain.Latitude,
		&mountain.Longitude,
		&mountain.Difficulty,
		&mountain.Description,
		&createdAt) == sql.ErrNoRows {
		return nil
	}

	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)

	mountain.CreatedAt = createdAtProto

	return &mountain
}

func (arg *AddRecommendationGear) initRecommendedGear(mt *mountain.Mountain, inEvent *events.AddRecommendedGear, outEvent *events.RecommendedGearAdded) *mountain.RecommendedGear {
	var recommendedGear mountain.RecommendedGear
	recommendedGear.CreatedAt = ptypes.TimestampNow()
	recommendedGear.Downvote = 0
	recommendedGear.Mountain = mt
	recommendedGear.Seen = 0
	recommendedGear.Upvote = 0
	recommendedGear.UserUuid = outEvent.Requester.Uuid
	recommendedGear.Uuid = uuid.New().String()

	return &recommendedGear
}

func (arg *AddRecommendationGear) initRecommendedGearItems(inEvent *events.AddRecommendedGear, recommendedGear *mountain.RecommendedGear) []*mountain.RecommendedGearItem {
	var recommendedGearItems []*mountain.RecommendedGearItem
	for i := 0; i < len(inEvent.RecommendedGearItems); i++ {
		recommendedGearItem := mountain.RecommendedGearItem{
			CreatedAt:       ptypes.TimestampNow(),
			ItemType:        inEvent.RecommendedGearItems[i].ItemType,
			RecommendedGear: recommendedGear,
			Total:           inEvent.RecommendedGearItems[i].Total,
			Uuid:            uuid.New().String()}
		recommendedGearItems = append(recommendedGearItems, &recommendedGearItem)
	}

	return recommendedGearItems
}
