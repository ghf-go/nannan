package common

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/smscode"
	"time"
)

//发送短信验证码
func SendSmsCodeAction(engine *web.EngineCtx) error {
	req := &ReqSendSmsCode{}
	engine.Verify(req)
	engine.LimitIP("sendsmscode", 1, time.Minute)
	smscode.SendCode(req.Mobile, req.SendType)
	return engine.JsonSuccess("OK")
}
