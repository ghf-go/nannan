package def

import (
	"fmt"
	"github.com/ghf-go/nannan/gutils"
	"runtime"
	"sync"
)

const (
	LOG_LEVEL_ERROR  = 0
	LOG_LEVEL_LOG    = 10
	LOG_LEVEL_WARING = 100
	LOG_LEVEL_INFO   = 1000
	LOG_LEVEL_DEBUG  = 10000
)

type LogDriver interface {
	Write(format string)
	Levels() []int
}
type GLog struct {
	Level int
	sync.Mutex
	Logs map[int][]LogDriver
}

func (l *GLog) Register(ll LogDriver) {
	ls := ll.Levels()
	l.Lock()
	defer l.Unlock()
	for _, lv := range ls {
		if ss, ok := l.Logs[lv]; ok {
			ss = append(ss, ll)
			l.Logs[lv] = ss
		} else {
			l.Logs[lv] = []LogDriver{ll}
		}
	}
}
func (l *GLog) format(level, format string, v ...interface{}) string {
	file := ""
	line := 0
	_, f, ln, ok := runtime.Caller(3)

	if ok {
		file = f
		line = ln
	}
	return fmt.Sprintf("[%s] %s %s:%d -> %s", gutils.FormatCurTimeMicroSeconds(), level, file, line, fmt.Sprintf(format, v...))
}
func (l *GLog) Error(format string, v ...interface{}) {
	l.Lock()
	defer l.Unlock()
	if ls, ok := l.Logs[LOG_LEVEL_ERROR]; ok {
		for _, ll := range ls {
			ll.Write(l.format("ERROR", format, v...))
		}
	}
}
func (l *GLog) Log(format string, v ...interface{}) {
	if l.Level < LOG_LEVEL_LOG {
		return
	}
	l.Lock()
	defer l.Unlock()
	if ls, ok := l.Logs[LOG_LEVEL_LOG]; ok {
		for _, ll := range ls {
			ll.Write(l.format("LOG", format, v...))
		}
	}
}
func (l *GLog) Waring(format string, v ...interface{}) {
	if l.Level < LOG_LEVEL_WARING {
		return
	}
	l.Lock()
	defer l.Unlock()
	if ls, ok := l.Logs[LOG_LEVEL_WARING]; ok {
		for _, ll := range ls {
			ll.Write(l.format("WARING", format, v...))
		}
	}
}
func (l *GLog) Info(format string, v ...interface{}) {
	if l.Level < LOG_LEVEL_INFO {
		return
	}
	l.Lock()
	defer l.Unlock()
	if ls, ok := l.Logs[LOG_LEVEL_INFO]; ok {
		for _, ll := range ls {
			ll.Write(l.format("INFO", format, v...))
		}
	}
}
func (l *GLog) Debug(format string, v ...interface{}) {
	if l.Level < LOG_LEVEL_DEBUG {
		return
	}
	l.Lock()
	defer l.Unlock()
	if ls, ok := l.Logs[LOG_LEVEL_DEBUG]; ok {
		for _, ll := range ls {
			ll.Write(l.format("DEBUG", format, v...))
		}
	}
}
