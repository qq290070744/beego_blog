package routers

import (
	"beego_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controller.MainController{})
	beego.Router("/category", &controller.CateController{})

	beego.Router("/topic", &controller.TopicController{})
	beego.AutoRouter(&controller.TopicController{})
	beego.Router("/reply", &controller.ReplyController{})
	beego.Router("/reply/add", &controller.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controller.ReplyController{}, "*:Delete")
	beego.Router("/login", &controller.LoginController{})
}
