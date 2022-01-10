package mod

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"strings"
)

var (
	_memcahceMap = map[string]*memcache.Client{}
)

// 创建新的Memcache链接
func NewMemcache(confKeyName string) *memcache.Client {
	conf := GetConf(confKeyName)
	slist := []string{fmt.Sprintf("%s:%d", conf.Host, conf.Port)}
	server := conf.GetArgs("servers")
	if server != "" {
		slist = append(slist, strings.Split(server, ",")...)
	}

	r := memcache.New(slist...)

	_memcahceMap[confKeyName] = r
	return r
}

// 获取memcache链接
func GetMemcache(confKeyName string) *memcache.Client {
	if r, ok := _memcahceMap[confKeyName]; ok {
		return r
	}
	return NewMemcache(confKeyName)
}
