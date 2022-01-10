package commentlogic

import (
	"context"
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/webbase/logic"
	"mime/multipart"
	"path/filepath"
)

func GetPathByFileKey(fk string) string {
	return logic.GetRedis().Get(context.Background(), redisUploadKey(fk)).String()
}

func SaveFile(uid int64, fk, path string, fh *multipart.FileHeader) {
	logic.GetTable(tb_common_upload).InsertMap(db.Data{
		"user_id":   uid,
		"file_key ": fk,
		"file_name": filepath.Base(fh.Filename),
		"file_size": fh.Size,
		"path":      path,
		"file_mine": fh.Header.Get("Content-Type"),
	})
}
