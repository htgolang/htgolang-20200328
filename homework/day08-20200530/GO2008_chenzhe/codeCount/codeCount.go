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

var (
	h bool
	filename string
	ch chan int
	wg sync.WaitGroup
	wg2 sync.WaitGroup
	lines int
)

func init() {
	flag.BoolVar(&h, "h", false, "Help information")
	flag.StringVar(&filename, "f", "", "dirname or filename end with .go")

}

func readLine(pathname string)int  {
	var total int = 0
	flag := false
	f,_ := os.OpenFile(pathname,os.O_RDONLY,os.ModePerm)
	bufferReader := bufio.NewReader(f)
	for {
		line,err := bufferReader.ReadString('\n')
		if err != nil{
			if err == io.EOF{
				break
			}
		}
		line = strings.TrimSpace(strings.TrimSuffix(line,"\n"))
		if strings.HasSuffix(line,`/*`){
			flag = true
		}
		if strings.HasSuffix(line,`*/`){
			flag = false
		}
		if line == ""{
			continue
		}
		if !flag{
			switch strings.HasPrefix(line,`//`) {
			case true:
			case false:
				total+=1

			}
		}
	}
	return total
}

func run()  {
	ch = make(chan int)
	flag.Parse()
	if h||filename ==""{
		flag.PrintDefaults()
	}
	//filename ="codeCount.go"
	filepath.Walk(filename,func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path,`.go`){
			wg.Add(1)
			go func() {
				num := readLine(path)
				ch <- num
				wg.Done()
			}()
	}
		return nil
	})

	go func() {
		wg2.Add(1)
		for i:= range ch{
			lines+=i
		}
		wg2.Done()
	}()

	wg.Wait()
	close(ch)

	wg2.Wait()
	fmt.Println(lines)

}
func main() {
	run()
}
