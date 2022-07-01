package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
)

var (
	sum   = 0
	mutex sync.Mutex
)

func main() {
	g := errgroup.Group{}
	for i := 1; i <= 100; i++ {
		g.Go(func() error {
			add(10)
			fmt.Println(sum)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Println(sum)
}

func add(i int) {
	/*mutex.Lock()
	defer mutex.Unlock()*/
	sum += i
}
