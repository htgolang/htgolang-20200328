package forms

type TaskForm struct {
	ID           int    `form: "ID"`
	Name         string `form: "name"`
	Status       int    `form: "status"`
	StartTime    string `form: "startTime"`
	CompleteTime string `form: "completeTime"`
	DeadlineTime string `form: "deadlineTime"`
	Content      string `form: "content"`
	User         int    `form: "user"`
}
