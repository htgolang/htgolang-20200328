package todo

import "time"

//创建一个todo结构体
type todo struct {
	Id string
	Name string
	StartTime *time.Time
	EndTime *time.Time
	Status string
	User string
}

//封装一个函数，返回todo类型指针对象
func Newtodo(id , name, status, user string, startTime,endTime *time.Time) *todo {
	return &todo{
		Id: id,
		Name: name,
		Status: status,
		User: user,
		StartTime: startTime,
		EndTime: endTime,
	}
}