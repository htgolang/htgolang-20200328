package cmds

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cobra"

	"cmdb/config"
)

var webCommand = &cobra.Command{
	Use:   "web",
	Short: "Web console",
	Long:  "Web console",
	RunE: func(cmd *cobra.Command, args []string) error {
		beego.SetLogger("file", `{"filename" : "logs/cmdb.log"}`)
		beego.SetLogFuncCall(true)
		beego.SetLevel(beego.LevelDebug)
		if !verbose {
			beego.BeeLogger.DelLogger("console")
		}

		config.Init("redis", `{"key":"cmdb:cache","conn":"10.0.0.2:6379","dbNum":"0","password":"golang@2020"}`)

		orm.Debug = verbose
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/cmdb?charset=utf8mb4&loc=PRC&parseTime=true",
			beego.AppConfig.DefaultString("mysql::User", "golang"),
			beego.AppConfig.DefaultString("mysql::Password", "golang@2020"),
			beego.AppConfig.DefaultString("mysql::Host", "127.0.0.1"),
			beego.AppConfig.DefaultInt("mysql::Port", 3306),
		)

		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", dsn)

		// 测试数据库连接
		if db, err := orm.GetDB("default"); err != nil {
			return err
		} else if err := db.Ping(); err != nil {
			return err
		}

		beego.Run()
		return nil
	},
}

func init() {
	rootCommand.AddCommand(webCommand)
}
