package log_driver

import (
	"fmt"
	"github.com/ghf-go/nannan/gutils"
	"github.com/segmentio/kafka-go"
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

type logDriver interface {
	Write(format string)
	Levels() []int
}
type GLog struct {
	Level int
	sync.Mutex
	logs map[int][]logDriver
}

func NewGLog(level int) *GLog {
	return &GLog{
		Level: level,
		logs:  map[int][]logDriver{},
	}
}
func NewLogStdDriver(level ...int) *LogStdDriver {
	return &LogStdDriver{leves: level}
}
func NewLogFileDriver(dir string, level ...int) *LogFileDriver {
	return &LogFileDriver{
		dir:   dir,
		leves: level}
}
func NewLogKafkaDriver(kw *kafka.Writer, level ...int) *LogKafkaDriver {
	return &LogKafkaDriver{
		kafkaWrite: kw,
		leves:      level}
}

func (l *GLog) Register(ll logDriver) {
	ls := ll.Levels()
	l.Lock()
	defer l.Unlock()
	for _, lv := range ls {
		if ss, ok := l.logs[lv]; ok {
			ss = append(ss, ll)
			l.logs[lv] = ss
		} else {
			l.logs[lv] = []logDriver{ll}
		}
	}
}
func (l *GLog) format(level, format string, v ...interface{}) string {
	file := ""
	line := 0
	_, f, ln, ok := runtime.Caller(2)

	if ok {
		file = f
		line = ln
	}
	return fmt.Sprintf("[%s] %s %s:%d -> %s", gutils.FormatCurTimeMicroSeconds(), level, file, line, fmt.Sprintf(format, v...))
}
func (l *GLog) Error(format string, v ...interface{}) {
	l.Lock()
	defer l.Unlock()
	if ls, ok := l.logs[LOG_LEVEL_ERROR]; ok {
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
	if ls, ok := l.logs[LOG_LEVEL_LOG]; ok {
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
	if ls, ok := l.logs[LOG_LEVEL_WARING]; ok {
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
	if ls, ok := l.logs[LOG_LEVEL_INFO]; ok {
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
	if ls, ok := l.logs[LOG_LEVEL_DEBUG]; ok {
		for _, ll := range ls {
			ll.Write(l.format("DEBUG", format, v...))
		}
	}
}
