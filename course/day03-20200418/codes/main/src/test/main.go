package main

import (
	"fmt"
	"testpkg"

)

var mainVar = "main Var"

func mainFunc() {
	fmt.Println("main func")
}

func main() {
	mainFunc()
	utilsFunc()
	fmt.Println(mainVar)
	fmt.Println(utilsVar)

	fmt.Println(pkg.UtilsVar)
}

// GO PATH 项目

// GOPATH GOMODULE
// GO包
// 1.同一个文件夹下所有go文件的包名 必须一致

// 2. 关闭了 GOMODULE
// GOPATH 在项目目录直接运行 go build 无文件名
// 将当前文件夹下的所有go文件进行编译
// 3. main 包编译为可执行程序
// 4. main包中只能有一个main函数

// GOPATH 环境变量信息, 定义多个目录
// src ==> 源文件
// pkg ==> 程序编译的包文件
// bin ==> 程序编译的可执行文件

// 编译程序 go build 项目文件路径(src)
// 导入包 使用包的文件路名名
// 调用函数/数据， 使用报名.变量名/函数名

// vendor
