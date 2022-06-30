package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 进程：包含了软件运行所需的所有资源，比如
// 内存空间、文件句柄
// 线程：是进程的执行空间，线程被操作系统调度执行
// 协程：被Go runtime所调度

// 有缓冲channel 内部有一个缓冲的队列，发送操作向队列的尾部插入元素，如果队列已满
// 则阻塞等待，直到另一个goroutine执行，接受操作释放队列的空间
// 接受操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个goroutine执行
// 发送操作插入新的元素
func main() {
	ch := make(chan string, 2)
	go func() {
		fmt.Println("in goroutine")
		for i := 0; i < 5; i++ {
			ch <- "goroutine send " + strconv.Itoa(i)
			time.Sleep(time.Second)
		}
		// channel被关闭，就不能向里面发送数据，如果发送会引起Panic异常,但是还可以接收channel中
		// 的数据，如果channel中没有数据，接受的数据就是元素类型的零值
		close(ch)
	}()
	fmt.Println("in main goroutine")
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("main goroutine end")

	cacheCh := make(chan int, 5)
	cacheCh <- 1
	cacheCh <- 2
	fmt.Println("cache容量:", cap(cacheCh), " 元素个数为：", len(cacheCh))

	// 单向channel
	// onlySend := make(chan<- int) 只能发送操作
	// onlyReceive := make(<-chan int) 只能接受操作
	firstCh := make(chan string)
	secondCh := make(chan string)
	thirdCh := make(chan string)

	go func() {
		firstCh <- downloadFile("firstCh")
	}()
	go func() {
		secondCh <- downloadFile("secondCh")
	}()
	go func() {
		thirdCh <- downloadFile("thirdCh")
	}()

	// 同时有多个case可以被执行，则随机选择一个
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-thirdCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	default:
		fmt.Println("default")
	}
	fmt.Printf("%#v\n", rand.Intn(10))
}

func downloadFile(chanName string) string {
	time.Sleep(time.Second * time.Duration(rand.Int63()))
	return chanName + ":filepath"
}
