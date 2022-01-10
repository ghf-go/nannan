package log_driver

import (
	"fmt"
)

type LogStdDriver struct {
	leves []int
}

func (l *LogStdDriver) Levels() []int {
	return l.leves
}

func (l *LogStdDriver) Write(format string) {
	fmt.Println(format)
}
