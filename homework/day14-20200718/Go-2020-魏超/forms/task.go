package forms

import (
	"cmdb/models"
	"time"

	"github.com/astaxie/beego/validation"
)

const (
	// DateTimeLayout 时间格式
	DateTimeLayout = "2006-01-02 15:04:05"
)

// TaskForm 任务表单
type TaskForm struct {
	ID           int    `form:"id" valid:"Numeric"`                // 不可以为空;必须是数字;
	Name         string `form:"name" valid:"Required;MaxSize(64)"` // 不可以为空;长度检查不可以超过64字节;
	StatusID     int    `form:"status_id" valid:"Numeric;"`        // 必须是数字;满足
	StartTime    string `form:"start_time"`                        // 当不为空时,时间必须满足时间格式
	CompleteTime string `form:"complete_time"`                     // 当不为空时,时间必须满足时间格式
	DeadlineTime string `form:"deadline_time" valid:"Required"`    // 不可以为空;时间必须满足时间格式
	UserID       int    `form:"user_id" valid:"Numeric"`           // 必须是数字;id的有效性
	Describe     string `form:"describe" valid:"MaxSize(1024)"`    // 长度检查不可以超过1024字节;
}

// Valid 做form表单自定义数据验证
func (t *TaskForm) Valid(v *validation.Validation) {
	// 检查statusID的有效性
	if _, ok := models.TaskStatusMap[t.StatusID]; !ok {
		v.SetError("StatusID", "task status code illegal")
	}

	// 检查开始时间的有效性
	if t.StartTime != "" {
		if _, err := time.Parse(DateTimeLayout, t.StartTime); err != nil {
			v.SetError("StartTime", "datetime format err")
		}
	}

	// 检查完成时间的有效性
	if t.CompleteTime != "" {
		if _, err := time.Parse(DateTimeLayout, t.CompleteTime); err != nil {
			v.SetError("CompleteTime", "datetime format err")
		}
	}

	// 检查截止时间的有效性
	if _, err := time.Parse(DateTimeLayout, t.DeadlineTime); err != nil {
		v.SetError("DeadlineTime", "datetime format err")
	}

	// 检查用户的id是否存在
	if u := models.GetUserByID(t.UserID); u == nil {
		v.SetError("UserID", "User not exist")
	}
}
