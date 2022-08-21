package externals

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-item-review-service/adapters"
	"github.com/ilhammhdd/kudaki-item-review-service/externals/mysql"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases"
	"github.com/ilhammhdd/kudaki-item-review-service/usecases/events"
)

type CommentReviewItem struct{}

func (cri *CommentReviewItem) Work() interface{} {
	usecase := usecases.CommentReviewItem{
		DBO: mysql.NewDBOperation(mysql.QueryDB)}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: cri,
		eventDrivenAdapter:  new(adapters.CommentReviewItem),
		eventDrivenUsecase:  &usecase,
		eventName:           events.ItemReviewServiceCommandTopic_COMMENT_ITEM_REVIEW.String(),
		inTopics:            []string{events.ItemReviewServiceCommandTopic_COMMENT_ITEM_REVIEW.String()},
		outTopic:            events.ItemReviewServiceEventTopic_ITEM_REVIEW_COMMENTED.String()}

	ede.handle()
	return nil
}

func (cri *CommentReviewItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.ItemReviewCommented)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dbo := mysql.NewDBOperation(mysql.CommandDB)
	_, err := dbo.Command("INSERT INTO kudaki_store.review_comments(uuid,item_review_uuid,user_uuid,comment,created_at) VALUES(?,?,?,?,UNIX_TIMESTAMP());",
		out.ReviewComment.Uuid, out.ReviewComment.ItemReview.Uuid, out.ReviewComment.UserUuid, out.ReviewComment.Comment)
	errorkit.ErrorHandled(err)
}
