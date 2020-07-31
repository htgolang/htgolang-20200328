package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Task struct {
	ID           int        `orm:"column(id)"`
	Name         string     `orm:"size(64)"`
	StartTime    *time.Time `orm:"column(start_time)"`
	CompleteTime *time.Time `orm:"column(complete_time);null"`
	DeadlineTime *time.Time `orm:"column(deadline_time);"`
	Status       int        `orm:""`
	Content      string     `orm:"null"`
	User         int        `orm:""`
	DeletedAt    *time.Time `orm:"column(deleted_at);null"`
}

func init() {
	orm.RegisterModel(new(Task))
}

func (task *Task) StatusText() string {
	switch task.Status {
	case 0:
		return "新建"
	case 1:
		return "开始"
	case 2:
		return "暂停"
	case 3:
		return "完成"
	}
	return "未知"
}

func (task *Task) RealUser() string {
	return GetNameById(task.User)
}
