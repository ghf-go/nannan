package app

import (
	"fmt"
	"github.com/ghf-go/nannan/mod"
	"github.com/ghf-go/nannan/web"
	"os"
	"strings"
)

func Run() {
	mod.Debug("程序启动")
	la := len(os.Args)
	if la == 1 {
		web.WebStart()
	} else {
		switch strings.ToLower(os.Args[1]) {
		case "cli":
			cli(os.Args[2:])
		case "service":
			service(os.Args[2:])
		case "crontab":
			crontab(os.Args[2:])
		default:
			fmt.Printf("参数错误 %s [cli|service|crontab]\n", os.Args[0])
		}

	}
}
