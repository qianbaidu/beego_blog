package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.html"

	//c.Ctx.WriteString("app name :" + beego.AppConfig.String("appname") + "app port " +  beego.AppConfig.String("httpport"))

	c.Data["TrueCond"] = true
	type u struct {
		Name string
		Age  int
		Sex  string
	}

	user := &u{"zhangSan", 28, "男"}
	c.Data["user"] = user
	nums := []int{1,2,3,3,4,5,67,7}
	c.Data["nums"] = nums
	c.Data["TplVar"] = "hello"
	c.Data["html"] = "<h3>hello</h3>"



	c.TplName = "default.tpl"
}
