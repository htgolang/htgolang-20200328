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

func copyfile(src,desc string) {
	srcfile,err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	} else {
		defer srcfile.Close()
		descfile,err := os.Create(desc)
		if err != nil {
			fmt.Println(err)
		}else {
			defer descfile.Close()
			if err != nil {
				fmt.Println(err)
			}else {
				bytes := make([]byte,1024*1024)
				reader := bufio.NewReader(srcfile)
				writer := bufio.NewWriter(descfile)
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
}

//copy目录
func copyDir(src ,dest string) {
	fileinfo,err := ioutil.ReadDir(src)
	//dest不存在，创建dest，存在忽略  不打印报错
	os.Mkdir(dest,0644)
	//首先判断是否有报错
	if err != nil {
		//判断报错是否是因为src目录或者文件不存在引起的
		if os.IsNotExist(err){
			fmt.Printf("%v不存在",src)
		}else {
			//如果不是因为不存在引起的那么打印报错
			fmt.Println(err)
		}
		//如果没有报错的话执行
	}else {
		//遍历目录信息
		for _,files := range fileinfo{
			//如果src/xxx  xxx是目录  是目录就递归调用
			if files.IsDir(){
				copyDir(filepath.Join(src,files.Name()),filepath.Join(dest,files.Name()))
			}else {
				//如果xxx不是目录 就复制文件
				copyfile(filepath.Join(src,files.Name()),filepath.Join(dest,files.Name()))
			}
		}
	}
}



func main() {
	src  := flag.String("s","","src file")
	desc := flag.String("d","","desc file")
	help := flag.Bool("h",false,"help")
	flag.Usage = func() {
		fmt.Println("Usage: copyfile -s xxx  -d xxx")
		flag.PrintDefaults()
	}
	//解析参数
	flag.Parse()
	if *src == "" || *desc == "" || *help  {
		flag.Usage()
	} else {
		file,err := os.Stat(*src)
		//判断src是否有报错
		if err != nil {
			//判断是否因为src不存在报错
			if os.IsNotExist(err){
				fmt.Printf("%v 不存在\n",*src)
			}else {
				//如果不是不存在打印报错信息
				fmt.Println(err)
			}
		}else {
			//src是目录
			if file.IsDir() {
				copyDir(*src,*desc)
			}else {
				//src是文件
				copyfile(*src,*desc)
			}
		}
	}
}
