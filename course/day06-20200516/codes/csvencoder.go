package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Task struct {
	id        int
	name      string
	status    int
	startTime *time.Time
	endTime   *time.Time
	user      string
}

func time2str(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:06")
}

func main() {

	now := time.Now()
	end := now.Add(time.Hour * 24)

	tasks := []*Task{
		{
			id:        1,
			name:      "整理课程笔记",
			status:    0,
			startTime: &now,
			endTime:   &end,
			user:      "kk",
		},
		{
			id:        2,
			name:      "整理课程笔记2",
			status:    0,
			startTime: &now,
			endTime:   &end,
			user:      "kk",
		},
	}

	// 创建文件
	file, _ := os.Create("task.csv")
	defer file.Close()

	// 创建带缓冲IO的写对象
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// 创建csv写对象
	csvWriter := csv.NewWriter(writer)

	// 写入csv头
	csvWriter.Write([]string{"ID", "名称", "状态", "开始时间", "结束时间", "执行者"})

	// 循环写入csv数据
	for _, task := range tasks {
		csvWriter.Write([]string{
			strconv.Itoa(task.id),
			task.name,
			strconv.Itoa(task.status),
			time2str(task.startTime),
			time2str(task.endTime),
			task.user,
		})
	}
}
