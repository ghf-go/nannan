package app

import (
	"fmt"
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/drivers/cli_driver"
	"os"
	"runtime"
)

var (
	_serviceMap = map[string]func(){}
)

func RegisterService(name string, callfunc func()) {
	_serviceMap[name] = callfunc
}
func service(args []string) {
	al := len(args)
	switch al {
	case 1:
		if f, ok := _serviceMap[args[0]]; ok {
			f()
			def.Wait()
		} else {
			fmt.Printf("%s 服务不存在", args[0])
		}
	case 2:
		if _, ok := _serviceMap[args[0]]; ok {
			_, p, _, _ := runtime.Caller(1)
			s := cli_driver.SystemService{
				Cmd:  fmt.Sprintf("%s service %s", p, args[0]),
				Name: args[0],
				Desc: args[0],
			}
			switch args[1] {
			case "start":
				s.Start()
			case "stop":
				s.Stop()
			case "restart":
				s.Restart()
			case "install":
				s.Install()
			case "uninstall":
				s.UnInstall()
			}
		} else {
			fmt.Printf("%s 服务不存在", args[0])
		}
	default:
		fmt.Printf("参数错误 -> %s service sername [start|stop|restart|install|uninstall]", os.Args[0])
	}
}
