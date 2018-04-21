package models
import (
	"github.com/astaxie/beego/orm"
	"time"
)
type ArtNo struct {
	Id             int64
	GmtCreate      time.Time `orm:"auto_now_add;type(datetime)"`
	GmtModifier    time.Time `orm:"auto_now;type(datetime)"`
	ArtNo          string `orm:"size(20);unique"`
	ProductName    string
	RetailPrice    float64 `orm:"null"`
	FactoryPrice   float64 `orm:"null"`
	Lining         *Lining  `orm:"rel(fk)"`
	Fabric         *Fabric  `orm:"rel(fk)"`
	Lock           int     `orm:"default(0)"`
	Colors          []*Color  `orm:"rel(m2m);rel_table(art_no_color)"`
	Sizes           []*Size `orm:"rel(m2m);rel_table(art_no_size)"`
	InStockDetails []*InStockDetail `orm:"reverse(many)"`
	Skus           []*Sku                 `orm:"reverse(many)"`
}





func (this *ArtNo)Delete() (int64, error) {
	return orm.NewOrm().QueryTable("art_no").Filter("id", this.Id).Update(orm.Params{
		"lock": 1,
	})
}
func (this *ArtNo) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *ArtNo) Read() (error) {
	return orm.NewOrm().QueryTable("art_no").Filter("id", this.Id).RelatedSel().One(this)
}


func (this *ArtNo) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}
