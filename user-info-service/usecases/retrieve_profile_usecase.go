package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-user-info-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type RetrieveProfile struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rp *RetrieveProfile) Handle(in proto.Message) (out proto.Message) {
	_, outEvent := rp.initInOutEvent(in)

	profile := rp.retrieveProfile(outEvent.Requester.Uuid)
	if profile == nil {
		return nil
	}

	outEvent.Profile = profile
	outEvent.Result = rp.ResultSchemer.SetResultSources(rp.parseToTemp(profile)).ParseToResult()
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (rp *RetrieveProfile) initInOutEvent(in proto.Message) (inEvent *events.RetrieveProfile, outEvent *events.ProfileRetrieved) {
	inEvent = in.(*events.RetrieveProfile)

	outEvent = new(events.ProfileRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (rp *RetrieveProfile) retrieveProfile(userUuid string) *user.Profile {
	row, err := rp.DBO.QueryRow("SELECT p.id,p.uuid,p.full_name,p.photo,u.phone_number,p.created_at FROM kudaki_user.profiles p JOIN kudaki_user.users u ON u.uuid=p.user_uuid WHERE u.uuid=?;", userUuid)
	errorkit.ErrorHandled(err)

	var profile user.Profile
	profile.User = new(user.User)
	var createdAt int64

	if row.Scan(&profile.Id, &profile.Uuid, &profile.FullName, &profile.Photo, &profile.User.PhoneNumber, &createdAt) == sql.ErrNoRows {
		return nil
	}

	createdAtProtoTime, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	profile.CreatedAt = createdAtProtoTime

	return &profile
}

func (rp *RetrieveProfile) parseToTemp(profile *user.Profile) *RetrieveProfileTemp {
	return &RetrieveProfileTemp{
		Id:          profile.Id,
		Uuid:        profile.Uuid,
		FullName:    profile.FullName,
		Photo:       profile.Photo,
		PhoneNumber: profile.User.PhoneNumber,
		CreatedAt:   profile.CreatedAt.GetSeconds()}
}

type RetrieveProfileTemp struct {
	Id          int64  `json:"id"`
	Uuid        string `json:"uuid"`
	FullName    string `json:"full_name"`
	Photo       string `json:"photo"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   int64  `json:"created_at"`
}
