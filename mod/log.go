package mod

import "github.com/ghf-go/nannan/drivers/log_driver"

func Debug(format string, args ...interface{}) {
	log_driver.NewGLog(_logLevle).Debug(format, args...)

}
func Error(format string, args ...interface{}) {
	log_driver.NewGLog(_logLevle).Error(format, args...)
}
func Waring(format string, args ...interface{}) {
	log_driver.NewGLog(_logLevle).Waring(format, args...)
}
func Info(format string, args ...interface{}) {
	log_driver.NewGLog(_logLevle).Info(format, args...)
}
func Log(format string, args ...interface{}) {
	log_driver.NewGLog(_logLevle).Log(format, args...)
}
