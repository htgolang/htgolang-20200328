package todo

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Task struct {
	Id string
	Name string
	Status string
	StartTime *time.Time
	EndTime *time.Time
	User string
}

func time2str(time *time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func (msTodo *msTodo) csvSaveToFile(f string) {
		//打开文件，读写，追加，如果文件不存在则创建
		file, _ := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		defer file.Close()
		//创建一个io读缓冲流
		reader := bufio.NewReader(file)
		//读取一行
		line, _, _ := reader.ReadLine()

		//创建一个io写缓冲流
		writer := bufio.NewWriter(file)
		defer writer.Flush()
		//创建csv写对象
		csvWriter := csv.NewWriter(writer)

		//如果表头信息存在，则...
		if string(line) == "id,name,startTime,endTime,status,user" {
			//获取 map[string][]*todo
			todolist := msTodo.todoItems
			//遍历 map[string][]*todo
			for _, v := range todolist { //得到每一个 []*todo
				for _, v1 := range v { //得到每一个todo结构体实例
					csvWriter.Write([]string{
						v1.Id,
						v1.Name,
						time2str(v1.StartTime),
						time2str(v1.EndTime),
						v1.Status,
						v1.User,
					})
				}

			}
		} else {
			//如果表头信息不存在则...
			//写入表头信息
			csvWriter.Write([]string{"id","name","startTime","endTime","status","user"})
			//获取 map[string][]*todo
			todolist := msTodo.todoItems
			//遍历 map[string][]*todo
			for _, v := range todolist { //得到每一个 []*todo
				for _, v1 := range v { //得到每一个todo结构体实例
					csvWriter.Write([]string{
						v1.Id,
						v1.Name,
						time2str(v1.StartTime),
						time2str(v1.EndTime),
						v1.Status,
						v1.User,
					})
				}

			}

		}
}

func (msTodo *msTodo) csvReadToFile(f string) {
	//打开文件，读写，追加，如果文件不存在则创建
	file, err := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Println("----------------------------")
	}
	//创建一个io读缓冲流
	reader := bufio.NewReader(file)
	//读取行首
	line, _, _ := reader.ReadLine()
	fmt.Printf("csv文件行首内容：%s\n",string(line))

	//读取io流并将数据添加到msTodo.todoItems中
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		s := string(line)
		sliece1 := strings.Split(s,",")
		startTime, _ := time.Parse("2006-01-02 15:04:05",sliece1[2])
		endTime, _ := time.Parse("2006-01-02 15:04:05",sliece1[3])
		msTodo.todoItems[sliece1[0]] = append(msTodo.todoItems[sliece1[0]],Newtodo(sliece1[0],sliece1[1],sliece1[4],sliece1[5],&startTime,&endTime))
	}
}