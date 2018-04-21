package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type InStock struct {
	Id          int64
	GmtCreate   time.Time        `orm:"auto_now_add;type(datetime)"`
	GmtModifier time.Time        `orm:"auto_now;type(datetime)"`
	StorageTime time.Time `orm:"type(datetime)"`
	InStockNum  string        `orm:"size(20);unique"` //入库编号
	WareHouse   *WareHouse  `orm:"rel(fk)"`           //所入仓库
	Remark      string        `orm:"type(text);null"` //备注
	Status      int       `orm:"default(0)"`          //是否提交
	Source      string        `orm:"size(255)"`       //来源
	Lock        int        `orm:"default(0)"`
	User        *User  `orm:"rel(fk)"`
	InStockDetails []*InStockDetail `orm:"reverse(many)"`
}

func (this *InStock)Delete() (int64, error) {
	return orm.NewOrm().QueryTable("in_stock").Filter("id", this.Id).Update(orm.Params{
		"lock": 1,
	})
}

func (this *InStock) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *InStock) Read() (error) {
	return orm.NewOrm().QueryTable("in_stock").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *InStock) Update() (int64, error) {
	return orm.NewOrm().Update(this,"GmtModifier","StorageTime","InStockNum","WareHouse","Remark","Source")
}
