package main

import (
	"fmt"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"sync"
)

func main() {
	builds := buy(100)
	pack1 := build(builds)
	pack2 := build(builds)
	pack3 := build(builds)

	packs := merge2(pack1, pack2, pack3)
	goods := pack(packs)
	for out := range goods {
		fmt.Println(out)
	}
}

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i <= n; i++ {
			out <- "配件" + cast.ToString(i)
		}
	}()
	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for ins := range in {
			out <- "组装：（" + ins + ")"
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for ins := range in {
			out <- "打包（" + ins + ")"
		}
	}()
	return out
}

func merge(ins ...<-chan string) <-chan string {
	g := errgroup.Group{}
	out := make(chan string)

	p := func(in <-chan string) {
		for i := range in {
			out <- i
		}
	}

	for _, inOne := range ins {
		g.Go(func() error {
			//inV := inOne
			p(inOne)
			return nil
		})
	}
	g.Go(func() error {
		defer close(out)
		return g.Wait()
	})
	return out
}

func merge2(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup

	out := make(chan string)

	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}

	wg.Add(len(ins))
	for _, cs := range ins {
		go p(cs)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
