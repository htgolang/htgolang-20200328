package task
import (
	"time"
)

type Task struct {
	id int
	name string
	startTime *time.Time
	endTime *time.Time
	status int
	*User
}

func NewTask(id int,name string,startTime *time.Time,status int,user *User) *Task  {
	return &Task{
		id:id,
		name:name,
		startTime: startTime,
		status: status,
		User:user,
	}
}

func (task *Task)GetIdTask() int  {
	return task.id
}
func (task *Task)SetNameTask(name string)  {
	task.name = name
}
func (task *Task)GetNameTask() string  {
	return task.name
}
func (task *Task)SetstartTimeTask(startTime *time.Time)  {
	task.startTime = startTime
}
func (task *Task)GetstartTimeTask() *time.Time  {
	return task.startTime
}
func (task *Task)SetendTime(endTime *time.Time)   {
	task.endTime = endTime
}
func (task *Task)GetendTimeTask() *time.Time  {
	return task.endTime
}
func (task *Task)SetStatusTask(status int)  {
	task.status = status
}
func (task *Task)GetStatusTask () int  {
	return task.status
}