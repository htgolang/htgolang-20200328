package main

import (
	"html/template"
	"os"
)

func main() {
	tplText := `
		{{ .ID }}-{{ .Name }}
		{{ if eq .Sex 1 }}男{{ else }}女{{ end }}
	`
	tpl := template.Must(template.New("tpl").Parse(tplText))
	// 结构体
	tpl.Execute(os.Stdout, struct {
		ID   int
		Name string
		Sex  int // 1=>男 0 => 女
	}{1, "kk", 0})
}
