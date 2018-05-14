package controller

import (
	"github.com/astaxie/beego"
	"beego_blog/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["uname"], _ = this.GetSecureCookie("uname", "uname")
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic.html"
	topics, err := models.GetAllTopics(this.Input().Get("cate"), false)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}
func (this *TopicController) Add() {
	this.Data["uname"], _ = this.GetSecureCookie("uname", "uname")
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_add.html"
	//this.Ctx.WriteString("add")
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicController) Post() {
	this.Data["uname"], _ = this.GetSecureCookie("uname", "uname")
	this.Data["IsTopic"] = true
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")
	category := this.Input().Get("category")

	if len(tid) == 0 {
		err := models.AddTopic(title, content, category)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.ModifyTopic(tid, title, content, category)
		if err != nil {
			beego.Error(err)
		}
	}

	//fmt.Println("---------------------------------------------------------------------------")

	this.Redirect("/topic", 302)
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	tid := this.Ctx.Input.Param("0")
	this.Data["Tid"] = tid
	this.Data["uname"], _ = this.GetSecureCookie("uname", "uname")
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	replices, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		//this.Redirect("/", 302)
		return
	}
	this.Data["Replices"] = replices

}

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "topic_modify.html"
	this.Data["uname"], _ = this.GetSecureCookie("uname", "uname")
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	//this.Data["Tid"] = tid

	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	//this.TplName = "topic_delete.html"
	this.Data["uname"], _ = this.GetSecureCookie("uname", "uname")
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	err := models.DeleteTopic(this.Input().Get("category"), this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 302)

}
