package main

import (
	"github.com/ghf-go/nannan/app"
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase"
	"os"
)

func main() {
	os.Setenv("app.web", ":9081")
	web.RegisterRouterGroup("/abc", func(group *web.RouterGroup) {
		group.GET("zz", func(ctx *web.EngineCtx) error {
			return ctx.JsonSuccess(map[string]interface{}{
				"a":  1,
				"dd": "阿萨德",
			})
		})
	})
	webbase.RegisterRouter()
	app.Run()
}
