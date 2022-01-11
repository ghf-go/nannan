package drivers

import (
	"github.com/ghf-go/nannan/def"
)

func Debug(format string, args ...interface{}) {
	//mod.Debug(format, args...)
}
func Error(format string, args ...interface{}) {
	//mod.Error(format, args...)
}
func Waring(format string, args ...interface{}) {
	//mod.Waring(format, args...)
}
func Info(format string, args ...interface{}) {
	//mod.Info(format, args...)
}
func Log(format string, args ...interface{}) {
	//mod.Log(format, args...)
}
func GetConf(key string) def.Conf {
	return def.Conf{} //mod.GetConf(key)
}
