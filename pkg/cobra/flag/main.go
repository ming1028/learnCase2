package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(os.Args)
	if len(args) <= 0 {
		fmt.Println("Usage: admin-cli[command]")
		return
	}

	switch args[0] {
	case "help":
		fmt.Println("help command")
	case "export":
		fmt.Println("export command")
		if len(args) == 3 {
			fmt.Println("export 3")
		} else if len(args) == 2 {
			fmt.Println("export 2")
		}
	default:
		fmt.Println("please input command")
	}
}
