package confdriver

import (
	"net/url"
	"strconv"
	"strings"
)

type Conf struct {
	Scheme   string
	Host     string
	Port     int
	UserName string
	PassWord string
	Path     string
	Args     url.Values
}
type ConfDriver interface {
	All() map[string]string
	Get(key string) string
	GetInt(key string) int64
	GetBool(key string) bool
	GetFloat(key string) float64
	GetConf(key string) Conf
	Del(key string)
	Set(key, val string)
	SetInt(key string, val int64)
	SetFloat(key string, val float64)
	SetBool(key string, val bool)
	SetConf(key string, val Conf)
}

func NewConfDriver(data string) ConfDriver {
	conf := BuildConf(data)
	switch conf.Scheme {
	case "ini":
		return NewIniDriver(conf.Path)
	case "etcd":
		return NewEtcdDriverByConf(conf)
	default:
		return NewEnvDriver()

	}
}
func BuildConf(data string) Conf {
	u, e := url.Parse(data)
	if e != nil {
		return Conf{}
	}
	ret := Conf{
		Scheme:   u.Scheme,
		Host:     u.Hostname(),
		UserName: u.User.Username(),
		Args:     u.Query(),
	}
	if len(u.Path) > 0 {
		ret.Path = u.Path[1:]
	}
	if u.Port() != "" {
		ret.Port, _ = strconv.Atoi(u.Port())
	}
	if p, e := u.User.Password(); e {
		ret.PassWord = p
	}

	return ret
}

func (c Conf) String() string {
	r := strings.Builder{}
	r.WriteString(c.Scheme + "://")
	if c.UserName != "" {
		r.WriteString(c.UserName + ":" + c.PassWord + "@")
	}
	r.WriteString(c.Host)
	if c.Port > 0 {
		r.WriteString(":" + strconv.Itoa(c.Port))
	}
	r.WriteString("/" + c.Path + "?")
	r.WriteString(c.Args.Encode())
	return r.String()
}
