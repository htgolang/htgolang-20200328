package main

import (
	"html/template"
	"os"
	"strings"
)

type Addr struct {
	Street string
	No     int
}

func main() {
	tplText := `{{ upper . }}-{{ title . }}`

	// 自定义函数
	// key 字符串 函数名称 在模板中调用使用
	// value 函数
	funcs := template.FuncMap{
		"upper": strings.ToUpper,
		"title": func(text string) string {
			if len(text) == 0 {
				return ""
			} else if len(text) == 1 {
				return strings.ToUpper(text)
			}
			return strings.ToUpper(text[:1]) + text[2:]
		},
	}

	tpl := template.Must(template.New("tpl").Funcs(funcs).Parse(tplText))
	// 结构体
	tpl.Execute(os.Stdout, "a")
}
