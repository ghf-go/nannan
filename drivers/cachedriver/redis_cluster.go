package cachedriver

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClusterCacheDriver struct {
	redis *redis.ClusterClient
}

func NewRedisClusterCacheDriver(redis *redis.ClusterClient) *RedisClusterCacheDriver {
	return &RedisClusterCacheDriver{redis: redis}
}

func (r *RedisClusterCacheDriver) Del(key ...string) bool {
	return r.redis.Del(context.Background(), key...).Err() == nil
}
func (r *RedisClusterCacheDriver) Set(key, val string) bool {
	return r.redis.Set(context.Background(), key, val, time.Hour*86400*365).Err() == nil
}
func (r *RedisClusterCacheDriver) Get(key string) string {
	return r.redis.Get(context.Background(), key).Val()
}
func (r *RedisClusterCacheDriver) SetX(key, val string, timeout int) bool {
	return r.redis.Set(context.Background(), key, val, time.Second*time.Duration(timeout)).Err() == nil
}
