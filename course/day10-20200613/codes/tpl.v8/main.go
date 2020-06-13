package main

import (
	"html/template"
	"os"
)

type Addr struct {
	Street string
	No     int
}

func main() {
	tplText := `输入内容: {{ block "content" . }} {{ . }} {{ end }}`

	tpl := template.Must(template.New("tpl").Parse(tplText))

	tpl2, _ := template.Must(tpl.Clone()).Parse(`{{ define "content" }} {{ len . }} {{ end }}`)
	// 结构体
	tpl.Execute(os.Stdout, "abcdef")
	tpl2.Execute(os.Stdout, "abcdef")
}
