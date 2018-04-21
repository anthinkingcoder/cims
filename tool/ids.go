package tool

import (
	"time"
	"strconv"
)

//入库单号格式 RC + 日期(20160502) + 4位数字(当天入库数量+1,不足补0） 14位
func GetInStockNum(count int64) string {
	var inStockNum string
	inStockNum += "RC";
	inStockNum += time.Now().Format("20060102");
	scount := strconv.FormatInt(count,10)
	if (len(scount) == 1) {
		scount = "000" + scount;
	}
	if (len(scount) == 2) {
		scount = "00" + scount;
	}
	if (len(scount) == 3) {
		scount = "0" + scount;
	}
	inStockNum += scount
	return inStockNum

}

//出库单号格式 CK + 日期(20160502) + 4位数字(当天出库数量+1,不足补0）14位
func GetOutStockNum(count int64) string {
	var inStockNum string
	inStockNum += "CK";
	inStockNum += time.Now().Format("20060102");
	scount := strconv.FormatInt(count,10)
	if (len(scount) == 1) {
		scount = "000" + scount;
	}
	if (len(scount) == 2) {
		scount = "00" + scount;
	}
	if (len(scount) == 3) {
		scount = "0" + scount;
	}
	inStockNum += scount
	return inStockNum

}
