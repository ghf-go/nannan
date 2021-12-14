package cachedriver

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type MemCacheCacheDriver struct {
	memcache *memcache.Client
}

func NewMemCacheCacheDriver(memcache *memcache.Client) *MemCacheCacheDriver {
	return &MemCacheCacheDriver{memcache: memcache}
}

func (r *MemCacheCacheDriver) Del(key ...string) bool {
	ret := true
	for _, k := range key {
		if r.memcache.Delete(k) != nil {
			ret = false
		}
	}
	return ret
}
func (r *MemCacheCacheDriver) Set(key, val string) bool {
	return r.memcache.Set(&memcache.Item{Key: key, Value: []byte(val)}) == nil
}
func (r *MemCacheCacheDriver) Get(key string) string {
	row, e := r.memcache.Get(key)
	if e != nil {
		return ""
	}
	return string(row.Value)
}
func (r *MemCacheCacheDriver) SetX(key, val string, timeout int) bool {
	return r.memcache.Set(&memcache.Item{Key: key, Value: []byte(val), Expiration: int32(timeout)}) == nil
}
