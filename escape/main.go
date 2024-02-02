package main

import (
	"fmt"
)

// 逃逸分析
func main() {
	fmt.Println("往往") // 对传入变量的类型未知，统一分配到堆上面
	// 被已经逃逸的指针引用的变量也会发生逃逸
	m := map[int]*string{}
	s := "飞雪无情"
	m[0] = &s // map/slice/chan引用的指针发生逃逸
	// spew.Dump(m)
	e1 := escape()
	e1()
	e1()
	data := []interface{}{100, 200}
	data[0] = 10 // []interface{}数据类型，通过索引（[idx]）赋值必定出现逃逸
	d := []int{1, 2, 3}
	d[0] = 11
	dataM := make(map[string]interface{})
	dataM["key"] = 20 // map[string]interface{}通过赋值必定出现逃逸
	dataM2 := make(map[string]int)
	dataM2["key"] = 20
	a := "111"
	foo(&a)
	/*switch alwaysFalse()
	{
	case true:
		println(true)
	case false:
		println(false)
	}*/
	// 等价于
	switch alwaysFalse(); true {
	case true:
		println(true)
	case false:
		println(false)
	}
}
func alwaysFalse() bool {
	return false
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

func foo(a *string) {
	return
}
