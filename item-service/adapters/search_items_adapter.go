package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-service/usecases"
	"github.com/ilhammhdd/kudaki-item-service/usecases/events"
)

type SearchItems struct {
	Result ItemsRetrievedSchema
}

func (si *SearchItems) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.SearchItems
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (si *SearchItems) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.ItemsSearched)

	outByte, err := proto.Marshal(out)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

func (si *SearchItems) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	si.Result = ItemsRetrievedSchema{
		Items: i[0].([]*usecases.ItemTemp)}

	return si
}

func (si *SearchItems) ParseToResult() []byte {
	resultJSON, err := json.Marshal(si.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
