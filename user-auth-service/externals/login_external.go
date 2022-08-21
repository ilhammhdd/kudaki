package externals

import (
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-user-auth-service/adapters"
	"github.com/ilhammhdd/kudaki-user-auth-service/entities/aggregates/user"
	"github.com/ilhammhdd/kudaki-user-auth-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type Login struct{}

func (l *Login) Work() interface{} {
	usecase := &usecases.Login{DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: l,
		eventDrivenAdapter:  new(adapters.Login),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserAuthServiceCommandTopic_LOGIN.String(),
		inTopics:            []string{events.UserAuthServiceCommandTopic_LOGIN.String()},
		outTopic:            events.UserAuthServiceEventTopic_LOGGEDIN.String()}
	ede.handle()
	return nil
}

func (l *Login) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.Loggedin)
	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}
	l.updateUserAuthToken(out.User)
	// l.reindexUser(out.User)
}

func (l *Login) updateUserAuthToken(usr *user.User) {
	/* grpcConn, err := grpc.Dial(os.Getenv("USER_REPO_SERVICE_GRPC_ADDRESS"), grpc.WithInsecure())
	errorkit.ErrorHandled(err)
	defer grpcConn.Close()

	userRepoClient := kudakigrpc.NewUserRepoClient(grpcConn)
	_, err = userRepoClient.UpdateUser(context.Background(), usr)
	errorkit.ErrorHandled(err) */

	dbo := mysql.NewDBOperation(mysql.CommandDB)

	_, err := dbo.Command("UPDATE users SET token=? WHERE uuid=?;", usr.Token, usr.Uuid)
	errorkit.ErrorHandled(err)
}

// func (l *Login) reindexUser(usr *user.User) {
// 	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
// 	client.CreateIndex(kudakiredisearch.User.Schema())

// 	sanitizer := new(kudakiredisearch.RedisearchText)

// 	sanitizer.Set(usr.Uuid)
// 	doc := redisearch.NewDocument(sanitizer.Sanitize(), 1.0)
// 	doc.Set("user_token", usr.Token)
// 	client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
// }
