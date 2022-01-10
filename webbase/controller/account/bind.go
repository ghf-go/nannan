package account

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/webbase/logic/accountlogic"
	"github.com/ghf-go/nannan/webbase/logic/smscode"
)

func BindMobileAction(engine *web.EngineCtx) error {
	uid := engine.ForceUID()
	req := &reqBindMobile{}
	engine.Verify(req)
	if smscode.VerifyCode(req.Mobile, req.Code, smscode.SMS_TYPE_BIND) {
		if accountlogic.BindMobile(uid, req.Mobile) {
			return engine.JsonSuccess("OK")
		} else {
			return engine.JsonFail(123, "绑定失败")
		}
	} else {
		return engine.JsonFail(123, "验证码错误")
	}
}
func BindEmailAction(engine *web.EngineCtx) error {
	return nil
}
func BindWxAction(engine *web.EngineCtx) error {
	return nil
}
