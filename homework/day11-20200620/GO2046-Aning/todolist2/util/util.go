package utils

import (
	"database/sql"
	"fmt"
	"log"
	"todolist/config"

	_ "github.com/go-sql-driver/mysql"
)

func MySQL() *sql.DB {
	//连接数据库库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=PRC", config.DbUser, config.DbPasswd, config.DbHost, config.DbPort, config.DbName)
	db, err := sql.Open(config.DbDriver, dsn)

	if err != nil {
		log.Fatal(err)
	}
	//判断数据库连通性
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
