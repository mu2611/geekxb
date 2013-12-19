package controllers

import (
	"geekxb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Admin_AuthorityController struct {
	beego.Controller
}

//list
func (this *Admin_AuthorityController) Get() {

	// 获取当前页码
	pagenumInt64, err := this.GetInt("page")
	//Check(err)
	pagenum := int64(pagenumInt64)
	if err != nil {
		pagenum = 1
	}

	o := orm.NewOrm()
	o.Using("geekxbweb")

	var items []*models.DbAuthority
	//记录总数
	itemCount, err := o.QueryTable("DbAuthority").Count()

	//页数超出后重新计算定位到最大页数
	var maxPagenum int64
	if itemCount/ItemsPerPage*ItemsPerPage < itemCount {
		maxPagenum = itemCount/ItemsPerPage + 1
	} else {
		maxPagenum = itemCount / ItemsPerPage
	}
	if pagenum > maxPagenum {
		pagenum = maxPagenum
	}

	//记录list
	itemList, err := o.QueryTable("DbAuthority").Limit(ItemsPerPage).Offset((pagenum - 1) * ItemsPerPage).All(&items)

	if err != nil || itemList <= 0 {
		Debug(" func (this *Admin_AuthorityController) Get() Error!!!")
		this.Redirect("/", 404)
	}

	this.Data["Authitems"] = items
	this.Data["PageDiv"] = ResponsePageDiv(itemCount, pagenum)
	PreUrl = this.Ctx.Request.URL.String()

	this.Layout = "Admin_Layer.tpl"
	this.TplNames = "Admin_Authority.tpl"
}

//item edit
func (this *Admin_AuthorityController) ActionEdit() {
	tempEditID, _ := this.GetInt("editID")
	this.Data["editID"] = tempEditID

	if tempEditID != -1 {
		o := orm.NewOrm()
		o.Using("geekxbweb")
		author := models.DbAuthority{Id: int(tempEditID)}
		if o.Read(&author) == nil {
			this.Data["authGroup"] = author.Authgroup
			this.Data["authName"] = author.Authname
			this.Data["authAction"] = author.Authaction
		}
	}
	this.Layout = "Admin_Layer.tpl"
	this.TplNames = "Admin_Authority_Edit.tpl"
}

//item update
func (this *Admin_AuthorityController) ActionUpdate() {
	thisAuthID, _ := this.GetInt("authid")
	thisAuthGroup, _ := this.GetInt("authgroup")
	thisAuthName := this.GetString("authname")
	thisAuthAction := this.GetString("authaction")

	o := orm.NewOrm()
	o.Using("geekxbweb")
	if int(thisAuthID) == -1 { //添加
		insertAuthor := models.DbAuthority{Id: int(thisAuthID), Authgroup: int(thisAuthGroup), Authname: thisAuthName, Authaction: thisAuthAction}
		o.Insert(&insertAuthor)
	} else { //修改
		author := models.DbAuthority{Id: int(thisAuthID)}
		if o.Read(&author) == nil {
			author.Authgroup = int(thisAuthGroup)
			author.Authname = thisAuthName
			author.Authaction = thisAuthAction
			o.Update(&author)
		}
	}

	this.Redirect(PreUrl, 302)
}

//item del
func (this *Admin_AuthorityController) ActionDel() {

	//
	delID, err := this.GetInt("delID")
	if err != nil {
		return
	}

	o := orm.NewOrm()
	o.Using("geekxbweb")
	o.Delete(&models.DbAuthority{Id: int(delID)})
	this.Redirect(PreUrl, 302)
}
