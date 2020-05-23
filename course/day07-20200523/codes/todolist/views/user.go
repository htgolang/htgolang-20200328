package views

import (
	"strings"
	"todolist/utils/ioutils"
)

func LoginTitle() {
	ioutils.Output(strings.Repeat("*", 25))
	ioutils.Output("*       用户登录        *")
	ioutils.Output(strings.Repeat("*", 25))
}
