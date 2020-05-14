package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var (
		src  string
		dest string
		help bool
		h    bool
	)
	flag.StringVar(&src, "s", "", "src file")
	flag.StringVar(&dest, "d", "", "dest file")
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&help, "help", false, "help")

	flag.Usage = func() {
		fmt.Println("Copy -s srcfile -d destfile")
		flag.PrintDefaults()
	}

	flag.Parse()
	if help || h || src == "" || dest == "" {
		flag.Usage()
		return
	}
	if _, err := os.Stat(dest); err == nil {
		fmt.Println("目的文件已经存在")
		return
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("目的文件获取错误", err)
		}
	}
	if info, err := os.Stat(src); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("源文件不存在")
		} else {
			fmt.Println("源文件获取错误：", err)
		}
	} else {
		if info.IsDir() {
			copyDir(src, dest)
		} else {
			copyFile(src, dest)
		}
	}
}

func copyFile(src, dest string) {
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	} else {
		defer srcfile.Close()
		destfile, err := os.Create(dest)
		if err != nil {
			fmt.Println(err)
		} else {
			defer destfile.Close()

			reader := bufio.NewReader(srcfile)
			writer := bufio.NewWriter(destfile)

			bytes := make([]byte, 1024)

			for {
				n, err := reader.Read(bytes)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
					}
					break
				}
				writer.Write(bytes[:n])
				writer.Flush()
			}
		}
	}
}

func copyDir(src, dest string) {
	files, err := ioutil.ReadDir(src)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				copyDir(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
			} else {
				copyFile(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
			}
		}
	}
}
