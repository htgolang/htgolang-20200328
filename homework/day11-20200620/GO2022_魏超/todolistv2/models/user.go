package models

import "time"

type User struct {
	ID           int    `gorm:"primary_key"`
	Name         string `gorm:"type:varchar(32);not null;default:''"`
	Account      string `gorm:"type:varchar(32);not null;default:''"`
	Tel          int64
	Passwd       string `gorm:"type:varchar(64);not null;default:''"`
	Address      string `gorm:"type:varchar(255);default:''"`
	RegisterTime *time.Time
}

func (u User) CreateUser() {
	db.Create(&u)
}

func (u User) UpdateUser() error {
	if err := db.Where("id = ?", u.ID).Error; err == nil {
		return db.Save(&u).Error
	} else {
		return err
	}
}

func (u User) DeleteTask() error {
	if err := db.Where("id = ?", u.ID).Error; err == nil {
		return db.Delete(&u).Error
	} else {
		return err
	}
}

func (u *User) GetTaskById() error {
	return db.Where("id = ?", u.ID).First(u).Error
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}
