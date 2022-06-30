package main

import "fmt"

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	panic("panic")
}
