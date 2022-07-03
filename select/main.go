package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string)

	go func() {
		time.Sleep(time.Second * 8)
		result <- "result"
	}()

	select {
	case res := <-result:
		fmt.Println(res)
	case <-time.After(time.Second * 5):
		fmt.Println("超时")
	}
}
