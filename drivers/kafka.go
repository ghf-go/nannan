package drivers

import (
	kafka "github.com/segmentio/kafka-go"
)

// 生成kafak的生产者
func NewKafkaWrite(url, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(url),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
