package drivers

import (
	"fmt"
	"github.com/ghf-go/nannan/gconf"
	"github.com/ghf-go/nannan/glog"
	"github.com/bradfitz/gomemcache/memcache"
	"strings"
)

var(
	_memcahceMap = map[string]*memcache.Client{}
)
func GetMemcached(conf gconf.GConf) *memcache.Client {
	if conf.GetScheme() != "memcache"{
		glog.Error("获取配置错误 %s",conf.GetBase())
		panic("Memcache配置类型错误")
	}
	if r ,ok := _memcahceMap[conf.GetBase()];ok{
		return r
	}
	slist := []string{fmt.Sprintf("%s:%d",conf.GetHost(),conf.GetPort())}
	server := conf.GetArgs("servers");
	if server != ""{
		slist = append(slist,strings.Split(server,",")...)
	}

	r := memcache.New(slist...)

	_memcahceMap[conf.GetBase()] = r
	return r
}