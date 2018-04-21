package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type WareHouse struct {
	Id           int64
	GmtCreate    time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModifier  time.Time `orm:"auto_now;type(datetime)"`
	WareHouseNum string  `orm:"size(20);unique"`
	Name         string
	Tel          string `orm:"size(20);null"`
	Contact      string `orm:"size(20);null"`
	Capacity        int `orm:"default(0)"`
	Stock           int `orm:"default(0)"`
	Lock         int  `orm:"default(0)"`
}

func (this *WareHouse)Delete() (int64, error) {
	return orm.NewOrm().QueryTable("ware_house").Filter("id", this.Id).Update(orm.Params{
		"lock": 1,
	})
}

func (this *WareHouse) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *WareHouse) Read() (error) {
	return orm.NewOrm().QueryTable("ware_house").Filter("id", this.Id).Filter("lock",0).RelatedSel().One(this)
}

func (this *WareHouse) Update() (int64, error) {
	return orm.NewOrm().Update(this,"GmtModifier","WareHouseNum","WareHouseNum","Tel","Capacity","Contact")
}