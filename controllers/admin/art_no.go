package admin

import (
	"cims/models"
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)


//@router /admin/art_no/:id/sku/list [get]
func (this *AjaxController)ListSku() {
	artNoId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var skus []*models.Sku
	orm.NewOrm().QueryTable("sku").Filter("art_no_id", artNoId).Filter("lock", 0).RelatedSel().All(&skus)
	this.ReturnData(skus)
}



//@router /admin/art_no/create [post]
func (this *AjaxController)CreateArtNo() {
	artno := this.GetString("artNo")
	productName := this.GetString("productName")
	liningId, _ := this.GetInt64("liningId")
	fabricId, _ := this.GetInt64("fabricId")
	retailPrice, _ := this.GetFloat("retailPrice")
	factoryPrice, _ := this.GetFloat("factoryPrice")
	colorIds := this.GetStrings("colors[]")
	sizeIds := this.GetStrings("sizes[]")

	var artNo models.ArtNo
	artNo.GmtCreate = time.Now()
	artNo.GmtModifier = time.Now()
	artNo.Fabric = &models.Fabric{Id:fabricId}
	artNo.Lining = &models.Lining{Id:liningId}
	artNo.RetailPrice = retailPrice
	artNo.FactoryPrice = factoryPrice
	artNo.ArtNo = artno
	artNo.ProductName = productName

	//事务控制
	o := orm.NewOrm()
	//开启事务
	o.Begin()
	//如果插入货号失败，则回滚
	if _, err := o.Insert(&artNo); err != nil {
		o.Rollback()
		beego.BeeLogger.Error("插入art_no表失败,错误详情:%s", err.Error())
		this.ReturnError("新增失败")
		return
	}

	//插入销售属性 颜色和储存
	cm2m := o.QueryM2M(&artNo, "Colors")
	for _, colorId := range colorIds {
		colorId, _ := strconv.ParseInt(colorId, 10, 64)
		if _, err := cm2m.Add(models.Color{Id:colorId}); err != nil {
			o.Rollback()
			this.ReturnError("新增失败")
			return
		}
	}
	sm2m := o.QueryM2M(&artNo, "Sizes")
	for _, sizeId := range sizeIds {
		sizeId, _ := strconv.ParseInt(sizeId, 10, 64)
		if _, err := sm2m.Add(models.Size{Id:sizeId}); err != nil {
			o.Rollback()
			this.ReturnError("新增失败")
			return
		}
	}
	//提交事务
	o.Commit()
	this.ReturnSuccess()
}

// @router /admin/art_no/:id/update [post]
func (this *AjaxController)UpdateArtNo() {
	artNoId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var artNo models.ArtNo
	artNo.Id = artNoId
	if err := artNo.Read(); err != nil {
		this.ReturnError("artno not found")
		return
	}

	artno := this.GetString("artNo")
	productName := this.GetString("productName")
	liningId, _ := this.GetInt64("liningId")
	fabricId, _ := this.GetInt64("fabricId")
	retailPrice, _ := this.GetFloat("retailPrice")
	factoryPrice, _ := this.GetFloat("factoryPrice")
	colorIds := this.GetStrings("colors[]")
	sizeIds := this.GetStrings("sizes[]")

	artNo.GmtModifier = time.Now()
	artNo.Fabric = &models.Fabric{Id:fabricId}
	artNo.Lining = &models.Lining{Id:liningId}
	artNo.RetailPrice = retailPrice
	artNo.FactoryPrice = factoryPrice
	artNo.ArtNo = artno
	artNo.ProductName = productName
	artNo.Id = artNoId

	o := orm.NewOrm()
	o.Begin()



	//如果更新货号失败，则回滚
	if _, err := o.Update(&artNo); err != nil {
		o.Rollback()
		beego.BeeLogger.Error("更新art_no表失败,错误详情:%s", err.Error())
		this.ReturnError("更新失败")
		return
	}

	//获取旧颜色和尺寸属性
	o.LoadRelated(&artNo, "Colors")
	o.LoadRelated(&artNo, "Sizes")

	//删除原先的销售属性
	if _, err := o.QueryTable("art_no_color").Filter("art_no_id", artNoId).Delete(); err != nil {
		o.Rollback()
		this.ReturnError("更新失败")
		return
	}

	if _, err := o.QueryTable("art_no_size").Filter("art_no_id", artNoId).Delete(); err != nil {
		o.Rollback()
		this.ReturnError("更新失败")
		return
	}

	//重新插入销售属性 颜色和储存
	cm2m := o.QueryM2M(&artNo, "Colors")
	for _, colorId := range colorIds {
		colorId, _ := strconv.ParseInt(colorId, 10, 64)
		if _, err := cm2m.Add(models.Color{Id:colorId}); err != nil {
			o.Rollback()
			this.ReturnError("更新失败")
			return
		}
	}
	sm2m := o.QueryM2M(&artNo, "Sizes")
	for _, sizeId := range sizeIds {
		sizeId, _ := strconv.ParseInt(sizeId, 10, 64)
		if _, err := sm2m.Add(models.Size{Id:sizeId}); err != nil {
			o.Rollback()
			this.ReturnError("更新失败")
			return
		}
	}
	o.Commit()
	this.ReturnSuccess()
}

// @router /admin/art_no/:id/delete [delete]
func (this *AjaxController) DeleteArtNo() {
	artNoId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	artNo := models.ArtNo{Id:artNoId}
	if _, err := artNo.Delete(); err != nil {
		this.ReturnError("artNo not found")
		return
	}
	o := orm.NewOrm();
	o.QueryTable("sku").Filter("art_no_id", artNoId).Update(orm.Params{"lock": 1})
	this.ReturnSuccess()
}


// @router /admin/art_no/sku/:id/delete [delete]
func (this *AjaxController) DeleteSku() {
	skuId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	sku := models.Sku{Id:skuId}
	if _, err := sku.Delete(); err != nil {
		this.ReturnError("sku not found")
		return
	}
	this.ReturnSuccess()
}

// @router /admin/art_no/:id/sku/create [post]
func (this *AjaxController) CreateSku() {
	artNoId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	retailPrice, _ := this.GetFloat("retailPrice")
	factoryPrice, _ := this.GetFloat("factoryPrice")
	colorId, _ := this.GetInt64("colorId")
	sizeId, _ := this.GetInt64("sizeId")
	if count, err := orm.NewOrm().QueryTable("sku").
		Filter("color_id", colorId).
		Filter("size_id", sizeId).
		Filter("art_no_id", artNoId).Filter("lock", 0).Count();
		err != nil {
		this.ReturnError("create sku error")
		return
	} else {
		if (count >= 1) {
			this.ReturnError("sku已经存在")
			return
		}
	}
	var sku models.Sku
	sku.GmtCreate = time.Now()
	sku.GmtModifier = time.Now()
	sku.RetailPrice = retailPrice
	sku.FactoryPrice = factoryPrice
	sku.Color = &models.Color{Id:colorId}
	sku.Size = &models.Size{Id:sizeId}
	sku.ArtNo = &models.ArtNo{Id:artNoId}
	if _, err := sku.Insert(); err != nil {
		this.ReturnError("sku已经存在");
		return;
	}
	this.ReturnSuccess()
}

// @router /admin/art_no/:art_no_id/sku/:id/update [post]
func (this *AjaxController) UpdateSku() {
	skuId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	retailPrice, _ := this.GetFloat("retailPrice")
	factoryPrice, _ := this.GetFloat("factoryPrice")
	colorId, _ := this.GetInt64("colorId")
	sizeId, _ := this.GetInt64("sizeId")
	art_no_id, _ := strconv.ParseInt(this.Ctx.Input.Param(":art_no_id"), 10, 64)
	var sku models.Sku
	sku.Id = skuId

	if count, err := orm.NewOrm().QueryTable("sku").
		Filter("color_id", colorId).
		Filter("size_id", sizeId).
		Filter("art_no_id", art_no_id).Filter("lock", 0).Count();
		err != nil {
		this.ReturnError("update sku error")
		return
	} else {
		if (count >= 1) {
			this.ReturnError("sku已经存在")
			return
		}
	}
	sku.GmtModifier = time.Now()
	sku.RetailPrice = retailPrice
	sku.FactoryPrice = factoryPrice
	sku.Color = &models.Color{Id:colorId}
	sku.Size = &models.Size{Id:sizeId}
	sku.Id = skuId
	if result, err := sku.Update(); err != nil {
		this.ReturnError("sku update error");
		return;
	} else {
		if result == 0 {
			this.ReturnError("sku not found");
			return;
		}
	}
	this.ReturnSuccess()
}


// @router /admin/art_no/:id/getColorsAndSizes [get]
func (this *AjaxController) GetColorsAndSizes() {
	artNoId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	result := make(map[string]interface{})
	artNo := models.ArtNo{Id:artNoId}
	artNo.Read()
	orm.NewOrm().LoadRelated(&artNo, "Colors")
	orm.NewOrm().LoadRelated(&artNo, "Sizes")
	result["colors"] = artNo.Colors
	result["sizes"] = artNo.Sizes
	this.ReturnData(result)
}




