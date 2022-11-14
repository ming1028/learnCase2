package main

import (
	"fmt"
	"io"
	"reflect"
)

func main() {
	i := 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	fmt.Println(iv.CanSet())
	fmt.Println(iv, it)

	i1 := iv.Interface().(int)
	fmt.Println(i1)

	ipv := reflect.ValueOf(&i)
	fmt.Println(ipv.Elem().CanSet())
	ipv.Elem().SetInt(4)
	fmt.Println(i)

	p := person{
		Name: "张三",
		Age:  20,
	}
	ppv := reflect.ValueOf(&p)
	ppv.Elem().Field(1).SetInt(33)
	fmt.Println(p)

	fmt.Println(ppv.Kind() == reflect.Ptr)
	pv := reflect.ValueOf(p)
	fmt.Println(pv.Kind())

	pt := reflect.TypeOf(p)
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段：", pt.Field(i).Name)
	}

	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法：", pt.Method(i).Name)
	}

	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现fmt.Stringer:", pt.Implements(stringerType))
	fmt.Println("是否实现io.Writer:", pt.Implements(writerType))
	fmt.Printf("stringerType:%v\nwriterType:%v\n", stringerType, writerType)
}

type person struct {
	Name string
	Age  int32
}

// 反射修改一个值的规则
// 可被寻址，reflect.ValueOf函数传递一个指针作为参数
// struct结构体字段值需要是可导出的，
// 使用Elem方法获得指针指向的值，才可以调用Set系列方法进行修改

func (p person) String() string {
	return fmt.Sprintf("Name is %s, Age is %d", p.Name, p.Age)
}
