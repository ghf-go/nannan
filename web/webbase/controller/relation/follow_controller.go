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
