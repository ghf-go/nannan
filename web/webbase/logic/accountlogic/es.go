package accountlogic

import (
	"github.com/ghf-go/nannan/glog"
	"github.com/ghf-go/nannan/gutils"
	"github.com/ghf-go/nannan/web/webbase/logic"
	"strconv"
)

const (
	_es_tb_user = "user"
)

//保存数据到es
func EsSaveUserInfo(uinfo *UserInfo) bool {
	r, e := logic.GetEsClient().Update(_es_tb_user, strconv.FormatInt(uinfo.UserId, 10), uinfo)
	if e != nil {
		glog.Error("ES 更新 %s %s 失败 %s ", _es_tb_user, uinfo.UserId, e.Error())
		return false
	}
	return r.IsSuccess()
}

//从es中查找用户信息
func EsFindUserInfo(uid int64) *UserInfo {
	r := &esFindUser{}
	e := logic.GetEsClient().Find(_es_tb_user, strconv.FormatInt(uid, 10), r)
	if e != nil {
		glog.AppDebug("ES 查询 %s 失败 %s", _es_tb_user, e.Error())
		return nil
	}
	if r.Found {
		return r.Data
	}
	return nil
}

// Es 删除账号信息
func EsDelUserInfo(uid int64) bool {
	r, e := logic.GetEsClient().Delete(_es_tb_user, strconv.FormatInt(uid, 10))
	if e != nil {
		glog.Error("ES 删除 %s  %d 失败 %s ", _es_tb_user, uid, e.Error())
		return false
	}
	return r.IsSuccess()
}

//批量获取用户信息
func EsFetchUserInfoByIds(uids ...int64) map[int64]*UserInfo {
	ret := map[int64]*UserInfo{}
	rep := &esMgetUser{}
	e := logic.GetEsClient().MGet(_es_tb_user, rep, gutils.SlicInt64String(uids)...)
	if e != nil {
		glog.Error("ES MGET %s 错误 %s", _es_tb_user, e.Error())
		return ret
	}
	if rep.ErrorCode != 0 {
		return ret
	}
	for _, val := range rep.Docs {
		if val.Found {
			ret[val.Data.UserId] = val.Data
		}
	}
	return ret
}
