package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Fabric struct {
	Id          int64
	GmtCreate   time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModifier time.Time `orm:"auto_now;type(datetime)"`
	Fabric        string   `orm:"size(20);unique"`
}


func (this *Fabric) Delete() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *Fabric) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Fabric) Read() (error) {
	return orm.NewOrm().QueryTable("fabric").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *Fabric) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

