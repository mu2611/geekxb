package controllers

import (
	"github.com/astaxie/beego"
)

type Admin_UserController struct {
	beego.Controller
}

func (this *Admin_UserController) Get() {

	this.Layout = "Admin_Layer.tpl"
	this.TplNames = "Admin_User.tpl"
}
