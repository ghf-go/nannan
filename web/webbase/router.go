package webbase

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/controller/account"
	"github.com/ghf-go/nannan/web/webbase/controller/comment"
	"github.com/ghf-go/nannan/web/webbase/controller/common"
	"github.com/ghf-go/nannan/web/webbase/controller/relation"
)

func RegisterRouter() {
	web.RegisterRouterGroup("/api/account", func(group *web.RouterGroup) {
		group.POST("profile", account.GetProfileAction)     //获取用户资料
		group.POST("profile_set", account.SetProfileAction) //设置用户资料
		group.POST("bind_mobile", account.BindMobileAction) //绑定手机号
		group.POST("bind_email", account.BindEmailAction)   //绑定邮箱
		group.POST("bind_wx", account.BindWxAction)         //绑定微信
		group.POST("login", account.LoginAction)            //登录
		group.POST("login_wx", account.LoginByH5WxAction)   //微信登录
		group.POST("set_pass", account.SetPassAction)       //设置密码
	})
	web.RegisterRouterGroup("/api/relation", func(group *web.RouterGroup) {
		group.POST("follow", relation.FollowAction)            //关注，取消关注
		group.POST("backlist", relation.BackListAction)        //添加，取消黑名单
		group.POST("apply_friend", relation.ApplyFriendAction) //申请添加好友
		group.POST("audit_friend", relation.AuditFriendAction) //确认或者删除好友
	})
	web.RegisterRouterGroup("/api/msg", func(group *web.RouterGroup) {

	})
	web.RegisterRouterGroup("/api/comment", func(group *web.RouterGroup) {
		group.POST("/create_comment", comment.NewCommentAction) //发布评论
		group.POST("/list_comment", comment.CommentListAction)  //评论列表
		group.POST("/praise", comment.PraiseAction)             //赞，取消点赞
		group.POST("/create_feed", comment.NewFeedAction)       //发布动态
		group.POST("/list_feed", comment.FeedList)              //动态列表
	})
	web.RegisterRouterGroup("/api/common", func(group *web.RouterGroup) {
		group.POST("send_sms_code", common.SendSmsCodeAction)
	})

}
