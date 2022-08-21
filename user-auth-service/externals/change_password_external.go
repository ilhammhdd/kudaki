package externals

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/adapters"
	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-auth-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type ChangePassword struct{}

func (cp *ChangePassword) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.PasswordChanged)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	cp.updateUsersPassword(out.User)
}

func (cp *ChangePassword) updateUsersPassword(usr *user.User) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("UPDATE users SET password = ? WHERE uuid = ?;", usr.Password, usr.Uuid)
	errorkit.ErrorHandled(err)
}

func (cp *ChangePassword) Work() interface{} {
	usecase := &usecases.ChangePassword{DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: cp,
		eventDrivenAdapter:  new(adapters.ChangePassword),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserAuthServiceCommandTopic_CHANGE_PASSWORD.String(),
		inTopics:            []string{events.UserAuthServiceCommandTopic_CHANGE_PASSWORD.String()},
		outTopic:            events.UserAuthServiceEventTopic_PASSWORD_CHANGED.String()}

	ede.handle()
	return nil
}
