package accountlogic

import (
	"github.com/ghf-go/nannan/gutils"
	"github.com/ghf-go/nannan/mod"
	"github.com/ghf-go/nannan/webbase/logic"
	"strconv"
)

const (
	_es_tb_user = "user"
)

//保存数据到es
func EsSaveUserInfo(uinfo *UserInfo) bool {
	return logic.GetEsClient().Update(_es_tb_user, strconv.FormatInt(uinfo.UserId, 10), uinfo)
}

//从es中查找用户信息
func EsFindUserInfo(uid int64) *UserInfo {
	r := &UserInfo{}
	e := logic.GetEsClient().Find(_es_tb_user, strconv.FormatInt(uid, 10), r)
	if e != nil {
		mod.Debug("ES 查询 %s 失败 %s", _es_tb_user, e.Error())
		return nil
	}
	return r
}

// Es 删除账号信息
func EsDelUserInfo(uid int64) bool {
	return logic.GetEsClient().Delete(_es_tb_user, strconv.FormatInt(uid, 10))
}

//批量获取用户信息
func EsFetchUserInfoByIds(uids ...int64) map[int64]*UserInfo {
	r := map[int64]*UserInfo{}
	e := logic.GetEsClient().MGet(_es_tb_user, r, gutils.SlicInt64String(uids)...)
	if e != nil {
		mod.Error("ES MGET %s 错误 %s", _es_tb_user, e.Error())
		return r
	}
	return r
}
