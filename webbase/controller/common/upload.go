package common

import (
	"github.com/ghf-go/nannan/gutils"
	"github.com/ghf-go/nannan/mod"
	"github.com/ghf-go/nannan/web"
)

func UploadVideo(ctx *web.EngineCtx) error {
	return nil
}
func UploadImg(ctx *web.EngineCtx) error {
	f, h, e := ctx.Req.FormFile("img")
	if e != nil {
		gutils.Error(500, "上传文件失败")
		return e
	}
	store := mod.GetNetStore("static")
	//store.UploadFile()
	return nil
}
