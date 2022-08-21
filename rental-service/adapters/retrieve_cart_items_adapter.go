package adapters

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-rental-service/usecases"
	"github.com/ilhammhdd/kudaki-rental-service/usecases/events"
)

type RetrieveCartItems struct {
	Result RetrieveCartItemsResultSchema
}

func (rci *RetrieveCartItems) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.RetrieveCartItems
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (rci *RetrieveCartItems) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CartItemsRetrieved)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type RetrieveCartItemsResultSchema struct {
	Cart        *usecases.CartTemp         `json:"cart"`
	Storefronts []*usecases.StorefrontTemp `json:"storefronts"`
}

func (rci *RetrieveCartItems) SetResultSources(i ...interface{}) usecases.ResultSchemer {
	rci.Result = RetrieveCartItemsResultSchema{
		Cart:        i[0].(*usecases.CartTemp),
		Storefronts: i[1].([]*usecases.StorefrontTemp)}

	return rci
}

func (rci *RetrieveCartItems) ParseToResult() []byte {
	resultJSON, err := json.Marshal(rci.Result)
	errorkit.ErrorHandled(err)

	return resultJSON
}
