package admin

import (
	"cims/models"
	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	Base
}

//@router /admin/index [get]
func (this *AdminController) Home() {
	this.TplName = "admin/index.tpl"
}

//@router /admin/ [get]
func (this *AdminController) Home1() {
	this.TplName = "admin/index.tpl"
}



//@router /admin/user [get]
func (this *AdminController) UserManage() {
	var users []*models.User
	orm.NewOrm().QueryTable("user").Filter("flag", 0).Filter("lock", 0).All(&users)
	this.Data["users"] = users
	this.TplName = "admin/user.tpl"
}

//@router /admin/art_no [get]
func (this *AdminController) ArtNo() {
	var artnos []*models.ArtNo
	//获取货号列表
	o := orm.NewOrm();
	o.QueryTable("art_no").Filter("lock", 0).OrderBy("-id").RelatedSel().All(&artnos)
	//载入销售属性颜色和尺寸
	for _, artno := range artnos {
		o.LoadRelated(artno, "Colors")
		o.LoadRelated(artno, "Sizes")
	}
	//获取所有里料
	var linings []*models.Lining
	o.QueryTable("lining").All(&linings)
	//获取所有面料
	var fabrics []*models.Fabric
	o.QueryTable("fabric").All(&fabrics)
	//获取所有尺码
	var sizes []*models.Size
	o.QueryTable("size").All(&sizes)
	//获取所有颜色
	var colors []*models.Color
	o.QueryTable("color").All(&colors)

	this.Data["artnos"] = artnos
	this.Data["linings"] = linings
	this.Data["fabrics"] = fabrics
	this.Data["sizes"] = sizes
	this.Data["colors"] = colors
	this.TplName = "admin/artno.tpl"

}

//@router /admin/ware_house [get]
func (this *AdminController) WareHouse() {
	var warehouses []*models.WareHouse
	o := orm.NewOrm()
	o.QueryTable("ware_house").Filter("lock", 0).OrderBy("-id").All(&warehouses)
	this.Data["warehouses"] = warehouses
	this.TplName = "admin/warehouse.tpl"
}

//@router /admin/in_stock [get]
func (this *AdminController) InStock() {
	var inStocks []*models.InStock
	o := orm.NewOrm()
	o.QueryTable("in_stock").Filter("lock", 0).OrderBy("-gmt_create").RelatedSel().All(&inStocks)
	//获取货号以及他的颜色和尺寸
	var artNos []*models.ArtNo
	o.QueryTable("art_no").Filter("lock", 0).All(&artNos)
	for _, artNo := range artNos {
		o.LoadRelated(artNo, "Sizes")
		o.LoadRelated(artNo, "Colors")
		o.QueryTable("sku").Filter("art_no_id", artNo.Id).Filter("lock", 0).RelatedSel().All(&artNo.Skus)
	}
	//获取所有仓库
	var wareHouses []*models.WareHouse
	o.QueryTable("ware_house").Filter("lock", 0).OrderBy("-id").All(&wareHouses)

	this.Data["instocks"] = inStocks
	this.Data["warehouses"] = wareHouses
	this.Data["artnos"] = artNos

	this.TplName = "admin/instock.tpl"
}

//@router /admin/out_stock [get]
func (this *AdminController) OutStock() {
	var outStocks []*models.OutStock
	o := orm.NewOrm()
	o.QueryTable("out_stock").Filter("lock", 0).OrderBy("id").RelatedSel().All(&outStocks)
	//获取货号以及他的颜色和尺寸
	var artNos []*models.ArtNo
	o.QueryTable("art_no").Filter("lock", 0).All(&artNos)
	for _, artNo := range artNos {
		o.LoadRelated(artNo, "Sizes")
		o.LoadRelated(artNo, "Colors")
		o.QueryTable("sku").Filter("art_no_id", artNo.Id).Filter("lock", 0).RelatedSel().All(&artNo.Skus)
	}
	//获取所有仓库
	var wareHouses []*models.WareHouse
	o.QueryTable("ware_house").Filter("lock", 0).OrderBy("-id").All(&wareHouses)

	this.Data["outstocks"] = outStocks
	this.Data["warehouses"] = wareHouses
	this.Data["artnos"] = artNos

	this.TplName = "admin/outstock.tpl"
}
//@router /admin/new [get]
func (this *AdminController) Article()  {
	this.TplName = "admin/article.tpl"

}



