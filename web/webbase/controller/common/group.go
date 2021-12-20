package common

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/commonlogic"
)

func GroupListAction(ctx *web.EngineCtx) error {
	return ctx.JsonSuccess(commonlogic.GetGroupAll())
}
