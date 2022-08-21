package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-order-service/usecases/events"
)

type OwnerOrderDisapproved struct{}

func (ood *OwnerOrderDisapproved) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.OwnerOrderDisapproved
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}
