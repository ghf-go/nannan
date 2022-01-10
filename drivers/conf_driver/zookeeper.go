package conf_driver

import (
	"github.com/go-zookeeper/zk"
	"strconv"
)

type ZookeeperDriver struct {
	path string
	con  *zk.Conn
}

func NewZookeeperDriver() *ZookeeperDriver {
	return &ZookeeperDriver{}
}
func (c *ZookeeperDriver) Get(key string) string {
	d, _, e := c.con.Get(c.path + key)
	if e != nil {
		return ""
	}
	return string(d)
}
func (c *ZookeeperDriver) All() map[string]string {
	ret := map[string]string{}
	list, _, e := c.con.Children(c.path)
	if e != nil {
		return ret
	}
	for _, v := range list {
		ret[v] = c.Get(v)
	}
	return ret
}
func (c *ZookeeperDriver) GetInt(key string) int64 {
	ret := c.Get(key)
	r, e := strconv.ParseInt(ret, 10, 64)
	if e != nil {
		return 0
	}
	return r
}
func (c *ZookeeperDriver) GetBool(key string) bool {
	ret := c.Get(key)
	r, e := strconv.ParseBool(ret)
	if e != nil {
		return false
	}
	return r
}
func (c *ZookeeperDriver) GetFloat(key string) float64 {
	ret := c.Get(key)
	r, e := strconv.ParseFloat(ret, 10)
	if e != nil {
		return 0.0
	}
	return r
}
func (c *ZookeeperDriver) Del(key string) {
	c.Set(key, "")
}
func (c *ZookeeperDriver) GetConf(key string) Conf {
	return BuildConf(c.Get(key))
}
func (c *ZookeeperDriver) Set(key, val string) {
	if ok, s, _ := c.con.Exists(c.path + key); ok {
		c.con.Set(c.path+key, []byte(val), s.Version+1)
	} else {
		c.con.Create(c.path+key, []byte(val), 1, []zk.ACL{})
	}
}
func (c *ZookeeperDriver) SetInt(key string, val int64) {
	c.Set(key, strconv.FormatInt(val, 10))
}
func (c *ZookeeperDriver) SetFloat(key string, val float64) {
	c.Set(key, strconv.FormatFloat(val, 'E', -1, 64))
}
func (c *ZookeeperDriver) SetBool(key string, val bool) {
	c.Set(key, strconv.FormatBool(val))
}
func (c *ZookeeperDriver) SetConf(key string, val Conf) {
	c.Set(key, val.String())
}
