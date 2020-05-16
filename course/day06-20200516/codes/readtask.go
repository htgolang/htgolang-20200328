package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 定义常量 时间格式字符串
const TimeLayout = "2006-01-02 15:04:05"

func time2str(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(TimeLayout)
}

// 定义task结构体
type Task struct {
	id        int
	name      string
	status    int
	startTime *time.Time
	endTime   *time.Time
	user      string
}

// 将Task写入到file对象中
func (task *Task) WriteTo(file *os.File) error {
	_, err := fmt.Fprintf(file,
		"%d,%s,%d,%s,%s,%s\n",
		task.id,
		task.name,
		task.status,
		time2str(task.startTime),
		time2str(task.endTime),
		task.user,
	)
	return err
}

// 解析字符串为task结构体指针对象， 如果解析错误返回error
func ParseTask(line string) (*Task, error) {
	// 字符串分割
	nodes := strings.Split(line, ",")
	if len(nodes) != 6 {
		return nil, errors.New("数据量不正确")
	}

	// 字符串转换为int
	id, err := strconv.Atoi(nodes[0])
	if err != nil {
		return nil, err
	}
	name := nodes[1]

	// 字符串转换为int
	status, err := strconv.Atoi(nodes[2])
	if err != nil {
		return nil, err
	}
	var startTime, endTime *time.Time
	if nodes[3] != "" {
		// 字符串转化为时间类型
		if t, err := time.Parse(TimeLayout, nodes[3]); err != nil {
			return nil, err
		} else {
			startTime = &t
		}
	}
	if nodes[4] != "" {

		// 字符串转化为时间类型
		if t, err := time.Parse(TimeLayout, nodes[4]); err != nil {
			return nil, err
		} else {
			endTime = &t
		}
	}

	// 创建结构体指针并返回
	return &Task{
		id:        id,
		name:      name,
		status:    status,
		startTime: startTime,
		endTime:   endTime,
		user:      nodes[5],
	}, nil
}

func main() {
	// 打开文件，读方式
	file, _ := os.Open("user.txt")
	defer file.Close()

	tasks := make([]*Task, 0, 100)

	//定义scanner读取每一行内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // 扫描一行

		//读取一行并通过ParseTask进行解析
		if task, err := ParseTask(scanner.Text()); err == nil {
			tasks = append(tasks, task) // 解析任务成功，加入到tasks
		}
	}

	// 打开文件
	outfile, _ := os.Create("user2.txt")
	defer outfile.Close()

	// 将task写入到文件及标准输出中
	for _, task := range tasks {
		task.WriteTo(outfile)
		task.WriteTo(os.Stdout)
	}
}
