package main

import "time"

func main() {

}

type Mutex struct {
	ch chan struct{}
}

// 初始化
func NewMutex() *Mutex {
	mu := &Mutex{
		make(chan struct{}, 1),
	}
	mu.ch <- struct{}{}
	return mu
}

func (m *Mutex) Lock() {
	<-m.ch // 能输出说明抢到锁，依靠有值输出控制锁的获取,负责会一直阻塞
}

func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}: // 解锁，让Lock可以有值输出
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
	timer := time.NewTimer(timeout) // 定时器
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:

	}
	return false
}
