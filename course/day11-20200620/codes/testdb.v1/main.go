package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// user:password@tcp(host:port)/database?charset=utf8mb4&loc=PRC&parseTime=true
	dsn := "golang:golang@2020@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&loc=PRC&parseTime=true" // 字符串的格式由对应的驱动进行定义
	db, err := sql.Open("mysql", dsn)
	fmt.Println(db, err)
	fmt.Println(db.Ping())

	// 执行
	// sql => go 字符串
	fmt.Println(db.Exec(`
		create table if not exists testkk(
			id bigint primary key auto_increment,
			name varchar(32) not null default '' comment 'testkk名字'
		) engine=innodb default charset utf8mb4;
	`))

	status := "1"
	sql := `update task set status = ` + status

	result, err := db.Exec(sql)
	fmt.Println(sql, err)
	fmt.Println(result.RowsAffected())

	id := "2 or 1=1"
	result, _ = db.Exec(`update task set status = 2 where id=` + id)

	fmt.Println(result.RowsAffected())

	// result, _ = db.Exec(`delete from task where id=1`)
	// fmt.Println(result.RowsAffected())

	// tname := "买个电视"
	// content := ""
	// deadline := "2020-10-10')"

	// sql = fmt.Sprintf(`insert into task(name, content, deadline_time) values('%s', '%s', '%s')`, tname, content, deadline)

	// fmt.Println(sql)
	// result, err = db.Exec(sql)
	// fmt.Println(err)
	// fmt.Println(result.LastInsertId())
	// fmt.Println(result.RowsAffected())
	// // 读

	// rows, err := db.Query("select id,name from task")
	// var (
	// 	id   int
	// 	name string
	// )
	// for rows.Next() {
	// 	rows.Scan(&id, &name)
	// 	fmt.Println(id, name)
	// }

	// 1. 配些字符串太麻烦
	// 2. 拼写字符串会错 ' 2 or 1=1(漏洞)
}
