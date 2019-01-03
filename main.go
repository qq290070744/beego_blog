package main

import (
	"beego_blog/models"
	_ "beego_blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDb_SQLTTE3()
}

func main() {

	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
