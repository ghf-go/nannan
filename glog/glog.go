package glog

import "log"

var (
	_logs = []*log.Logger{log.Default()}
)

func RegisterLogger(l *log.Logger) {
	_logs = append(_logs, l)
}
func Debug(format string, args ...interface{}) {
	for _, l := range _logs {
		l.Printf(" DEBUG "+format, args...)
	}
}
func Error(format string, args ...interface{}) {
	for _, l := range _logs {
		l.Fatalf(" Error "+format, args...)
	}
}
func Info(format string, args ...interface{}) {
	for _, l := range _logs {
		l.Printf(" INFO "+format, args...)
	}
}
func Sql(format string, args ...interface{}) {
	for _, l := range _logs {
		l.Printf(" INFO "+format, args...)
	}
}
