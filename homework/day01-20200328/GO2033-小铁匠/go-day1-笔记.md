go： 易于开发，快速编译，高效执行

特性：
- 静态类型并且有丰富的内置类型
- 函数多返回值
- 错误处理机制
- 语言层并发
- 面向对象： 使用类型、组合、接口来实现面向对象思想
- 反射
- CGO：用于调用C语言实现的模块
- 自动垃圾回收
- 静态编译
- 交叉编译
- 易于部署


go开源项目列表：
https://github.com/golang/go/wiki/Projects

## go环境安装
下载地址：https://golang.org/dl/


配置第三方代理：
https://goproxy.io/zh/

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```
使用go version和go env命令查看安装的版本及环境变量

#### linux环境安装

```
[root@localhost local]# tar -C /usr/local -zxvf go1.13.9.linux-amd64.tar.gz
[root@localhost local]# tail -n3 /etc/profile
## go
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
[root@localhost local]# source /etc/profile
[root@localhost local]# go version
go version go1.13.9 linux/amd64
[root@localhost local]# go env -w GO111MODULE=on
[root@localhost local]# go env -w GOPROXY=https://goproxy.io,direct

```

## hello world

```
//声明包名main
package main
//导入fmt包
import "fmt"
//定义main函数
func main() {
	fmt.Println("Hello World!")//调用fmt的Println打印字符串到控制台
}

```
```
PS C:\Users\hqq\Desktop\go\course\day1> go run .\helloworld.go
Hello World!
PS C:\Users\hqq\Desktop\go\course\day1> go build .\helloworld.go
PS C:\Users\hqq\Desktop\go\course\day1> .\helloworld.exe
Hello World!
```
编译：

```
go build filename

-x 查看过程
-o 指定编译后的文件名
-work 打印临时目录
```
go run：

```
-x
-n 不会真正执行
```
其他相关命令，go help查看

## 注释
- 行注释： //
- 块注释： /*  */


```go
//声明包名main
package main
//导入fmt包
import "fmt"
//定义main函数
func main() {
	/*
	  这是一个注释
	  我正在学习golang
	*/
	fmt.Println("Hello World!") //调用fmt的Println打印字符串到控制台
}
```

## 变量

变量定义: var + VAR_NAME + TYPE = VALUE

简单示例：
```go
package main

import "fmt"

func main() {

	var msg string = "hello hqq"
	fmt.Println(msg)
	fmt.Println(msg)

}
```

### 变量的作用域
- 包级别的变量
- 函数级别的变量
- 块级别的变量

注意：
- 大括号定义变量的作用域
- 子块可以使用父块的变量
- 父块不能使用子块的变量
- 子块变量会覆盖父块变量
- 函数和块内的变量必须使用，不使用则报错

示例1：
```go
package main

import (
	"fmt"
)

//包级别
var packageVar = "package var"

func main() {
	// 函数级别
	var funcVar string = "func var"

	fmt.Println(packageVar, funcVar)

}
```
示例2：父块不能使用子块的变量
```go
package main

import (
	"fmt"
)

//包级别
var packageVar = "package var"

func main() {
	// 函数级别
	var funcVar string = "func var"
	// 块级别
	{
		var blockVar string = "block var"
		fmt.Println(packageVar, funcVar, blockVar)
		{
			var innerBlockVar string = "innerBlock var"
			fmt.Println(packageVar, funcVar, blockVar, innerBlockVar)
		}
	}
	fmt.Println(packageVar, funcVar, blockVar, innerBlockVar)

}

父调用子变量，执行报错未定义：
PS C:\Users\hqq\Desktop\go\course\day1> go run .\block.go
# command-line-arguments
.\block.go:22:35: undefined: blockVar
.\block.go:22:45: undefined: innerBlockVar
```

示例3：子块变量会覆盖父块的变量


```
package main

import (
	"fmt"
)

//包级别
var packageVar = "package var"

func main() {
	// 函数级别
	var funcVar string = "func var"
	var packageVar string = "func package var"
	// 块级别
	{
		var blockVar string = "block var"
		fmt.Println(packageVar, funcVar, blockVar)
		{
			var innerBlockVar string = "innerBlock var"
			fmt.Println(packageVar, funcVar, blockVar, innerBlockVar)
		}
	}
	fmt.Println(packageVar, funcVar)

}
```


```go
package main

import (
	"fmt"
)

//包级别
var packageVar = "package var"

func main() {
	// 函数级别
	var packageVar string = "func package var"
	// 块级别
	{
		var packageVar string = "block package var"

		fmt.Println("1", packageVar)
	}
	fmt.Println("2", packageVar)

}
执行结果：
1 block package var
2 func package var
```

### 变量的定义方式
最完整的定义方式：
 var + VAR_NAME + TYPE = VALUE
 
- 定义变量并初始化值
- 定义变量但不初始化值
- 可以省略类型，但不能省略初始化值。根据值推导变量的类型
- 短声明： VAR_NAME := VALUE

注意：
- 短声明必须在函数中使用，不能在包级别中使用

示例：

```
package main

import (
	"fmt"
)

func main() {
	// 定义了类型并初始化值
	var name string = "hqq"
	// 定义了类型但不初始化值,值为空字符串("")
	var zeroString string

	// 定义变量省略类型，但不能省略值
	var typeString = "sky" //通过值类型推导变量的类型

	//短声明，通过值类型推导变量的类型，只能在函数中使用
	shortString := "world"

	fmt.Println(name, zeroString, typeString, shortString)

}
```
### 变量合并
```go
package main

import "fmt"

func main() {
	/*
		var name string = "hqq"
		var msg string
		var desc = "haha"
	*/
	var (
		name string = "hqq"
		msg  string
		desc = "haha"
	)
	/*
		x := "x"
		y := "y"
	*/
	x, y := "x", "y"

	fmt.Println(name, msg, desc, x, y)
}
```
### 变量的赋值（更新）

```
package main

import "fmt"

func main() {

	var msg string = "hello hqq"
	fmt.Println(msg)
	msg = "hello world"
	fmt.Println(msg)

}

hello hqq
hello world
```

子块级别可修改父块的变量值
```
package main

import "fmt"

func main() {

	var msg string = "hello hqq"
	fmt.Println(msg)
	msg = "hello world"
	fmt.Println(msg)

	{   //此处为赋值，修改了父级别的变量值
		msg = "hello block"
		//如果此处为定义变量，则不会影响父级别
		// var msg = "hello block"
	}
	fmt.Println(msg)
}
```
## 基本组成元素
### 标识符

标识符： 程序中定义的名字，变量名，常量名字，函数名字，自定义类型，接口，包名

规范：
1. 必须满足： 组成只能由非空的Unicode编码字符串、数字、下划线组成
2. 必须满足： 必须以unicode编码的字符串或下划线开头（不能以数字开头）
3. 必须满足：不能与go的关键字冲突（package，func，var... 25个）

建议：
1. Ascill编码（a-z,A-Z）,数字，下划线
2. 变量使用驼峰式
3. 与go内置的标识符不要冲突

说明： 标识符区分大小写


```go
var 我的名字 = "hqq" //可以，但不建议使用
var _myName = "hqq" // 可以
var 0Name = "hqq" // 不可以数字开头
var package = "hqq" // 与go内置标识符冲突

```
Go语言提供一些预先定义的标识符用来表示内置的常量、类型、函数 在自定义标识符时应避免使用： 
1. 内置常量：true、false、nil、iota 
2. 内置类型：bool、byte、rune、int、int8、int16、int32、int64、uint、uint8、unit16、 unit32、unit64、uintptr、float32、float64、complex64、complex128、string、error
3. 内置函数：make、len、cap、new、append、copy、close、delete、complex、real、 imag、panic、recover 
4. 空白标识符:_

### 关键字
关键字用于特定的语法结构 uGo语言定义25关键字：
- 声明：import、package 
- 实体声明和定义：char、const、func、interface、map、struct、type、var 
- 流程控制：break、case、continue、default、defer、else、fallthrough、 for、go、goto、if、range、return、select、switch

### 字面量
### 常量
1. 常量一旦定义则不能修改
2. 常量定义之后在func中可不使用
```
package main

import (
	"fmt"
)

func main() {

	const msg string = "hello hqq"
	fmt.Println(msg)

	msg = "silence"

}
```

```
package main

import (
	"fmt"
)

const (
	packageName string = "kk"
	packageMsg         = "gg"
)

func main() {

	const msg string = "hello hqq"
	fmt.Println(msg)

	//msg = "silence"
	fmt.Println(packageName, packageMsg)
}
```
定义多个变量并进行初始化，批量复制中 变量类型可省略，并且除了第一个常量值 外其他常量可同时省略类型和值，表示使 用前一个常量的初始化表达式。

```
package main

import (
	"fmt"
)
// const A  第一个不能省略
func main() {
	const (
		A = "test"
		B
		C
		D = "testD"
		E
		F
	)

	fmt.Println(B, C)
	fmt.Println(E, F)
}

test test
testD testD
```
### 枚举类型


```
package main

import (
	"fmt"
)

func main() {
	const (
		A = iota // 在常量中使用iota，初始化值为0，每次调用+1
		B
		C
		E
		F
	)

	fmt.Println(A, B, C, E, F)
}

0 1 2 3 4
```


## 问题跟踪

- Println 会打印换行符
- Print 只打印变量不加换行
- Printf 
  - %T 打印类型
  - %v 打印变量值
  - %#v 
  - %b：二进制 
  - %c：字符 
  - %d：十进制 
  - %+d表示对正整数带+符号
  - %nd表示最小占位n个宽度且右对齐
  - %-nd表示最小占位n个宽度且左对齐
  - %0nd表示最小占位n个宽度且右对齐，空字符使用0填充 
  - %o：八进制，%#o带0的前缀
  - %x、%X：十六进制,%#x(%#X)带0x(0X)的前缀 
  - %U: Unicode码点，%#U带字符的Unicode码点
  - %q：带单引号的字符


```
package main

import (
	"fmt"
)

func main(){
	var name = "kk"
	fmt.Println("*") 
	fmt.Println(name) // 打印变量加换行符
	fmt.Println("*")
	fmt.Print(name) // 打印变量不加换行符
	fmt.Print("*")
	fmt.Println(" ")
	fmt.Printf("%T,%v,%#v",name,name,name)

}

*
kk
*
kk*
string,kk,"kk"
```
## 数据类型
- 布尔类型
  - true
  - false
- 整数类型
  - int, uint
  - int8,int16,int32,int64
  - uint8,uint16,uint32,uint64
  - byte,rune
  - uintptr
- 浮点类型
  - float32 
  - float64
- 复数类型
- 字符串
### 布尔类型
布尔类型的0值是false
```
package main

import (
	"fmt"
)

func main() {
	//isGirl := false
	var isGirl bool = true
	fmt.Printf("%T,%#v", isGirl, isGirl)
}
```
布尔类型常用的操作：
逻辑运算，与，或，非和关系运算


```
package main

import (
	"fmt"
)

func main() {
	//isGirl := false
	var isGirl bool = true
	fmt.Printf("%T,%#v", isGirl, isGirl)

	a, b, c, d := true, true, false, false
	//与
	fmt.Println("a, b:", a && b) // true && true : true
	fmt.Println("a,c:", a && c)  // true && false : false
	fmt.Println("c,b:", c && b)  // false && true: false
	fmt.Println("c,d:", c && d)  // false && false: false

	fmt.Println("a,b:", a || b) // true || true : true
	fmt.Println("a,c:", a || c) // true || false : true
	fmt.Println("c,b:", c || b) // false || true: true
	fmt.Println("c,d:", c || d) // false || false: false

	fmt.Println("a:", !a) // !true: false
	fmt.Println("b:", !b) // !false: true
	fmt.Println("c:", !c)
	fmt.Println("d:", !d)
	// 关系运算

	fmt.Println(a == b)
	fmt.Println(a != c)
}
```

### 整数类型
int类型零值为0


定义
```
package main

import "fmt"

func main() {
	var age8 int8 = 31
	var age int = 31
	var myAge = 31

	fmt.Printf("%T, %#v,%d\n", age8, age8, age8)
	fmt.Printf("%T, %#v,%d\n", age, age, age)
	fmt.Printf("%T, %#v,%d", myAge, myAge, myAge)
}

int8, 31,31
int, 31,31
int, 31,31
```

算数运算

```go
//算数运算 + - * / % ++ --
	a, b := 2, 4
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++
	b--
	fmt.Println(a)
	fmt.Println(b)
```

关系运算
```go
//关系运算 >,<, >=, <=, !=, ==
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a == b)
	fmt.Println(a != b)
```
类型转换
```
//类型转换 type(value) int32(i) int(i)

	var (
		i   int   = 1
		i32 int32 = 1
		i64 int64 = 1
	)

	fmt.Println(i + int(i32))
	fmt.Println(i + int(i64))
	fmt.Println(int32(i) + i32)
	fmt.Println(int64(i) + i64)

}
```


### 原码、反码、补码，负数表示法
- 原码： 二进制
- 反码： 正数将原码按位取反，负数符号位不变其余按位取反
- 补码： 正数的补码是其本身，负数符号位不变其余按位取反后+1
- 负数表示法
  - 数字电路中cpu中的运算器实现了加法器，没有减法，减法是转换成加法
  - 在二进制码中,为了区分正负数,采用最高位是符号位的方法来区分,正数的符号位为0、负数的符号位为1
  - 负数在计算机中是补码存储
  - 计算机的相加都是补码相加
  - print(~12): 12 取反 再补码结果就是-13
  - 5-1=5+(-1)=4: 5的补码是本身0b101,-1为1取反再补码为0b11111111，再二进制相加


### 字节，码点

```
// 字节，码点
	var (
		achar        byte = 'A'
		aint         byte = 65
		unicodePoint rune = '中'
	)
	fmt.Println(achar, aint)
	fmt.Println(unicodePoint)

	fmt.Printf("%d %b %o %x %U %c %c", achar, 15, 15, 15, unicodePoint, achar, 65)
```

```
65 1111 17 f U+4E2D A A
```
### 赋值 += -+ *= /=
```
//赋值 += -+ *= /=
	a += a // a=+a
	fmt.Println(a)
```

## 浮点数
定义
默认是float64
```
package main

import "fmt"

func main() {
	var height float32 = 1.68
	var wight = 75.2

	fmt.Printf("%T %#v %f\n", height, height, height)
	fmt.Printf("%T %#v %f\n", wight, wight, wight)
}
```

```
float32 1.68 1.680000
float64 75.2 75.200000
```
运算

```go
//运算

	var (
		f1 = 1.2
		f2 = 2.36
	)

	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)
	fmt.Println(f1 / f2)
	fmt.Println(f1 + f2)

	f1++
	f2--

	fmt.Println(f1, f2)
```

关系运算

```
// 关系运算
	fmt.Println(f1 > f2)
	fmt.Println(f1 < f2)
	fmt.Println(f1 >= f2)
	fmt.Println(f1 <= f2)
	fmt.Println(f1 != f2)
```

```
true
false
true
false
true
```
## 字符串类型

特殊字符：
- \\\：反斜线 
-  \\'：单引号 
- \\"：双引号 
- \\a：响铃 
- \\b：退格 
- \\f：换页 
- \\n：换行 
- \\r：回车 
- \\t：制表符 
- \\v：垂直制表符
- \\ooo：3个8位数字给定的八进制码点的Unicode字符（不能超过\377） 
- \\uhhhh：4个16位数字给定的十六进制码点的Unicode字符 
- \\Uhhhhhhhh：8个32位数字给定的十六进制码点的Unicode字符 
- \\xhh：2个8位数字给定的十六进制码点的Unicode字符

示例：
```
package main

import (
	"fmt"
)

func main() {
	var msg = "我的名字是hqq"

	fmt.Printf("%T %s\n", msg, msg)
}
```
``中的反斜杠不会解析为转义，原生打印

```
package main

import (
	"fmt"
)

func main() {
	var msg = `我的名字\n是hqq`

	fmt.Printf("%T %s\n", msg, msg)
}
```

```
string 我的名字\n是hqq
```

字符串拼接，+

```
var msg = "我的名字是hqq"
	//拼接
	var msgRaw = "，今天天气不错"

	fmt.Println(msg + msgRaw)
```
关系运算

```
// 关系运算 >,>=,<=,<,!=,==

	fmt.Println("abc" > "acd")
	fmt.Println("abc" == "abc")
```
索引切片

```
// 索引，切片
	msg = "abcdef"
	//索引
	fmt.Printf("%T %#v %c\n", msg[0], msg[0], msg[0])
	//切片
	fmt.Println(msg[1:3])
```

```
uint8 0x61 a
bc
```
len 计算字节大小

```
//len字节大小，不是计算字符的数量
    msg = "abcdef"
	fmt.Println(len(msg)) 
	fmt.Println(len(msgRaw))
```

```
6
21
```

## 数据类型转换
string不能通过TYPE()转换为int或float，需要用strconv方法转换


```go
package main

import (
	"fmt"
	"strconv"
)

func main(){
	var (
		intVal=65
		float64Val=2.2
		stringVal="3.3"
	)
	fmt.Println(intVal,float64Val,stringVal)
	//整型转换为浮点型
	fmt.Printf("%T %#v\n", float64(intVal),float64(intVal) )
	//浮点型转换为整型
	fmt.Printf("%T %#v\n", int(float64Val),int(intVal))

	//整型转换为字符串
	fmt.Println(string(intVal))
	fmt.Printf("%T %#v\n", string(intVal),string(intVal))

	//string不能通过TYPE()转换为int或float，需要用strconv方法转换
	//v, err := strconv.Atoi(stringVal)
	//fmt.Println(err,v)

	vv, err := strconv.ParseFloat(stringVal,64)
	fmt.Println(err, vv)

	fmt.Println(strconv.Itoa(intVal))
	fmt.Println(strconv.FormatFloat(float64Val, 'f', 10, 64))
}
```

```
65 2.2 3.3
float64 65
int 65
A
string "A"
<nil> 3.3
65
2.2000000000
```


## 指针

### 什么是指针
一个指针变量指向了一个值的内存地址。

类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针
```
var var_name *var-type
```
以下是有效的指针声明：

```
var ip *int        /* 指向整型*/
var fp *float32    /* 指向浮点型 */
```

### 如何使用指针
指针使用流程：

1. 定义指针变量
2. 为指针变量赋值
3. 访问指针变量中指向地址的值

示例：


```go
package main

import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
```

```go
package main

import (
	"fmt"
)

func main() {
	//定义指针
	//零值 nil
	var (
		pointerInt    *int
		pointerString *string
	)

	fmt.Printf("%T %#v\n", pointerInt, pointerInt)
	fmt.Printf("%T %#v\n", pointerString, pointerString)

	//赋值
	//取地址 &VAR_NAME
	age := 32
	pointerInt = &age // 使用&取地址并赋值给指针变量pointerInt

	fmt.Printf("%T %#v\n", &age, &age)
	fmt.Printf("%T %#v\n", pointerInt, pointerInt)
	fmt.Printf("%T %#v\n", *pointerInt, *pointerInt) // 在指针类型前面加上 * 号（前缀）来获取指针所指向的内容(字面量的值)

	age2 := age // 将age赋值为age2,age2和age地址不一样
	fmt.Printf("%T %#v\n", &age2, &age2)

	*pointerInt = 33000 // 可通过对*指针变量赋值的方式来修改原变量字面量的值

	fmt.Println(age, age2)
}

*int (*int)(nil)
*string (*string)(nil)
*int (*int)(0xc0000100a8)
*int (*int)(0xc0000100a8)
int 32
*int (*int)(0xc0000100f0)
33000 32
```
### 空指针
当一个指针被定义后没有分配到任何变量时，它的值为 nil。

nil 指针也称为空指针。

nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

一个指针变量通常缩写为 ptr。

查看以下实例：

```
package main

import "fmt"

func main() {
	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)
	fmt.Printf("ptr 的类型为 : %T\n", ptr)
	fmt.Println(ptr)
}
```
```
ptr 的值为 : 0
ptr 的类型为 : *int
<nil>
```
### 通过new赋值

```go
pointerString = new(string)
	fmt.Println(pointerString)
	fmt.Printf("%#v %#v", pointerString, *pointerString)
```
结果为：
```
0xc0000401f0
(*string)(0xc0000401f0) ""
```

### 指针的指针

```
pp := &pointerString
	fmt.Printf("%T\n", pp)
	fmt.Println(**pp)
	**pp = "hqq"
	fmt.Println(**pp)
```
```
(*string)(0xc0000561e0) ""
**string

hqq
```

### scan用户输入
Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
```
package main

import (
	"fmt"
)

func main() {
	name := ""
	var age int
	fmt.Print("请输入你的名字：")
	fmt.Scan(&name)

	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age)

	fmt.Println("你输入的名字是：", name)
	fmt.Println("你的年龄是：", age)
}

```

```
请输入你的名字：hqq
请输入你的年龄：26
你输入的名字是： hqq
你的年龄是： 26
```

## 流程控制

### if else
语法如下：
```
if condition {
   /* 在布尔表达式为 true 时执行 */
}else if condition{
  /* 在布尔表达式为 true 时执行 */
}else {
  /* 在布尔表达式为 false 时执行 */
}
```

示例
```
package main

import (
	"fmt"
)

func main() {
	var y string
	fmt.Print("有卖西瓜的吗：")
	fmt.Scan(&y)

	if y == "yes" {
		fmt.Println("买一个包子")

	} else {
		fmt.Println("买十个包子")
	}

}
```

```
package main

import (
	"fmt"
)

func main() {
	var score float32
	fmt.Print("请输入分数：")
	fmt.Scan(&score)
	fmt.Println("你输入的分数是：", score)

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}
```

### switch
1.switch 语句用于基于不同条件执行不同动作，每一个 case分支都是唯一的，从上至下逐一测试，直到匹配为止。

2.switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。

3.switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。

```
package main

import (
	"fmt"
)

func main() {
	var y string
	fmt.Print("有卖西瓜的吗：")
	fmt.Scan(&y)

	switch y {
	case "yes", "y", "Y":
		fmt.Println("买一个包子")
	case "no", "n", "N":
		fmt.Println("买十个包子")
	default:
		fmt.Println("输入错误")
	}
}
```

```
package main

import (
	"fmt"
)

func main() {
	var score float32
	fmt.Print("请输入分数：")
	fmt.Scan(&score)
	fmt.Println("你输入的分数是：", score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("C")
	case score <= 60:
		fmt.Println("D")
	default:
		fmt.Println("输入错误")
	}
}
```


## 循环
### for

```
package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		//sum = sum + i
		sum += i

	}
	fmt.Println(sum)
}
```

## break continue
break 退出所有循环

continue 退出本次循环


```
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

结果：
1
2
3
4
```

```
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
结果：
1
2
3
4
6
7
8
9
```

### 遍历字符串
打印字符串的每个元素
```
for i := 0; i < len(letters); i++ {
		fmt.Printf("%c\n", letters[i])
	}
```

一般用for range的方式

```
	msg := "今天天气不错"

	for _, v := range msg { //“_”是特殊标识符，用来忽略结果。
		//fmt.Printf("%T %#v %T %#v\n", k, k, v, v)
		fmt.Printf("%q\n", v)
	}
```
匿名变量
1. 下划线为匿名变量
2. 匿名变量不占用命名空间，不会分配内存。
3. 匿名变量与匿名变量之间不会因为多次声明而无法使用。
### loop

### 模拟while

```
package main

import (
	"fmt"
)

func main() {
	var index = 1
	for index <= 10 {
		fmt.Println(index)
		index++
```
}
}











