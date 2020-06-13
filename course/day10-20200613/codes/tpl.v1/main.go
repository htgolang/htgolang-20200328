package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	tplText := "{{ . }}"
	tpl := template.Must(template.New("tpl").Parse(tplText))
	tpl.Execute(os.Stdout, "kk")

	fmt.Println()

	// 切片
	tpl.Execute(os.Stdout, []int{1, 2, 3, 4, 5})

	fmt.Println()
	// 映射
	tpl.Execute(os.Stdout, map[string]int{"one": 1, "two": 2})

	fmt.Println()

	// 结构体
	tpl.Execute(os.Stdout, struct {
		ID   int
		Name string
	}{1, "kk"})
}
