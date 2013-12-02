package main

import (
	"geekxb/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	// 路由设置
	beego.Router("/", &controllers.Admin_IndexController{})

	beego.Router("/Admin_Login", &controllers.Admin_LoginController{})
	beego.Router("/Admin_Index", &controllers.Admin_IndexController{})
	beego.Router("/Admin_User", &controllers.Admin_UserController{})
	//后台管理:用户组
	beego.Router("/Admin_UserGroup", &controllers.Admin_UserGroupController{})
	beego.Router("/Admin_UserGroup/Edit", &controllers.Admin_UserGroupController{}, "get:ActionEdit")
	beego.Router("/Admin_UserGroup/Update", &controllers.Admin_UserGroupController{}, "post:ActionUpdate")
	beego.Router("/Admin_UserGroup/Del", &controllers.Admin_UserGroupController{}, "get:ActionDel")
	//后台管理:权限基本配置
	beego.Router("/Admin_Authority", &controllers.Admin_AuthorityController{})
	beego.Router("/Admin_Authority/Edit", &controllers.Admin_AuthorityController{}, "get:ActionEdit")
	beego.Router("/Admin_Authority/Update", &controllers.Admin_AuthorityController{}, "post:ActionUpdate")
	beego.Router("/Admin_Authority/Del", &controllers.Admin_AuthorityController{}, "get:ActionDel")

	//静态地址设置( 前一个参数必须开头"/"    )
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")

	// 配置参数......
	beego.SessionOn = true

	//连接数据库
	orm.Debug = true
	controllers.InitDb()

	//Cache
	controllers.InitCache()

	beego.Run()
}
