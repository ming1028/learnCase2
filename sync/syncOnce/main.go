package main

import (
	"fmt"
	"sync"
)

func main() {
	doOnce()
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	done := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- struct{}{}
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
