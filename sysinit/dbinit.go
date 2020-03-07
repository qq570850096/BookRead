package sysinit

import (
	_ "Mybook/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func dbinit(aliases... string) {

	// 是否为开发模式
	isDev := ("dev" == beego.AppConfig.String("runmode"))
	orm.Debug = isDev

	if len(aliases) > 0 {
		for _,alias := range aliases {
			registerDB(alias)
			// 主库自动建表
			if "w" == alias {
				orm.RunSyncdb("default",false,isDev)
			}
		}
	} else {
		registerDB("w")
		// 主库自动建表
		orm.RunSyncdb("default",false,isDev)
	}
}

func registerDB(alias string) {
	if len(alias) == 0 {
		return
	}
	dbAlias := alias // default
	if "w" == alias || "default" == alias {
		dbAlias = "default"
		alias = "w"
	}

	// 数据库配置
	dbName := beego.AppConfig.String("db_"+alias+"_database")
	dbUser := beego.AppConfig.String("db_"+alias+"_username")
	dbPwd := beego.AppConfig.String("db_"+alias+"_password")
	dbHost := beego.AppConfig.String("db_"+alias+"_host")
	dbPort := beego.AppConfig.String("db_"+alias+"_port")

	orm.RegisterDataBase(dbAlias,"mysql",dbUser+":"+dbPwd+
		"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8")

}
