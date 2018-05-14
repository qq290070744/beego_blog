package controller

import (
	"github.com/astaxie/beego"
	"beego_blog/models"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid")
	nickname := this.Input().Get("nickname")
	content := this.Input().Get("content")
	err := models.AddReply(tid, nickname, content)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	tid := this.Input().Get("tid")

	err := models.DeleteReply(tid,this.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
}
