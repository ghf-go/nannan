package main

import (
	"github.com/ghf-go/nannan/drivers"
	"github.com/ghf-go/nannan/drivers/log_driver"
)

func main() {
	l := log_driver.NewGLog(log_driver.LOG_LEVEL_DEBUG)
	l.Register(log_driver.NewLogKafkaDriver(drivers.NewKafkaWrite("127.0.0.1:9092", "test001"), log_driver.LOG_LEVEL_DEBUG))
	//l.Register(log_driver.NewLogStdDriver(log_driver.LOG_LEVEL_DEBUG))
	//l.Register(log_driver.NewLogFileDriver("/tmp", log_driver.LOG_LEVEL_DEBUG))
	for i := 0; i < 1000; i++ {
		l.Debug("你好 %d", i)
	}

	//str := []rune("你好不好")
	//fmt.Println(len(str), string(str[1:]))
	//os.Setenv("app.web", ":9081")
	//os.Setenv("db.default", "mysql://admin:123456@(127.0.0.1:3306)/dev_gay?parseTime=true")
	//os.Setenv("redis.default", "redis://127.0.0.1:6379")
	//os.Setenv("limiter.token", "mem://mem:10/?time_window=100")
	//os.Setenv("limiter.ip", "mem://limitip")
	//os.Setenv("es.test", "mem://dev_gay/us.ggvjj.ml:9200")
	//
	//web.RegisterMiddleWare(web.JWTMiddleWare)
	//web.RegisterMiddleWare(web.WxEchoStrMiddkeWare)
	//
	//webbase.RegisterRouter()
	//app.Run()
}
