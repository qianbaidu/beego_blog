package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/common/log"
	"time"
)

const (
	_DB_NAME        = "data/beego.db"
	_SQLLTE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Name            string
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
	Content         string `orm:"size(5000)"`
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
	dbtype := beego.AppConfig.String("dbtype")
	log.Info(fmt.Sprintf("db_type:%s", dbtype))
	//if dbtype == "mysql" {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	dns := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dns)

	//} else if (dbtype == "sqlite") {
	//	if !com.IsExist(_DB_NAME) {
	//		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
	//		os.Create(_DB_NAME)
	//	}
	//	orm.RegisterDriver(_SQLLTE3_DRIVER, orm.DRSqlite)
	//	orm.RegisterDataBase("default", _SQLLTE3_DRIVER, _DB_NAME, 10)
	//} else {
	//	fmt.Errorf("undefine db type")
	//}

	orm.RegisterModel(new(Category), new(Topic), new(Article), new(User), new(Conf), new(Tag))
}

func AddCate(title string) error {
	o := orm.NewOrm()
	cate := &Category{
		Name:            title,
		Created:         time.Now(),
		Views:           1,
		TopicTime:       time.Now(),
		TopicCount:      1,
		TopicLastUserId: 1,
	}
	_, err := o.Insert(cate)
	return err
}

func GetCategoryList() (cate []*Category) {
	o := orm.NewOrm()
	cate = make([]*Category, 0)
	qs := o.QueryTable("category")
	var err error
	_, err = qs.All(&cate)
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
	}
	return
}
