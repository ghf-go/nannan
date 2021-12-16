package common

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/smscode"
)
//发送短信验证码
func SendSmsCodeAction(engine *web.EngineCtx) error  {
	req := &ReqSendSmsCode{}
	engine.Verify(req)
	smscode.SendCode(req.Mobile,req.SendType)
	return engine.JsonSuccess("OK")
}
