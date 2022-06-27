package main

import "fmt"

// 调用方法的时候，传递的接受者本质是副本
// 只不过一个是值的副本，一个是指向这个值的指针的副本
// 如果使用一个值类型变量调用指针类型接受者的方法，GO编译器会自动取指针调用
// 同理指针调用值类型接受者的方法，也会自动帮我们解引用调用

func main() {
	cl := colsure()
	// 匿名函数引用了外部变量，形成一个闭包，在闭包的声明周期内，引用变量一直有效，会常驻与内存
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	// 强制类型转换
	age := Age(25)
	age.String("zhangsan")

	fmt.Printf("%#v\n", Age.String)

	ageStr := Age.String
	ageStr(age, "zhangsan") // 不管方法有没有参数，通过方法表达式调用，第一个参数必须是接受者
	// 然后才是方法自身的参数
}

func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Age uint

func (age Age) String(name string) {
	fmt.Println(name, "the age is", age)
}
