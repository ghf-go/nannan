package mod

func Debug(format string, args ...interface{}) {
	if _log != nil {
		_log.Debug(format, args...)
	}

}
func Error(format string, args ...interface{}) {
	if _log != nil {
		_log.Error(format, args...)
	}
}
func Waring(format string, args ...interface{}) {
	if _log != nil {
		_log.Waring(format, args...)
	}
}
func Info(format string, args ...interface{}) {
	if _log != nil {
		_log.Info(format, args...)
	}
}
func Log(format string, args ...interface{}) {
	if _log != nil {
		_log.Log(format, args...)
	}
}
