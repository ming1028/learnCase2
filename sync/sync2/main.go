package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum   = 0
	mutex sync.RWMutex
)

func main() {
	for i := 0; i < 100; i++ {
		go add(10)
	}
	for i := 0; i < 10; i++ {
		go fmt.Println(readSum())
	}
	time.Sleep(time.Second * 2)
}

func readSum() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

func add(i int) {
	sum += i
}
