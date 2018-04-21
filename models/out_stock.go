package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type OutStock struct {
	Id           int64
	GmtCreate    time.Time        `orm:"auto_now_add;type(datetime)"`
	GmtModifier  time.Time        `orm:"auto_now;type(datetime)"`
	OutStockNum  string        `orm:"size(20);unique"` //入库编号
	DeliveryTime time.Time `orm:"type(datetime)"`
	WareHouse    *WareHouse  `orm:"rel(fk)"`           //所出仓库
	ReceiverName string     `orm:"null"`               //接收人
	ReceiverTel  string     `orm:"null"`               //接收人电话
	Address      string     `orm:"null"`               //地址
	Remark       string      `orm:"type(text);null"`   //备注
	Status       int       `orm:"default(0)"`          //是否提交
	Lock         int        `orm:"default(0)"`
}

func (this *OutStock)Delete() (int64, error) {
	return orm.NewOrm().QueryTable("out_stock").Filter("id", this.Id).Update(orm.Params{
		"lock": 1,
	})
}

func (this *OutStock) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *OutStock) Read() (error) {
	return orm.NewOrm().QueryTable("out_stock").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *OutStock) Update() (int64, error) {
	return orm.NewOrm().Update(this,"GmtModifier","OutStockNum","DeliveryTime","WareHouse","ReceiverName","ReceiverTel","Address","Remark")
}
