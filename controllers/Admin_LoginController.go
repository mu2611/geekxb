package controllers

import (
	"github.com/astaxie/beego"
)

type Admin_LoginController struct {
	beego.Controller
}

func (this *Admin_LoginController) Get() {

	this.TplNames = "Admin_Login.tpl"
}
