package admin

import (
	"time"
	"github.com/astaxie/beego/orm"
	"cims/tool"
	"fmt"
	"cims/models"
	"strconv"
)
//@router /admin/out_stock/:id/update [post]
func (this *AjaxController) UpdateOutStock() {
	if !this.Verity() {
		this.ReturnUnAuth()
		return
	}
	outStockId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	outStockNum := this.GetString("outStockNum")
	wareHouseId, _ := this.GetInt64("warehouseId")
	receiverName := this.GetString("receiverName")
	receiverTel := this.GetString("receiverTel")
	receiverAddress := this.GetString("address")
	remark := this.GetString("remark")
	deliveryTime, _ := time.Parse("2006-01-02 03:04:05", this.GetString("deliveryTime"))
	outStock := models.OutStock{Id:outStockId, OutStockNum:outStockNum,
		ReceiverName:receiverName,
		ReceiverTel:receiverTel,
		Address:receiverAddress,
		WareHouse:&models.WareHouse{Id:wareHouseId},
		Remark:remark,
		GmtModifier:time.Now(),
		DeliveryTime:deliveryTime}
	if result, err := outStock.Update(); err != nil {
		fmt.Println(err.Error())
		this.ReturnError("更新失败")
		return
	} else {
		if result == 0 {
			fmt.Println(err.Error())
			this.ReturnError("outstock not found ")
			return
		}
	}
	this.ReturnSuccess()
}

//@router /admin/out_stock/create [post]
func (this *AjaxController) CreateOutStock() {
	if !this.Verity() {
		this.ReturnUnAuth()
		return
	}
	outStockId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	outStockNum := this.GetString("outStockNum")
	wareHouseId, _ := this.GetInt64("warehouseId")
	receiverName := this.GetString("receiverName")
	receiverTel := this.GetString("receiverTel")
	receiverAddress := this.GetString("address")
	remark := this.GetString("remark")
	deliveryTime, _ := time.Parse("2006-01-02 03:04:05", this.GetString("deliveryTime"))
	outStock := models.OutStock{Id:outStockId, OutStockNum:outStockNum,
		ReceiverName:receiverName,
		ReceiverTel:receiverTel,
		Address:receiverAddress,
		WareHouse:&models.WareHouse{Id:wareHouseId},
		Remark:remark,
		GmtModifier:time.Now(), GmtCreate:time.Now(),
		DeliveryTime:deliveryTime}
	if _, err := outStock.Insert(); err != nil {
		fmt.Println(err.Error())
		this.ReturnError("新增失败")
		return
	}
	this.ReturnSuccess()
}

//@router /admin/out_stock/:id/delete [delete]
func (this *AjaxController) DeleteOutStock() {
	outStockId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	outStock := models.OutStock{Id:outStockId}
	if _, err := outStock.Delete(); err != nil {
		this.ReturnError("outstock not found,delete fail")
		return
	}
	orm.NewOrm().QueryTable("out_stock_detail").Filter("out_stock_id", outStockId).Update(orm.Params{"lock":1})
	this.ReturnSuccess()
}

//@router /admin/out_stock/:id/finish [post]
func (this *AjaxController) FinishOutStock() {
	outStockId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	outStock := models.OutStock{Id:outStockId}
	//检查出库单是否存在
	if err := outStock.Read(); err != nil {
		this.ReturnError(" out_stock not found")
		return
	}
	//检查仓库是否存在
	if outStock.WareHouse.Lock == 1 {
		this.ReturnError("仓库名为" + outStock.WareHouse.Name +"的仓库已经被删除,无法提交订单");
		return
	}

	o := orm.NewOrm()
	var outStockDetails []*models.OutStockDetail
	if _, err := o.QueryTable("out_stock_detail").Filter("out_stock_id", outStockId).Filter("lock", 0).RelatedSel().All(&outStockDetails); err != nil {
		this.ReturnError(" out_stock not found")
		return
	}

	//检查出库单里是否有出库详情
	if len(outStockDetails) == 0 {
		this.ReturnError("该出库单中没有任何出库详情,请先添加出库详情再提交")
		return
	}

	num := 0
	for _, outStockDetail := range outStockDetails {
		num += outStockDetail.Count
	}

	o.Begin()
	//更新sku库存
	for _, outStockDetail := range outStockDetails {
		if result, err := o.QueryTable("sku").Filter("art_no_id", outStockDetail.ArtNo).Filter("color_id", outStockDetail.Color.Id).Filter("size_id", outStockDetail.Size.Id).Filter("lock", 0).Filter("stock__gte", outStockDetail.Count).Update(orm.Params{"stock":orm.ColValue(orm.ColMinus, outStockDetail.Count)});
			err != nil {
			o.Rollback()
			this.ReturnError("finish out_stock error")
			return
		} else {
			if (result == 0) {
				o.Rollback()
				this.ReturnError("货号为" + outStockDetail.ArtNo.ArtNo + ",尺寸为" + strconv.Itoa(outStockDetail.Size.Size) + ",颜色为" + outStockDetail.Color.Color + "的商品已失效或者库存不足，请重新编辑再提交订单");
				return
			}
		}
	}
	//更新仓库库存
	if index, err := o.QueryTable("ware_house").Filter("id", outStock.WareHouse.Id).Filter("lock", 0).Filter("stock__gte", num).Update(orm.Params{"gmt_modifier":time.Now(), "stock" :orm.ColValue(orm.ColMinus, num)}); err != nil {
		o.Rollback()
		this.ReturnError("finish out_stock error")
		return
	} else {
		if (index == 0) {
			o.Rollback()
			this.ReturnError("仓库库存异常，提交失败")
			return
		}
	}

	//更新订单状态
	if _, err := o.QueryTable("out_stock").Filter("id", outStockId).Filter("status", 0).Update(orm.Params{"status":1});
		err != nil {
		o.Rollback()
		this.ReturnError("finish out_stock error")
		return
	}
	o.Commit()
	this.ReturnSuccess()
}



//@router /admin/out_stock_detail/create [post]
func (this *AjaxController) CreateOutStockDetail() {
	artNoId, _ := this.GetInt64("artNoId")
	colorId, _ := this.GetInt64("colorId")
	sizeId, _ := this.GetInt64("sizeId")
	count, _ := this.GetInt("count")
	outStockId, _ := this.GetInt64("outStockId")
	var outStockDetail models.OutStockDetail
	outStockDetail.GmtCreate = time.Now()
	outStockDetail.GmtModifier = time.Now()
	outStockDetail.OutStock = &models.OutStock{Id:outStockId}
	outStockDetail.ArtNo = &models.ArtNo{Id:artNoId}
	outStockDetail.Color = &models.Color{Id:colorId}
	outStockDetail.Size = &models.Size{Id:sizeId}
	outStockDetail.Count = count
	if _, err := outStockDetail.Insert(); err != nil {
		this.ReturnError("新增失败")
		return
	}

	this.ReturnSuccess()
}

//@router /admin/out_stock_detail/:id/update [post]
func (this *AjaxController) UpdateOutStockDetail() {
	outStockDetailId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	artNoId, _ := this.GetInt64("artNoId")
	colorId, _ := this.GetInt64("colorId")
	sizeId, _ := this.GetInt64("sizeId")
	count, _ := this.GetInt("count")
	outStockId, _ := this.GetInt64("outStockId")
	var outStockDetail models.OutStockDetail
	outStockDetail.GmtModifier = time.Now()
	outStockDetail.Id = outStockDetailId
	outStockDetail.OutStock = &models.OutStock{Id:outStockId}
	outStockDetail.ArtNo = &models.ArtNo{Id:artNoId}
	outStockDetail.Color = &models.Color{Id:colorId}
	outStockDetail.Size = &models.Size{Id:sizeId}
	outStockDetail.Count = count
	if _, err := outStockDetail.Update(); err != nil {
		this.ReturnError("更新失败")
		return
	}
	this.ReturnSuccess()
}

//@router /admin/out_stock_detail/:id/delete [delete]
func (this *AjaxController) DeleteOutStockDetail() {
	outStockDetailId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	outStockDetail := models.OutStockDetail{Id:outStockDetailId}
	if _, err := outStockDetail.Delete(); err != nil {
		this.ReturnError("outstock not found,delete fail")
		return
	}
	this.ReturnSuccess()
}


//@router /admin/out_stock/:id/out_stock_detail/list [get]
func (this *AjaxController) ListOutStockDetails() {
	outStockId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var outStockDetails []*models.OutStockDetail
	outStock := models.OutStock{Id:outStockId}
	if err := outStock.Read(); err != nil {
		this.ReturnError("outstock not found")
		return
	}
	orm.NewOrm().QueryTable("out_stock_detail").Filter("out_stock_id", outStockId).Filter("lock", 0).RelatedSel().All(&outStockDetails)
	//如果入库单未提交，检查入库详情中商品是否有效
	if (outStock.Status == 0) {
		for _, outStockDetail := range outStockDetails {
			var sku models.Sku
			if err := orm.NewOrm().QueryTable("sku").Filter("lock", 0).Filter("color_id", outStockDetail.Color.Id).Filter("size_id", outStockDetail.Size.Id).Filter("art_no_id", outStockDetail.ArtNo.Id).One(&sku);
				err != nil {
				//商品失效
				outStockDetail.Can = -1

			}else {
				//商品库存不足
				if(sku.Stock < outStockDetail.Count) {
					outStockDetail.Can = -2
				}
			}
		}
	}
	this.ReturnData(outStockDetails)
}





//@router /admin/out_stock/getOutStockNum [get]
func (this *AjaxController) GetOutStockNum() {
	time := time.Now().Format("2006-01-02")
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("count(*)").
		From("out_stock").
		Where("gmt_create like '" + time + "%'").
		Limit(1);
	var count int64;
	// 导出 SQL 语句
	sql := qb.String()
	fmt.Println(sql)
	// 执行 SQL 语句
	o := orm.NewOrm()
	if err := o.Raw(sql).QueryRow(&count); err != nil {
		fmt.Println(err.Error())
	}
	this.ReturnData(tool.GetOutStockNum(count + 1))
}
