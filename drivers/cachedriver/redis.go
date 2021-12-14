package cachedriver

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCacheDriver struct {
	redis *redis.Client
}

func NewRedisCacheDriver(redis *redis.Client) *RedisCacheDriver {
	return &RedisCacheDriver{redis: redis}
}

func (r *RedisCacheDriver) Del(key ...string) bool {
	return r.redis.Del(context.Background(), key...).Err() == nil
}
func (r *RedisCacheDriver) Set(key, val string) bool {
	return r.redis.Set(context.Background(), key, val, time.Hour*86400*365).Err() == nil
}
func (r *RedisCacheDriver) Get(key string) string {
	return r.redis.Get(context.Background(), key).Val()
}
func (r *RedisCacheDriver) SetX(key, val string, timeout int) bool {
	return r.redis.Set(context.Background(), key, val, time.Second*time.Duration(timeout)).Err() == nil
}
