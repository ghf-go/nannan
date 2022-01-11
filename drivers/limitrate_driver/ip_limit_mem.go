package limitrate_driver

import (
	"sync"
	"time"
)

type ipLimiterMemIP struct {
	sync.Mutex
	lastTime time.Time
	val      int
}
type ipLimiterMemGroup struct {
	sync.Mutex
	data map[string]*ipLimiterMemIP
}
type IpLimiterMemDriver struct {
	sync.Mutex
	data map[string]*ipLimiterMemGroup
}

func (rate *IpLimiterMemDriver) Check(groupName, ip string, maxReq int, timeWindow time.Duration) bool {
	var group *ipLimiterMemGroup
	var ipItem *ipLimiterMemIP
	if g, ok := rate.data[groupName]; ok {
		group = g

	} else {
		group = &ipLimiterMemGroup{data: map[string]*ipLimiterMemIP{ip: &ipLimiterMemIP{lastTime: time.Now().Add(timeWindow), val: 1}}}
		rate.Lock()
		rate.data[groupName] = group
		rate.Unlock()
		return true
	}
	if ipI, ok := group.data[ip]; ok {
		ipItem = ipI
	} else {
		group.Lock()
		ipItem = &ipLimiterMemIP{lastTime: time.Now().Add(timeWindow), val: 1}
		group.data[ip] = ipItem
		group.Unlock()
		return true
	}
	ipItem.Lock()
	defer ipItem.Unlock()
	if ipItem.lastTime.Nanosecond() > time.Now().Nanosecond() {
		if ipItem.val >= maxReq {
			return false
		}
		ipItem.val += 1
	} else {
		ipItem.val = 1
		ipItem.lastTime = time.Now().Add(timeWindow)
	}

	return true
}

func (rate *IpLimiterMemDriver) Start() {
	rate.Lock()
	for k, g := range rate.data {
		g.Lock()
		for ip, item := range g.data {
			if time.Now().Nanosecond() > item.lastTime.Nanosecond() {
				delete(g.data, ip)
			}
		}
		g.Unlock()
		if len(g.data) == 0 {
			delete(rate.data, k)
		}
	}
	rate.Unlock()
	time.Sleep(time.Second * 600)
	rate.Start()
}
