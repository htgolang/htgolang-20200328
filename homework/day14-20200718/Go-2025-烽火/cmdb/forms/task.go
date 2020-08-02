package forms

type FormTask struct {
	ID           int    `form:"id"`
	Name         string `form:"name"`
	StartTime    string `form:"start_time"`
	CompleteTime string `form:"complete_time"`
	DeadlineTime string `form:"deadline_time"`
	Status       int    `form:"status"`
	Content      string `form:"content"`
	User         int    `form:"user"`
}
