package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	i := 10
	iString := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", iString, iString)

	stringInt, err := strconv.Atoi(iString)
	fmt.Printf("%T, %v\n", stringInt, err)

	// 数字类型之间可以通过强制转换的方式
	stringFloat := "32.33"
	strConFloat, err := strconv.ParseFloat(stringFloat, 32)
	fmt.Printf("%T, %v\n", strConFloat, err)

	strConFloat64, err := strconv.ParseFloat(stringFloat, 64)
	fmt.Printf("%T, %v\n", strConFloat64, err)

	boolStr := "true"
	strConBool, err := strconv.ParseBool(boolStr)
	fmt.Printf("%T, %v\n", strConBool, err)

	str := 3.141592832423923239480238432849238482394
	strFormat := strconv.FormatFloat(str, 'f', 3, 64)
	fmt.Printf("%T %v\n", strFormat, strFormat)

	s1, s2 := "hello", "lalala"
	fmt.Println("前缀", strings.HasPrefix(s1, "H"), strings.HasPrefix(s1, "h"))
	fmt.Println("索引", strings.Index(s2, "k")) // strings.Index() 不存在返回-1
	fmt.Println("前缀", strings.HasPrefix(s1, ""), s1[0:len("")], len(""))
	fmt.Println("后缀", strings.HasSuffix(s1, ""))

	fmt.Println("大小写", strings.ToUpper(s1), strings.ToLower(s1), strings.ToTitle(s1))
	fmt.Println("包含", strings.Contains(s1, "h"), strings.Contains(s1, "rr"))
	fmt.Println("包含任意一个", strings.ContainsAny(s1, "ol"), strings.ContainsAny(s1, "re"))
	fmt.Println("包含任意一个", strings.ContainsAny("", "")) // str为空直接返回-1
	fmt.Println("包含任意一个", strings.ContainsAny("hello", "re"))

	fmt.Println("出现次数：", strings.Count(s1, "r"), strings.Count(s1, ""))

	fmt.Println("字符首次出现的位置：", strings.IndexByte(s1, 'a'), strings.IndexByte(s1, 'h'))
	fmt.Println("字符最后出现的位置：", strings.LastIndex(s1, "ss"), strings.LastIndex(s2, "ll"),
		strings.LastIndexByte(s1, 'a'), strings.LastIndexByte(s1, 'o'),
	)

	// 字符串切割
	fmt.Println("字符串切割", strings.Fields(s1), strings.Fields("hello world 123 2 3   f"))
	f := func(r rune) bool {
		fmt.Println(string(r))
		return true
	}
	// 返回：如果str中的所有代码点均满足f(c)或字符串为空，则返回空片。
	fmt.Printf("字符串闭包切割: %q\n", strings.FieldsFunc("ABC123PQR", f))

	f = func(c rune) bool {
		return unicode.IsSpace(c) || c == '.'
	}

	s := "We are humans. We are social animals."
	fmt.Printf("Fields are:%q\n", strings.FieldsFunc(s, f))

	// 字符串分割
	fmt.Println(strings.Split(s1, "ll"))       // 不包含ll
	splitAfter := strings.SplitAfter(s1, "ll") // 包含ll
	fmt.Println(splitAfter)
	splitAfterN := strings.SplitAfterN(s1, "l", 4) // 使用l分割，n分割多少段
	fmt.Println(splitAfterN)

	// 大小写转换
	fmt.Println(strings.ToTitle("hello world"))
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.Title("hello world"))
	fmt.Println(strings.ToLower(s1))

	// 修改字符串
	fmt.Println(strings.Trim("hello world", "ed"))
	fmt.Println(strings.TrimSpace(" hello world   "))

	// 字符串比较
	s1, s2, s3 := "a", "b", ""
	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s3))
	fmt.Println(strings.Compare(s1, "a"), strings.Compare("12", "23"))

	fmt.Println(strings.Repeat(s1, 2)) // 重复多少次
	fmt.Println(strings.Replace("hellollllll", "l", "a", 2))

	fmt.Println(strings.Join([]string{"h", "a", "b"}, ""))
}
