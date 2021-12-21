package comment

import "github.com/ghf-go/nannan/web"

func NewFeedAction(ctx *web.EngineCtx) error {
	req := &reqFeedAdd{}
	ctx.Verify(req)

	return ctx.JsonSuccess("OK")
}
func FeedList(ctx *web.EngineCtx) error {
	return ctx.JsonSuccess("OK")
}
