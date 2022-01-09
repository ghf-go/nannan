package log_driver

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type LogKafkaDriver struct {
	leves      []int
	kafkaWrite *kafka.Writer
}

func (l *LogKafkaDriver) Levels() []int {
	return l.leves
}

func (l *LogKafkaDriver) Write(format string) {
	e := l.kafkaWrite.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(format),
	})
	if e != nil {
		//panic(e)
	}
}
