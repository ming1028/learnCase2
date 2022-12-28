package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)

	var x, y int
	go func() {
		x = 1
		r1 := y
		fmt.Println(r1)
		wg.Done()
	}()

	go func() {
		y = 1
		r2 := x
		fmt.Println(r2)
		wg.Done()
	}()

	wg.Wait()
}
