package smscode

import (
	"context"
	"fmt"
	"github.com/ghf-go/nannan/gerr"
	"github.com/ghf-go/nannan/web/webbase/logic"
	"time"
)

const (
	SMS_TYPE_LOGIN = 1
	SMS_TYPE_BIND  = 2
)

var (
	_rkformat         = "sms:%d:%s"
	_smsTypeFormatMap = map[int]string{
		1: "321",
		2: "432",
	}
)

func RegisterSmsTypeFormat(data map[int]string) {
	_smsTypeFormatMap = data
}
func SendCode(mobile string, sendType int) {
	rk := getRedisKey(mobile, sendType)
	ctx := context.Background()
	redis := logic.GetRedis()
	if redis.TTL(ctx, rk).Val() > 540 {
		gutils.Error(401, "你发送的太快了")
	}
	if format, ok := _smsTypeFormatMap[sendType]; ok {
		code := fmt.Sprintf("%d", time.Now().UnixNano()%100000)
		_ = fmt.Sprintf(format, code)

		redis.Set(ctx, rk, code, time.Second*600).Result()

		return
	}
	gutils.Error(401, "参数错误")
}

//验证短信验证码
func VerifyCode(mobile, code string, sendType int) bool {
	rk := getRedisKey(mobile, sendType)
	//glog.Debug("m : %s c: (%s)  r:(%s)", mobile, code, logic.GetRedis().Get(context.Background(), rk).Val())
	if logic.GetRedis().Get(context.Background(), rk).Val() == code {
		logic.GetRedis().Del(context.Background(), rk)
		return true
	}
	return false
}

func getRedisKey(mobile string, sendType int) string {
	return fmt.Sprintf(_rkformat, sendType, mobile)
}
