package init

import (
	"promagent/task"
	"promagent/task/plugins/profile"
	"promagent/task/plugins/register"
)

func init() {
	task.Register("register", &register.Register{})
	task.Register("profile", &profile.Profile{})
}
