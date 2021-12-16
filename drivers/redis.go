package drivers

import (
	"fmt"
	"github.com/ghf-go/nannan/gconf"
	"github.com/ghf-go/nannan/glog"
	"github.com/go-redis/redis/v8"
	"strings"
)

var (
	_redisMap         = map[string]*redis.Client{}
	_redisClusterMap  = map[string]*redis.ClusterClient{}
	_redisSentinelMap = map[string]*redis.Client{}
)

func GetRedisByKey(confName string) *redis.Client {
	conf := gconf.GetConf("cache." + confName)
	switch conf.GetScheme() {
	case "redis":
		return GetRedis(conf)
	case "redis_sentinel":
		return GetRedisSentinel(conf)
	}
	panic("redis参数错误")
}

// 获取redis
func GetRedis(conf gconf.GConf) *redis.Client {
	if conf.GetScheme() != "redis" {
		glog.Error("获取配置错误 %s", conf.GetBase())
		panic("redis配置类型错误")
	}
	if r, ok := _redisMap[conf.GetBase()]; ok {
		return r
	}
	opt := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", conf.GetHost(), conf.GetPort()),
	}
	if conf.GetArgInt("db") > 0 {
		opt.DB = conf.GetArgInt("db")
	}

	if conf.GetArgInt("retries") > 0 {
		opt.MaxRetries = conf.GetArgInt("retries")
	}
	if conf.GetUserName() != "" {
		opt.Username = conf.GetUserName()
	}
	if conf.GetPassWord() != "" {
		opt.Password = conf.GetPassWord()
	}
	r := redis.NewClient(opt)
	_redisMap[conf.GetBase()] = r
	return r
}

//获取redis集群
func GetRedisCluster(conf gconf.GConf) *redis.ClusterClient {
	if conf.GetScheme() != "redis_cluster" {
		glog.Error("获取配置错误 %s", conf.GetBase())
		panic("redis配置类型错误")
	}
	if r, ok := _redisClusterMap[conf.GetBase()]; ok {
		return r
	}
	server := conf.GetArgs("servers")
	if server == "" {
		glog.Error("获取配置错误,没有配置服务器 %s", conf.GetBase())
		panic("redis配置类型错误")
	}
	opt := &redis.ClusterOptions{
		Addrs: strings.Split(server, ","),
	}

	if conf.GetArgInt("retries") > 0 {
		opt.MaxRetries = conf.GetArgInt("retries")
	}
	if conf.GetUserName() != "" {
		opt.Username = conf.GetUserName()
	}
	if conf.GetPassWord() != "" {
		opt.Password = conf.GetPassWord()
	}
	r := redis.NewClusterClient(opt)
	_redisClusterMap[conf.GetBase()] = r
	return r
}

//获取redis哨兵
func GetRedisSentinel(conf gconf.GConf) *redis.Client {
	if conf.GetScheme() != "redis_sentinel" {
		glog.Error("获取配置错误 %s", conf.GetBase())
		panic("redis配置类型错误")
	}
	if r, ok := _redisMap[conf.GetBase()]; ok {
		return r
	}
	server := conf.GetArgs("servers")
	if server == "" {
		glog.Error("获取配置错误,没有配置服务器 %s", conf.GetBase())
		panic("redis配置类型错误")
	}
	opt := &redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: strings.Split(server, ","),
	}
	if conf.GetArgInt("db") > 0 {
		opt.DB = conf.GetArgInt("db")
	}

	if conf.GetArgInt("retries") > 0 {
		opt.MaxRetries = conf.GetArgInt("retries")
	}
	if conf.GetUserName() != "" {
		opt.Username = conf.GetUserName()
	}
	if conf.GetPassWord() != "" {
		opt.Password = conf.GetPassWord()
	}
	r := redis.NewFailoverClient(opt)
	_redisSentinelMap[conf.GetBase()] = r
	return r
}
