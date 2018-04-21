package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id          int64
	GmtCreate   time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModifier time.Time `orm:"auto_now;type(datetime)"`
	Name        string `orm:"null"`
	Username    string `orm:"size(10);unique"`
	Password    string
	Profile     string `orm:"type(text)"`
	Flag        int `orm:"default(0)"`
	Lock        int `orm:"default(0)"`
	InStock     [] *InStock `orm:"reverse(many)"`
}

func (this *User) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *User) Update() (int64, error) {
	return orm.NewOrm().Update(this,"GmtModifier","Name","Username","Password","Profile")
}

func (this *User) Delete() (int64, error) {
	return orm.NewOrm().QueryTable("user").Filter("id", this.Id).Update(orm.Params{
		"lock": 1,
	})
}

func (this *User) GetByUsername() (error) {
	return orm.NewOrm().QueryTable("user").Filter("username", this.Username).RelatedSel().One(this)
}

func (this *User) Read() (error) {
	return orm.NewOrm().QueryTable("user").Filter("id", this.Id).RelatedSel().One(this)
}

