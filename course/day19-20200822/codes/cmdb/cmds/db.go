package cmds

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cobra"
)

var (
	db    string
	force bool
)

var dbCommand = &cobra.Command{
	Use:   "db",
	Short: "db console",
	Long:  "db console",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		return orm.RunSyncdb(db, force, verbose)
	},
}

func init() {
	rootCommand.AddCommand(dbCommand)
	dbCommand.Flags().StringVarP(&db, "database", "d", "default", "database")
	dbCommand.Flags().BoolVarP(&force, "force", "f", false, "force syncdb")
}
