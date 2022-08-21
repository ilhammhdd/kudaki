package usecases

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-item-review-service/entities/aggregates/user"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-review-service/entities/aggregates/store"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type CommentReviewItem struct {
	DBO DBOperator
}

func (cri *CommentReviewItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := cri.initInOutEvent(in)

	itemReview := cri.retrieveItemReview(inEvent)
	if itemReview == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"item review with the given uuid not found"}

		return outEvent
	}

	outEvent.ReviewComment = cri.initItemReviewComment(outEvent.Requester, inEvent, itemReview)

	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (cri *CommentReviewItem) initInOutEvent(in proto.Message) (inEvent *events.CommentItemReview, outEvent *events.ItemReviewCommented) {
	inEvent = in.(*events.CommentItemReview)

	outEvent = new(events.ItemReviewCommented)
	outEvent.CommentItemReview = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.Uid = inEvent.Uid

	return
}

func (cri *CommentReviewItem) retrieveItemReview(inEvent *events.CommentItemReview) *store.ItemReview {
	row, err := cri.DBO.QueryRow("SELECT id,uuid,user_uuid,item_uuid,review,rating,created_at FROM kudaki_store.item_reviews WHERE uuid=?;", inEvent.ItemReviewUuid)
	errorkit.ErrorHandled(err)

	var itemReview store.ItemReview
	itemReview.Item = new(store.Item)
	var createdAt int64
	if row.Scan(
		&itemReview.Id,
		&itemReview.Uuid,
		&itemReview.UserUuid,
		&itemReview.Item.Uuid,
		&itemReview.Review,
		&itemReview.Rating,
		&createdAt) == sql.ErrNoRows {
		return nil
	}

	createdAtProto, err := ptypes.TimestampProto(time.Unix(createdAt, 0))
	errorkit.ErrorHandled(err)

	itemReview.CreatedAt = createdAtProto

	return &itemReview
}

func (cri *CommentReviewItem) initItemReviewComment(usr *user.User, inEvent *events.CommentItemReview, itemReview *store.ItemReview) *store.ReviewComment {
	return &store.ReviewComment{
		Comment:    inEvent.Comment,
		CreatedAt:  ptypes.TimestampNow(),
		ItemReview: itemReview,
		UserUuid:   usr.Uuid,
		Uuid:       uuid.New().String()}
}
