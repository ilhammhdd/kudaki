package externals

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-payment-service/adapters"
	"github.com/ilhammhdd/kudaki-event-payment-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases/events"
)

type PaymentRequestDoku struct{}

func (prd *PaymentRequestDoku) Work() interface{} {
	adapter := &adapters.PaymentRequestDoku{}
	usecase := &usecases.PaymentRequestDoku{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: prd,
		eventDrivenAdapter:  adapter,
		eventDrivenUsecase:  usecase,
		eventName:           events.EventPaymentServiceCommandTopic_PAYMENT_REQUEST_DOKU.String(),
		inTopics:            []string{events.EventPaymentServiceCommandTopic_PAYMENT_REQUEST_DOKU.String()},
		outTopic:            events.EventPaymentServiceEventTopic_PAYMENT_REQUESTED_DOKU.String()}

	ede.handle()
	return nil
}

func (prd *PaymentRequestDoku) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.PaymentRequestedDoku)

	dbo := mysql.NewDBOperation(mysql.CommandDB)

	_, err := dbo.Command("UPDATE kudaki_event.doku_invoices SET words = ?, transaction_id_merchant = ? WHERE uuid = ?;",
		out.DokuInvoice.Words, out.DokuInvoice.TransactionIdMerchant, out.DokuInvoice.Uuid)
	errorkit.ErrorHandled(err)
}
