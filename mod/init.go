package mod

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/drivers/log_driver"
	"os"
)

var (
	_conf def.ConfDriver
	_log  *log_driver.GLog
)

func init() {
	_conf = NewConfDriver(os.Getenv("envinit"))
	_log = log_driver.NewGLog(log_driver.LOG_LEVEL_DEBUG)

}
