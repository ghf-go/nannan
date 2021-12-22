package web

import "github.com/ghf-go/nannan/gerr"

type session map[string]interface{}

func (ses session) SetSession(key string, val interface{}) {
	ses[key] = val
}

func (ses session) GetSession(key string) interface{} {
	if r, ok := ses[key]; ok {
		return r
	}
	return nil
}
func (ses session) UID() int64 {
	r := ses.GetSession("uid")
	if r == nil {
		return 0
	}
	return r.(int64)
}
func (ses session) SetUID(uid int64) {
	ses.SetSession("uid", uid)
}
func (ses session) AdminID() int64 {
	r := ses.GetSession("admin_id")
	if r == nil {
		return 0
	}
	return r.(int64)
}
func (ses session) SetAdminID(uid int64) {
	ses.SetSession("admin_id", uid)
}
func (ses session) ForceUID() int64 {
	uid := ses.UID()
	if uid == 0 {
		gerr.Error(123, "账号没有登录")
	}
	return uid

}
func (ses session) ForceAdminID() int64 {
	uid := ses.AdminID()
	if uid == 0 {
		gerr.Error(123, "账号没有登录")
	}
	return uid

}
