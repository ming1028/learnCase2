package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	_, err = add(-1, 2)
	if err != nil {
		fmt.Printf("%#v\n", err.Error())
	}
	_, err = add(11, 2)
	fmt.Printf("%#v\n", err.Error())
	// 类型断言
	if emError, ok := err.(*commonError); ok {
		fmt.Println(emError.ErrorCode, emError.ErrorMsg, emError.Error())
	}

	newErr := MyError{err, "错误!"}
	//fmt.Printf("%#v\n", newErr)
	// fmt.Println类型断言，查看是否实现error接口和string接口，如果实现就会调用
	fmt.Println(newErr)
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("不能为负！")
	}

	if a > 10 {
		return 0, &commonError{
			ErrorMsg:  "error",
			ErrorCode: 200,
		}
	}
	return a + b, nil
}

type commonError struct {
	ErrorCode int
	ErrorMsg  string
}

func (ce *commonError) Error() string {
	return ce.ErrorMsg
}

type MyError struct {
	err error
	msg string
}

func (me MyError) Error() string {
	return me.err.Error() + me.msg
}
