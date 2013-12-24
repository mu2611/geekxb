package controllers

import (
	"geekxb/models"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var (
	Cache_All cache.Cache
)

//初始化缓存
func InitCache() {

	//初始化权限配置缓存
	var err error

	Cache_All, err = cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		//抛出异常
		Debug(err)
		return
	}

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
		Debug("获取权限对照数据 Error!!!")
		return
	}

	if err = Cache_All.Put("cacheAuthority", items, 60); err != nil {
		Debug("cache err!!")
		return
	}

	//Debug("init : %s", Cache_All.Get("cacheAuthority"))
}
