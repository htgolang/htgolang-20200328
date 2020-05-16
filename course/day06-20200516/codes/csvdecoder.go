package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

type Task struct {
	id        int
	name      string
	status    int
	startTime *time.Time
	endTime   *time.Time
	user      string
}

func ParseTask(nodes []string) (*Task, error) {
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
	return nil, nil
}

func main() {

	tasks := make([]*Task, 0, 100)

	// 打开文件
	file, _ := os.Open("task.csv")
	defer file.Close()

	// 定义带缓冲IO 读对象
	reader := bufio.NewReader(file)

	// 创建csv读对象
	csvReader := csv.NewReader(reader)
	csvReader.Read() // 读取头信息

	// 循环读取文件内容
	for {
		// 读取csv一行数据
		line, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(line, err)
			break
		}

		// 解析为task
		if task, err := ParseTask(line); err == nil {
			tasks = append(tasks, task)
		} else {
			fmt.Println(err)
		}

	}

	// 打印task信息
	for _, task := range tasks {
		fmt.Printf("%#v\n", task)
	}
}
