package comment

import (
	"github.com/ghf-go/nannan/gutils"
	"github.com/ghf-go/nannan/web"
	commentlogic2 "github.com/ghf-go/nannan/webbase/logic/commentlogic"
	"time"
)

func NewFeedAction(ctx *web.EngineCtx) error {
	ctx.LimitIP("newFeed", 1, time.Minute*3)
	req := &reqFeedAdd{}
	ctx.Verify(req)
	switch req.FeedType {
	case commentlogic2.FEED_TYPE_YUEHUI: //约会
		return newFeedYehui(ctx)
	case commentlogic2.FEED_TYPE_VOTE: //投票
		return newFeedVote(ctx)
	default:
		return newFeedBase(ctx) //基本
	}
}
func FeedList(ctx *web.EngineCtx) error {
	return ctx.JsonSuccess("OK")
}

//新增动态
func newFeedBase(ctx *web.EngineCtx) error {
	req := &reqFeedBaseAdd{}
	ctx.Verify(req)
	commentlogic2.NewFeed(
		ctx.ForceUID(),
		commentlogic2.FEED_TYPE_BASE,
		gutils.StrMaxLen(req.Content, 100),
		req.FeedImgs,
		req.X, req.Y,
		req.City,
		"{}",
		req.Content)
	return nil
}

//新增约会
func newFeedYehui(ctx *web.EngineCtx) error {
	req := &reqFeedYehuiAdd{}
	ctx.Verify(req)
	_ = ctx.ForceUID()
	return nil
}

//新增投票
func newFeedVote(ctx *web.EngineCtx) error {
	req := &reqFeedVoteAdd{}
	ctx.Verify(req)
	_ = ctx.ForceUID()
	return nil
}
