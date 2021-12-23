package limitrate

import (
	"context"
	"fmt"
	"github.com/ghf-go/nannan/drivers"
	"time"
)

type IpLimiterRedisDriver struct {
	redisConfName string
}

func (rate *IpLimiterRedisDriver) Check(groupName, ip string, maxReq int, timeWindow time.Duration) bool {
	r := drivers.GetRedisByKey(rate.redisConfName)
	ctx := context.Background()
	rk := fmt.Sprintf("limit:%s:%s", groupName, ip)
	if r.Exists(ctx, rk).Val() > 0 {
		return r.Incr(ctx, rk).Val() < int64(maxReq)
	} else {

		r.SetEX(ctx, rk, 1, timeWindow)
	}
	return true
}
