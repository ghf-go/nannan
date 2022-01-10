package webbase

import (
	"github.com/ghf-go/nannan/web"
	account2 "github.com/ghf-go/nannan/webbase/controller/account"
	comment2 "github.com/ghf-go/nannan/webbase/controller/comment"
	common2 "github.com/ghf-go/nannan/webbase/controller/common"
	relation2 "github.com/ghf-go/nannan/webbase/controller/relation"
)

func RegisterRouter() {
	web.RegisterRouterGroup("/api/account", func(group *web.RouterGroup) {
		group.POST("profile", account2.GetProfileAction)     //获取用户资料
		group.POST("profile_set", account2.SetProfileAction) //设置用户资料
		group.POST("bind_mobile", account2.BindMobileAction) //绑定手机号
		group.POST("bind_email", account2.BindEmailAction)   //绑定邮箱
		group.POST("bind_wx", account2.BindWxAction)         //绑定微信
		group.POST("login", account2.LoginAction)            //登录
		group.POST("login_wx", account2.LoginByH5WxAction)   //微信登录
		group.POST("set_pass", account2.SetPassAction)       //设置密码
	})
	web.RegisterRouterGroup("/api/relation", func(group *web.RouterGroup) {
		group.POST("follow", relation2.FollowAction)            //关注，取消关注
		group.POST("backlist", relation2.BackListAction)        //添加，取消黑名单
		group.POST("apply_friend", relation2.ApplyFriendAction) //申请添加好友
		group.POST("audit_friend", relation2.AuditFriendAction) //确认或者删除好友
	})
	web.RegisterRouterGroup("/api/msg", func(group *web.RouterGroup) {

	})
	web.RegisterRouterGroup("/api/comment", func(group *web.RouterGroup) {
		group.POST("/comment/new", comment2.NewCommentAction)   //发布评论
		group.POST("/comment/list", comment2.CommentListAction) //评论列表
		group.POST("/praise", comment2.PraiseAction)            //赞，取消点赞
		group.POST("/feed/create", comment2.NewFeedAction)      //发布动态
		group.POST("/list_feed", comment2.FeedList)             //动态列表
		group.POST("/upload", comment2.UploadFileAction)
	})
	web.RegisterRouterGroup("/api/common", func(group *web.RouterGroup) {
		group.POST("send_sms_code", common2.SendSmsCodeAction) //发送短信验证码
		group.POST("/tags", common2.TagListAction)             //标签列表
		group.POST("/tag/new", common2.NewTagAction)           //标签添加
		group.POST("/groups", common2.GroupListAction)         //分组列表
		group.POST("/group/new", common2.NewGroupAction)       //添加分组
		group.POST("/confs", common2.ConfListAction)           //配置列表
		group.POST("/conf/new", common2.NewConfAction)         //配置添加

	})

}
