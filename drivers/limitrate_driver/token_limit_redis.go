package limitrate_driver

import (
	"github.com/ghf-go/nannan/def"
	"sync"
	"time"
)

type TokenLimitMemDriver struct {
	sync.Mutex
	maxReq     int
	timeWindow time.Duration
	data       map[string]int
}

func NewTokenLimitMemDriver(conf def.Conf) *TokenLimitMemDriver {
	return &TokenLimitMemDriver{
		maxReq:     conf.Port,
		timeWindow: time.Duration(conf.GetArgInt("time_window")) * time.Second,
		data:       map[string]int{},
	}
}
func (t *TokenLimitMemDriver) GetToken(key string) bool {
	if k, ok := t.data[key]; ok {
		if k > 0 {
			t.Lock()
			t.data[key] -= 1
			t.Unlock()
			return true
		} else {
			return false
		}
	} else {
		t.Lock()
		t.data[key] = t.maxReq - 1
		t.Unlock()
		return true
	}
	return true
}
func (t *TokenLimitMemDriver) Start() {
	for def.IsRun() {
		t.Lock()
		for k := range t.data {
			t.data[k] = t.maxReq
		}
		t.Unlock()
		time.Sleep(t.timeWindow)
	}
}
