package controllers

import (
	"github.com/astaxie/beego"
	"github.com/qianbaidu/beego_blog/models"
	"strconv"
	"math"
)

type TagController struct {
	BaseIndexController
}


func (this *TagController) Get() {
	var (
		ArticleList []*models.Article
		pageSize    int
		page        int
		tid         int
		err         error
	)

	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}
	if tid, err = strconv.Atoi(this.Ctx.Input.Param(":id")); err != nil {
		tid ,_ = strconv.Atoi(this.Ctx.Input.Param(":id"))
	}
	if pageSize, err = beego.AppConfig.Int("page_size"); err != nil {
		pageSize, _ = beego.AppConfig.Int("page_size")
	}

	query := new(models.Article).Query().Filter("status", 0).Filter("tag_id", tid)
	count, _ := query.Count()
	var totalPageInt64 float64 = 0
	if count > 0 {
		query.OrderBy("-id").Limit(pageSize, (page-1)*pageSize).All(&ArticleList)
		totalPage := float64(count / int64(pageSize))
		totalPageInt64 = math.Ceil(totalPage)
	}


	//articleList := models.GetArticleListcate()
	category := models.GetCategoryList()
	cateList := make(map[int64]string)
	for _,v:=range category  {
		cateList[v.Id] = v.Name
	}

	tags := models.GetHotTagList()
	tag,_ := models.GetTagById()
	this.Data["totalPage"] = totalPageInt64
	this.Data["article"] = ArticleList
	this.Data["current_page"] = page
	this.Data["next_page"] = page + 1
	this.Data["tid"] = tid
	this.Data["tag"] = tag

	category_name := ""

	this.Data["category"] = category
	this.Data["category_name"] = category_name
	this.Data["username"] = this.username
	this.Data["tags"] = tags
	this.Data["cateList"] = cateList

	this.TplName = "tag.html"

}
