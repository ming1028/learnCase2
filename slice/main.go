package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	sli := []int{1, 2, 3, 4}
	fmt.Println(len(sli), cap(sli))

	nameMap := make(map[string]int)
	nameMap["name"] = 12
	nameMap["name2"] = 33
	delete(nameMap, "name")
	fmt.Println(nameMap)
	// UTF8编码下一个汉字对应三个字节

	// 计算汉字长度
	chinese := "啦嗷啊"
	fmt.Println(utf8.RuneCountInString(chinese))
	for idx, s := range chinese {
		st := s
		fmt.Println(idx, st, string(s)) // s unicode字符对应的码点、
		// for range自动隐式解码unicode字符串
	}
}
