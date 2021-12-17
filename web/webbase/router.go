package webbase

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/controller/account"
	"github.com/ghf-go/nannan/web/webbase/controller/common"
	"github.com/ghf-go/nannan/web/webbase/controller/relation"
)

func RegisterRouter() {
	web.RegisterRouterGroup("/api/account", func(group *web.RouterGroup) {
		group.POST("profile", account.GetProfileAction)
		group.POST("profile_set", account.SetProfileAction)
		group.POST("bind_mobile", account.BindMobileAction)
		group.POST("bind_email", account.BindEmailAction)
		group.POST("bind_wx", account.BindWxAction)
		group.POST("login", account.LoginAction)
		group.POST("login_wx", account.LoginByH5WxAction)
	})
	web.RegisterRouterGroup("/api/relation", func(group *web.RouterGroup) {
		group.POST("follow", relation.FollowAction)
		group.POST("backlist", relation.BackListAction)
		group.POST("apply_friend", relation.ApplyFriendAction)
		group.POST("audit_friend", relation.AuditFriendAction)
	})
	web.RegisterRouterGroup("/api/msg", func(group *web.RouterGroup) {

	})
	web.RegisterRouterGroup("/api/comment", func(group *web.RouterGroup) {

	})
	web.RegisterRouterGroup("/api/common", func(group *web.RouterGroup) {
		group.POST("send_sms_code", common.SendSmsCodeAction)
	})

}
