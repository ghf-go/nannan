package app

import (
	"fmt"
	"os"
)

var (
	_cliMap = map[string]func([]string){}
)

func RegisterCli(name string, callfunc func([]string)) {
	_cliMap[name] = callfunc
}

func cli(args []string) {
	al := len(args)
	if al == 0 {
		fmt.Printf("参数错误 %s cli cliname", os.Args[0])
	} else {
		if f, ok := _cliMap[args[0]]; ok {
			f(args[1:])
		} else {
			fmt.Printf("%s  cli %s 命令不存在", os.Args[0], args[0])
		}
	}
}
