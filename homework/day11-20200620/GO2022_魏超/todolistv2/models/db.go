package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func MySQLDB(device, dsn string) error {
	var (
		err error
	)
	db, err = gorm.Open(device, dsn)
	if err != nil {
		return err
	} else if db.DB().Ping() != nil {
		return errors.New("数据库无法链接")
	} else {
		db.AutoMigrate(&User{}, &Task{})
	}
	return nil
}
