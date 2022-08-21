package usecases

import (
	"database/sql"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-recommendation-service/entities/aggregates/mountain"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type RetrieveRecommendedGearItems struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rrgi *RetrieveRecommendedGearItems) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rrgi.initInOutEvent(in)

	recommendedGear := rrgi.retrieveRecommendedGear(inEvent)
	if recommendedGear == nil {
		outEvent.EventStatus.Errors = []string{"recommended gear with the given uuid not found"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	recommendedGearItems := rrgi.retrieveRecommendedGearItems(inEvent)

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.Result = rrgi.ResultSchemer.SetResultSources(recommendedGear, recommendedGearItems).ParseToResult()

	return outEvent
}

func (rrgi *RetrieveRecommendedGearItems) initInOutEvent(in proto.Message) (inEvent *events.RetrieveRecommendedGearItems, outEvent *events.RecommendedGearItemsRetrieved) {
	inEvent = in.(*events.RetrieveRecommendedGearItems)

	outEvent = new(events.RecommendedGearItemsRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.RetrieveRecommendedGearItems = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (rrgi *RetrieveRecommendedGearItems) retrieveRecommendedGearItems(inEvent *events.RetrieveRecommendedGearItems) []*mountain.RecommendedGearItem {
	rows, err := rrgi.DBO.Query("SELECT id,uuid,item_type,total FROM kudaki_mountain.recommended_gear_items WHERE recommended_gear_uuid = ?;", inEvent.RecommendedGearUuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var rgis []*mountain.RecommendedGearItem
	for rows.Next() {
		var rgi mountain.RecommendedGearItem
		rows.Scan(
			&rgi.Id,
			&rgi.Uuid,
			&rgi.ItemType,
			&rgi.Total)

		rgis = append(rgis, &rgi)
	}

	return rgis
}

func (rrgi *RetrieveRecommendedGearItems) retrieveRecommendedGear(inEvent *events.RetrieveRecommendedGearItems) *RecommendedGearTemp {
	rgRow, err := rrgi.DBO.QueryRow("SELECT rg.id, rg.uuid, rg.upvote, rg.downvote, rg.seen, rg.created_at, p.full_name, u.email, m.name, m.uuid AS mountain_uuid FROM kudaki_mountain.recommended_gears rg JOIN kudaki_user.users u ON rg.user_uuid = u.uuid JOIN kudaki_user.profiles p ON p.user_uuid = u.uuid JOIN kudaki_mountain.mountains m ON rg.mountain_uuid = m.uuid WHERE rg.uuid = ?;", inEvent.RecommendedGearUuid)
	errorkit.ErrorHandled(err)

	var rgt RecommendedGearTemp
	err = rgRow.Scan(
		&rgt.Id,
		&rgt.Uuid,
		&rgt.UpVote,
		&rgt.DownVote,
		&rgt.Seen,
		&rgt.CreatedAt,
		&rgt.CreatorFullName,
		&rgt.CreatorEmail,
		&rgt.MountainName,
		&rgt.mountainUuid)
	if err == sql.ErrNoRows {
		return nil
	}

	mfRows, err := rrgi.DBO.Query("SELECT file_path FROM kudaki_mountain.mountain_files WHERE mountain_uuid = ?;", rgt.mountainUuid)
	errorkit.ErrorHandled(err)
	defer mfRows.Close()

	for mfRows.Next() {
		var mf string
		mfRows.Scan(&mf)

		rgt.MountainFiles = append(rgt.MountainFiles, &mf)
	}

	rgt.Seen++

	return &rgt
}
