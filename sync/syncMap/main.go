package main

import (
	"fmt"
	"sync"
)

func main() {
	map1 := sync.Map{}
	map1.Store("name", "张三")
	fmt.Println(map1.Load("name"))
	fmt.Println(map1.LoadOrStore("address", "上海"))
	fmt.Println(map1)
	map1.Delete("name")
	fmt.Println(map1.Load("name"))
	fmt.Println("====")
	map1.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
