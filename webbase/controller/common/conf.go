package common

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/webbase/logic/commonlogic"
)

func ConfListAction(ctx *web.EngineCtx) error {
	req := &reqBaseGroup{}
	ctx.Verify(req)
	return ctx.JsonSuccess(commonlogic.GetConfByGroupID(req.GroupID))
}

func NewConfAction(ctx *web.EngineCtx) error {
	req := &reqNewConf{}
	ctx.Verify(req)
	commonlogic.NewConf(req.GroupID, req.ValType, req.Key, req.Desc, req.Val)
	return ctx.JsonSuccess("OK")
}
