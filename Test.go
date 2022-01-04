package main

import (
	"fmt"
	"github.com/ghf-go/nannan/app"
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/web/webbase"
	"os"
)

func main() {
	str := []rune("你好不好")
	fmt.Println(len(str), string(str[1:]))
	os.Setenv("app.web", ":9081")
	os.Setenv("db.default", "mysql://admin:123456@(127.0.0.1:3306)/dev_gay?parseTime=true")
	os.Setenv("redis.default", "redis://127.0.0.1:6379")
	os.Setenv("limiter.token", "mem://mem:10/?time_window=100")
	os.Setenv("limiter.ip", "mem://limitip")
	os.Setenv("es.test", "mem://dev_gay/us.ggvjj.ml:9200")

	web.RegisterMiddleWare(web.JWTMiddleWare)
	web.RegisterMiddleWare(web.WxEchoStrMiddkeWare)

	webbase.RegisterRouter()
	app.Run()
}
