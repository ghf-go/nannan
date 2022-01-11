package mod

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/drivers/log_driver"
	"os"
)

var (
	_conf def.ConfDriver
	_log  *def.GLog
)

func init() {
	_conf = NewConfDriver(os.Getenv("init"))
	_log = log_driver.NewGLog(def.LOG_LEVEL_DEBUG)

}
