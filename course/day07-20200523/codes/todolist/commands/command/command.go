package command

type Callback func()

type LoginCallback func() bool

type Command struct {
	Name     string
	Callback Callback
}

func New(name string, callback Callback) *Command {
	return &Command{name, callback}
}
