package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	fmt.Println(fp)
	*fp = *fp * 3
	fmt.Println(i)

	p := new(person)
	pName := (*string)(unsafe.Pointer(p)) // p在内存中的起始位置也是Name的内存位置
	*pName = "张三"
	// 内存偏移量计算 uintptr(unsafe.Pointer(p)) p在内存起始位置
	// unsafe.Offsetof(p.Age) Age内存大小
	// 进行指针运算，先通过unsafe.Pointer转换为uintptr类型的指针
	// 赋值或者取值操作还需要通过unsafe.Pointer转换为真实的指针类型
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 22
	fmt.Println(*p)

	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(1000000)))

	fmt.Println(unsafe.Alignof(p))
	fmt.Println(unsafe.Alignof(true))
}

type person struct {
	Name string
	Age  int
}
