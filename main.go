package main

import (
	"github.com/astaxie/beego"
	"github.com/qianbaidu/beego_blog/models"
	"github.com/astaxie/beego/orm"
	"github.com/qianbaidu/beego_blog/controllers"
)

func init() {
	models.RegisterDb()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, false)
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Run()
}
