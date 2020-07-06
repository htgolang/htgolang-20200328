package models

import (
	"errors"
	"fmt"
	"log"

	"todolist/global"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	if err := ConnDB(); err != nil {
		log.Fatal(err)
	}
}

func ConnDB() error {
	var (
		err error
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
			global.Config.Database.UserName,
			global.Config.Database.Password,
			global.Config.Database.Host,
			global.Config.Database.Port,
			global.Config.Database.DBName,
		)
	)
	db, err = gorm.Open(global.Config.Database.Device, dsn)
	db.LogMode(true)
	if err != nil {
		return err
	} else if db.DB().Ping() != nil {
		return errors.New("数据库无法链接")
	} else {
		// 针对不同
		switch global.Config.Database.Device {
		case "mysql":
			err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET = utf8mb4").AutoMigrate(&User{}, &Task{}).Error
		default:
			err = nil
		}
		if err != nil {
			log.Printf("auto migrate data failed, %s\n", err)
		}

	}
	log.Printf("connection %s database %s:%d\n", global.Config.Database.Device, global.Config.Database.Host, global.Config.Database.Port)
	return nil
}
