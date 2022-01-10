package common

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/webbase/logic/commonlogic"
)

func GroupListAction(ctx *web.EngineCtx) error {
	return ctx.JsonSuccess(commonlogic.GetGroupAll())
}
func NewGroupAction(ctx *web.EngineCtx) error {
	req := &reqGroupNew{}
	ctx.Verify(req)
	commonlogic.NewGroup(req.GroupName)
	return ctx.JsonSuccess("OK")
}
