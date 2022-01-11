package app

import (
	"fmt"
	"github.com/ghf-go/nannan/drivers/cli_driver"
	"os"
	"runtime"
)

var (
	_crontabMap = map[string]cli_driver.CrontabDriver{}
)

func RegisterCrontab(name, timer string, isLock bool, callfunc func()) {
	_, p, _, _ := runtime.Caller(1)
	_crontabMap[name] = cli_driver.CrontabDriver{
		Timer:    timer,
		IsLock:   false,
		Cmd:      fmt.Sprintf("%s crontab %s", p, name),
		CallFunc: callfunc,
	}
}
func crontab(args []string) {
	al := len(args)
	switch al {
	case 1:
		if f, ok := _crontabMap[args[0]]; ok {
			f.CallFunc()
			Wait()
		} else {
			fmt.Printf("%s 不存在", args[0])
		}
	case 2:
		if t, ok := _crontabMap[args[0]]; ok {
			switch args[1] {
			case "save":
				t.Save()
			case "remove":
				t.Remove()
			case "check":
				t.Current()
			}
		} else {
			fmt.Printf("%s 服务不存在", args[0])
		}
	default:
		fmt.Printf("参数错误 -> %s crontab sername [save|remove|check]", os.Args[0])
	}
}
