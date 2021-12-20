package common

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/commonlogic"
)

func ConfListAction(ctx *web.EngineCtx) error {
	req := &reqBaseGroup{}
	ctx.Verify(req)
	return ctx.JsonSuccess(commonlogic.GetConfByGroupID(req.GroupID))
}
