package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/qianbaidu/beego_blog/routers"
	"github.com/qianbaidu/beego_blog/models"
)

func init() {
	models.RegisterDb()
}

func main() {
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	orm.RunSyncdb("default", false, false)
	beego.Run()
}
