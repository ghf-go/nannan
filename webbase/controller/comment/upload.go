package comment

import (
	"github.com/ghf-go/nannan/gutils"
	"github.com/ghf-go/nannan/secret"
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/webbase/logic/commentlogic"
	"path/filepath"
)

func UploadFileAction(ctx *web.EngineCtx) error {
	f, fh, e := ctx.Req.FormFile("file")
	if e != nil {
		return ctx.JsonFail(500, e.Error())
	}
	fk := secret.MD5HttpFile(f)
	path := commentlogic.GetPathByFileKey(fk)
	if path != "" {
		return ctx.JsonSuccess(map[string]interface{}{
			"path": path,
			"url":  gutils.GetNetStore("default").CdnUrl(path),
		})
	} else {
		path := fk[:2] + "/" + fk[2:2] + "/" + fk[4:] + filepath.Ext(fh.Filename)
		if gutils.GetNetStore("default").UploadReader(path, f) != nil {
			return ctx.JsonFail(500, "上传失败")
		} else {
			commentlogic.SaveFile(ctx.UID(), fk, path, fh)
			return ctx.JsonSuccess(map[string]interface{}{
				"path": path,
				"url":  gutils.GetNetStore("default").CdnUrl(path),
			})
		}
	}
}
