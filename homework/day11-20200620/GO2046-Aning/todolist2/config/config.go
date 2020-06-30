package config

import "time"

//连接数据库信息
const (
	DbDriver = "mysql"
	DbUser   = "gostudy"
	DbPasswd = "123456q!"
	DbName   = "gostudy"
	DbHost   = "120.79.60.117"
	DbPort   = 3306
)

//web服务端口
const (
	ListenAdd = ":8989"
)

//查询数据库SQL
const (
	SqlTask     = "select task.id,task.name,task.status,task.start_time,task.complete_time,task.deadline_time,user.name as user  from task left join user on task.user=user.id"
	SqlCreate   = "insert into task(name,content,deadline_time) value(?,?,?)"
	SqlDelete   = "delete from task where id = ?"
	SqlEditTask = "update task set name=?, start_time=?, deadline_time=?, content=? where id =?"
	SqlGetTask  = "select id, name , content, start_time, complete_time, deadline_time,user from task where id = ?"
)

//时间格式
const (
	DatatimeLayout = "2006-01-02 15:04:05"
)

var (
	StatusMap = map[int]string{
		0: "新建",
		1: "正在进行",
		2: "暂停",
		3: "完成",
	}
)

type TaskForm struct {
	ID           int
	Name         string
	Status       int
	StartTime    string
	DeadlineTime string
	Content      string
	User         int
}

//实体化task
type Task struct {
	ID           int
	Name         string
	Status       int
	StartTime    *time.Time
	CompleteTime *time.Time
	DeadlineTime *time.Time
	User         *string
	Content      string
}
