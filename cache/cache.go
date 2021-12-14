package cache

import (
	"github.com/ghf-go/nannan/drivers"
	"github.com/ghf-go/nannan/drivers/cachedriver"
	"github.com/ghf-go/nannan/gconf"
	"github.com/ghf-go/nannan/glog"
)

type CacheDriver interface {
	Del(key ...string) bool
	Set(key, val string) bool
	Get(key string) string
	SetX(key, val string, timeout int) bool
}

var (
	_cacheMap = map[string]CacheDriver{}
)

func GetCache(confName string) CacheDriver {
	if r, ok := _cacheMap[confName]; ok {
		return r
	}
	conf := gconf.GetConf(confName)
	var r CacheDriver
	switch conf.GetScheme() {
	case "memcache":
		r = cachedriver.NewMemCacheCacheDriver(drivers.GetMemcached(conf))
	case "redis":
		r = cachedriver.NewRedisCacheDriver(drivers.GetRedis(conf))
	case "redis_cluster":
		r = cachedriver.NewRedisClusterCacheDriver(drivers.GetRedisCluster(conf))
	case "redis_sentinel":
		r = cachedriver.NewRedisSentinelCacheDriver(drivers.GetRedisSentinel(conf))
	default:
		glog.Error("没有配置缓存")
		panic("没有配置缓存")
	}
	_cacheMap[confName] = r
	return r
}
