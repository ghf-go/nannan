package drivers

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/drivers/log_driver"
)

func Debug(format string, args ...interface{}) {
	log_driver.GetGLog().Debug(format, args...)
}
func Error(format string, args ...interface{}) {
	log_driver.GetGLog().Error(format, args...)
}
func Waring(format string, args ...interface{}) {
	log_driver.GetGLog().Waring(format, args...)
}
func Info(format string, args ...interface{}) {
	log_driver.GetGLog().Info(format, args...)
}
func Log(format string, args ...interface{}) {
	log_driver.GetGLog().Log(format, args...)
}
func GetConf(key string) def.Conf {
	return def.Conf{} //mod.GetConf(key)
}
