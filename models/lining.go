package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Lining struct {
	Id          int64
	GmtCreate   time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModifier time.Time `orm:"auto_now;type(datetime)"`
	Lining      string    `orm:"size(20);unique"`
}

func (this *Lining) Delete() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Lining) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Lining) Read() (error) {
	return orm.NewOrm().QueryTable("lining").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *Lining) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

