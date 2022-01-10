package comment

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/webbase/logic/commentlogic"
)

func PraiseAction(ctx *web.EngineCtx) error {
	req := &reqPraise{}
	ctx.Verify(req)
	if req.IsUnPraise {
		return ctx.JsonSuccess(commentlogic.Praise(ctx.ForceUID(), req.TargetID, req.TargetType))
	} else {
		return ctx.JsonSuccess(commentlogic.UnPraise(ctx.ForceUID(), req.TargetID, req.TargetType))
	}
}
