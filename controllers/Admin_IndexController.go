package controllers

import (
	"database/sql"
	"fmt"
	"geekxb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Admin_IndexController struct {
	beego.Controller
}

func (this *Admin_IndexController) Get() {

	o := orm.NewOrm()
	o.Using("geekxbweb")
	author := models.DbAuthority{Id: 1}

	err := o.Read(&author)

	if err == sql.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(author.Authname)
	}

	this.Data["Username"] = author.Authname

	//insertAuthor := DbAuthority{Id: 1, Authname: "test", Authaction: "jjjjjjjj/kkkkkk"}
	//o.Insert(&insertAuthor)

	//o.Delete(&DbAuthority{Id: 2})

	//testaubb := models.DbAuthority{Id: 1}

	//if o.Read(&testaubb) == nil {
	//	testaubb.Authaction = "1223"
	//	o.Update(&testaubb)
	//}

	this.Layout = "Admin_Layer.tpl"
	this.TplNames = "Admin_Index.tpl"

}
