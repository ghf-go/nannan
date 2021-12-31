package comment

import (
	"github.com/ghf-go/nannan/glog"
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/commentlogic"
	"time"
)

func NewFeedAction(ctx *web.EngineCtx) error {
	ctx.Req.FormValue("123")
	for k, v := range ctx.Req.Form {
		for i, vv := range v {
			glog.Debug("提交的内容 %s -> [%d %s]", k, i, vv)
		}
	}

	ctx.LimitIP("newFeed", 1, time.Minute*3)
	req := &reqFeedAdd{}
	ctx.Verify(req)
	switch req.FeedType {
	case commentlogic.FEED_TYPE_YUEHUI: //约会
		return newFeedYehui(ctx)
	case commentlogic.FEED_TYPE_VOTE: //投票
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
	_ = ctx.ForceUID()
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
