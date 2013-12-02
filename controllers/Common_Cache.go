package controllers

import (
	"geekxb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var (
	Cache_All *beego.BeeCache //权限配置缓存
)

//初始化缓存
func InitCache() {

	//初始化权限配置缓存
	Cache_All = beego.NewBeeCache()
	Cache_All.Every = 0 //不过期
	Cache_All.Start()

	setCacheAdminConfigData()
}

//缓存权限数据
func setCacheAdminConfigData() {

	o := orm.NewOrm()
	o.Using("geekxbweb")

	var items []*models.DbAuthority

	//记录list
	itemList, err := o.QueryTable("DbAuthority").All(&items)

	if err != nil || itemList <= 0 {
		Debug("缓存权限数据 Error!!!")
		return
	}

	Cache_All.Put("cacheAuthority", items, 0)
}
