package main

import (
	// "fmt"
	"os"
	"html/template"
)

func main() {
	// 定义一个名为template test 的模板
	t := template.New("template test")

	// 解析模板
	t = template.Must(t.Parse("demo1: {{ if `` }} none {{ end }}\n"))
	t.Execute(os.Stdout, nil)

	// 定义一个名为template test 的模板
	t = template.New("template test")

	// 解析模板
	t = template.Must(t.Parse("demo2: {{ if `true` }} anything {{ end }}\n"))
	t.Execute(os.Stdout, nil)
}