package log_driver

import (
	"github.com/ghf-go/nannan/def"
	"github.com/segmentio/kafka-go"
)

var (
	_glog *def.GLog
)

func GetGLog() *def.GLog {
	if _glog == nil {
		_glog = &def.GLog{
			Level: 0,
			Logs:  map[int][]def.LogDriver{},
		}
		//_glog = NewGLog(def.LOG_LEVEL_DEBUG)
	}
	return _glog
}
func NewGLog(level int) *def.GLog {

	if _glog != nil {
		if _glog.Level != level {
			_glog.Level = level
		}
		return _glog
	}

	_glog = &def.GLog{
		Level: level,
		Logs:  map[int][]def.LogDriver{},
	}
	_glog.Register(NewLogStdDriver(def.LOG_LEVEL_DEBUG))
	return _glog
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
