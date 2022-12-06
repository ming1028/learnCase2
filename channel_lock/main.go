package main

import "time"

func main() {

}

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{
		make(chan struct{}, 1),
	}
	mu.ch <- struct{}{} // 获取不到一直阻塞
	return mu
}

func (m *Mutex) Lock() {
	<-m.ch // 能输出说明抢到锁
}

func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

func (m *Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:

	}
	return false
}
