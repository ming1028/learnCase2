package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock() // wait方法中会解锁，需要先加锁
			defer cond.L.Unlock()

			// 公用同一个锁，需要先解锁，使其他goroutine进入临界区。
			cond.Wait()
			fmt.Println(num, "号开始跑")
		}(i)
	}
	time.Sleep(time.Second * 2)

	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() // 唤醒所有等待的协程
	}()
	wg.Wait()
}
