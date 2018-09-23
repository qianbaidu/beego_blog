package models

import "github.com/astaxie/beego/orm"

type Conf struct {
	Id    int
	Name  string
	Value string
}

func (m *Conf) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
