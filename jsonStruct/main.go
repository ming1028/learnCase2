package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	p := person{
		Name: "张三",
		Age:  22,
	}
	jsonP, _ := json.Marshal(p)
	respJson := "{\"name:\":\"李四\",\"age\":50}"
	json.Unmarshal([]byte(respJson), &p)
	fmt.Println(string(jsonP), p) // 调用对应的String包

	pt := reflect.TypeOf(p)
	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Println(sf.IsExported())
		fmt.Printf("字段%s, json tag为：%s\n", sf.Name, sf.Tag.Get("json"))
	}

	pv := reflect.ValueOf(p)
	pt = reflect.TypeOf(p)
	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")
	num := pt.NumField()
	for i := 0; i < num; i++ {
		jsonTag := pt.Field(i).Tag.Get("json")
		jsonBuilder.WriteString("\"" + jsonTag + "\"")
		jsonBuilder.WriteString(":")
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i)))
		if i < num-1 {
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String())

	mPrint := pv.MethodByName("Print")
	fmt.Println(mPrint)
	args := []reflect.Value{reflect.ValueOf("登陆")}
	mPrint.Call(args)
}

func (p person) String() string {
	return fmt.Sprintf("Name is %s, Age is %d", p.Name, p.Age)
}

type person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func (p person) Print(prefix string) {
	fmt.Printf("%s:Name is %s,Age is %d\n", prefix, p.Name, p.Age)
}
