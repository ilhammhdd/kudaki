package usecases

import (
	"encoding/json"
	"net/http"

	"github.com/ilhammhdd/kudaki-user-info-service/entities/aggregates/user"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type RetrieveAddresses struct {
	DBO DBOperator
}

func (ra *RetrieveAddresses) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := ra.initInOutEvent(in)

	outEvent.User = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Result = ra.parseResultRows(ra.retrieveAddresses(outEvent.User))
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (ra *RetrieveAddresses) initInOutEvent(in proto.Message) (inEvent *events.RetrieveAddresses, outEvent *events.AddressesRetrieved) {
	inEvent = in.(*events.RetrieveAddresses)

	outEvent = new(events.AddressesRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid

	return
}

type resultRowSchema struct {
	Id                  int64   `json:"id"`
	Uuid                string  `json:"uuid"`
	FullAddress         string  `json:"full_address"`
	ReceiverName        string  `json:"receiver_name"`
	ReceiverPhoneNumber string  `json:"receiver_phone_number"`
	ZipCode             string  `json:"zip_code"`
	Latitude            float32 `json:"latitude"`
	Longitude           float32 `json:"longitude"`
}

func (ra *RetrieveAddresses) retrieveAddresses(usr *user.User) []*resultRowSchema {
	rows, err := ra.DBO.Query("SELECT a.id,a.uuid,a.full_address,a.receiver_name,a.receiver_phone_number,a.zip_code,a.latitude,a.longitude FROM kudaki_user.addresses a JOIN kudaki_user.profiles p ON a.profile_uuid = p.uuid JOIN kudaki_user.users u ON p.user_uuid = u.uuid WHERE u.uuid = ?;", usr.Uuid)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var resultRows []*resultRowSchema

	for rows.Next() {
		var resultRow resultRowSchema
		err := rows.Scan(
			&resultRow.Id,
			&resultRow.Uuid,
			&resultRow.FullAddress,
			&resultRow.ReceiverName,
			&resultRow.ReceiverPhoneNumber,
			&resultRow.ZipCode,
			&resultRow.Latitude,
			&resultRow.Longitude)
		errorkit.ErrorHandled(err)

		resultRows = append(resultRows, &resultRow)
	}

	return resultRows
}

func (ra *RetrieveAddresses) parseResultRows(resultRows []*resultRowSchema) []byte {
	resultMapJSON, err := json.Marshal(resultRows)
	errorkit.ErrorHandled(err)

	return resultMapJSON
}
