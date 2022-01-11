package limitrate_driver

import (
	"github.com/ghf-go/nannan/mod"
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
	conf := mod.GetConf("limiter.token")
	switch conf.Scheme {
	case "redis":
		_tokenLimiter = &tokenLimitRedisDriver{
			redisConfName: conf.Host,
			maxReq:        conf.Port,
			timeWindow:    time.Duration(conf.GetArgInt("time_window")) * time.Second,
		}
	default:
		_tokenLimiter = &tokenLimitMemDriver{
			maxReq:     conf.Port,
			timeWindow: time.Duration(conf.GetArgInt("time_window")) * time.Second,
			data:       map[string]int{},
		}

	}
	go _tokenLimiter.Start()
	return nil
}
