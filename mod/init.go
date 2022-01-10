package mod

import (
	"github.com/ghf-go/nannan/drivers/conf_driver"
	"github.com/ghf-go/nannan/drivers/log_driver"
	"os"
)

var (
	_conf conf_driver.ConfDriver
	_log  *log_driver.GLog
)

func init() {
	_conf = conf_driver.NewConfDriver(os.Getenv("envinit"))
}
