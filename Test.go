package main

import (
	"github.com/ghf-go/nannan/gconf"
	"github.com/ghf-go/nannan/glog"
	"os"
)

func main() {
	os.Setenv("app.web", ":9081")
	os.Setenv("db.default", "mysql://admin:123456@(127.0.0.1:3306)/dev_gay?parseTime=true")
	os.Setenv("redis.default", "redis://127.0.0.1:6379")
	os.Setenv("limiter.token", "mem://mem:10/?time_window=100")
	os.Setenv("limiter.ip", "mem://limitip")
	os.Setenv("es.aaa", "mem:///127.0.0.1:9200,127.0.0.1:9201,127.0.0.1:9204")
	aa := gconf.GetConf("es.aaa")
	glog.Debug("aa %s", aa.GetPath())
	//web.RegisterMiddleWare(web.JWTMiddleWare)
	//web.RegisterMiddleWare(web.WxEchoStrMiddkeWare)
	//
	//webbase.RegisterRouter()
	//app.Run()
}
