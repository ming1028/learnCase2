package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	ctx = context.WithValue(ctx, "userId", "10010")
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		ctxWithValue(ctx)
		return nil
	})
	if err := g.Wait(); err != nil {
		panic(fmt.Errorf("goroutine出错：%#v\n", err))
	}
}

// 作为参数时放在第一位
//要使用 context.Background 函数生成根节点的 Context，也就是最顶层的 Context。
// 携带traceId串联整个请求链路
func ctxWithValue(
	ctx context.Context,
) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("取消")
			return
		default:
			userId := ctx.Value("userId")
			fmt.Println("用户id:", userId)
		}
		time.Sleep(time.Second * 1)
	}
}
