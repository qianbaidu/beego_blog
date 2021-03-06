package controllers

import (
	"github.com/qianbaidu/beego_blog/models"
	"strconv"
	"github.com/astaxie/beego"
	"math"
)

type IndexController struct {
	BaseIndexController
}

func (this *IndexController) Get() {
	var (
		ArticleList []*models.Article
		pageSize    int
		page        int
		err         error
	)

	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}
	if pageSize, err = beego.AppConfig.Int("page_size"); err != nil {
		pageSize, _ = beego.AppConfig.Int("page_size")
	}
	query := new(models.Article).Query().Filter("status", 0)
	count, _ := query.Count()
	var totalPageInt64 float64 = 0
	if count > 0 {
		query.OrderBy("-id").Limit(pageSize, (page-1)*pageSize).All(&ArticleList)
		totalPage := float64(count / int64(pageSize))
		totalPageInt64 = math.Ceil(totalPage)
	}

	category := models.GetCategoryList()
	cateList := make(map[int64]string)
	for _,v:=range category  {
		cateList[v.Id] = v.Name
	}
	tags := models.GetHotTagList()
	this.Data["totalPage"] = totalPageInt64
	this.Data["article"] = ArticleList
	this.Data["current_page"] = page
	this.Data["next_page"] = page + 1
	this.Data["category"] = category
	this.Data["username"] = this.username
	this.Data["tags"] = tags
	this.Data["cateList"] = cateList
	this.TplName = "index.html"

}
