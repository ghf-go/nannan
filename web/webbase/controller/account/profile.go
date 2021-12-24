package account

import (
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic/accountlogic"
)

//设置用资料
func SetProfileAction(engine *web.EngineCtx) error {
	engine.Req.ParseMultipartForm(999999)
	uid := engine.ForceUID()
	pData := map[string]interface{}{}
	for k := range engine.Req.Form {
		pData[k] = engine.Req.FormValue(k)
	}

	accountlogic.SetProfile(uid, pData)
	return engine.JsonSuccess(accountlogic.GetProfileByUid(uid))
}

//获取用户资料
func GetProfileAction(engine *web.EngineCtx) error {
	uid := engine.ForceUID()
	return engine.JsonSuccess(accountlogic.GetProfileByUid(uid))
}
