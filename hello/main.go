package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// 变量声明
	var vari int = 10
	fmt.Println(vari)

	var var2 = 20
	fmt.Println(var2)

	var (
		var3 int = 30
		var4 int = 40

		var5 = 50
		var6 = 60
	)
	fmt.Println(var3, var4, var5, var6)
}
