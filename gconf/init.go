package gconf

import (
	"github.com/ghf-go/nannan/drivers/confdriver"
	"os"
)

var (
	_conf confdriver.ConfDriver
)

func init() {
	_conf = confdriver.NewConfDriver(os.Getenv("envinit"))
}
func GetConf(key string) confdriver.Conf {
	return _conf.GetConf(key)
}
