package log_driver

import (
	"fmt"
	"github.com/ghf-go/nannan/gutils"
	"os"
	"sync"
)

type LogFileDriver struct {
	sync.Mutex
	dir     string
	curFile *os.File
	leves   []int
	curDate int64
}

func (l *LogFileDriver) Write(format string) {
	l.Lock()
	defer l.Unlock()
	if l.curDate == 0 || l.curDate < gutils.CurDateUnix() {
		l.curDate = gutils.CurDateUnix() + 86399
		if l.curFile != nil {
			l.curFile.Close()
		}

		f, ee := os.Create(fmt.Sprintf("%s/%s.log", l.dir, gutils.CurDate()))
		if ee != nil {
			panic(ee)
		}
		l.curFile = f
	}

	if l.curFile != nil {
		l.curFile.WriteString(format + "\n")
		l.curFile.Sync()
	}

}
func (l *LogFileDriver) Levels() []int {
	return l.leves
}
