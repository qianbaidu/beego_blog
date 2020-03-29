package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
	"time"
)

type Article struct {
	Id          int64
	Author      int
	CreatedTime time.Time `orm:"index"`
	UpdatedTime time.Time `orm:"index"`
	Content     string    `orm:"type(text)"`
	Title       string
	Status      int
	ArticleType int
	CategoryId  int
	Times       int64
	ThubImg     string
	TagId       int64
	//CategoryInfo Category

}

func GetArticleById(id int) (article Article, err error) {
	o := orm.NewOrm()
	article = Article{Id: int64(id)}
	err = o.Read(&article)

	//if err == orm.ErrNoRows {
	//
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println(article.Id, article.Title)
	//}
	return
}

func GetArticleListcate() (article []*Article) {
	o := orm.NewOrm()
	article = make([]*Article, 0)
	qs := o.QueryTable("article")
	var err error
	_, err = qs.All(&article)
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
	}
	return
}

func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
