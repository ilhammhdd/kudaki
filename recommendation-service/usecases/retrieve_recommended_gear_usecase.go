package usecases

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-recommendation-service/usecases/events"
)

type RetrieveRecommendationGear struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rrg *RetrieveRecommendationGear) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rrg.initInOutEvent(in)

	recommendedGears := rrg.retrieveRecommendedGears(inEvent)
	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.Result = rrg.ResultSchemer.SetResultSources(recommendedGears).ParseToResult()

	return outEvent
}

func (rrg *RetrieveRecommendationGear) initInOutEvent(in proto.Message) (inEvent *events.RetrieveRecommendedGears, outEvent *events.RecommendedGearsRetrieved) {
	inEvent = in.(*events.RetrieveRecommendedGears)

	outEvent = new(events.RecommendedGearsRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.RetrieveRecommendedGears = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

type RecommendedGearTemp struct {
	Id              int64     `json:"id"`
	Uuid            string    `json:"uuid"`
	UpVote          int32     `json:"up_vote"`
	DownVote        int32     `json:"down_vote"`
	Seen            int32     `json:"seen"`
	CreatedAt       int64     `json:"created_at"`
	CreatorFullName string    `json:"creator_full_name"`
	CreatorEmail    string    `json:"creator_email"`
	mountainUuid    string    `json:"-"`
	MountainName    string    `json:"mountain_name"`
	MountainFiles   []*string `json:"mountain_files"`
}

func (rrg *RetrieveRecommendationGear) retrieveRecommendedGears(inEvent *events.RetrieveRecommendedGears) []*RecommendedGearTemp {
	rgRows, err := rrg.DBO.Query("SELECT rg.id, rg.uuid, rg.upvote, rg.downvote, rg.seen, rg.created_at, p.full_name, u.email, m.name FROM kudaki_mountain.recommended_gears rg JOIN kudaki_user.users u ON rg.user_uuid = u.uuid JOIN kudaki_user.profiles p ON p.user_uuid = u.uuid JOIN kudaki_mountain.mountains m ON rg.mountain_uuid = m.uuid WHERE rg.mountain_uuid = ?;", inEvent.MountainUuid)
	errorkit.ErrorHandled(err)
	defer rgRows.Close()

	var rgts []*RecommendedGearTemp
	for rgRows.Next() {
		var rgt RecommendedGearTemp
		rgRows.Scan(
			&rgt.Id,
			&rgt.Uuid,
			&rgt.UpVote,
			&rgt.DownVote,
			&rgt.Seen,
			&rgt.CreatedAt,
			&rgt.CreatorFullName,
			&rgt.CreatorEmail,
			&rgt.MountainName)

		mfRows, err := rrg.DBO.Query("SELECT file_path FROM kudaki_mountain.mountain_files WHERE mountain_uuid = ?;", inEvent.MountainUuid)
		errorkit.ErrorHandled(err)
		defer mfRows.Close()

		for mfRows.Next() {
			var mf string
			mfRows.Scan(&mf)

			rgt.MountainFiles = append(rgt.MountainFiles, &mf)
		}

		rgts = append(rgts, &rgt)
	}

	return rgts
}
