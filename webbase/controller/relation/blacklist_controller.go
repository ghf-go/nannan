package relation

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/webbase/logic/relationlogic"
)

//添加或者删除黑名单
func BackListAction(ctx *web.EngineCtx) error {
	req := &reqBlackList{}
	ctx.Verify(req)
	if req.IsDel {
		if relationlogic.DelBlackList(ctx.ForceUID(), req.TargetUid) {
			return ctx.JsonSuccess("OK")
		} else {
			return ctx.JsonFail(123, "删除黑名单失败")
		}
	} else {
		if relationlogic.AddBlackList(ctx.ForceUID(), req.TargetUid) {
			return ctx.JsonSuccess("OK")
		} else {
			return ctx.JsonFail(123, "添加黑名单失败")
		}
	}
}
