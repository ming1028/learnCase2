package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, _ := os.Create("./trace/trace.out")
	defer f.Close()

	_ = trace.Start(f)
	defer trace.Stop()

	fmt.Println("hello world")
}
