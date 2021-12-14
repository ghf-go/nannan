package cachedriver

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisSentinelCacheDriver struct {
	redis *redis.Client
}
func NewRedisSentinelCacheDriver(redis *redis.Client) *RedisSentinelCacheDriver {
	return &RedisSentinelCacheDriver{redis: redis}
}

func (r *RedisSentinelCacheDriver) Del(key ...string) bool {
	return r.redis.Del(context.Background(),key...).Err() == nil
}
func (r *RedisSentinelCacheDriver) Set(key, val string) bool {
	return r.redis.Set(context.Background(),key,val,time.Hour * 86400 * 365).Err() == nil
}
func (r *RedisSentinelCacheDriver) Get(key string) string {
	return r.redis.Get(context.Background(),key).Val()
}
func (r *RedisSentinelCacheDriver) SetX(key, val string, timeout int) bool{
	return r.redis.Set(context.Background(),key,val,time.Second * time.Duration(timeout)).Err() == nil
}