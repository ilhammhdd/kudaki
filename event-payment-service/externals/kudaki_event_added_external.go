package externals

import (
	"log"
	"net/http"
	"os"

	"github.com/ilhammhdd/kudaki-event-payment-service/entities/aggregates/kudaki_event"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-event-payment-service/adapters"
	"github.com/ilhammhdd/kudaki-event-payment-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases/events"
)

type KudakiEventAdded struct{}

func (ae *KudakiEventAdded) Work() interface{} {
	usecase := usecases.KudakiEventAdded{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: ae,
		eventDrivenAdapter:  new(adapters.KudakiEventAdded),
		eventDrivenUsecase:  &usecase,
		eventName:           events.EventServiceEventTopic_KUDAKI_EVENT_ADDED.String(),
		inTopics:            []string{events.EventServiceEventTopic_KUDAKI_EVENT_ADDED.String()},
		outTopic:            events.EventPaymentServiceEventTopic_EVENT_DOKU_INVOICE_ISSUED.String()}

	ede.handle()
	return nil
}

func (ae *KudakiEventAdded) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.KudakiEventDokuInvoiceIssued)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	log.Println(out.DokuInvoice.Words)

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_event.doku_invoices(uuid,kudaki_event_uuid,mall_id,chain_merchant,amount,purchase_amount,transaction_id_merchant,words,request_date_time,currency,purchase_currency,session_id,name,email,basket,status,created_at) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,UNIX_TIMESTAMP());",
		out.DokuInvoice.Uuid,
		out.DokuInvoice.KudakiEventUuid,
		os.Getenv("MALLID"),
		0,
		out.DokuInvoice.Amount,
		out.DokuInvoice.PurchaseAmount,
		out.DokuInvoice.TransactionIdMerchant,
		out.DokuInvoice.Words,
		out.DokuInvoice.RequestDateTime,
		out.DokuInvoice.Currency,
		out.DokuInvoice.PurchaseCurrency,
		out.DokuInvoice.SessionId,
		out.DokuInvoice.Name,
		out.DokuInvoice.Email,
		out.DokuInvoice.Basket,
		kudaki_event.DokuInvoiceStatus_NEW.String())
	errorkit.ErrorHandled(err)
}
