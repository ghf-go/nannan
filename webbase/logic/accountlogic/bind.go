package accountlogic

import (
	"context"
	"fmt"
	"github.com/ghf-go/nannan/gnet"
	"github.com/ghf-go/nannan/gutils"
	"github.com/ghf-go/nannan/webbase/logic"
	"strconv"
	"time"
)

const (
	redisKeyBindUid    = "u:b:%s"
	redisExpireBindUid = time.Second * 86400 * 365
)

var (
	tb_user_mobile = "tb_user_mobile"
	tb_user_email  = "tb_user_email"
	tb_user_wx     = "tb_user_wx"
)

func GetUidByName(name string) int64 {
	re := logic.GetRedis()
	rk := getbindkey(name)
	v, e := re.Get(context.Background(), rk).Result()
	if e == nil && v != "" {
		uid, e := strconv.ParseInt(v, 10, 64)
		if e == nil {
			return uid
		}
	}
	if gutils.IsMobile(name) {
		uid := GetUidByMobile(name)
		if uid > 0 {
			re.Set(context.Background(), rk, uid, redisExpireBindUid)
			return uid
		}
	} else if gutils.IsEmail(name) {
		uid := GetUidByEmail(name)
		if uid > 0 {
			re.Set(context.Background(), rk, uid, redisExpireBindUid)
			return uid
		}
	} else if gutils.IsWxOpenID(name) {
		uid := GetUidByWxOpenID(name)
		if uid > 0 {
			re.Set(context.Background(), rk, uid, redisExpireBindUid)
			return uid
		}
	}
	return 0
}

func GetUidByMobile(mobile string) int64 {
	um := &modelUserMobile{}
	if logic.CreateQuery(tb_user_mobile).Where("mobile=?", mobile).Frist(um) == nil {
		return um.UserID
	}
	return 0
}

func GetUidByEmail(email string) int64 {
	var um *modelUserEmail
	if logic.CreateQuery(tb_user_wx).Where("email=?", email).Frist(um) == nil {
		return um.UserID
	}
	return 0
}
func GetUidByWxOpenID(openid string) int64 {
	var um *modelUserWx
	if logic.CreateQuery(tb_user_email).Where("open_id=?", openid).Frist(um) == nil {
		return um.UserID
	}
	return 0
}
func BindMobile(uid int64, mobile string) bool {
	tuid := GetUidByMobile(mobile)
	if tuid > 0 {
		logic.CreateQuery(tb_user_mobile).Where("mobile=?", mobile).Delete()
	}
	if tuid == uid {
		return true
	}
	ret := logic.GetTable(tb_user_mobile).InsertMap(map[string]interface{}{
		"user_id": uid,
		"mobile":  mobile,
	}) > 0
	if ret {
		logic.GetRedis().Set(context.Background(), getbindkey(mobile), uid, redisExpireBindUid)
	}
	return ret
}

func BindEmail(uid int64, email string) bool {
	tuid := GetUidByEmail(email)
	if tuid > 0 {
		logic.CreateQuery(tb_user_email).Where("email=?", email).Delete()
	}
	if tuid == uid {
		return true
	}
	ret := logic.GetTable(tb_user_email).InsertMap(map[string]interface{}{
		"user_id": uid,
		"email":   email,
	}) > 0
	if ret {
		logic.GetRedis().Set(context.Background(), getbindkey(email), uid, redisExpireBindUid)
	}
	return ret
}
func BindWx(uid int64, info *gnet.WxUserInfo) bool {
	tuid := GetUidByWxOpenID(info.Openid)
	if tuid > 0 {
		logic.CreateQuery(tb_user_wx).Where("open_id=?", info.Openid).Delete()
	}
	if tuid == uid {
		return true
	}
	ret := logic.GetTable(tb_user_wx).InsertMap(map[string]interface{}{
		"user_id": uid,
		"open_id": info.Openid,
		"avatar":  info.Headimgurl,
		"name":    info.Nickname,
	}) > 0
	if ret {
		logic.GetRedis().Set(context.Background(), getbindkey(info.Openid), uid, redisExpireBindUid)
	}
	return ret
}
func getbindkey(name string) string {
	return fmt.Sprintf(redisKeyBindUid, name)
}
