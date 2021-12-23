package main

import (
	"github.com/ghf-go/nannan/app"
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase"
	"os"
)

func main() {
	os.Setenv("app.web", ":9081")
	os.Setenv("db.default", "mysql://admin:123456@(127.0.0.1:3306)/dev_gay?parseTime=true")
	os.Setenv("redis.default", "redis://127.0.0.1:6379")
	web.RegisterMiddleWare(web.WxEchoStrMiddkeWare)
	web.RegisterMiddleWare(func(ctx *web.EngineCtx, f func(*web.EngineCtx)) {
		ctx.LimitToken("/app.all")
		f(ctx)
	})
	webbase.RegisterRouter()
	app.Run()
}
