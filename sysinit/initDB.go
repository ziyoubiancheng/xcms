package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/ziyoubiancheng/xcms/models"
)

func initDB() {
	//连接名称
	dbAlias := beego.AppConfig.String("db_alias")
	//数据库名称
	dbName := beego.AppConfig.String("db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String("db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String("db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String("db_host")
	//数据库端口
	dbPort := beego.AppConfig.String("db_port")
	//数据库编码
	dbCharset := beego.AppConfig.String("db_charset")

	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset, 30)

	isDev := (beego.AppConfig.String("runmode") == "dev")
	if isDev {
		orm.Debug = isDev
	}
}
