package controllers

import (
	"strconv"
	"github.com/qianbaidu/beego_blog/models"
)

type ArticleController struct {
	BaseIndexController
}

func (this *ArticleController) Get() {
	var (
		article models.Article
		id  int
		err error
	)
	if id, err = strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil  {
		id, _ = strconv.Atoi(this.Ctx.Input.Param(":id"))
	}

	article,err = models.GetArticleById(id)
	if err != nil{
		this.Redirect("/404", 302)
		return
	}

	category := models.GetCategoryList()
	tags := models.GetHotTagList()
	category_name := ""
	for _, v := range category {
		if v.Id == int64(article.CategoryId) {
			category_name = v.Name
			break
		}
	}
	this.Data["article"] = article
	this.Data["category"] = category
	this.Data["username"] = this.username
	this.Data["tags"] = tags
	this.Data["category_name"] = category_name

	this.TplName = "article.html"
}
