package utils

import (

	// "cmdb/models"

	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	TimeLayout = "2006-01-02 15:04"
	DateLayout = "2006-01-02 15:04:05"
)

// // 检查任务名称
// func CheckTaskName(txt string) error {
// 	nameLength := utf8.RuneCountInString(txt)
// 	if nameLength == 0 {
// 		return errors.New("任务名不能为空")
// 	}
// 	if nameLength > 32 {
// 		return errors.New("任务名称长度不能超过32位")
// 	}
// 	return nil
// }

// // 检查截止日期
// func CheckDeadline(deadlineTime string) error {
// 	dt, err := time.Parse(base.TimeLayout, deadlineTime)
// 	if err == nil {
// 		if dt.Before(time.Now()) {
// 			return errors.New("截止日期不能小于当前时间")
// 		}
// 	} else {
// 		return errors.New("日期格式不能为空!")
// 	}
// 	return nil
// }

// 检查任务描述
// func CheckContent(cnt string) error {
// 	contentLength := utf8.RuneCountInString(cnt)
// 	if contentLength > 256 {
// 		return errors.New("任务描述不能超过256个字符")
// 	}
// 	return nil
// }

// func CheckUserName(name string) error {
// 	var username string

// 	if len(name) <= 4 {
// 		return errors.New("用户名长度小于4位")
// 	}

// 	row := models.TDB.QueryRow(base.SqlQueryUser, name)
// 	err := row.Scan(&username)
// 	if err == nil {
// 		return errors.New("用户名已存在")
// 	} else {
// 		return nil
// 	}
// }

// func CheckUserPassword(pass1, pass2 string) (string, error) {
// 	if pass1 == pass2 {
// 		password := fmt.Sprintf("%x", md5.Sum([]byte(pass1)))
// 		return password, nil

// 	} else {
// 		return "false", errors.New("两次输入的密码不匹配")
// 	}
// }

func String2Time(t string) *time.Time {
	tt, err := time.Parse(TimeLayout, strings.ReplaceAll(t, "T", " "))
	if err != nil {
		return nil
	}
	return &tt
}

func GeneratePassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		panic(err.Error())
	}
	return string(hashPassword)
}

func CheckPassword(password, hashPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)) == nil
}
