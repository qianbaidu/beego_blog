package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//err := models.AddCate("test add ")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//cate := models.GetList()
	//fmt.Println(cate)
	//this.Data["category"] = cate
	this.TplName = "login.html"
}
