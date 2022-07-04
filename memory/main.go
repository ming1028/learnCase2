package main

import "fmt"

// Go语言程序所管理的虚拟内存空间会被分为两部分，堆内存和栈内存
// 栈内存主要由Go语言来管理，
// Go语言的内存垃圾回收是针对堆内存的垃圾回收
func main() {
	// var s string 只是声明了一个变量，类型为string，没有初始化，值为类型的零值
	// var sp *string 未被初始化，是*string类型的零值 nil（指针零值）

	// 值类型的变量没有初始化（申请内存），会分配好内存，可以直接赋值
	var s *string
	fmt.Printf("%p\n", s) // 默认地址

	// 一个变量必须经过声明、内存分配才能赋值，才可以在声明的时候进行初始化。指针类型
	// 在声明的时候，Go语言并没有自动分配内存，所以不能对其进行赋值操作。

	// new根据传入的类型申请一块内存，返回这块内存的指针，指向的数据就是该类型的零值

}
