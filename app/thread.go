package app

import "sync"

var (
	_run = true
	_w   = sync.WaitGroup{}
)

func IsRun() bool {
	return _run
}
func Exit() {
	_run = false
}

func GoStart() {
	_w.Add(1)
}
func Done() {
	_w.Done()
}
func Wait() {
	_w.Wait()
}
