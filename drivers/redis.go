package drivers

import (
	"fmt"
	"github.com/ghf-go/nannan/gconf"
	"github.com/ghf-go/nannan/glog"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
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

	opt := &redis.Options{
		Addr:               fmt.Sprintf("%s:%d", conf.GetHost(), conf.GetPort()),
		DialTimeout:        time.Second * 30,
		ReadTimeout:        time.Second * 30,
		IdleTimeout:        time.Second * 30,
		PoolTimeout:        time.Second * 30,
		WriteTimeout:       time.Second * 30,
		IdleCheckFrequency: time.Second * 3,
		PoolSize:           5,
		MinIdleConns:       1,
		MaxConnAge:         time.Second * 45,
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
	return r
}

//获取redis集群
func GetRedisCluster(conf gconf.GConf) *redis.ClusterClient {
	if conf.GetScheme() != "redis_cluster" {
		glog.Error("获取配置错误 %s", conf.GetBase())
		panic("redis配置类型错误")
	}

	server := conf.GetArgs("servers")
	if server == "" {
		glog.Error("获取配置错误,没有配置服务器 %s", conf.GetBase())
		panic("redis配置类型错误")
	}
	opt := &redis.ClusterOptions{
		Addrs:              strings.Split(server, ","),
		DialTimeout:        time.Second * 30,
		ReadTimeout:        time.Second * 30,
		IdleTimeout:        time.Second * 30,
		PoolTimeout:        time.Second * 30,
		WriteTimeout:       time.Second * 30,
		IdleCheckFrequency: time.Second * 3,
		PoolSize:           5,
		MinIdleConns:       1,
		MaxConnAge:         time.Second * 45,
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
	return r
}

//获取redis哨兵
func GetRedisSentinel(conf gconf.GConf) *redis.Client {
	if conf.GetScheme() != "redis_sentinel" {
		glog.Error("获取配置错误 %s", conf.GetBase())
		panic("redis配置类型错误")
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
	return r
}
