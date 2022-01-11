package mod

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/drivers/log_driver"
	"os"
)

var (
	_conf        def.ConfDriver
	_logLevle    = def.LOG_LEVEL_DEBUG
	_limit_ip    def.IpLimiterDriver
	_limit_token def.TokenLimiterDriver
)

func init() {
	_conf = NewConfDriver(os.Getenv("init"))
	_logLevle = def.LOG_LEVEL_DEBUG
	log_driver.NewGLog(_logLevle)
}
