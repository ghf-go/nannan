package mod

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/drivers/log_driver"
	"os"
)

var (
	_conf def.ConfDriver
	//_log  *def.GLog
	_logLevle = def.LOG_LEVEL_DEBUG
)

func init() {
	_conf = NewConfDriver(os.Getenv("init"))
	_logLevle = def.LOG_LEVEL_DEBUG
	log_driver.NewGLog(_logLevle)
	//_log = log_driver.NewGLog(def.LOG_LEVEL_DEBUG)

}
