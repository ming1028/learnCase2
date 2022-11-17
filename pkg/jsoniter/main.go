package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	p1 := Person{
		Name: "张三",
		Age:  22,
		Hobby: Hobby{
			Names: "basketball",
		},
	}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))
	p2 := new(Person)
	/*err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(p1Json, p2)
	if err != nil {
		return
	}*/
	jsoniter.Unmarshal(p1Json, p2)
	fmt.Println(p2)
	fmt.Println(jsoniter.Get(p1Json, "names").ToString())
	nilStruct, _ := json.Marshal(struct{}{})
	fmt.Println(string(nilStruct))
}

type Person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Hobby
}

type Hobby struct {
	Names string `json:"names"`
}
