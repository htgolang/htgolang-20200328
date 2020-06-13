package main

import (
	"html/template"
	"os"
)

func main() {
	tplText := "{{ .ID }}-{{ .Name }}"
	tpl := template.Must(template.New("tpl").Parse(tplText))
	// 结构体
	tpl.Execute(os.Stdout, struct {
		ID   int
		Name string
	}{1, "kk"})
}
