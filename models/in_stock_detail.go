package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type InStockDetail struct {
	Id          int64
	GmtCreate   time.Time                `orm:"auto_now_add;type(datetime)"`
	GmtModifier time.Time                `orm:"auto_now;type(datetime)"`
	ArtNo       *ArtNo                `orm:"rel(fk)"` //货号
	Size        *Size                `orm:"rel(fk)"`  //尺码
	Color       *Color                `orm:"rel(fk)"` //颜色
	Count       int                                   //数量
	Can 	    int `orm:"-"`
	Lock        int                        `orm:"default(0)"`
	InStock     *InStock                         `orm:"rel(fk)"`
}

func (this *InStockDetail)Delete() (int64, error) {
	return orm.NewOrm().QueryTable("in_stock_detail").Filter("id", this.Id).Update(orm.Params{
		"lock": 1,
	})
}

func (this *InStockDetail) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *InStockDetail) Read() (error) {
	return orm.NewOrm().QueryTable("size").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *InStockDetail) Update() (int64, error) {
	return orm.NewOrm().Update(this,"GmtModifier","ArtNo","Size","Color","Count")
}
