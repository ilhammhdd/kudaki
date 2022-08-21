package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-info-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type UpdateAddress struct {
	DBO DBOperator
}

func (ua *UpdateAddress) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := ua.initInOutEvent(in)

	outEvent.User = GetUserFromKudakiToken(inEvent.KudakiToken)
	profile := ua.retrieveProfile(outEvent.User)

	beforeAddress := ua.retrieveAddress(profile, inEvent.AddressUuid)
	if beforeAddress == nil {
		outEvent.EventStatus.Errors = []string{"address doesn't exists"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		return outEvent
	}

	outEvent.AddressAfter = ua.parseNewAddress(profile, inEvent)
	outEvent.AddressBefore = beforeAddress

	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (ua *UpdateAddress) initInOutEvent(in proto.Message) (inEvent *events.UpdateAddress, outEvent *events.AddressUpdated) {
	inEvent = in.(*events.UpdateAddress)

	outEvent = new(events.AddressUpdated)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid
	outEvent.UpdateAddress = inEvent

	return
}

func (ua *UpdateAddress) retrieveProfile(usr *user.User) *user.Profile {
	row, err := ua.DBO.QueryRow("SELECT id,uuid,full_name,photo,created_at from kudaki_user.profiles WHERE user_uuid = ?;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var profile user.Profile
	var createdAt sql.NullInt64
	errorkit.ErrorHandled(row.Scan(&profile.Id, &profile.Uuid, &profile.FullName, &profile.Photo, &createdAt))
	createdAtProtoTime, err := ptypes.TimestampProto(time.Unix(createdAt.Int64, 0))
	errorkit.ErrorHandled(err)
	profile.CreatedAt = createdAtProtoTime

	return &profile
}

func (ua *UpdateAddress) parseNewAddress(profile *user.Profile, inEvent *events.UpdateAddress) *user.Address {
	return &user.Address{
		FullAddress:         inEvent.FullAddress,
		Latitude:            inEvent.Latitude,
		Longitude:           inEvent.Longitude,
		Profile:             profile,
		ReceiverName:        inEvent.ReceiverName,
		ReceiverPhoneNumber: inEvent.ReceiverPhoneNumber,
		Uuid:                inEvent.AddressUuid}
}

func (ua *UpdateAddress) retrieveAddress(profile *user.Profile, addressUUID string) *user.Address {
	row, err := ua.DBO.QueryRow("SELECT id,uuid,full_address,receiver_name,receiver_phone_number,zip_code,latitude,longitude,created_at FROM kudaki_user.addresses WHERE uuid = ? AND profile_uuid = ?;", addressUUID, profile.Uuid)
	errorkit.ErrorHandled(err)

	var address user.Address
	var createdAt sql.NullInt64
	err = row.Scan(&address.Id, &address.Uuid, &address.FullAddress, &address.ReceiverName, &address.ReceiverPhoneNumber, &address.ZipCode, &address.Latitude, &address.Longitude, &createdAt)
	if err == sql.ErrNoRows {
		return nil
	}
	errorkit.ErrorHandled(err)
	createdAtProtoTime, err := ptypes.TimestampProto(time.Unix(createdAt.Int64, 0))
	errorkit.ErrorHandled(err)
	address.CreatedAt = createdAtProtoTime
	address.Profile = profile

	return &address
}
