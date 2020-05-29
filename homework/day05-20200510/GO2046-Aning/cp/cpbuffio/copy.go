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

//复制到文件夹及文件
func copydir(src, dst string) {
	filesrc, err := ioutil.ReadDir(src)
	if err == nil {
		for _, file := range filesrc {
			if file.IsDir() {
				copydir(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			} else {
				copyfile(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			}
		}
	}
}

//拷贝文件内容
func copyfile(src, dst string) {
	//第一步打开来文件
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	} else {
		defer srcfile.Close()
		//创建目标文件
		dstfile, err := os.Create(dst)
		if err != nil {
			fmt.Println(err)
		} else {
			defer dstfile.Close()
			//reader  wirter 读取写入
			reader := bufio.NewReader(srcfile)
			writer := bufio.NewWriter(dstfile)

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
func main() {
	//定义参数
	var (
		src string
		dst string
		h   bool
	)
	flag.StringVar(&src, "s", "", "src file")
	flag.StringVar(&dst, "d", "", "dst file")
	flag.BoolVar(&h, "h", false, "help")

	//提示信息
	flag.Usage = func() {
		fmt.Println("cp -s [srcfile] -d [dstfile]")
		flag.PrintDefaults()
	}

	//参数解析
	flag.Parse()
	if h || src == " " || dst == " " {
		flag.Usage()
		return
	}

	//判断目的文件
	if _, err := os.Stat(dst); err == nil {
		fmt.Println("目的文件已经存在")
		return
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("目的文件获取错误", err)
		}
	}

	//判断源文件
	if srcfiles, err := os.Stat(src); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("源文件不存在")
		} else {
			fmt.Println("源文件获取错误：", err)
		}
	} else {
		if srcfiles.IsDir() {
			copydir(src, dst)
		} else {
			copyfile(src, dst)
		}
	}

	// dstfile, errd := os.Stat(dst)
	// if errd == nil {
	// 	if os.IsNotExist(errd) {
	// 		fmt.Printf("%v 不存在\n", dstfile)
	// 	} else {
	// 		if !os.IsNotExist(errd) {
	// 			fmt.Println("dst file 获取失败", errd)
	// 		}
	// 	}
	// }
}
