package externals

import (
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-user-info-service/adapters"
	"github.com/ilhammhdd/kudaki-user-info-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases"
	"github.com/ilhammhdd/kudaki-user-info-service/usecases/events"
)

type AddAddress struct{}

func (ad *AddAddress) Work() interface{} {
	usecase := usecases.AddAddress{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: ad,
		eventDrivenAdapter:  new(adapters.AddAddress),
		eventDrivenUsecase:  &usecase,
		eventName:           events.UserInfoServiceCommandTopic_ADD_ADDRESS.String(),
		inTopics:            []string{events.UserInfoServiceCommandTopic_ADD_ADDRESS.String()},
		outTopic:            events.UserInfoServiceEventTopic_ADDRESS_ADDED.String()}

	ede.handle()
	return nil
}

func (ad *AddAddress) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.AddressAdded)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_user.addresses(full_address,latitude,longitude,profile_uuid,receiver_name,receiver_phone_number,uuid,zip_code,created_at) VALUES(?,ROUND(?,8),ROUND(?,8),?,?,?,?,?,UNIX_TIMESTAMP());",
		out.UsersAddress.FullAddress,
		out.UsersAddress.Latitude,
		out.UsersAddress.Longitude,
		out.UsersAddress.Profile.Uuid,
		out.UsersAddress.ReceiverName,
		out.UsersAddress.ReceiverPhoneNumber,
		out.UsersAddress.Uuid,
		out.UsersAddress.ZipCode)
	errorkit.ErrorHandled(err)
}
