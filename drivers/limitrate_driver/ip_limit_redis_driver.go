package limitrate_driver

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type IpLimiterRedisDriver struct {
	redisCon *redis.Client
}

func NewIpLimiterRedisDriver(r *redis.Client) *IpLimiterRedisDriver {
	return &IpLimiterRedisDriver{redisCon: r}
}

func (r *IpLimiterRedisDriver) Check(groupName, ip string, maxReq int, timeWindow time.Duration) bool {
	ctx := context.Background()
	rk := fmt.Sprintf("limit:%s:%s", groupName, ip)
	if r.redisCon.Exists(ctx, rk).Val() > 0 {
		return r.redisCon.Incr(ctx, rk).Val() < int64(maxReq)
	} else {
		r.redisCon.SetEX(ctx, rk, 1, timeWindow)
	}
	return true
}
