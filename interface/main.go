package main

import (
	"fmt"
)

func main() {
	p := person{
		name: "echo",
		age:  33,
		addr: address{
			province: "shang hai",
			city:     "jing an",
		},
		address: address{
			province: "he nan",
			city:     "zheng zhou",
		},
	}
	fmt.Println(p.province)
	printString(p)
	printString(p.addr)

	var si fmt.Stringer = address{
		province: "河南",
		city:     "三门峡",
	}
	printString(si) // si interface
	/*sip := &si type *fmt.Stringer is pointer to interface, not interface
	printString(sip)*/
	// 指向具体类型的指针可以实现一个接口，但是指向接口的指针永远不可能实现该接口

	fmt.Printf("内存地址为%p\n", &p)
	// Go语言中的函数传递是值传递，将原有数据拷贝
	// 值类型：浮点型、整型、字符串、布尔、数组

	// 不管字面量还是make函数最终都是调用runtime.makemap函数，返回一个*hmap

	// chan 调用runtime.makechan函数，返回一个*hchan
	// map、chan、函数、接口、slice都可以称为引用类型

	// make、new函数属于显示声明并初始化，如果没有显式声明初始化，那么该变量的默认值
	// 就是对应类型的零值
}

/*type Stringer interface {
	String() string
}*/

type person struct {
	name string
	age  uint
	addr address
	address
}

type address struct {
	province string
	city     string
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

// 以值类型接受者实现接口的时候，不管是类型本身还是指针类型，都实现了该接口
func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func (a address) String() string {
	return fmt.Sprintf("the addr is %s%s", a.province, a.city)
}
