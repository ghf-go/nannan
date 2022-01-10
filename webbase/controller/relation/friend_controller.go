package relation

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/webbase/logic/relationlogic"
)

//申请添加好友
func ApplyFriendAction(ctx *web.EngineCtx) error {
	req := &reqFriendApply{}
	ctx.Verify(req)
	if relationlogic.ApplyFriend(ctx.ForceUID(), req.TargetUid, req.Msg) {
		return ctx.JsonSuccess(relationlogic.StatFriend(ctx.ForceUID(), req.TargetUid))
	}
	return ctx.JsonFail(123, "申请失败")
}

//确实好友
func AuditFriendAction(ctx *web.EngineCtx) error {
	req := &reqFriendAudit{}
	ctx.Verify(req)
	if relationlogic.AuditFriend(ctx.ForceUID(), req.TargetUid, req.IsDel) {
		return ctx.JsonSuccess(relationlogic.StatFriend(ctx.ForceUID(), req.TargetUid))
	}
	return ctx.JsonFail(123, "处理失败")
}
