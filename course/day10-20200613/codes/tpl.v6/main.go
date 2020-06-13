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
	tplText := `
		{{ range .}}
			{{ .ID }}-{{ .Name }}-{{ if eq .Sex 1 }}男{{ else }}女{{ end }} {{ .Addr.Street}}
		{{ end }}
	`
	tpl := template.Must(template.New("tpl").Parse(tplText))
	// 结构体
	tpl.Execute(os.Stdout, []struct {
		ID   int
		Name string
		Sex  int // 1=>男 0 => 女
		Addr Addr
	}{{1, "kk", 1, Addr{"西安", 1111}}, {2, "何森", 1, Addr{"海淀区", 1000}}})
}
