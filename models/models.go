package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const (
	_DB_NAME        = "data/beego.db"


	//_SQLLTE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string    `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64
	Auther          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDb() {
	//if !com.IsExist(_DB_NAME) {
	//	os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
	//	os.Create(_DB_NAME)
	//}
	//
	orm.RegisterModel(new(Category), new(Topic))
	//orm.RegisterDriver(_SQLLTE3_DRIVER, orm.DRSqlite)
	//orm.RegisterDataBase("default", _SQLLTE3_DRIVER, _DB_NAME, 10)


	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@/beego_blog?charset=utf8")

}

func AddCate(title string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:           title,
		Created:         time.Now(),
		Views:           1,
		TopicTime:       time.Now(),
		TopicCount:      1,
		TopicLastUserId: 1,
	}
	_, err := o.Insert(cate)
	return err
}

func GetList() (cate []*Category) {
	o := orm.NewOrm()
	cate = make([]*Category, 0)
	qs := o.QueryTable("category")
	var err error
	_, err = qs.All(&cate)
	if err != nil {
		fmt.Println(err)
	}
	return
}
