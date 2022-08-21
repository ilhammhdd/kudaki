package externals

import (
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-user-auth-service/adapters"
	"github.com/ilhammhdd/kudaki-user-auth-service/externals/mysql"
	kudakimysql "github.com/ilhammhdd/kudaki-user-auth-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type VerifyUser struct{}

func (vu *VerifyUser) Work() interface{} {
	usecase := &usecases.VerifyUser{DBO: kudakimysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: vu,
		eventDrivenAdapter:  new(adapters.VerifyUser),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserAuthServiceCommandTopic_VERIFY_USER.String(),
		inTopics:            []string{events.UserAuthServiceCommandTopic_VERIFY_USER.String()},
		outTopic:            events.UserAuthServiceEventTopic_USER_VERIFIED.String()}
	ede.handle()
	return nil
}

func (vu *VerifyUser) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.UserVerified)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	/* grpcConn, err := grpc.Dial(os.Getenv("USER_REPO_SERVICE_GRPC_ADDRESS"), grpc.WithInsecure())
	errorkit.ErrorHandled(err)
	defer grpcConn.Close()

	unverifUser := &user.UnverifiedUser{User: out.User}

	unverifUserRepoClient := kudakigrpc.NewUnverifiedUserRepoClient(grpcConn)
	_, err = unverifUserRepoClient.DeleteUnverifiedUser(context.Background(), unverifUser)
	errorkit.ErrorHandled(err) */

	dbo := kudakimysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("DELETE FROM unverified_users WHERE user_uuid = ?;", out.User.Uuid)
	errorkit.ErrorHandled(err)
}
