package limitrate_driver

import (
	"context"
	"github.com/ghf-go/nannan/def"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type TokenLimitRedisDriver struct {
	sync.Mutex
	maxReq     int
	timeWindow time.Duration
	redisCon   *redis.Client
}

func NewTokenLimitRedisDriver(conf def.Conf, r *redis.Client) *TokenLimitRedisDriver {
	return &TokenLimitRedisDriver{
		redisCon:   r,
		maxReq:     conf.GetArgInt("max_req"),
		timeWindow: time.Duration(conf.GetArgInt("time_window")) * time.Second,
	}
}
func (t *TokenLimitRedisDriver) GetToken(key string) bool {
	if t.redisCon.HExists(context.Background(), "limit:tokens", key).Val() {
		return t.redisCon.HIncrBy(context.Background(), "limit:tokens", key, -1).Val() > 0
	} else {
		t.redisCon.HSet(context.Background(), "limit:tokens", key, t.maxReq-1)
	}
	return true
}
func (t *TokenLimitRedisDriver) Start() {
	for def.IsRun() {
		ctr := context.Background()
		rk := "limit:tokens"
		d := t.redisCon.HGetAll(ctr, rk).Val()
		args := []interface{}{}
		for k := range d {
			args = append(args, k, t.maxReq)
		}
		t.redisCon.HSet(ctr, rk, args...)
		time.Sleep(t.timeWindow)
	}

}
