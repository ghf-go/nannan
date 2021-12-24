package account

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/accountlogic"
	"github.com/ghf-go/nannan/web/webbase/logic/smscode"
	"time"
)

func LoginAction(engine *web.EngineCtx) error {
	req := &reqLogin{}
	engine.Verify(req)
	if req.Passws == "" && req.Code == "" {
		return engine.JsonFail(401, "请填写密码或者验证码")
	} else if req.Passws != "" {
		engine.LimitIP("loginname", 1, time.Second*15)
		uid := accountlogic.LoginByPass(req.Name, req.Passws)
		if uid > 0 {
			engine.SetUID(uid)
			return engine.JsonSuccess("登录成功")
		}
		return engine.JsonFail(401, "账号或者密码错误")
	} else if req.Code != "" {
		uid := accountlogic.LoginByMobile(req.Name, req.Code, smscode.SMS_TYPE_LOGIN)
		if uid > 0 {
			engine.SetUID(uid)
			return engine.JsonSuccess("登录成功")
		}
		return engine.JsonFail(401, "账号或者密码错误")
	} else {
		return engine.JsonFail(401, "参数错误")
	}
	return nil
}
func LoginByH5WxAction(engine *web.EngineCtx) error {
	req := &reqLogH5Wx{}
	engine.Verify(req)
	//accountlogic.LoginByWxUserInfo()
	return nil
}
