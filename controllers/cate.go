package controller

import (
	"github.com/astaxie/beego"
	"beego_blog/models"
)

type CateController struct {
	beego.Controller
}

func (this *CateController) Get() {
	this.Data["uname"], _ = this.GetSecureCookie("uname", "uname")
	op := this.Input().Get("op")

	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategories(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	}

	this.Data["IsCategory"] = true
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "category.html"

}
