package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var(
	h bool
	src string
	dest string
	wg sync.WaitGroup
	st [][]string
)

func init() {
	flag.BoolVar(&h, "h", false, "Help information")
	flag.StringVar(&src, "s", "", "source filenanme")
	flag.StringVar(&dest, "d", "", "destination filename")

}


func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func fileCopy(src, dest string)  {
	f1,_ := os.OpenFile(src,os.O_RDONLY,os.ModePerm)
	f2,_ := os.OpenFile(dest,os.O_WRONLY|os.O_CREATE|os.O_APPEND,os.ModePerm)
	newWriter := bufio.NewWriter(f2)
	io.Copy(newWriter,f1)
	newWriter.Flush()
}

func run()  {
	flag.Parse()
	if h||src ==""||dest==""{
		flag.PrintDefaults()
	}
	//src=`G:\rds\homework`
	//dest=`G:\rds\homework\test`

	if srcExist,_ := PathExists(src);!srcExist{
		fmt.Printf("%s don't exist\n",src)
		flag.PrintDefaults()
	}
	if destExist,_ := PathExists(dest);destExist{
		if fileStat,_ := os.Stat(dest);fileStat.IsDir(){
			dest = filepath.Join(dest,filepath.Base(src))
		}else {
			fmt.Printf("%s has existing and not dir\n",dest)
			flag.PrintDefaults()
		}
	}
	filepath.Walk(src,func(path string, info os.FileInfo, err error) error {
		dest := strings.Replace(path,src,dest,-1)
		if fileStat,_ := os.Stat(path);fileStat.IsDir(){
			os.MkdirAll(dest,os.ModePerm)
		}else {
			wg.Add(1)
			go func() {
				fileCopy(path,dest)
				wg.Done()
			}()
		}

		return nil
	})

	wg.Wait()

}
func main() {
	run()
}
