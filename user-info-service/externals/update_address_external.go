package externals

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-info-service/adapters"
	"github.com/ilhammhdd/kudaki-user-info-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type UpdateAddress struct{}

func (ua *UpdateAddress) Work() interface{} {
	usecase := usecases.UpdateAddress{DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: ua,
		eventDrivenAdapter:  new(adapters.UpdateAddress),
		eventDrivenUsecase:  &usecase,
		eventName:           events.UserInfoServiceCommandTopic_UPDATE_ADDRESS.String(),
		inTopics:            []string{events.UserInfoServiceCommandTopic_UPDATE_ADDRESS.String()},
		outTopic:            events.UserInfoServiceEventTopic_ADDRESS_UPDATED.String()}

	ede.handle()
	return nil
}

func (ua *UpdateAddress) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.AddressUpdated)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_user.addresses SET full_address=?,receiver_name=?,receiver_phone_number=?,zip_code=?,latitude=?,longitude=? WHERE uuid = ?;",
		out.AddressAfter.FullAddress,
		out.AddressAfter.ReceiverName,
		out.AddressAfter.ReceiverPhoneNumber,
		out.AddressAfter.ZipCode,
		out.AddressAfter.Latitude,
		out.AddressAfter.Longitude,
		out.AddressBefore.Uuid)
	errorkit.ErrorHandled(err)
}
