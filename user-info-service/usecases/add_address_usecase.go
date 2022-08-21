package usecases

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/golang/protobuf/ptypes"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-info-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type AddAddress struct {
	DBO DBOperator
}

func (ad *AddAddress) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := ad.initInOutEvent(in)
	outEvent.User = GetUserFromKudakiToken(inEvent.KudakiToken)

	if ad.addressExistsByLatLong(outEvent.User, inEvent) {
		outEvent.EventStatus.Errors = []string{"address with this coordinate already exists"}
		outEvent.EventStatus.HttpCode = http.StatusConflict
		return outEvent
	}

	outEvent.UsersAddress = ad.initNewAddress(ad.retrieveProfile(outEvent.User.Uuid), inEvent)
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (ad *AddAddress) initInOutEvent(in proto.Message) (inEvent *events.AddAddress, outEvent *events.AddressAdded) {
	inEvent = in.(*events.AddAddress)

	outEvent = new(events.AddressAdded)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid

	return
}

func (ad *AddAddress) addressExistsByLatLong(usr *user.User, inEvent *events.AddAddress) bool {
	row, err := ad.DBO.QueryRow("SELECT a.id FROM kudaki_user.addresses a JOIN profiles p ON p.uuid = a.profile_uuid WHERE p.user_uuid = ? AND a.latitude = ? AND a.longitude = ?;", usr.Uuid, fmt.Sprintf("%.8f", inEvent.Latitude), fmt.Sprintf("%.8f", inEvent.Longitude))
	errorkit.ErrorHandled(err)

	var addressID int64
	if row.Scan(&addressID) == sql.ErrNoRows {
		return false
	}
	return true
}

func (ad *AddAddress) retrieveProfile(userUUID string) *user.Profile {
	row, err := ad.DBO.QueryRow("SELECT id,uuid,full_name,photo,created_at FROM kudaki_user.profiles WHERE user_uuid = ?;", userUUID)
	errorkit.ErrorHandled(err)

	var profile user.Profile
	var createdAt int64
	row.Scan(&profile.Id, &profile.Uuid, &profile.FullName, &profile.Photo, &createdAt)
	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)
	profile.CreatedAt = createdAtProto

	return &profile
}

func (ad *AddAddress) initNewAddress(profile *user.Profile, inEvent *events.AddAddress) (newAddress *user.Address) {
	newAddress = new(user.Address)

	newAddress.CreatedAt = ptypes.TimestampNow()
	newAddress.FullAddress = inEvent.FullAddress
	newAddress.Latitude = inEvent.Latitude
	newAddress.Longitude = inEvent.Longitude
	newAddress.Profile = profile
	newAddress.ReceiverName = inEvent.ReceiverName
	newAddress.ReceiverPhoneNumber = inEvent.ReceiverPhoneNumber
	newAddress.Uuid = uuid.New().String()

	return newAddress
}
