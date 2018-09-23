package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	UserName   string    `orm:"unique;size(15)"`
	Password   string    `orm:"size(32)"`
	Email      string    `orm:"size(50)"`
	LastLogin  time.Time `orm:"auto_now_add;type(datetime)"`
	LoginCount int
	LastIp     string    `orm:"size(32)"`
	Authkey    string    `orm:"size(10)"`
	Active     int8
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}