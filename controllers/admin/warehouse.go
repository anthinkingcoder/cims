package admin

import (
	"cims/models"
	"strconv"
	"time"
)


//@router /admin/ware_house/create [post]
func (this *AjaxController) CreateWareHouse() {
	wareHouseNum := this.GetString("wareHouseNum")
	name := this.GetString("name")
	tel := this.GetString("tel")
	contact := this.GetString("contact")
	capacity, _ := this.GetInt("capacity")
	wareHouse := models.WareHouse{WareHouseNum:wareHouseNum, Name:name, Tel:tel, Contact:contact, Capacity:capacity}
	if _, err := wareHouse.Insert(); err != nil {
		this.ReturnError("插入失败,仓库编号已经存在")
		return
	}
	this.ReturnSuccess()
}

//@router /admin/ware_house/:id/delete [delete]
func (this *AjaxController) DeleteWareHouse() {
	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	wareHouse := models.WareHouse{Id:id}
	if _, err := wareHouse.Delete(); err != nil {
		this.ReturnError("warehouse not found")
		return
	}
	this.ReturnSuccess()
}

//@router /admin/ware_house/:id/update [post]
func (this *AjaxController) UpdateWareHouse() {
	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	wareHouseNum := this.GetString("wareHouseNum")
	name := this.GetString("name")
	tel := this.GetString("tel")
	contact := this.GetString("contact")
	capacity, _ := this.GetInt("capacity")
	wareHouse := models.WareHouse{Id:id}
	if err := wareHouse.Read(); err != nil {
		this.ReturnError("该仓库不存在")
		return;
	}
	if (wareHouse.Stock > capacity) {
		this.ReturnError("仓储量不能小于当前库存量");
		return;
	}
	wareHouse.Capacity = capacity
	wareHouse.Tel = tel
	wareHouse.Contact = contact
	wareHouse.Name = name
	wareHouse.WareHouseNum = wareHouseNum
	wareHouse.GmtModifier = time.Now()
	if _, err := wareHouse.Update(); err != nil {
		this.ReturnError("更新失败,该仓库编号已存在")
		return
	}
	this.ReturnSuccess()
}
