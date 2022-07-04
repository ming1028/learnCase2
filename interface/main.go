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
