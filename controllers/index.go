package controllers

import (
	"github.com/astaxie/beego"
	"github.com/qianbaidu/beego_blog/models"
	"fmt"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	err := models.AddCate("test add ")
	if err != nil {
		fmt.Println(err)
	}
	cate := models.GetList()
	fmt.Println(cate)
	this.Data["category"] = cate
	this.TplName = "index.html"
}
