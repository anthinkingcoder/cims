package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type OutStockDetail struct {
	Id          int64
	GmtCreate   time.Time                `orm:"auto_now_add;type(datetime)"`
	GmtModifier time.Time                `orm:"auto_now;type(datetime)"`
	ArtNo       *ArtNo                `orm:"rel(fk)"` //货号
	Size        *Size                `orm:"rel(fk)"`  //尺码
	Color       *Color                `orm:"rel(fk)"` //颜色
	Count       int                                   //数量
	Can 	    int `orm:"-"`
	Lock        int                        `orm:"default(0)"`
	OutStock     *OutStock  			 `orm:"rel(fk)"`
}

func (this *OutStockDetail)Delete() (int64, error) {
	return orm.NewOrm().Delete(this)
}

func (this *OutStockDetail) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *OutStockDetail) Read() (error) {
	return orm.NewOrm().QueryTable("size").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *OutStockDetail) Update() (int64, error) {
	return orm.NewOrm().Update(this,"GmtModifier","ArtNo","Size","Color","Count")
}
