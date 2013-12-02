package controllers

import (
	"geekxb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Admin_UserGroupController struct {
	beego.Controller
}

//list
func (this *Admin_UserGroupController) Get() {

	if err := Cache_All.IsExist("cacheAuthority"); err != nil {

		//l := Cache_All.Get("cacheAuthority").([]*models.DbAuthority)
		//for _, temp := range l {
		//	Debug("%s", temp.Authname)
		//}

	} else {
		Debug("don't get cacheAuthority")
	}

	// 获取当前页码
	pagenumInt64, err := this.GetInt("page")
	//Check(err)
	pagenum := int64(pagenumInt64)
	if err != nil {
		pagenum = 1
	}

	o := orm.NewOrm()
	o.Using("geekxbweb")

	var items []*models.DbUsergroup
	//记录总数
	itemCount, _ := o.QueryTable("DbUsergroup").Count()

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
	itemList, err := o.QueryTable("DbUsergroup").Limit(ItemsPerPage).Offset((pagenum - 1) * ItemsPerPage).All(&items)

	if err != nil || itemList <= 0 {
		this.Redirect("/", 302)
	}

	this.Data["UserGroupitems"] = items
	this.Data["PageDiv"] = ResponsePageDiv(itemCount, pagenum)
	PreUrl = this.Ctx.Request.URL.String()

	this.Layout = "Admin_Layer.tpl"
	this.TplNames = "Admin_UserGroup.tpl"
}

//item edit
func (this *Admin_UserGroupController) ActionEdit() {
	tempEditID, _ := this.GetInt("editID")
	this.Data["editID"] = tempEditID

	if tempEditID != -1 {
		o := orm.NewOrm()
		o.Using("geekxbweb")
		item := models.DbUsergroup{Id: int(tempEditID)}
		if o.Read(&item) == nil {
			this.Data["Ugname"] = item.Ugname
			this.Data["Ugmanageauthority"] = item.Ugmanageauthority
			Debug(item.Ugmanageauthority)
		}
	}
	this.Layout = "Admin_Layer.tpl"
	this.TplNames = "Admin_UserGroup_Edit.tpl"
}

//item update
func (this *Admin_UserGroupController) ActionUpdate() {
	thisUgID, _ := this.GetInt("editid")
	thisUgname := this.GetString("Ugname")
	thisUgmanageauthority := this.GetString("Ugmanageauthority")

	o := orm.NewOrm()
	o.Using("geekxbweb")
	if int(thisUgID) == -1 { //添加
		insertUserGroup := models.DbUsergroup{Id: int(thisUgID), Ugname: thisUgname, Ugmanageauthority: thisUgmanageauthority}
		o.Insert(&insertUserGroup)
	} else { //修改
		item := models.DbUsergroup{Id: int(thisUgID)}
		if o.Read(&item) == nil {
			item.Ugname = thisUgname
			item.Ugmanageauthority = thisUgmanageauthority
			o.Update(&item)
		}
	}

	this.Redirect(PreUrl, 302)
}

//item del
func (this *Admin_UserGroupController) ActionDel() {

	//
	delID, err := this.GetInt("delID")
	if err != nil {
		return
	}

	o := orm.NewOrm()
	o.Using("geekxbweb")
	o.Delete(&models.DbUsergroup{Id: int(delID)})
	this.Redirect(PreUrl, 302)
}
