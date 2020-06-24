package utils

import (
	"time"
	"unicode/utf8"
)

const (
	TimeLayout = "2006-01-02 15:04"
	DateLayout = "2006-01-02T15:04:05Z"
)

// 检查任务名称
func CheckTaskName(txt string) (string, bool) {
	var msg string
	ok := true
	nameLength := utf8.RuneCountInString(txt)
	if nameLength == 0 {
		msg = "任务名不能为空"
		ok = false
	}
	if nameLength > 32 {
		msg = "任务名称长度不能超过32位"
		ok = false
	}
	return msg, ok
}

// 检查截止日期
func CheckDeadline(deadlineTime string) (string, bool) {
	var msg string
	ok := true
	dt, err := time.Parse(TimeLayout, deadlineTime)
	if err != nil {
		msg = "日期格式不能为空!"
		ok = false
	} else {
		if dt.Before(time.Now()) {
			msg = "截止日期不能小于当前时间"
			ok = false
		}
	}
	return msg, ok
}

// 检查任务描述
func CheckContent(cnt string) (string, bool) {
	var msg string
	ok := true
	contentLength := utf8.RuneCountInString(cnt)
	if contentLength > 256 {
		msg = "任务描述不能超过256个字符"
		ok = false
	}
	return msg, ok
}
