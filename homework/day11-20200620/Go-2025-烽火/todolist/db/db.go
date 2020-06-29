package db

import (
	"database/sql"
	"fmt"
	"os"
	"todolist/config"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

func newMySQL() *MySQL {
	// 数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	db, err := sql.Open(config.DbDriver, dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 测试数据库连接
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &MySQL{
		DB: db,
	}

}

var Config = newMySQL()
