package task
import (
	"time"
)

type Task struct {
	Id int
	Name string
	StartTime *time.Time
	EndTime *time.Time
	Status int
	*User
}

func NewTask(id int,name string,startTime *time.Time,status int,user *User) *Task  {
	return &Task{
		Id:id,
		Name:name,
		StartTime: startTime,
		Status: status,
		User:user,
	}
}

func (task *Task)GetIdTask() int  {
	return task.Id
}
func (task *Task)SetNameTask(name string)  {
	task.Name = name
}
func (task *Task)GetNameTask() string  {
	return task.Name
}
func (task *Task)SetstartTimeTask(startTime *time.Time)  {
	task.StartTime = startTime
}
func (task *Task)GetstartTimeTask() *time.Time  {
	return task.StartTime
}
func (task *Task)SetendTime(endTime *time.Time)   {
	task.EndTime = endTime
}
func (task *Task)GetendTimeTask() *time.Time  {
	return task.EndTime
}
func (task *Task)SetStatusTask(status int)  {
	task.Status = status
}
func (task *Task)GetStatusTask () int  {
	return task.Status
}