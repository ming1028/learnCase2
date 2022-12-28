package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 一个goroutine中写数据，一个goroutine中读数据
func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(2)

	var x, y int

	go func() {
		x = 1
		y = 1
		wg.Done()
	}()

	go func() {
		r1 := x
		r2 := y

		fmt.Println(r1, r2)
		wg.Done()
	}()

	wg.Wait()
}
