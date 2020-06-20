package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// id 增加 money
func changeMoney(tx *sql.Tx, id int, money float64) error {
	if money < 0 {
		// 检查
		var accountMoney float64
		err := tx.QueryRow("select money from account where id=?", id).Scan(&accountMoney)
		if err != nil {
			return err
		}
		if accountMoney < -money {
			return fmt.Errorf("没有足够的金额")
		}
	}
	_, err := tx.Exec("update account set money=money+? where id=?", money, id)
	return err
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=PRC",
		"golang", "golang@2020", "127.0.0.1", 3306, "todolist")

	/*
		create table account(
			id bigint primary key auto_increment,
			name varchar(32) not null default '',
			money decimal(10, 5) not null default 0
		) engine=innodb default charset utf8mb4;

		insert into account(name, money) values("kk", 1000);
		insert into account(name, money) values("烽火", 1000);
	*/
	db, _ := sql.Open("mysql", dsn)

	// 转账

	var a, b = 1, 2
	// a => b money
	// a - money
	// b + money

	var money float64 = 100.0

	// 同时成功同时失败
	// 事务
	tx, _ := db.Begin()

	err1 := changeMoney(tx, a, -money)
	err2 := changeMoney(tx, b, money)

	fmt.Println(err1, err2)
	if err1 == nil && err2 == nil {
		// 提交事务
		tx.Commit()
	} else {
		// 回滚
		tx.Rollback()
	}

}
