package limitrate

import (
	"github.com/ghf-go/nannan/gconf"
	"time"
)

var (
	_limitIp IpLimiter
)

type IpLimiter interface {
	Check(groupName, ip string, maxReq int, timeCslip time.Duration) bool
}

func GetIpLimiter() IpLimiter {
	if _limitIp != nil {
		return _limitIp
	}
	conf := gconf.GetConf("limiter:ip")
	switch conf.GetScheme() {
	case "redis":
		_limitIp = &IpLimiterRedisDriver{redisConfName: conf.GetHost()}
	default:
		ipm := &IpLimiterMemDriver{data: map[string]*ipLimiterMemGroup{}}
		go ipm.Start()
		_limitIp = ipm
	}
	return _limitIp
}
