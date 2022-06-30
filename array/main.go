package main

import "fmt"

// slice map不能比较
// go不存在隐式的类型转换，类型不同不能比较
// 数组的长度是类型的一部分，长度不同，无法比较
// 逐个元素比较类型和值，如果元素是不可比较的类型，则不能比较报错
// slice、map类型不能比较，只能与nil做比较
func main() {
	fmt.Println("1" == 1)
}
