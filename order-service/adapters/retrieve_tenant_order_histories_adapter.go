package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-order-service/usecases"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type RetrieveTenantOrderHistories struct {
	Result RetrieveTenantOrderHistoriesResultSchema
}

func (rtoh *RetrieveTenantOrderHistories) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveTenantsOrderHistories
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rtoh *RetrieveTenantOrderHistories) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.TenantOrderHistoriesRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveTenantOrderHistoriesResultSchema struct {
	Orders []*usecases.OrderTemp `json:"orders"`
}

func (rtoh *RetrieveTenantOrderHistories) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rtoh.Result = RetrieveTenantOrderHistoriesResultSchema{
		Orders: i[0].([]*usecases.OrderTemp)}

	return rtoh
}

func (rtoh *RetrieveTenantOrderHistories) ParseToResult() []byte {
	resultJSON, err := json.Marshal(rtoh.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
