package comment

import (
	"github.com/ghf-go/nannan/web"
	"time"
)

func NewFeedAction(ctx *web.EngineCtx) error {
	req := &reqFeedAdd{}
	ctx.Verify(req)
	_ = ctx.ForceUID()
	ctx.LimitIP("newFeed", 1, time.Minute*3)
	//commentlogic.NewFeed()
	return ctx.JsonSuccess("OK")
}
func FeedList(ctx *web.EngineCtx) error {
	return ctx.JsonSuccess("OK")
}
