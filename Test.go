package main

import (
	"fmt"
	"github.com/ghf-go/nannan/app"
	"github.com/ghf-go/nannan/mod"
	"github.com/ghf-go/nannan/web"
	"github.com/ghf-go/nannan/webbase"
	"os"
)

//import _ "github.com/ghf-go/nannan/test"

func main() {
	fmt.Println(os.Getenv("gay_init"))
	mod.NewConfDriver(os.Getenv("gay_init"))
	web.RegisterMiddleWare(web.JWTMiddleWare)
	web.RegisterMiddleWare(web.WxEchoStrMiddkeWare)

	webbase.RegisterRouter()
	app.Run()
}
