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

type UpdateProfile struct{}

func (up *UpdateProfile) Work() interface{} {
	usecase := usecases.UpdateProfile{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: up,
		eventDrivenAdapter:  new(adapters.UpdateProfile),
		eventDrivenUsecase:  &usecase,
		eventName:           events.UserInfoServiceCommandTopic_UPDATE_PROFILE.String(),
		inTopics:            []string{events.UserInfoServiceCommandTopic_UPDATE_PROFILE.String()},
		outTopic:            events.UserInfoServiceEventTopic_PROFILE_UPDATED.String()}

	ede.handle()
	return nil
}

func (up *UpdateProfile) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.ProfileUpdated)
	in := inEvent.(*events.UpdateProfile)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE kudaki_user.profiles SET photo=?,full_name=? WHERE uuid=?;", out.ProfileAfter.Photo, out.ProfileAfter.FullName, out.ProfileAfter.Uuid)
	errorkit.ErrorHandled(err)

	dboUser := mysql.NewDBOperation(mysql.CommandDB)
	_, err = dboUser.Command("UPDATE kudaki_user.users SET phone_number=? WHERE uuid=?;", in.PhoneNumber, out.User.Uuid)
	errorkit.ErrorHandled(err)
}
