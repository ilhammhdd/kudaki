package usecases

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-review-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type RetrieveItemReviews struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rir *RetrieveItemReviews) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rir.initInOutEvent(in)

	item := rir.retrieveItem(inEvent)
	if item == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"item with the given uuid not found"}

		return outEvent
	}
	itemReviewTemps := rir.retrieveItemReviews(inEvent, item)

	outEvent.Result = rir.ResultSchemer.SetResultSources(itemReviewTemps).ParseToResult()
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (rir *RetrieveItemReviews) initInOutEvent(in proto.Message) (inEvent *events.RetrieveItemReviews, outEvent *events.ItemReviewsRetrieved) {
	inEvent = in.(*events.RetrieveItemReviews)

	outEvent = new(events.ItemReviewsRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.RetrieveItemReviews = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (rir *RetrieveItemReviews) retrieveItem(inEvent *events.RetrieveItemReviews) *store.Item {
	row, err := rir.DBO.QueryRow("SELECT id,uuid,storefront_uuid,name,amount,unit,price,price_duration,description,photo,rating,length,width,height,color,unit_of_measurement,created_at FROM kudaki_store.items WHERE uuid=?", inEvent.ItemUuid)
	errorkit.ErrorHandled(err)

	var item store.Item
	item.Storefront = new(store.Storefront)
	var priceDuration string
	item.ItemDimension = new(store.ItemDimension)
	var createdAt int64
	var unitOfMeasurement string

	err = row.Scan(&item.Id, &item.Uuid, &item.Storefront.Uuid, &item.Name, &item.Amount, &item.Unit, &item.Price, &priceDuration, &item.Description, &item.Photo, &item.Rating, &item.ItemDimension.Length, &item.ItemDimension.Width, &item.ItemDimension.Height, &item.Color, &unitOfMeasurement, &createdAt)
	errorkit.ErrorHandled(err)

	item.ItemDimension.UnitOfMeasurement = store.UnitofMeasurement(store.UnitofMeasurement_value[unitOfMeasurement])

	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	item.CreatedAt = createdAtProto

	item.PriceDuration = store.PriceDuration(store.PriceDuration_value[priceDuration])

	return &item
}

func (rir *RetrieveItemReviews) retrieveItemReviews(inEvent *events.RetrieveItemReviews, item *store.Item) []*ItemReviewTemp {
	rows, err := rir.DBO.Query("SELECT ir.id,ir.uuid,ir.review,ir.rating,ir.created_at,u.uuid AS reviewer_uuid,u.email AS reviewer_email,p.full_name AS reviewer_full_name FROM (SELECT id FROM kudaki_store.item_reviews LIMIT ?,?) ir_ids JOIN kudaki_store.item_reviews ir ON ir.id=ir_ids.id JOIN kudaki_user.users u ON ir.user_uuid=u.uuid JOIN kudaki_user.profiles p ON u.uuid=p.user_uuid WHERE ir.item_uuid=?;", inEvent.Offset, inEvent.Limit, inEvent.ItemUuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var itemReviewTemps []*ItemReviewTemp

	for rows.Next() {
		var itemReviewTemp ItemReviewTemp

		err = rows.Scan(
			&itemReviewTemp.Id,
			&itemReviewTemp.Uuid,
			&itemReviewTemp.Review,
			&itemReviewTemp.RatingT,
			&itemReviewTemp.CreatedAtT,
			&itemReviewTemp.ReviewerUuid,
			&itemReviewTemp.ReviewerEmail,
			&itemReviewTemp.ReviewerFullName)
		errorkit.ErrorHandled(err)

		itemReviewTemps = append(itemReviewTemps, &itemReviewTemp)
	}

	return itemReviewTemps
}

type ItemReviewTemp struct {
	store.ItemReview
	RatingT          float64 `json:"rating"`
	CreatedAtT       int64   `json:"created_at"`
	ReviewerUuid     string  `json:"reviewer_uuid"`
	ReviewerEmail    string  `json:"reviewer_email"`
	ReviewerFullName string  `json:"reviewer_full_name"`
	ReviewerPhoto    string  `json:"reviewer_photo,omitempty"`
}
