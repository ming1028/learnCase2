package main

import "fmt"

// defer 后进先出，go返回机制中，执行return语句后，Go会创建一个临时变量保存返回值

func cal(str string, a, b int) int {
	ret := a + b
	fmt.Println(str, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer cal("11", a, cal("22", a, b))
	a = 3
	defer cal("33", a, cal("44", a, b))
	// 22 1 2 3
	// 44 3 2 5
	// 33 3 5 8
	// 11 1 3 4
	fmt.Println("main return1", test1())
	fmt.Println("main return2", test2())

	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func test1() int {
	i := 0
	defer func() {
		i += 1
		fmt.Println("defer test1")
	}()
	return i //创建一个临时变量保存返回值
}

func test2() (i int) {
	defer func() {
		i += 1
		fmt.Println("defer test2")
	}()
	return i // 有名返回值
}
