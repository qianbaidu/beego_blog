package controllers

import (
	"github.com/astaxie/beego"
	"github.com/qianbaidu/beego_blog/models"
	"strconv"
	"math"
)

type CategoryController struct {
	BaseIndexController
}

func (this *CategoryController) Get() {
	var (
		ArticleList []*models.Article
		pageSize    int
		page        int
		cid         int
		err         error
	)

	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}
	if cid, err = strconv.Atoi(this.Ctx.Input.Param(":cid")); err != nil {
		cid ,_ = strconv.Atoi(this.Ctx.Input.Param(":cid"))
	}
	if pageSize, err = beego.AppConfig.Int("page_size"); err != nil {
		pageSize, _ = beego.AppConfig.Int("page_size")
	}

	query := new(models.Article).Query().Filter("status", 0).Filter("category_id", cid)
	count, _ := query.Count()
	var totalPageInt64 float64 = 0
	if count > 0 {
		query.OrderBy("-id").Limit(pageSize, (page-1)*pageSize).All(&ArticleList)
		totalPage := float64(count / int64(pageSize))
		totalPageInt64 = math.Ceil(totalPage)
	}

	//articleList := models.GetArticleListcate()
	category := models.GetCategoryList()
	tags := models.GetHotTagList()
	this.Data["totalPage"] = totalPageInt64
	this.Data["article"] = ArticleList
	this.Data["current_page"] = page
	this.Data["next_page"] = page + 1
	this.Data["cid"] = cid

	category_name := ""
	for _, v := range category {
		if v.Id == int64(cid) {
			category_name = v.Name
		}
	}
	this.Data["category"] = category
	this.Data["category_name"] = category_name
	this.Data["username"] = this.username
	this.Data["tags"] = tags

	this.TplName = "category.html"

}
