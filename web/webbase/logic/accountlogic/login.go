package accountlogic

import (
	"github.com/ghf-go/nannan/gerr"
	"github.com/ghf-go/nannan/gnet"
)

func LoginByMobile(mobile, code string, smstype int) int64 {
	//if !smscode.VerifyCode(mobile, code, smstype) {
	//	gutils.Error(401, "验证码错误")
	//}
	uid := GetUidByMobile(mobile)
	if uid > 0 {
		return uid
	}
	uid = NewUserID()
	BindMobile(uid, mobile)
	return uid
}
func LoginByPass(name, pass string) int64 {
	uid := GetUidByName(name)
	if uid == 0 {
		gutils.Error(401, "账号或密码错误")
	}
	if !CheckPasswd(uid, pass) {
		gutils.Error(401, "账号或密码错误")
	}
	return uid
}

func LoginByWxUserInfo(info *gnet.WxUserInfo) int64 {
	uid := GetUidByWxOpenID(info.Openid)
	if uid > 0 {
		return uid
	}
	uid = NewUserID()
	BindWx(uid, info)
	return uid
}
