package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

// 逃逸分析
func main() {
	fmt.Println("往往") // 被已经逃逸的指针引用的变量也会发生逃逸
	m := map[int]*string{}
	s := "飞雪无情"
	m[0] = &s
	spew.Dump(m)
	e1 := escape()
	e1()
	e1()
}

// 指针作为函数返回值的时候，一定会发生逃逸
func newString() *string {
	s := new(string)
	*s = "赵钱孙李"
	return s
}

func newString2() string {
	s := new(string)
	*s = "赵钱孙李"
	return *s
}

func escape() func() int {
	a := 6
	return func() int {
		a += 1
		return a
	}
}
