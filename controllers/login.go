package controller

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"
	//fmt.Println(isExit,"exit")
	if isExit {
		this.Ctx.SetSecureCookie("uname", "uname", "", -1, "/")
		this.Ctx.SetSecureCookie("pwd", "pwd", "", -1, "/")
		this.Redirect("/login", 301)
	}
	this.TplName = "login.html"
	return
}

func (this *LoginController) Post() {
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<21 - 1
		}
		fmt.Println(maxAge)
		this.Ctx.SetSecureCookie("uname", "uname", uname, "/")
		this.Ctx.SetSecureCookie("pwd", "pwd", uname, "/")
	}

	this.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	uname, _ := ctx.GetSecureCookie("uname", "uname")

	pwd, _ := ctx.GetSecureCookie("pwd", "pwd")

	fmt.Println(uname, "----", pwd)
	if len(uname) == 0 || len(pwd) == 0 {
		return false
	}
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		return true
	}
	return false
}
