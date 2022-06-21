package main

import "fmt"

func main() {
	// 基础类型：整型、浮点数、布尔型和字符串

	// 整型：int、uint可能是32位或者64位，具体和cpu有关
	// 有符号整型：int、int8、int16、int32和int64
	// 无符号整型：uint、unit8、uint16、uint32和uint64
	// 字节类型byte等价于uint8，是uint8的别名

	// 浮点数 float32 float64
	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Println("f32 is", f32, "f64 is", f64)

	// 布尔型只有true和false
	var bf bool = false
	var bt bool = true
	fmt.Println("bf is", bf, "bt is", bt)

	// 字符串
	var s1 string = "Hello"
	var s2 string = "世界"
	fmt.Println("s1 is", s1, "s2 is", s2)
	fmt.Println("s1+s2=", s1+s2)
	s1 += s2
	fmt.Println("s1 += s2:", s1)

	// 零值
	var (
		zInt     int
		zFloat64 float64
		zFloat32 float32
		zBool    bool
		zString  string
	)
	fmt.Println(zInt, zFloat32, zFloat64, zBool, zString)

	// 常量 const name = "" 只允许布尔型、字符串、数字类型

	// itoa 常量生成器
	const (
		one = iota
		two
		three = "位"
		four
		five
		six, seven = iota, iota
	)
	fmt.Println(one, two, three, four, five, six, seven)
}
