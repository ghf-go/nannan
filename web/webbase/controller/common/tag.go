package common

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/commonlogic"
)

func TagListAction(ctx *web.EngineCtx) error {
	req := &reqBaseGroup{}
	ctx.Verify(req)
	return ctx.JsonSuccess(commonlogic.GetAllTagByGroupID(req.GroupID))
}
