package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Size struct {
	Id          int64
	GmtCreate   time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModifier time.Time `orm:"auto_now;type(datetime)"`
	Size        int
	ArtNo	    []*ArtNo `orm:"reverse(many)"`
}

func (this *Size)Delete() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Size) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Size) Read() (error) {
	return orm.NewOrm().QueryTable("size").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *Size) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}
