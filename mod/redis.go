package mod

import (
	"fmt"
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/gutils"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

var (
	_redisMap     = map[string]*redis.Client{}
	_redisCluster = map[string]*redis.ClusterClient{}
)

// 创建新的redis链接
func NewReids(confKeyName string) *redis.Client {
	conf := GetConf(confKeyName)
	switch strings.ToLower(conf.Scheme) {
	case "redis":
		return newRedis(conf)
	case "sentinel":
		return newSentinel(conf)
	default:
		Error("创建Redis失败 配置信息 %s", conf.Raw)
		gutils.Error(500, "redis配置错误")

	}
	return nil
}

// 获取redis链接
func GetRedis(confKeyName string) *redis.Client {
	if r, ok := _redisMap[confKeyName]; ok {
		return r
	}
	r := NewReids(confKeyName)
	_redisMap[confKeyName] = r
	return r
}

// 创建redis 群集链接
func NewRedisCluster(confKeyName string) *redis.ClusterClient {
	conf := GetConf(confKeyName)
	server := conf.GetArgs("servers")
	if server == "" {
		Error("创建Redis失败 配置信息 %s", conf.Raw)
		gutils.Error(500, "redis 配置错误")
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
	if conf.UserName != "" {
		opt.Username = conf.UserName
	}
	if conf.PassWord != "" {
		opt.Password = conf.PassWord
	}
	r := redis.NewClusterClient(opt)
	return r
}

// 获取redis 群集链接
func GetRedisCluster(confKeyName string) *redis.ClusterClient {
	if r, ok := _redisCluster[confKeyName]; ok {
		return r
	}
	return NewRedisCluster(confKeyName)
}

func newRedis(conf def.Conf) *redis.Client {
	opt := &redis.Options{
		Addr:               fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		DialTimeout:        time.Second * 30,
		ReadTimeout:        time.Second * 30,
		IdleTimeout:        time.Second * 30,
		PoolTimeout:        time.Second * 30,
		WriteTimeout:       time.Second * 30,
		IdleCheckFrequency: time.Second * 30,
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
	if conf.UserName != "" {
		opt.Username = conf.UserName
	}
	if conf.PassWord != "" {
		opt.Password = conf.PassWord
	}
	r := redis.NewClient(opt)
	return r
}
func newSentinel(conf def.Conf) *redis.Client {
	server := conf.GetArgs("servers")
	if server == "" {
		Error("创建Redis失败 配置信息 %s", conf.Raw)
		gutils.Error(500, "redisSentinel 没有配置服务器")
	}
	opt := &redis.FailoverOptions{
		MasterName:         "master",
		SentinelAddrs:      strings.Split(server, ","),
		DialTimeout:        time.Second * 30,
		ReadTimeout:        time.Second * 30,
		IdleTimeout:        time.Second * 30,
		PoolTimeout:        time.Second * 30,
		WriteTimeout:       time.Second * 30,
		IdleCheckFrequency: time.Second * 30,
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
	if conf.UserName != "" {
		opt.Username = conf.UserName
	}
	if conf.PassWord != "" {
		opt.Password = conf.PassWord
	}
	r := redis.NewFailoverClient(opt)
	return r
}
