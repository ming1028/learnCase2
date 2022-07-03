package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	g := errgroup.Group{}
	//ctx, cancel := context.WithCancel(context.Background()) // 主动取消
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5) // 定时取消
	g.Go(func() error {
		watchDog(ctx, "watchDog")
		return nil
	})
	time.Sleep(time.Second * 15)
	//cancel()
	if err := g.Wait(); err != nil {
		return
	}
}

func watchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("监控暂停")
			return
		default:
			fmt.Println(name, "监控中")
		}
		time.Sleep(time.Second * 1)
	}
}
