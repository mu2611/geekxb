package models

import (
	"github.com/astaxie/beego/orm"
)

type DbUsergroup struct {
	Id                int `orm:"auto"`
	Ugname            string
	Ugmanageauthority string
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(DbUsergroup))
}
