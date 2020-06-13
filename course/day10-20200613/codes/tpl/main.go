package main

import (
	"html/template"
	"os"
)

func main() {
	tpl := template.Must(template.ParseGlob("html/*.html"))
	tpl.ExecuteTemplate(os.Stdout, "index.html", []int{1, 2, 3})
}
