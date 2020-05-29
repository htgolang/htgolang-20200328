package commands

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"myTodolist/commands/command"
	"myTodolist/utils/ioutils"
	"os"
	"strconv"
)

//定义一个manager结构体，结构体的属性cmds是一个map，map内的key是command对应的选择项，value则是具体的command
type manager struct {
	cmds map[int]*command.Command
	loginCallBack command.LoginCallBack
}

//定义一个New函数,返回一个manager实例类型指针
func NewMgr() *manager {
	return &manager{
		//给属性cmds初始化
		cmds: make(map[int]*command.Command),
	}
}
//定义一个 注册 方法，函数功能：将LoginCallBack添加到mgr中
func (mgr *manager) registerLoginCallBack(callback command.LoginCallBack)  {
	mgr.loginCallBack = callback
}

//创建一个管理器
var mgr = NewMgr()

//定义一个 注册 方法，函数功能：将command添加到manager.cmds中
func (mgr *manager) register(name string, callback command.CallBack) {
	//key是mgr.cmds数据结构内的元素长度+1，value就是Command
	mgr.cmds[len(mgr.cmds)+1] = command.NewCommand(name,callback)
}


//定义一个提示(说明)函数,打印出选项数字和选项名字
func (mgr *manager) prompt() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"编号","功能说明"})
	for i := 1;i<=len(mgr.cmds);i++ {
		//fmt.Printf("%d. %s\n",i, mgr.cmds[i].Name)
		table.Append([]string{strconv.Itoa(i),mgr.cmds[i].Name})
	}
	table.Render()
}

//定义一个函数，返回的是 回调函数 和可能发生的错误,key有可能不存在，此时返回错误
func (mgr *manager) getFunc(key int) (command.CallBack,error) {
	//获取到对应的回调函数并返回
	if cmd, ok := mgr.cmds[key];ok {
		return cmd.CallBack,nil
	}else {
		return nil,fmt.Errorf("不存在此选项!")
	}
}

/*
定义一个运行方法,接收器是manager类型指针，然后对接收器做一系列的操作。
	1. 获取manager.cmds的key(command的name)
	2. 根据1中的name来执行回调函数，这样就可以操作manager了。
*/

//根据用户的输入获取(getFunc())回调函数并执行
func (mgr *manager) run() {
	if mgr.loginCallBack != nil {
		if !mgr.loginCallBack() {
			//ioutils.Input("密码输入错误！")
			return
		}
	}

	for {
		//调用提示(说明)函数
		mgr.prompt()
		//打印出说明后，提示用户输入选项（选项数字）,调用utils.ioutils.Input函数
		key, err := strconv.Atoi(ioutils.Input("请输入功能编号："))
		if err != nil {
			//当发生错误时则输出，调用自定义的output输出
			ioutils.Error("输入错误!!!")
			continue
		}
		//如果err为空则说明用户输出的正确，那么则根据用户的输入取出对应的回调函数,并执行
		if callback, err := mgr.getFunc(key); err != nil {
			/*
			err不为空时，调用自定义ioutils.Error
			err.Error()返回的是字符串
			源码:
			type error interface {
				Error() string
			}
			*/
			ioutils.Error(err.Error())
		}else {
			//执行回调函数
			callback()
		}

	}
}


//定义一个注册函数，注意：此函数是暴漏在外，给别人(包外)调用的。实际依然调用的是程序内部的register函数
func Register(name string, callback command.CallBack) {
	//实际调用的依然是内部register函数
	mgr.register(name,callback)
}

//定义一个注册函数，注意：此函数是暴漏在外，给别人(包外)调用的。
func RegisterLoginCallBack(callback command.LoginCallBack) {
	//实际调用的依然是内部registerLoginCallBack函数
	mgr.registerLoginCallBack(callback)
}


//同理，也需要暴漏一个外部使用的Run()函数，给别人(包外)调用
func Run() {
	//实际调用的依然是内部run函数
	mgr.run()
}