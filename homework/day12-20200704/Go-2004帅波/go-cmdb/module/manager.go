package module

import (
	"github.com/jinzhu/gorm"
)


type Manager struct {
	gorm.Model
	Name string

}

type Job struct {
	gorm.Model
	Name string
	Status int
}
