package main

import (
	"html/template"
	"os"
)

func main() {
	tplText := "{{ .one }}-{{ .two }}-{{ .three }}"
	tpl := template.Must(template.New("tpl").Parse(tplText))

	// 映射
	tpl.Execute(os.Stdout, map[string]int{"one": 1, "two": 2})
}
