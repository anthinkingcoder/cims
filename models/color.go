package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Color struct {
	Id          int64
	GmtCreate   time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModifier time.Time `orm:"auto_now;type(datetime)"`
	Color       string
	ArtNo	    []*ArtNo `orm:"reverse(many)"`
}

func (this *Color)Delete() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Color) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Color) Read() (error) {
	return orm.NewOrm().QueryTable("color").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *Color) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}
