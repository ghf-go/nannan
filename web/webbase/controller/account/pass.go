package account

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/accountlogic"
)

func SetPassAction(engine *web.EngineCtx) error {
	uid := engine.ForceUID()
	req := &reqSetPass{}
	engine.Verify(req)
	if accountlogic.SetPasswd(uid, req.Pass) {
		return engine.JsonSuccess("OK")
	}
	return engine.JsonFail(123, "设置失败")
}
