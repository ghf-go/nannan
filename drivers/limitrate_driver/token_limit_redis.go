package limitrate_driver

import (
	"sync"
	"time"
)

type tokenLimitMemDriver struct {
	sync.Mutex
	maxReq     int
	timeWindow time.Duration
	data       map[string]int
}

func (t *tokenLimitMemDriver) GetToken(key string) bool {
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
func (t *tokenLimitMemDriver) Start() {
	t.Lock()
	for k, _ := range t.data {
		t.data[k] = t.maxReq
	}
	t.Unlock()
	time.Sleep(t.timeWindow)
	t.Start()
}
