package externals

import (
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-user-auth-service/adapters"
	"github.com/ilhammhdd/kudaki-user-auth-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases"
	"github.com/ilhammhdd/kudaki-user-auth-service/usecases/events"
)

type ResetPasswordSendEmail struct{}

func (rpse *ResetPasswordSendEmail) Work() interface{} {
	usecase := usecases.ResetPasswordSendEmail{DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: rpse,
		eventDrivenAdapter:  new(adapters.ResetPasswordSendEmail),
		eventDrivenUsecase:  &usecase,
		eventName:           events.UserAuthServiceCommandTopic_SEND_RESET_PASSWORD_EMAIL.String(),
		inTopics:            []string{events.UserAuthServiceCommandTopic_SEND_RESET_PASSWORD_EMAIL.String()},
		outTopic:            events.UserAuthServiceEventTopic_RESET_PASSWORD_EMAIL_SENT.String()}

	ede.handle()
	return nil
}

func (rpse *ResetPasswordSendEmail) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.ResetPasswordEmailSent)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	/* grpcConn, err := grpc.Dial(os.Getenv("USER_REPO_SERVICE_GRPC_ADDRESS"), grpc.WithInsecure())
	errorkit.ErrorHandled(err)

	resetPassword := &user.ResetPassword{
		Token: out.RestToken,
		User:  out.User}

	resetPassRepoClient := kudakigrpc.NewResetPasswordRepoClient(grpcConn)
	_, err = resetPassRepoClient.AddResetPassword(context.Background(), resetPassword)
	errorkit.ErrorHandled(err) */

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO reset_passwords(user_uuid,token) VALUES(?,?) ON DUPLICATE KEY UPDATE token = ?;", out.User.Uuid, out.RestToken, out.User.Uuid)
	errorkit.ErrorHandled(err)
}

type ResetPassword struct{}

func (rp *ResetPassword) Work() interface{} {
	usecase := &usecases.ResetPassword{DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: rp,
		eventDrivenAdapter:  new(adapters.ResetPassword),
		eventDrivenUsecase:  usecase,
		eventName:           events.UserAuthServiceCommandTopic_RESET_PASSWORD.String(),
		inTopics:            []string{events.UserAuthServiceCommandTopic_RESET_PASSWORD.String()},
		outTopic:            events.UserAuthServiceEventTopic_PASSWORD_RESETED.String()}

	ede.handle()
	return nil
}

func (rp *ResetPassword) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.PasswordReseted)
	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	rp.updateUser(out)
	// rp.reIndexUser(out)
}

func (rp *ResetPassword) updateUser(out *events.PasswordReseted) {
	/* grpcConn, err := grpc.Dial(os.Getenv("USER_REPO_SERVICE_GRPC_ADDRESS"), grpc.WithInsecure())
	errorkit.ErrorHandled(err)

	userRepoClient := kudakigrpc.NewUserRepoClient(grpcConn)
	_, err = userRepoClient.UpdateUser(context.Background(), out.User)
	errorkit.ErrorHandled(err) */

	dbo := mysql.NewDBOperation(mysql.CommandDB)

	_, err := dbo.Command("UPDATE users SET password=? WHERE uuid = ?;", out.User.Password, out.User.Uuid)
	errorkit.ErrorHandled(err)
}

// func (rp *ResetPassword) reIndexUser(out *events.PasswordReseted) {
// 	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.User.Name())
// 	client.CreateIndex(kudakiredisearch.User.Schema())

// 	sanitizer := new(kudakiredisearch.RedisearchText)

// 	sanitizer.Set(out.User.Uuid)
// 	doc := redisearch.NewDocument(sanitizer.Sanitize(), 1.0)
// 	doc.Set("user_password", out.User.Password)

// 	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, doc)
// 	errorkit.ErrorHandled(err)
// }
