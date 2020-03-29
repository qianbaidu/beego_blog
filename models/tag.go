package models

import (
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id    int
	Name  string `orm:"size(20);index"`
	Count int
}

func GetHotTagList() (tags []orm.Params) {
	//var tags []orm.ParamsList
	//var tags []orm.Params
	//var tags []Tag

	o := orm.NewOrm()
	//var tagList []Tag
	o.Raw("select id as Id,name as Name,count as Count from tag where count > 0 order by count desc limit 100").Values(&tags)
	//for _, v := range tags {
	//	fmt.Println(v)
	//for kk,vv:=range v{
	//	fmt.Println(kk,vv)
	//}
	//tag:=Tag{v.Id,v.Name,v.Count}
	//tagList = append(tagList, v)
	//}
	return
}

//func GetTagList() () {
//	o := orm.NewOrm()
//	var tags []orm.Params
//	num, err := o.QueryTable("tag").Values(&tags, "name", "count")
//	if err == nil {
//		log.Error(fmt.Sprintf("%s", err))
//		//fmt.Printf("Result Nums: %d\n", num)
//		//for _, row := range tags {
//		//fmt.Println(row["Name"])
//		//fmt.Printf("Name: %s, Age: %s\m", row[0], row[1])
//		//}
//	}
//}

func GetTagById() (tag Tag, err error) {
	o := orm.NewOrm()
	//var tag Tag
	err = o.QueryTable("tag").Filter("Id", 1).One(&tag)
	//if err == nil {
	//	fmt.Println(tag)
	//}
	return
}

func (m *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
