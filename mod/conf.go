package mod

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/drivers/conf_driver"
)

func GetConf(key string) def.Conf {
	return _conf.GetConf(key)
}

func NewConfDriver(data string) def.ConfDriver {
	conf := def.BuildConf(data)
	switch conf.Scheme {
	case "ini":
		return conf_driver.NewIniDriver(conf.Path)
	case "etcd":
		return conf_driver.NewEtcdDriverByConf(conf)
	default:
		return conf_driver.NewEnvDriver()
	}
}
