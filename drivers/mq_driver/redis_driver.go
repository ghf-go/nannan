package mq_driver

import (
	"context"
	"github.com/ghf-go/nannan/def"
	"github.com/go-redis/redis/v8"
	"time"
)

type MqRedisDriver struct {
	redis *redis.Client
}

func NewMqRedisDriver(r *redis.Client) *MqRedisDriver {
	return &MqRedisDriver{
		redis: r,
	}
}
func (m *MqRedisDriver) Send(key, msg string) {
	m.redis.LPush(context.Background(), key, msg)
}
func (m *MqRedisDriver) SendDelay(key, msg string, delay time.Duration) {
	m.redis.ZAdd(context.Background(), key, &redis.Z{
		Score:  float64(time.Now().Add(delay).Unix()),
		Member: msg,
	})
}
func (m *MqRedisDriver) SendTopic(key, msg string) {
	m.redis.Publish(context.Background(), key, msg)
}

func (m *MqRedisDriver) ConsumerMq(callfunc func(string, string, def.MqDriver), topics ...string) {
	for def.IsRun() {
		r, e := m.redis.BRPop(context.Background(), time.Second*60, topics...).Result()
		if e != nil {
			continue
		}
		mlen := len(r)
		for i := 0; i < mlen; i++ {
			callfunc(r[i], r[i+1], m)
			i++
		}
	}
}
func (m *MqRedisDriver) ConsumerMqDelay(callfunc func(string, string, def.MqDriver), topics ...string) {
	for def.IsRun() {
		for _, topic := range topics {
			r, e := m.redis.ZRangeWithScores(context.Background(), topic, 0, time.Now().Unix()).Result()
			if e != nil {
				continue
			}
			for _, z := range r {
				callfunc(topic, z.Member.(string), m)
			}
		}

	}
}
func (m *MqRedisDriver) ConsumerTopic(callfunc func(string, string, def.MqDriver), topics ...string) {
	r := m.redis.Subscribe(context.Background(), topics...)
	for def.IsRun() {
		msg, e := r.ReceiveMessage(context.Background())
		if e != nil {
			continue
		}
		callfunc(msg.Channel, msg.Payload, m)
	}
	r.Unsubscribe(context.Background(), topics...)
}
