package usecases

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-mountain-service/entities/aggregates/mountain"
	"github.com/ilhammhdd/kudaki-mountain-service/usecases/events"
)

type RetrieveMountains struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rm *RetrieveMountains) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rm.initInOutEvent(in)

	mts := rm.retrieveMountains(inEvent)
	rm.retrieveAndPasteMountainFiles(mts)

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.Result = rm.ResultSchemer.SetResultSources(mts).ParseToResult()

	return outEvent
}

func (rm *RetrieveMountains) initInOutEvent(in proto.Message) (inEvent *events.RetrieveMountains, outEvent *events.MountainsRetrieved) {
	inEvent = in.(*events.RetrieveMountains)

	outEvent = new(events.MountainsRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.RetrieveMountains = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (rm *RetrieveMountains) retrieveMountains(inEvent *events.RetrieveMountains) []*MountainTemp {
	rows, err := rm.DBO.Query("SELECT m.uuid, m.name, m.height, m.latitude, m.longitude, m.difficulty, m.description, m.created_at FROM (SELECT m_i.id FROM kudaki_mountain.mountains m_i LIMIT ?, ? ) m_ids JOIN kudaki_mountain.mountains m ON m_ids.id = m.id;",
		inEvent.Offset, inEvent.Limit)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var mountains []*MountainTemp
	for rows.Next() {
		var mt MountainTemp
		var mtfs []*MountainPhotoTemp

		rows.Scan(
			&mt.Uuid,
			&mt.Name,
			&mt.Height,
			&mt.Latitude,
			&mt.Longitude,
			&mt.Difficulty,
			&mt.Description,
			&mt.CreatedAt)
		mt.Photos = mtfs
		mountains = append(mountains, &mt)
	}

	return mountains
}

func (rm *RetrieveMountains) retrieveAndPasteMountainFiles(mountains []*MountainTemp) {
	for i := 0; i < len(mountains); i++ {
		rows, err := rm.DBO.Query("SELECT mf.file_path FROM kudaki_mountain.mountain_files mf WHERE mf.mountain_uuid = ?;",
			mountains[i].Uuid)
		errorkit.ErrorHandled(err)
		defer rows.Close()

		var mtf MountainPhotoTemp
		for rows.Next() {
			rows.Scan(&mtf.FilePath)
			(*mountains[i]).Photos = append((*mountains[i]).Photos, &mtf)
		}
	}
}

type MountainPhotoTemp struct {
	FilePath string `json:"file_path"`
}

type MountainTemp struct {
	mountain.Mountain
	CreatedAt int64                `json:"created_at"`
	Photos    []*MountainPhotoTemp `json:"photos"`
}
