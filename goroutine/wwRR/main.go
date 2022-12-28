package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(4)

	var x, y int
	go func() {
		x = 1
		wg.Done()
	}()

	go func() {
		y = 1
		wg.Done()
	}()

	go func() {
		r1 := x
		r2 := y

		fmt.Println(r1, r2)
		wg.Done()
	}()

	go func() {
		r3 := x
		r4 := y

		fmt.Println(r3, r4)
		wg.Done()
	}()
	wg.Done()
}
