package accountlogic

import (
	"fmt"
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/gutils"
	"github.com/ghf-go/nannan/webbase/logic"
	"time"
)

const (
	tb_user_pass = "tb_user_passwd"
)

func SetPasswd(uid int64, pass string) bool {
	row := &modelUserPasswd{}
	sign := fmt.Sprintf("%d", time.Now().UnixNano()%10000)
	setData := db.Data{
		"sign":   sign,
		"passwd": buildPass(pass, sign),
	}
	if logic.CreateQuery(tb_user_pass).Where("id=?", uid).Frist(row) != nil && row.ID > 0 {
		return logic.CreateQuery(tb_user_pass).Where("id=?", uid).UpdateMap(setData) > 0
	} else {
		setData["id"] = uid
		return logic.GetTable(tb_user_pass).InsertMap(setData) > 0
	}
}
func CheckPasswd(uid int64, pass string) bool {
	row := &modelUserPasswd{}
	if logic.CreateQuery(tb_user_pass).Where("id=?", uid).Frist(row) != nil {
		return false
	}
	return row.Passwd == buildPass(pass, row.Sign)
}

func buildPass(pass, sign string) string {
	return gutils.MD5String(pass + sign)
}
