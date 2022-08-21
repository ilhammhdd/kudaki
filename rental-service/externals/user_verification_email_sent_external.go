package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-rental-service/adapters"
	"github.com/ilhammhdd/kudaki-rental-service/entities/aggregates/rental"
	"github.com/ilhammhdd/kudaki-rental-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-rental-service/usecases"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type UserVerificationEmailSent struct{}

func (uves *UserVerificationEmailSent) Work() interface{} {
	edde := EventDrivenDownstreamExternal{
		PostUsecaseExecutor: uves,
		eventDrivenAdapter:  new(adapters.UserVerificationEmailSent),
		eventDrivenUsecase:  new(usecases.UserVerificationEmailSent),
		eventName:           events.UserAuthServiceEventTopic_USER_VERIFICATION_EMAIL_SENT.String(),
		inTopics:            []string{events.UserAuthServiceEventTopic_USER_VERIFICATION_EMAIL_SENT.String()}}

	edde.handle()
	return nil
}

func (uves *UserVerificationEmailSent) ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse) {
	emptyCart := usecaseRes.Data.(*rental.Cart)

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_rental.carts(uuid,user_uuid,total_price,total_items,open,created_at) VALUES(?,?,?,?,1,UNIX_TIMESTAMP());",
		emptyCart.Uuid, emptyCart.UserUuid, emptyCart.TotalPrice, emptyCart.TotalItems)
	errorkit.ErrorHandled(err)
}
