package main

import (
	"github.com/ghf-go/nannan/app"
	"github.com/ghf-go/nannan/web"
	"os"
)

func main() {
	os.Setenv("app.web",":9081")
	web.RegisterRouterGroup("/abc", func(group *web.RouterGroup) {
		group.GET("zz", func(ctx *web.EngineCtx) error {
			return ctx.JsonSuccess(map[string]interface{}{
				"a" : 1,
				"dd" : "阿萨德",
			})
		})
	})
	web.RegisterMiddleWare(func(e *web.EngineCtx ,f func(*web.EngineCtx)){
		e.Header().Add("aa","bb")
		//glog.Debug("测试1")
		f(e)
		//glog.Debug("ceshi 2")
	})
	app.Run()
}
