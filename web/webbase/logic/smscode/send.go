package smscode

import (
	"context"
	"fmt"
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase/logic"
	"time"
)

const (
	SMS_TYPE_LOGIN = 1
	SMS_TYPE_BIND  = 2
)

var (
	_rkformat         = "sms:%d:%s"
	_smsTypeFormatMap = map[int]string{}
)

func RegisterSmsTypeFormat(data map[int]string) {
	_smsTypeFormatMap = data
}
func SendCode(mobile string, sendType int) {
	rk := getRedisKey(mobile, sendType)
	ctx := context.Background()
	redis := logic.GetRedis()
	if redis.TTL(ctx, rk).Val() > 540 {
		web.Error(20, "你发送的太快了")
	}
	if format, ok := _smsTypeFormatMap[sendType]; ok {
		code := fmt.Sprintf("%d", time.Now().UnixNano()%100000)
		_ = fmt.Sprintf(format, code)
		redis.Set(ctx, rk, code, time.Second*600)
		return
	}
	web.Error(20, "参数错误")
}

//验证短信验证码
func VerifyCode(mobile, code string, sendType int) bool {
	rk := getRedisKey(mobile, sendType)
	if logic.GetRedis().Get(context.Background(), rk).String() == code {
		logic.GetRedis().Del(context.Background(), rk)
		return true
	}
	return false
}

func getRedisKey(mobile string, sendType int) string {
	return fmt.Sprintf(_rkformat, sendType, mobile)
}
