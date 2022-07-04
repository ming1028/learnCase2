package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog(ctx, "监控1")
	}()

	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}

func watchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控暂停")
			return
		default:
			fmt.Println(name, "正在监控")
		}
		time.Sleep(time.Second)
	}
}
