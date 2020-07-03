package models

import "time"

type Task struct {
	ID           int    `gorm:"primary_key"`
	Name         string `gorm:"type:varchar(32);not null;default:''"`
	Status       int
	StartTime    *time.Time
	CompleteTime *time.Time
	DeadlineTime *time.Time
	UserID       int    `gorm:"type:int(11);"`
	Describe     string `gorm:"type:text"`
}

var StatusMap = map[int]string{0: "新建", 1: "正在进行", 2: "暂停", 3: "完成"}

func (t Task) CreateTask() error {
	return db.Create(&t).Error
}

func (t Task) UpdateTask() error {
	if err := db.Where("id = ?", t.ID).Error; err == nil {
		return db.Save(&t).Error
	} else {
		return err
	}
}

func (t Task) DeleteTask() error {
	if err := db.Where("id = ?", t.ID).Error; err == nil {
		return db.Delete(&t).Error
	} else {
		return err
	}
}

func (t *Task) GetTaskById(taskId int) error {
	return db.Where("id = ?", taskId).First(t).Error
}

func GetTasks() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}
