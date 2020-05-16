package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func FileExt(path string) string {
	pos := strings.LastIndex(path, ".")
	if pos < 0 {
		return ""
	}
	return path[strings.LastIndex(path, "."):]
}

func main() {
	path, _ := filepath.Abs("./filepath.go")
	fmt.Println(path)
	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.Dir(path))

	fmt.Println(filepath.Clean("./.../..../////abc/abc"))

	fmt.Println(filepath.Ext(path))

	fmt.Println(FileExt(path))
	fmt.Println(filepath.FromSlash("./.../..../////abc/abc"))
	fmt.Println(filepath.ToSlash("./.../..../////abc/abc"))

	path2, _ := filepath.Abs("./test")
	fmt.Println(filepath.HasPrefix(path, path2))
	fmt.Println(filepath.IsAbs(path))
	fmt.Println(filepath.IsAbs("."))

	dir, _ := filepath.Abs("c://opt//cmdb")
	fmt.Println(filepath.Join(dir, "etc", "app.ini"))

	fmt.Println(filepath.Split(path))
	paths := "c://test/a;c://test/b;c://test/c"
	fmt.Println(filepath.SplitList(paths))

	fmt.Println(filepath.Glob("./test/a*"))
	fmt.Println(filepath.Glob("./test/*.txt"))
	fmt.Println(filepath.Match("./test/a.go", "./test/a.go"))
	fmt.Println(filepath.Match("./test/a.*", "./test/a.go"))
	fmt.Println(filepath.Match("./test/b.*", "./test/a.go"))

}
