package limitrate

import (
	"github.com/ghf-go/nannan/gconf"
	"time"
)

type TokenLimiter interface {
	GetToken(key string) bool
	Start()
}

var _tokenLimiter TokenLimiter

func GetTokenLimiter() TokenLimiter {
	if _tokenLimiter != nil {
		return _tokenLimiter
	}
	conf := gconf.GetConf("limiter.token")
	switch conf.GetScheme() {
	case "redis":
		_tokenLimiter = &tokenLimitRedisDriver{
			redisConfName: conf.GetHost(),
			maxReq:        conf.GetPort(),
			timeWindow:    time.Duration(conf.GetArgInt("time_window")) * time.Second,
		}
	default:
		_tokenLimiter = &tokenLimitMemDriver{
			maxReq:     conf.GetPort(),
			timeWindow: time.Duration(conf.GetArgInt("time_window")) * time.Second,
			data:       map[string]int{},
		}

	}
	go _tokenLimiter.Start()
	return nil
}
