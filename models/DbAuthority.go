package models

import (
	"github.com/astaxie/beego/orm"
)

type DbAuthority struct {
	Id         int `orm:"auto"`
	Authname   string
	Authaction string
	Authgroup  int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(DbAuthority))
}
