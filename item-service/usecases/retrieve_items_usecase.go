package usecases

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-item-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-item-service/usecases/events"
)

type RetrieveItems struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (ri *RetrieveItems) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := ri.initInOutEvent(in)

	items := ri.retrieveItems(outEvent.Requester, inEvent)

	outEvent.Result = ri.ResultSchemer.SetResultSources(items).ParseToResult()
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (ri *RetrieveItems) initInOutEvent(in proto.Message) (inEvent *events.RetrieveItems, outEvent *events.ItemsRetrieved) {
	inEvent = in.(*events.RetrieveItems)

	outEvent = new(events.ItemsRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.Limit = inEvent.Limit
	outEvent.Offset = inEvent.Offset
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.RetrieveItems = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (ri *RetrieveItems) retrieveItems(usr *user.User, inEvent *events.RetrieveItems) []*ItemTemp {
	rows, err := ri.DBO.Query("SELECT i.id,i.uuid,i.storefront_uuid,i.name,i.amount,i.unit,i.price,i.price_duration,i.description,i.photo,i.rating,i.length,i.width,i.height,i.color,i.unit_of_measurement,i.created_at FROM (SELECT i_i.id FROM kudaki_store.items i_i LIMIT ?, ?) i_ids JOIN kudaki_store.items i ON i.id = i_ids.id;", inEvent.Offset, inEvent.Limit)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var items []*ItemTemp
	for rows.Next() {
		var item ItemTemp

		rows.Scan(&item.Id, &item.Uuid, &item.StorefrontUuidT, &item.Name, &item.Amount, &item.Unit, &item.Price, &item.PriceDurationT, &item.Description, &item.Photo, &item.RatingT, &item.LengthT, &item.WidthT, &item.HeightT, &item.Color, &item.UnitofMeasurementT, &item.CreatedAtT)

		items = append(items, &item)
	}
	return items
}

type ItemTemp struct {
	store.Item
	StorefrontUuidT    string  `json:"storefront_uuid"`
	LengthT            int32   `json:"length"`
	WidthT             int32   `json:"width"`
	HeightT            int32   `json:"height"`
	PriceDurationT     string  `json:"price_duration"`
	UnitofMeasurementT string  `json:"unit_of_measurement"`
	RatingT            float64 `json:"rating"`
	CreatedAtT         int64   `json:"created_at"`
}
