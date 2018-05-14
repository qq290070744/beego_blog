package controller

import (
	"github.com/astaxie/beego"
	"beego_blog/models"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.html"
//	c.Data["TrueCond"] = true
//	c.Data["FalseCond"] = false
//
//	type u struct {
//		Name string
//		Age  int
//		Sex  string
//	}
//	user := &u{
//		Name: "jwh",
//		Age:  20,
//		Sex:  "Male",
//	}
//	c.Data["user"] = user
//
//	nums := []int{1, 2, 3, 4, 5, 6, 7}
//	c.Data["nums"] = nums
//
//	c.Data["html"]="<h1>html<h1>"
//}

func (this *MainController) Get() {
	//this.Ctx.WriteString(" appname: " + beego.AppConfig.String("appname") +
	//	"\n httpport: " + beego.AppConfig.String("httpport") +
	//	"\n runmode: " + beego.AppConfig.String("runmode"))
	//this.Ctx.WriteString(" appname: " + beego.AppName +
	//	"\n httpport: " + string(beego.HttpPort) +
	//	"\n runmode: " + beego.RunMode)
	//beego.Trace("Trace test1")

	this.TplName = "home.html"
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["uname"], _ = this.GetSecureCookie("uname", "uname")
	topics, err := models.GetAllTopics(this.Input().Get("cate"),true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}
