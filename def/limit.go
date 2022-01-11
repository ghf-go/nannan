package def

import "time"

type IpLimiterDriver interface {
	Check(groupName, ip string, maxReq int, timeCslip time.Duration) bool
}

type TokenLimiterDriver interface {
	GetToken(key string) bool
	Start()
}
