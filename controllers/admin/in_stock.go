package admin

import (
	"time"
	"github.com/astaxie/beego/orm"
	"cims/tool"
	"fmt"
	"cims/models"
	"strconv"
)
//@router /admin/in_stock/:id/update [post]
func (this *AjaxController) UpdateInStock() {
	if !this.Verity() {
		this.ReturnUnAuth()
		return
	}
	inStockId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	intStockNum := this.GetString("inStockNum")
	wareHouseId, _ := this.GetInt64("warehouseId")
	remark := this.GetString("remark")
	source := this.GetString("source")
	storageTime, _ := time.Parse("2006-01-02 03:04:05", this.GetString("storageTime"))
	inStock := models.InStock{Id:inStockId, InStockNum:intStockNum,
		Remark:remark,
		Source:source,
		WareHouse:&models.WareHouse{Id:wareHouseId},
		GmtModifier:time.Now(), StorageTime:storageTime}
	if _, err := inStock.Update(); err != nil {
		fmt.Println(err.Error())
		this.ReturnError("新增失败")
		return
	}
	this.ReturnSuccess()
}

//@router /admin/in_stock/create [post]
func (this *AjaxController) CreateInStock() {
	if !this.Verity() {
		this.ReturnUnAuth()
		return
	}
	intStockNum := this.GetString("inStockNum")
	wareHouseId, _ := this.GetInt64("warehouseId")
	remark := this.GetString("remark")
	source := this.GetString("source")
	user := this.GetSession("userinfo").(models.User)
	storageTime, _ := time.Parse("2006-01-02 03:04:05", this.GetString("storageTime"))
	inStock := models.InStock{InStockNum:intStockNum,
		Remark:remark,
		Source:source,
		WareHouse:&models.WareHouse{Id:wareHouseId},
		GmtCreate:time.Now(), GmtModifier:time.Now(), StorageTime:storageTime, User:&user}
	if _, err := inStock.Insert(); err != nil {
		fmt.Println(err.Error())
		this.ReturnError("新增失败")
		return
	}
	this.ReturnSuccess()
}

//@router /admin/in_stock/:id/delete [delete]
func (this *AjaxController) DeleteInStock() {
	inStcokId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	inStock := models.InStock{Id:inStcokId}
	if _, err := inStock.Delete(); err != nil {
		this.ReturnError("instock not found,delete fail")
		return
	}
	orm.NewOrm().QueryTable("in_stock_detail").Filter("in_stock_id", inStcokId).Update(orm.Params{"lock":1})
	this.ReturnSuccess()
}

//@router /admin/in_stock/:id/finish [post]
func (this *AjaxController) FinishInStock() {
	inStockId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	inStock := models.InStock{Id:inStockId}
	//检查入库单是否存在
	if err := inStock.Read(); err != nil {
		this.ReturnError(" in_stock not found")
		return
	}
	//检查仓库是否存在
	if inStock.WareHouse.Lock == 1 {
		this.ReturnError("仓库名为" + inStock.WareHouse.Name +"的仓库已经被删除,无法提交订单");
		return
	}

	o := orm.NewOrm()
	var inStockDetails []*models.InStockDetail
	if _, err := o.QueryTable("in_stock_detail").Filter("in_stock_id", inStockId).Filter("lock", 0).RelatedSel().All(&inStockDetails); err != nil {
		this.ReturnError(" in_stock not found")
		return
	}

	//检查入库单中是否有入库详情
	if len(inStockDetails) == 0 {
		this.ReturnError("该入库单中没有任何入库详情,请先添加入库详情再提交")
		return
	}

	num := 0
	for _, inStockDetail := range inStockDetails {
		num += inStockDetail.Count
	}

	//开启事务
	o.Begin()
	//计算最新的仓库库存
	currentStock := num + inStock.WareHouse.Stock

	//更新仓库库存
	if index, err := o.QueryTable("ware_house").Filter("id", inStock.WareHouse.Id).Filter("lock", 0).Filter("capacity__gte", currentStock).Update(orm.Params{"gmt_modifier":time.Now(), "stock" :orm.ColValue(orm.ColAdd, num)}); err != nil {
		o.Rollback()
		this.ReturnError("finish in_stock error")
		return
	} else {
		if (index == 0) {
			o.Rollback()
			this.ReturnError(inStock.WareHouse.Name + "仓储量不足,请扩充仓储量。" + "当前仓储量为" + strconv.Itoa(inStock.WareHouse.Capacity) + "至少需要" + strconv.Itoa(currentStock) + "以上的仓储量")
			return
		}
	}
	//更新sku库存
	for _, instockDetail := range inStockDetails {
		if result, err := o.QueryTable("sku").Filter("art_no_id", instockDetail.ArtNo).Filter("color_id", instockDetail.Color.Id).Filter("size_id", instockDetail.Size.Id).Filter("lock", 0).Update(orm.Params{"stock":orm.ColValue(orm.ColAdd, instockDetail.Count)});
			err != nil {
			o.Rollback()
			this.ReturnError("finish in_stock error")
			return
		} else {
			if (result == 0) {
				o.Rollback()
				this.ReturnError("货号为" + instockDetail.ArtNo.ArtNo + ",尺寸为" + strconv.Itoa(instockDetail.Size.Size) + ",颜色为" + instockDetail.Color.Color + "的商品已不存或者已被修改，请先删除再提交订单");
				return
			}
		}
	}
	//更新订单状态
	if _, err := o.QueryTable("in_stock").Filter("id", inStockId).Filter("status", 0).Update(orm.Params{"status":1});
		err != nil {
		o.Rollback()
		this.ReturnError("finish in_stock error")
		return
	}
	o.Commit()
	this.ReturnSuccess()
}



//@router /admin/in_stock_detail/create [post]
func (this *AjaxController) CreateInStockDetail() {
	artNoId, _ := this.GetInt64("artNoId")
	colorId, _ := this.GetInt64("colorId")
	sizeId, _ := this.GetInt64("sizeId")
	count, _ := this.GetInt("count")
	inStockId, _ := this.GetInt64("inStockId")
	var inStockDetail models.InStockDetail
	inStockDetail.GmtCreate = time.Now()
	inStockDetail.GmtModifier = time.Now()
	inStockDetail.InStock = &models.InStock{Id:inStockId}
	inStockDetail.ArtNo = &models.ArtNo{Id:artNoId}
	inStockDetail.Color = &models.Color{Id:colorId}
	inStockDetail.Size = &models.Size{Id:sizeId}
	inStockDetail.Count = count
	if _, err := inStockDetail.Insert(); err != nil {
		this.ReturnError("新增失败")
		return
	}

	this.ReturnSuccess()
}

//@router /admin/in_stock_detail/:id/update [post]
func (this *AjaxController) UpdateInStockDetail() {
	inStockDetailId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	artNoId, _ := this.GetInt64("artNoId")
	colorId, _ := this.GetInt64("colorId")
	sizeId, _ := this.GetInt64("sizeId")
	count, _ := this.GetInt("count")
	inStockId, _ := this.GetInt64("inStockId")
	var inStockDetail models.InStockDetail
	inStockDetail.GmtModifier = time.Now()
	inStockDetail.Id = inStockDetailId
	inStockDetail.InStock = &models.InStock{Id:inStockId}
	inStockDetail.ArtNo = &models.ArtNo{Id:artNoId}
	inStockDetail.Color = &models.Color{Id:colorId}
	inStockDetail.Size = &models.Size{Id:sizeId}
	inStockDetail.Count = count
	if _, err := inStockDetail.Update(); err != nil {
		this.ReturnError("更新失败")
		return
	}
	this.ReturnSuccess()
}

//@router /admin/in_stock_detail/:id/delete [delete]
func (this *AjaxController) DeleteInStockDetail() {
	inStockDetailId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	inStockDetail := models.InStockDetail{Id:inStockDetailId}
	if _, err := inStockDetail.Delete(); err != nil {
		this.ReturnError("instock not found,delete fail")
		return
	}
	this.ReturnSuccess()
}


//@router /admin/int_stock/:id/int_stock_detail/list [get]
func (this *AjaxController) ListIntStockDetails() {
	inStcokId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var inStockDetails []*models.InStockDetail
	inStock := models.InStock{Id:inStcokId}
	if err := inStock.Read(); err != nil {
		this.ReturnError("instock not found")
		return
	}
	orm.NewOrm().QueryTable("in_stock_detail").Filter("in_stock_id", inStcokId).Filter("lock", 0).RelatedSel().All(&inStockDetails)

	//如果入库单未提交，检查入库详情中商品是否有效
	if (inStock.Status == 0) {
		for _, inStockDetail := range inStockDetails {
			var sku models.Sku
			if err := orm.NewOrm().QueryTable("sku").Filter("lock", 0).Filter("color_id", inStockDetail.Color.Id).Filter("size_id", inStockDetail.Size.Id).Filter("art_no_id", inStockDetail.ArtNo.Id).One(&sku);
			err != nil {
				//商品失效
				inStockDetail.Can = -1
			}

		}
	}

	this.ReturnData(inStockDetails)
}





//@router /admin/in_stock/getInStockNum [get]
func (this *AjaxController) GetInStockNum() {
	time := time.Now().Format("2006-01-02")

	//count1, _ := orm.NewOrm().QueryTable("in_stock").Filter("gmt_create__istartswith",time).Count();
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("count(*)").
		From("in_stock").
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
	this.ReturnData(tool.GetInStockNum(count + 1))
}
