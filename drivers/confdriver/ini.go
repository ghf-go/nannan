package confdriver

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strconv"
	"strings"
)

type IniDriver struct {
	filePath string
	data     map[string]string
}

func NewIniDriver(filePath string) *IniDriver {
	return &IniDriver{
		filePath: filePath,
		data:     paseriniFile(filePath),
	}
}
func (c *IniDriver) Get(key string) string {
	if v, o := c.data[key]; o {
		return v
	}
	return ""
}
func (c *IniDriver) All() map[string]string {
	return c.data
}
func (c *IniDriver) GetInt(key string) int64 {
	ret := c.Get(key)
	r, e := strconv.ParseInt(ret, 10, 64)
	if e != nil {
		return 0
	}
	return r
}
func (c *IniDriver) GetBool(key string) bool {
	ret := c.Get(key)
	r, e := strconv.ParseBool(ret)
	if e != nil {
		return false
	}
	return r
}
func (c *IniDriver) GetFloat(key string) float64 {
	ret := c.Get(key)
	r, e := strconv.ParseFloat(ret, 10)
	if e != nil {
		return 0.0
	}
	return r
}
func (c *IniDriver) Del(key string) {
	delete(c.data, key)
	c.save()
}
func (c *IniDriver) GetConf(key string) Conf {
	return BuildConf(c.Get(key))
}
func (c *IniDriver) Set(key, val string) {
	c.data[key] = val
	c.save()
}
func (c *IniDriver) SetInt(key string, val int64) {
	c.Set(key, strconv.FormatInt(val, 10))
}
func (c *IniDriver) SetFloat(key string, val float64) {
	c.Set(key, strconv.FormatFloat(val, 'E', -1, 64))
}
func (c *IniDriver) SetBool(key string, val bool) {
	c.Set(key, strconv.FormatBool(val))
}
func (c *IniDriver) SetConf(key string, val Conf) {
	c.Set(key, val.String())
}

func (c *IniDriver) save() {
	st := bytes.Buffer{}
	for k, v := range c.data {
		st.WriteString(fmt.Sprintf("%s = %s\n", k, v))
	}
	ioutil.WriteFile(c.filePath, st.Bytes(), fs.ModePerm)
}

func paseriniFile(file string) map[string]string {
	ret := map[string]string{}
	data, e := ioutil.ReadFile(file)
	if e != nil {
		return ret
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		i := strings.Index(line, "=")
		i2 := strings.Index(line, "#")
		key := ""
		val := ""
		if i <= 0 {
			continue
		}
		if i2 > 0 {
			if i2 < i {
				continue
			}
		}
		key = strings.Trim(line[:i], " ")
		vv := strings.Trim(line[i+1:], " ")
		if !strings.HasPrefix(vv, "\"") {
			i = strings.Index(vv, "#")
			if i > 0 {
				val = vv[:i]
			} else {
				val = vv
			}
		} else {
			i = strings.Index(vv[1:], "\"")
			if i > 0 {
				val = vv[1 : i-1]
			} else {
				val = val[1:]
			}
		}

		if key != "" {
			ret[key] = val
		}
	}
	return ret
}
