package utils

import (
	"time"

	"todolist/models"
)

const (
	DateTimeLayout = "2006-01-02 15:04:05"
)

func FormatDatetime(t *time.Time) string {
	if t == nil {
		return "--"
	} else {
		return t.Format(DateTimeLayout)
	}
}

func FormatStatus(statusCode int) string {
	return models.StatusMap[statusCode]
}

func FormatUserID(userID int) string {
	var (
		user models.User
		err  error
	)
	user.ID = userID
	err = user.GetTaskById()
	if err == nil {
		return user.Name
	}
	return "--"
}
