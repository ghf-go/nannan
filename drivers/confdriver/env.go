package confdriver

import (
	"os"
	"strconv"
	"strings"
)

type EnvDriver string

func NewEnvDriver() EnvDriver {
	return ""
}
func (c EnvDriver) Get(key string) string {
	return os.Getenv(key)
}
func (c EnvDriver) All() map[string]string {
	list := os.Environ()
	ret := map[string]string{}
	for _, v := range list {
		i := strings.Index(v, "=")
		ret[v[:i]] = v[i+1:]
	}
	return ret
}
func (c EnvDriver) GetInt(key string) int64 {
	ret := c.Get(key)
	r, e := strconv.ParseInt(ret, 10, 64)
	if e != nil {
		return 0
	}
	return r
}
func (c EnvDriver) GetBool(key string) bool {
	ret := c.Get(key)
	r, e := strconv.ParseBool(ret)
	if e != nil {
		return false
	}
	return r
}
func (c EnvDriver) GetFloat(key string) float64 {
	ret := c.Get(key)
	r, e := strconv.ParseFloat(ret, 10)
	if e != nil {
		return 0.0
	}
	return r
}
func (c EnvDriver) Del(key string) {
	c.Set(key, "")
}
func (c EnvDriver) GetConf(key string) Conf {
	return BuildConf(c.Get(key))
}
func (c EnvDriver) Set(key, val string) {
	os.Setenv(key, val)
}
func (c EnvDriver) SetInt(key string, val int64) {
	c.Set(key, strconv.FormatInt(val, 10))
}
func (c EnvDriver) SetFloat(key string, val float64) {
	c.Set(key, strconv.FormatFloat(val, 'E', -1, 64))
}
func (c EnvDriver) SetBool(key string, val bool) {
	c.Set(key, strconv.FormatBool(val))
}
func (c EnvDriver) SetConf(key string, val Conf) {
	c.Set(key, val.String())
}
