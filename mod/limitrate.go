package mod

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/drivers/limitrate_driver"
)

func GetLimitByIP() def.IpLimiterDriver {
	if _limit_ip != nil {
		return _limit_ip
	}
	conf := GetConf("limit.ip")
	switch conf.Scheme {
	case "redis":
		_limit_ip = limitrate_driver.NewIpLimiterRedisDriver(newRedis(conf))
	default:
		_limit_ip = limitrate_driver.NewIpLimiterMemDriver(conf)
	}
	return _limit_ip
}
func GetLimitByToken() def.TokenLimiterDriver {
	if _limit_token != nil {
		return _limit_token
	}
	conf := GetConf("limiter.token")
	switch conf.Scheme {
	case "redis":
		_limit_token = limitrate_driver.NewTokenLimitRedisDriver(conf, NewReids(conf.Host))
	default:
		_limit_token = limitrate_driver.NewTokenLimitMemDriver(conf)
	}
	go _limit_token.Start()
	return _limit_token
}
