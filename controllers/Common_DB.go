package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

//初始化数据库连接
func InitDb() {
	database := beego.AppConfig.String("mysqldb")
	username := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	dburls := beego.AppConfig.String("mysqlurls")

	//orm.RegisterDriver("mymysql", orm.DR_MySQL)

	orm.RegisterDataBase("default", "mysql", ""+username+":"+password+"@tcp("+dburls+")/"+database+"?charset=utf8", 30)

}
