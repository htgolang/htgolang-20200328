package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

const (
	id        = "id"
	name      = "name"
	startTime = "starttime"
	endTime   = "endtime"
	status    = "status"
	user      = "user"
)
const (
	statusNew     = "未执行"
	statusStart   = "准备开始"
	statusDoding  = "进行中"
	statusCompele = "完成"
)

const (
	// salt     = "pY5L3%"                           //salt
	// password = "ebc44309e71b754e60376ed2815e36ff" //centos
	password = "b0230d6ac1b3b3bcff28ace36d15ad5d" // hello
	salt     = "02P8bA"
)

var todos = []map[string]string{
	{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "2", "name": "备课", "startTime": "21:00", "endTime": "", "status": statusNew, "user": "kk"},
	{"id": "3", "name": "复习", "startTime": "09:00", "endTime": "", "status": statusNew, "user": "kk"},
}

//查询输入排序方式  升序
func sorttask(task map[string]string, way string) map[string]string {
	if way == "name" || way == "startTime" {
		sort.Slice(task, func(i, j int) bool { return task[i][way] < task[j][way] })
	}
	return task
}

//初始化种子
func init() {
	rand.Seed(time.Now().Unix())
}

//生产随机字符串  生成的字符当作salt
func randString(n int) string {
	//定义一个字节切片
	saltt := make([]byte, n, n)
	//随机字符集合
	chars := "!,./[]@#$%^&*()_+=-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 循环n次 每次生成随机数(切片范围内)，获取对应的字符
	for i := 0; i < n; i++ {
		saltt = append(saltt, chars[rand.Intn(len(chars))])
	}
	return string(saltt)

}

//md5 加密
func md5String(text string, salt string) string {
	//salt + ":" + text
	bytes := []byte(salt)
	bytes = append(bytes, ':')
	bytes = append(bytes, []byte(text)...)
	return fmt.Sprintf("%x", md5.Sum(bytes))
}

// 加盐MD5
func saltMd5(input string) string {
	hasher := md5.New()
	io.WriteString(hasher, input)
	io.WriteString(hasher, salt)

	cryptoPasswd := fmt.Sprintf("%x", hasher.Sum(nil))
	return cryptoPasswd
}

//隐藏密码
func getpasswd() string {
	fmt.Println("please input password: ")
	passwd, _ := gopass.GetPasswdMasked()
	return string(passwd)
}

//authpasswd  3次失败推出
func authpasswd() bool {
	for i := 0; i <= 3; i++ {
		input := getpasswd()
		if saltMd5(input) == password {
			return true
		} else {
			fmt.Println("密码验证错误，还剩%d次机会!\n", 3-i)
		}
	}
	return false
}
func getid() int {
	var big int
	for _, todo := range todos {
		todoid, _ := strconv.Atoi(todo["id"])
		if big < todoid {
			big = todoid
		}
	}
	return big + 1
}
func newtask() map[string]string {
	task := make(map[string]string)
	task["id"] = strconv.Itoa(getid())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}
func printtask(task map[string]string) {
	fmt.Printf("ID：%s\n", task[id])
	fmt.Printf("NAME：%s\n", task[name])
	fmt.Printf("STARTTIME：%s\n", task[startTime])
	fmt.Printf("ENDTIME：%s\n", task[endTime])
	fmt.Printf("STATUS：%s\n", task[status])
	fmt.Printf("USER：%s\n", task[user])
}
func input(inputt string) string {
	var inputtt string
	fmt.Println(inputt)
	fmt.Println()
	fmt.Scan(&inputtt)
	return strings.TrimSpace(inputtt)
}
func add() {
	task := newtask()
	fmt.Println("输入信息")
	task[name] = input("任务名:")
	task[startTime] = input("开始时间")
	task[user] = input("负责人")
	todos = append(todos, task)
	fmt.Println("创建成功")
}
func edit() {
	q := input("请输入需要修改的任务ID:")
	for _, task := range todos {
		if q == task["id"] {
			print(task)
			switch input("是否确认修改(y/yes):") {
			case "y", "yes":
				tempName := input("任务名称:")
				task[name] = tempName
				task[startTime] = input("开始时间:")
				qq := input("状态:")
				task[endTime] = time.Now().Format("2006-01-02 15:04:05")
				task[status] = qq
				print(task)
			default:
				break
			}
		}
	}
}

func del() {
	q := input("要删除的")
	for index, task := range todos {
		if q == task["id"] {
			print(task)
			switch input("是否进行删除(y/yes):") {
			case "y", "yes":
				copy(todos[index:], todos[index+1:])
				newTasks := todos[:len(todos)-1]
				for _, task := range newTasks {
					print(task)
				}
			}
		}
	}
}
func search() {
	//定义搜索排序切片
	sortt := make(map[string]string, 0)
	id := input("输入查询")
	if id == "all" {
		for _, j := range todos {
			fmt.Println(strings.Repeat("-", 20))
			printtask(j)
			fmt.Println(strings.Repeat("-", 20))
		}
	} else {
		for i, j := range todos {
			if id == j["id"] {
				printtask(j)
				break
			}
			if i == len(todos)-1 {
				fmt.Println("没有这个id")
			}
		}
		if len(sortt) == 0 {
			fmt.Println("无任务")
		} else {
			wayy := input("请输入排序方式[time |  startTime]:")
			wayytask := sorttask(sortt, wayy)
			printtask(wayytask)
		}
	}
}
func exit() {
	os.Exit(0)
}
func main() {
	if !authpasswd() {
		fmt.Println("3次密码验证错误，程序退出")
		os.Exit(1)
	}
	op := map[string]func(){
		"add":    add,
		"edit":   edit,
		"del":    del,
		"search": search,
		"exit":   exit,
	}
	for {

		inputt := input("1.add    2.edit    3.del	   4.search   5.exit")
		fmt.Scan(&inputt)
		if opp, ok := op[inputt]; ok {
			opp()
		} else {
			fmt.Println("没有选项")
		}
	}
}

/*
问题  ：
.\todopasswd.go:48:53: cannot use i (type int) as type string in map index
.\todopasswd.go:48:56: non-integer string index way
.\todopasswd.go:48:68: cannot use j (type int) as type string in map index
.\todopasswd.go:48:71: non-integer string index way
46 行的自己定义的函数 类型是string  sort.Slice 要求是int  咋转换。。。



我写的太乱了  奔溃
*/
