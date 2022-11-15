package main

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"unicode/utf8"
	"unsafe"
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

	a1 := [2]string{"张三", "李四"}
	s1 := a1[0:1]
	s2 := a1[:]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)
	sli1 := make([]string, 0, 4)
	// 待定
	expanSlice(sli1)
	fmt.Println(sli1)

	s := "张三"
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&b)).Data)
	s3 := string(b)
	fmt.Printf("b的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)

	// 大字符串内存开销大，转换思路：不重新申请内存的情况下实现类型转换
	s = "赵钱孙李"
	b = []byte(s)
	s4 := *(*string)(unsafe.Pointer(&b))
	fmt.Println(s4)
	fmt.Printf("%T, %v\n", s4, s4)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	b1 := *(*[]byte)(unsafe.Pointer(sh)) // b1 和b内容一样，b1没有申请内存，跟s使用同一内存
	fmt.Println(string(b1))              // b1不可修改，共用内存，string不可修改
	var sa, sc []string
	sb := []string{"ff", "aa"}
	fmt.Println(len(sa), sa == nil)
	fmt.Println(reflect.DeepEqual(sa, sb), reflect.DeepEqual(sa, sc))

	// 测试题
	var num1 []interface{}
	num2 := []int{1, 2, 3}
	num3 := append(num1, num2)
	fmt.Println(len(num3), num3)
}

func expanSlice(sli []string) {
	fmt.Println("扩容：", len(sli), cap(sli))
	for i := 0; i < 1; i++ {
		sli = append(sli, "王五"+cast.ToString(i))
	}
	// append 然后新slice
	fmt.Println(sli)
}

func StringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
