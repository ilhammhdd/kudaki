package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases"
	"github.com/ilhammhdd/kudaki-event-payment-service/usecases/events"
)

type RetrieveOrganizerTransaction struct {
	Result RetrieveOrganizerTransactionResultSchema
}

func (rot *RetrieveOrganizerTransaction) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveOrganizerInvoices
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rot *RetrieveOrganizerTransaction) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.OrganizerInvoicesRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveOrganizerTransactionResultSchema struct {
	Invoices []*usecases.DokuInvoiceTemp `json:"invoices"`
}

func (rot *RetrieveOrganizerTransaction) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rot.Result = RetrieveOrganizerTransactionResultSchema{
		Invoices: i[0].([]*usecases.DokuInvoiceTemp)}

	return rot
}

func (rot *RetrieveOrganizerTransaction) ParseToResult() []byte {
	resultJSON, err := json.Marshal(rot.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
