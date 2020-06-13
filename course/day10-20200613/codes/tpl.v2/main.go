package main

import (
	"html/template"
	"os"
)

func main() {
	tplText := "{{ index . 1 }}-{{ index . 0 }}"
	tpl := template.Must(template.New("tpl").Parse(tplText))

	// 切片
	tpl.Execute(os.Stdout, []int{1, 2, 3, 4, 5})
}
