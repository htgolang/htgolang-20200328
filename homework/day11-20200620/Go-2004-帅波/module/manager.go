package module

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Manager struct {
	gorm.Model
	Name string
	StartTime time.Time
	StopTime time.Time
	User User
	Job []Job
}

type Job struct {
	gorm.Model
	Name string
	Status int
	User User
	Manager Manager
}
