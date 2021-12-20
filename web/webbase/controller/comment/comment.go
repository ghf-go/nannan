package comment

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/commentlogic"
)

func NewCommentAction(ctx *web.EngineCtx) error {
	req := &reqNewComment{}
	ctx.Verify(req)
	return ctx.JsonSuccess(commentlogic.NewComment(ctx.ForceUID(), req.TargetID, req.ParentID, req.TargetType, req.Content))
}
func CommentListAction(ctx *web.EngineCtx) error {
	req := &reqCommentList{}
	ctx.Verify(req)
	return ctx.JsonSuccess(map[string]interface{}{
		"total": commentlogic.CommentTotal(ctx.UID(), req.TargetID, req.ParentID, req.TargetType),
		"list":  commentlogic.CommentList(ctx.UID(), req.TargetID, req.ParentID, req.TargetType, req.Start, req.PageSize),
	})
}
