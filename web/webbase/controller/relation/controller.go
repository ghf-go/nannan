package relation

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/relationlogic"
)

//关注或者取消关注
func FollowAction(ctx *web.EngineCtx) error {
	req := &reqFollow{}
	ctx.Verify(req)
	if req.IsDel {
		if relationlogic.UnFollow(ctx.ForceUID(), req.TargetUid) {
			return ctx.JsonSuccess(relationlogic.GetFanCount(req.TargetUid))
		} else {
			return ctx.JsonFail(123, "取消关注失败")
		}
	} else {
		if relationlogic.Follow(ctx.ForceUID(), req.TargetUid) {
			return ctx.JsonSuccess(relationlogic.GetFanCount(req.TargetUid))
		} else {
			return ctx.JsonFail(123, "关注失败")
		}
	}
}

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
