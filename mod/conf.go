package mod

import "github.com/ghf-go/nannan/drivers/conf_driver"

func GetConf(key string) conf_driver.Conf {
	return _conf.GetConf(key)
}
