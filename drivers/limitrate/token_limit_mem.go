package limitrate

import (
	"context"
	"github.com/ghf-go/nannan/drivers"
	"sync"
	"time"
)

type tokenLimitRedisDriver struct {
	sync.Mutex
	maxReq        int
	timeWindow    time.Duration
	redisConfName string
}

func (t *tokenLimitRedisDriver) GetToken(key string) bool {
	if drivers.GetRedisByKey(t.redisConfName).HExists(context.Background(), "limit:tokens", key).Val() {
		return drivers.GetRedisByKey(t.redisConfName).HIncrBy(context.Background(), "limit:tokens", key, -1).Val() > 0
	} else {
		drivers.GetRedisByKey(t.redisConfName).HSet(context.Background(), "limit:tokens", key, t.maxReq-1)
	}
	return true
}
func (t *tokenLimitRedisDriver) Start() {
	r := drivers.GetRedisByKey(t.redisConfName)
	ctr := context.Background()
	rk := "limit:tokens"
	d := r.HGetAll(ctr, rk).Val()
	args := []interface{}{}
	for k, _ := range d {
		args = append(args, k, t.maxReq)
	}
	r.HSet(ctr, rk, args...)
	time.Sleep(t.timeWindow)
}
