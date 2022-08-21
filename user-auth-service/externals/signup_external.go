package externals

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-user-auth-service/adapters"
	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-auth-service/externals/kafka"
	"github.com/ilhammhdd/kudaki-user-auth-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type Signup struct{}

func (s *Signup) Work() interface{} {
	usecase := &usecases.Signup{
		DBO:      mysql.NewDBOperation(mysql.QueryDB),
		Producer: kafka.NewProduction()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: s,
		eventDrivenAdapter:  new(adapters.Signup),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserAuthServiceCommandTopic_SIGN_UP.String(),
		inTopics:            []string{events.UserAuthServiceCommandTopic_SIGN_UP.String()},
		outTopic:            events.UserAuthServiceEventTopic_SIGNED_UP.String()}
	ede.handle()
	return nil
}

func (s *Signup) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.Signedup)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	s.insertUser(out.Profile.User)
	s.insertProfile(out.Profile)
}

func (s *Signup) insertUser(usr *user.User) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO users(uuid,email,password,token,role,phone_number,account_type,created_at) VALUES (?,?,?,?,?,?,?,UNIX_TIMESTAMP());",
		usr.Uuid, usr.Email, usr.Password, usr.Token, usr.Role.String(), usr.PhoneNumber, usr.AccountType.String())
	errorkit.ErrorHandled(err)

	_, err = dbo.Command("INSERT INTO unverified_users(user_uuid) VALUES (?);", usr.Uuid)
	errorkit.ErrorHandled(err)
}

func (s *Signup) insertProfile(profile *user.Profile) {
	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO profiles(uuid,user_uuid,full_name,photo,created_at) VALUES (?,?,?,?,UNIX_TIMESTAMP());",
		profile.Uuid, profile.User.Uuid, profile.FullName, profile.Photo)
	errorkit.ErrorHandled(err)
}
