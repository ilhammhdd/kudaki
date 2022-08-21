package usecases

import (
	"database/sql"
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-user-info-service/entities/aggregates/user"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type UpdateProfile struct {
	DBO DBOperator
}

func (up *UpdateProfile) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := up.initInOutEvent(in)

	outEvent.User = GetUserFromKudakiToken(inEvent.KudakiToken)

	profileBefore := up.retrieveProfile(outEvent.User.Uuid)
	if profileBefore == nil {
		outEvent.EventStatus.Errors = []string{"user's profile not exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	outEvent.ProfileBefore = profileBefore
	outEvent.ProfileAfter = up.newProfile(inEvent, outEvent.ProfileBefore)

	outEvent.EventStatus.HttpCode = http.StatusOK
	return outEvent
}

func (up *UpdateProfile) initInOutEvent(in proto.Message) (inEvent *events.UpdateProfile, outEvent *events.ProfileUpdated) {
	inEvent = in.(*events.UpdateProfile)

	outEvent = new(events.ProfileUpdated)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.UpdateProfile = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (up *UpdateProfile) retrieveProfile(userUUID string) *user.Profile {
	row, err := up.DBO.QueryRow("SELECT id,uuid,full_name,photo,created_at FROM kudaki_user.profiles WHERE user_uuid=?;", userUUID)
	errorkit.ErrorHandled(err)

	var profile user.Profile
	var createdAt int64
	if row.Scan(&profile.Id, &profile.Uuid, &profile.FullName, &profile.Photo, &createdAt) == sql.ErrNoRows {
		return nil
	}

	return &profile
}

func (up *UpdateProfile) newProfile(inEvent *events.UpdateProfile, beforeProfile *user.Profile) *user.Profile {
	var afterProfile user.Profile
	afterProfile = *beforeProfile

	afterProfile.FullName = inEvent.FullName
	afterProfile.Photo = inEvent.Photo

	return &afterProfile
}
