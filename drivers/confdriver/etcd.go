package confdriver

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"strconv"
	"strings"
)

type EtcdDriver struct {
	client *clientv3.Client
	ctx    context.Context
	conf   clientv3.Config
}

func NewEtcdDriver(config clientv3.Config) *EtcdDriver {
	client, e := clientv3.New(config)
	if e != nil {
		panic(e)
	}
	return &EtcdDriver{
		client: client,
		ctx:    context.Background(),
		conf:   config,
	}
}
func NewEtcdDriverByConf(config Conf) *EtcdDriver {
	ec := clientv3.Config{}
	client, e := clientv3.New(ec)
	if e != nil {
		panic(e)
	}
	return &EtcdDriver{
		client: client,
		ctx:    context.Background(),
		conf:   ec,
	}
}
func (c *EtcdDriver) Get(key string) string {
	r, e := c.client.Get(c.ctx, key)
	if e != nil {
		return ""
	}
	return string(r.Kvs[0].Value)
}
func (c *EtcdDriver) All() map[string]string {
	list := os.Environ()
	ret := map[string]string{}
	for _, v := range list {
		i := strings.Index(v, "=")
		ret[v[:i]] = v[i+1:]
	}
	return ret
}
func (c *EtcdDriver) GetInt(key string) int64 {
	ret := c.Get(key)
	r, e := strconv.ParseInt(ret, 10, 64)
	if e != nil {
		return 0
	}
	return r
}
func (c *EtcdDriver) GetBool(key string) bool {
	ret := c.Get(key)
	r, e := strconv.ParseBool(ret)
	if e != nil {
		return false
	}
	return r
}
func (c *EtcdDriver) GetFloat(key string) float64 {
	ret := c.Get(key)
	r, e := strconv.ParseFloat(ret, 10)
	if e != nil {
		return 0.0
	}
	return r
}
func (c *EtcdDriver) Del(key string) {
	c.Set(key, "")
}
func (c *EtcdDriver) GetConf(key string) Conf {
	return BuildConf(c.Get(key))
}
func (c *EtcdDriver) Set(key, val string) {
	c.client.Put(c.ctx, key, val)
	c.client.Sync(c.ctx)
}
func (c *EtcdDriver) SetInt(key string, val int64) {
	c.Set(key, strconv.FormatInt(val, 10))
}
func (c *EtcdDriver) SetFloat(key string, val float64) {
	c.Set(key, strconv.FormatFloat(val, 'E', -1, 64))
}
func (c *EtcdDriver) SetBool(key string, val bool) {
	c.Set(key, strconv.FormatBool(val))
}
func (c *EtcdDriver) SetConf(key string, val Conf) {
	c.Set(key, val.String())
}
