package account

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/accountlogic"
)

//设置用资料
func SetProfileAction(engine *web.EngineCtx) error {
	req := &reqBaseUid{}
	engine.Verify(req)
	if req.Uid == 0 {
		req.Uid = engine.ForceUID()
	}
	return engine.JsonSuccess(accountlogic.GetProfileByUid(req.Uid))
}

//获取用户资料
func GetProfileAction(engine *web.EngineCtx) error {
	return nil
}
