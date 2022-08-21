package usecases

import (
	"database/sql"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-review-service/entities/aggregates/store"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type RetrieveItemReviewComment struct {
	DBO           DBOperator
	ResultSchemer ResultSchemer
}

func (rirc *RetrieveItemReviewComment) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := rirc.initIntOutEvent(in)

	itemReview := rirc.retrieveItemReview(inEvent)
	if itemReview == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"item review with the given uuid not found"}

		return outEvent
	}

	reviewComments := rirc.retrieveItemReviewComment(&itemReview.ItemReview, inEvent)

	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.Result = rirc.ResultSchemer.SetResultSources(itemReview, reviewComments).ParseToResult()

	return outEvent
}

func (rirc *RetrieveItemReviewComment) initIntOutEvent(in proto.Message) (inEvent *events.RetrieveItemReviewComments, outEvent *events.ItemReviewCommentsRetrieved) {
	inEvent = in.(*events.RetrieveItemReviewComments)

	outEvent = new(events.ItemReviewCommentsRetrieved)
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Requester = GetUserFromKudakiToken(inEvent.KudakiToken)
	outEvent.RetrieveItemReviewComments = inEvent
	outEvent.Uid = inEvent.Uid

	return
}

func (rirc *RetrieveItemReviewComment) retrieveItemReview(inEvent *events.RetrieveItemReviewComments) *ItemReviewTemp {
	row, err := rirc.DBO.QueryRow("SELECT ir.id,ir.uuid,ir.review,ir.rating,ir.created_at,u.uuid AS reviewer_uuid,u.email AS reviewer_email,p.full_name AS reviewer_full_name,p.photo AS reviewer_photo FROM kudaki_store.item_reviews ir JOIN kudaki_user.users u ON ir.user_uuid=u.uuid JOIN kudaki_user.profiles p ON u.uuid=p.user_uuid WHERE ir.uuid=?;", inEvent.ItemReviewUuid)
	errorkit.ErrorHandled(err)

	var itemReviewTemp ItemReviewTemp

	if row.Scan(
		&itemReviewTemp.Id,
		&itemReviewTemp.Uuid,
		&itemReviewTemp.Review,
		&itemReviewTemp.RatingT,
		&itemReviewTemp.CreatedAtT,
		&itemReviewTemp.ReviewerUuid,
		&itemReviewTemp.ReviewerEmail,
		&itemReviewTemp.ReviewerFullName,
		&itemReviewTemp.ReviewerPhoto) == sql.ErrNoRows {
		return nil
	}

	return &itemReviewTemp
}

func (rirc *RetrieveItemReviewComment) retrieveItemReviewComment(itemReview *store.ItemReview, inEvent *events.RetrieveItemReviewComments) []*ReviewCommentTemp {
	rows, err := rirc.DBO.Query("SELECT rc.id,rc.uuid,rc.comment,rc.created_at,u.uuid AS commenter_uuid,u.email AS commenter_email,p.full_name AS commenter_full_name,p.photo AS commenter_photo FROM (SELECT id FROM kudaki_store.review_comments WHERE item_review_uuid=? LIMIT ?,?) rc_ids JOIN kudaki_store.review_comments rc ON rc_ids.id=rc.id JOIN kudaki_user.users u ON rc.user_uuid=u.uuid JOIN kudaki_user.profiles p ON p.user_uuid=u.uuid;",
		itemReview.Uuid, inEvent.Offset, inEvent.Limit)
	errorkit.ErrorHandled(err)
	defer rows.Close()

	var reviewCommentTemps []*ReviewCommentTemp
	for rows.Next() {
		var reviewCommentTemp ReviewCommentTemp

		rows.Scan(
			&reviewCommentTemp.Id,
			&reviewCommentTemp.Uuid,
			&reviewCommentTemp.Comment,
			&reviewCommentTemp.CreatedAtT,
			&reviewCommentTemp.CommenterUuid,
			&reviewCommentTemp.CommenterEmail,
			&reviewCommentTemp.CommenterFullName,
			&reviewCommentTemp.CommenterPhoto)

		reviewCommentTemps = append(reviewCommentTemps, &reviewCommentTemp)
	}

	return reviewCommentTemps
}

type ReviewCommentTemp struct {
	store.ReviewComment
	CreatedAtT        int64  `json:"created_at"`
	CommenterUuid     string `json:"commenter_uuid"`
	CommenterEmail    string `json:"commenter_email"`
	CommenterFullName string `json:"commenter_full_name"`
	CommenterPhoto    string `json:"commenter_photo"`
}
