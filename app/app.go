package app

import (
	"fmt"
	"github.com/ghf-go/nannan/web"
	"os"
	"strings"
)

var (
	_run = true
)

func Run() {
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
func IsRun() bool {
	return _run
}
func Exit() {
	_run = false
}
