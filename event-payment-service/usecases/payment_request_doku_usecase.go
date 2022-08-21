package usecases

import (
	"net/http"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-event-payment-service/entities/aggregates/kudaki_event"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases/events"
)

type PaymentRequestDoku struct {
	DBO DBOperator
}

func (prd *PaymentRequestDoku) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := prd.initInOutEvent(in)

	di := prd.retreiveDokuInvoice(inEvent)
	prd.updateExistingDokuInvoice(inEvent, di)

	outEvent.DokuInvoice = di
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (prd *PaymentRequestDoku) initInOutEvent(in proto.Message) (inEvent *events.PaymentRequestDoku, outEvent *events.PaymentRequestedDoku) {
	inEvent = in.(*events.PaymentRequestDoku)

	outEvent = new(events.PaymentRequestedDoku)
	outEvent.EventStatus = new(events.Status)
	outEvent.Organizer = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (prd *PaymentRequestDoku) retreiveDokuInvoice(inEvent *events.PaymentRequestDoku) *kudaki_event.DokuInvoice {
	row, err := prd.DBO.QueryRow("SELECT id,uuid,kudaki_event_uuid,amount,purchase_amount,transaction_id_merchant,words,request_date_time,currency,purchase_currency,session_id,name,email,basket,status FROM kudaki_event.doku_invoices WHERE session_id=?;",
		inEvent.SessionId)
	errorkit.ErrorHandled(err)

	var di kudaki_event.DokuInvoice
	var status string
	err = row.Scan(
		&di.Id,
		&di.Uuid,
		&di.KudakiEventUuid,
		&di.Amount,
		&di.PurchaseAmount,
		&di.TransactionIdMerchant,
		&di.Words,
		&di.RequestDateTime,
		&di.Currency,
		&di.PurchaseCurrency,
		&di.SessionId,
		&di.Name,
		&di.Email,
		&di.Basket,
		&status)
	if errorkit.ErrorHandled(err) {
		return nil
	}

	return &di
}

func (prd *PaymentRequestDoku) updateExistingDokuInvoice(inEvent *events.PaymentRequestDoku, di *kudaki_event.DokuInvoice) {
	(*di).TransactionIdMerchant = inEvent.TransactionIdMerchant
	(*di).Words = inEvent.HashedWords
}
