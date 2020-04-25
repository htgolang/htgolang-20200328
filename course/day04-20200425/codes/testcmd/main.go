package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "-n", "10", "www.baidu.com") // 定义
	// bytes, err := cmd.Output()                              // 执行
	// fmt.Println(string(bytes), err)
	output, err := cmd.StdoutPipe()

	cmd.Start()
	fmt.Println(err)
	io.Copy(os.Stdout, output)
	cmd.Wait()

	// linux /bin/bash -c "a | b| c|d"
}
