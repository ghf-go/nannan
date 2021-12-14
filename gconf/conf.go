package gconf

import (
	"errors"
	"net/url"
	"os"
	"strconv"
)

type GConf struct {
	_base  string
	scheme string
	user   string
	pass   string
	host   string
	port   int
	path   string
	args   url.Values
}

var (
	_confmap     = map[string]GConf{}
	ErrorNotConf = errors.New("配置不存在")
	ErrorConf    = errors.New("配置错误")
)

//  GetConf 获取配置
func GetConf(confName string) GConf {
	if r, ok := _confmap[confName]; ok {
		return r
	}
	rUrl := os.Getenv(confName)
	if rUrl == "" {
		panic(ErrorNotConf)
	}
	u, e := url.Parse(rUrl)
	if e != nil {
		panic(ErrorConf)
	}
	pass := ""
	if p, ok := u.User.Password(); ok {
		pass = p
	}
	port := 0
	p := u.Port()
	if p != "" {
		port, e = strconv.Atoi(p[1:])
	}
	r := GConf{
		_base:  rUrl,
		scheme: u.Scheme,
		user:   u.User.Username(),
		pass:   pass,
		host:   u.Host,
		port:   port,
		path:   u.Path,
		args:   u.Query(),
	}
	_confmap[confName] = r
	return r
}

func (conf GConf) GetScheme() string {
	return conf.scheme
}
func (conf GConf) GetUserName() string {
	return conf.user
}
func (conf GConf) GetPassWord() string {
	return conf.pass
}
func (conf GConf) GetHost() string {
	return conf.host
}
func (conf GConf) GetPort() int {
	return conf.port
}
func (conf GConf) GetPath() string {
	return conf.path
}
func (conf GConf) GetArgs(name string) string {
	return conf.args.Get(name)
}
func (conf GConf) GetArgInt(name string) int {
	s := conf.args.Get(name)
	if s == ""{
		return 0
	}
	r ,e := strconv.Atoi(s)
	if e != nil{
		return 0
	}
	return r
}
func (conf GConf) GetBase() string {
	return conf._base
}
