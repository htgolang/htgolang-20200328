package models

import "time"

type User struct {
	ID           int    `gorm:"primary_key"`
	Name         string `gorm:"type:varchar(32);not null;default:''"`
	Sex          int    `gorm:"type:integer(1);not null"`
	Account      string `gorm:"type:varchar(32);not null;unique;default:''"`
	Tel          string `gorm:"type:varchar(16);not null;default=''"`
	Passwd       string `gorm:"type:varchar(255);not null;default:''"`
	Address      string `gorm:"type:varchar(255);default:''"`
	RegisterTime *time.Time
}

var SexMap = map[int]string{0: "女", 1: "男"}

func (u User) TableName() string {
	return "users"
}

func (u User) CreateUser() error {
	return db.Create(&u).Error
}

func (u User) UpdateUser() error {
	if err := db.Where("id = ?", u.ID).Error; err == nil {
		return db.Save(&u).Error
	} else {
		return err
	}
}

func (u User) DeleteUser() error {
	if err := db.Where("id = ?", u.ID).Error; err == nil {
		return db.Delete(&u).Error
	} else {
		return err
	}
}

func (u *User) GetUserById() error {
	return db.Where("id = ?", u.ID).First(u).Error
}

func (u *User) GetUserByAccount() error {
	return db.Where("account = ?", u.Account).First(u).Error
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}
