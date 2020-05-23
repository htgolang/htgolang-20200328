package commands

import (
	"fmt"
	"strconv"
	"todolist/commands/command"
	"todolist/utils/ioutils"
	"todolist/views"
)

type manager struct {
	loginCallback command.LoginCallback
	cmds          map[int]*command.Command
}

func newManager() *manager {
	return &manager{
		cmds: make(map[int]*command.Command),
	}
}

func (mgr *manager) registerLoginCallback(callback command.LoginCallback) {
	mgr.loginCallback = callback
}

func (mgr *manager) register(name string, callback command.Callback) {
	mgr.cmds[len(mgr.cmds)+1] = command.New(name, callback)
}

func (mgr *manager) prompt() {
	views.CommandView.Menu(mgr.cmds)
}

func (mgr *manager) get(key int) (command.Callback, error) {
	if cmd, ok := mgr.cmds[key]; ok {
		return cmd.Callback, nil
	}
	return nil, fmt.Errorf("指令不存在")
}

func (mgr *manager) run() {
	if mgr.loginCallback != nil {
		if !mgr.loginCallback() {
			return
		}
	}
	for {
		mgr.prompt()
		key, err := strconv.Atoi(ioutils.Input("请输入指令: "))
		if err != nil {
			ioutils.Error("输入指令错误")
			continue
		}

		if callback, err := mgr.get(key); err != nil {
			ioutils.Error(err.Error())
		} else {
			func() {
				defer func() {
					if err := recover(); err != nil {
						ioutils.Error(fmt.Sprintf("%s", err))
					}
				}()
				callback()
			}()
		}
	}
}

var mgr = newManager()

func RegisterLoginCallback(callback command.LoginCallback) {
	mgr.registerLoginCallback(callback)
}

func Register(name string, callback command.Callback) {
	mgr.register(name, callback)
}

func Run() {
	mgr.run()
}
