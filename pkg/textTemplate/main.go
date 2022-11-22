package main

import (
	"os"
	"text/template"
)

var templateStr = `
package main
func main() {
	fmt.Println("Hello, {{.Name}}!")
}
`

func main() {
	// 创建template对象
	tpl := template.New("demo")
	// 解析模板内容
	parse, err := tpl.Parse(templateStr)
	if err != nil {
		panic(err)
	}
	data := map[string]string{
		"Name": "world",
	}
	parse.Execute(os.Stdout, data)
}
