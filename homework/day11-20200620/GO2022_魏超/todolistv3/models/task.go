package models

import "time"

type StatusRelation struct {
	Name     string
	Relation []int
}

var StatusMap = map[int]StatusRelation{
	1: StatusRelation{Name: "新建", Relation: []int{1, 2}},
	2: StatusRelation{Name: "正在进行", Relation: []int{2, 3, 4}},
	3: StatusRelation{Name: "暂停", Relation: []int{2, 3, 4}},
	4: StatusRelation{Name: "已完成", Relation: []int{4}},
}

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

func (t Task) TableName() string {
	return "tasks"
}

// var StatusMap = map[int]string{0: "新建", 1: "正在进行", 2: "暂停", 3: "完成"}

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

func (t *Task) GetTaskById() error {
	return db.Where("id = ?", t.ID).First(t).Error
}

func GetTasks() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}
