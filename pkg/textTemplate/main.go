package main

import (
	"os"
	"text/template"
)

var templateStr = `
package main
func main() {
	fmt.Println("Hello, {{.Name}}!")
	fmt.Println("Hello, {{key "val"}}")
}
`
var textFunc template.FuncMap = template.FuncMap{
	"key": func(key string) string {
		return key + "funcMap"
	},
}

func main() {
	// 创建template对象
	tpl := template.New("demo")
	// 解析模板内容
	parse, err := tpl.Funcs(textFunc).Parse(templateStr)
	if err != nil {
		panic(err)
	}
	data := map[string]string{
		"Name": "world",
	}
	parse.Execute(os.Stdout, data)
}
