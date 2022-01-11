package limitrate_driver

import (
	"context"
	"github.com/ghf-go/nannan/mod"
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
	if mod.GetRedis(t.redisConfName).HExists(context.Background(), "limit:tokens", key).Val() {
		return mod.GetRedis(t.redisConfName).HIncrBy(context.Background(), "limit:tokens", key, -1).Val() > 0
	} else {
		mod.GetRedis(t.redisConfName).HSet(context.Background(), "limit:tokens", key, t.maxReq-1)
	}
	return true
}
func (t *tokenLimitRedisDriver) Start() {
	r := mod.GetRedis(t.redisConfName)
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
