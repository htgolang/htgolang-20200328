package main

import (
	"fmt"
	"html/template"
	"os"
)

type Addr struct {
	Street string
	No     int
}

func main() {
	tplText := `
	{{ define "len" }} {{ len . }} {{ end }}
	{{ define "raw" }} {{ . }} {{ end }}

	{{ template "raw" . }}
	`

	tpl := template.Must(template.New("tpl").Parse(tplText))

	// 结构体
	tpl.Execute(os.Stdout, "abcdef")

	tpl.ExecuteTemplate(os.Stdout, "len", "abcdef")
	fmt.Println()
	tpl.ExecuteTemplate(os.Stdout, "raw", "abcdef")
}
