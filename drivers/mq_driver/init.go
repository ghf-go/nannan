package mq_driver

import (
	"time"
)

type MqDriver interface {
	Send(string, string)
	SendDelay(string, string, time.Duration)
	SendTopic(string, string)
	ConsumerMq(callfunc func(string, string, MqDriver), topics ...string)
	ConsumerMqDelay(callfunc func(string, string, MqDriver), topics ...string)
	ConsumerTopic(callfunc func(string, string, MqDriver), topics ...string)
}
