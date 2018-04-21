package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Sku struct {
	Id           int64
	GmtCreate    time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModifier  time.Time `orm:"auto_now;type(datetime)"`
	RetailPrice  float64 `orm:"null"`
	FactoryPrice float64 `orm:"null"`
	Stock        int   `orm:"default(0)"`
	Size         *Size `orm:"rel(fk)"`
	Color        *Color `orm:"rel(fk)"`
	ArtNo        *ArtNo        `orm:"rel(fk)"`
	Lock         int    `orm:"default(0)"`
}

//// 多字段唯一键
//func (this *ArtNo) TableUnique() [][]string {
//	return [][]string{
//		[]string{"ArtNo", "Color","Size"},
//	}
//}

func (this *Sku) Delete() (int64, error) {
	return orm.NewOrm().QueryTable("sku").Filter("id", this.Id).Update(orm.Params{
		"lock": 1,
	})
}

func (this *Sku) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *Sku) Read() (error) {
	return orm.NewOrm().QueryTable("sku").Filter("id", this.Id).Filter("lock",0).RelatedSel().One(this)
}

func (this *Sku) Update() (int64, error) {
	return orm.NewOrm().Update(this,"GmtModifier","RetailPrice","FactoryPrice","Size","Color")
}


