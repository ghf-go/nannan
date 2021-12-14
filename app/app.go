package app

import (
	"fmt"
	"github.com/ghf-go/nannan/web"
	"os"
	"strings"
)

func Run()  {
	la := len(os.Args)
	if la == 1{
		web.WebStart()
	}else{
		switch strings.ToLower(os.Args[1]){
		case "cli":
			cli(os.Args[1:])
		case "service":
			service(os.Args[1:])
		case "crontab":
			crontab(os.Args[1:])
		default:
			fmt.Printf("参数错误 %s [cli|service|crontab]\n",os.Args[0])
		}

	}
}
