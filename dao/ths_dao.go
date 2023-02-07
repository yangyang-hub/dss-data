package dao

import (
	db "dss-data/db"
	"log"

	"github.com/yangyang-hub/dss-common/constant"
	"github.com/yangyang-hub/dss-common/model"
)

// 查询股票所包含的概念
func QueryThsGnBySymbols(symbols []string) *[]model.ThsGnRelSymbol {
	rels := []model.ThsGnRelSymbol{}
	res := db.Mysql.Where("symbol IN ?", symbols).Find(&rels).Error
	if res != nil {
		log.Println(res.Error())
	}
	return &rels
}

// 查询概念所包含的股票代码
func QuerySymbolsByThsGn(gns []string) *[]model.ThsGnRelSymbol {
	rels := []model.ThsGnRelSymbol{}
	res := db.Mysql.Where("gn_name IN ?", gns).Find(&rels).Error
	if res != nil {
		log.Println(res.Error())
	}
	return &rels
}

// 查询所有同花顺概念
func QueryAllThsGn() *[]model.ThsGn {
	thsGns := []model.ThsGn{}
	res := db.Mysql.Find(&thsGns).Error
	if res != nil {
		log.Println(res.Error())
	}
	return &thsGns
}

// 新增同花顺概念
func InsertThsGn(thsGns *[]model.ThsGn) {
	res := db.Mysql.CreateInBatches(thsGns, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 删除所有同花顺概念
func DeleteAllThsGn() {
	res := db.Mysql.Where("1 = 1").Delete(&model.ThsGn{}).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 批量新增同花顺概念与股票代码关联关系
func InsertThsGnRelSymbols(thsGnRelSymbols *[]model.ThsGnRelSymbol) {
	res := db.Mysql.CreateInBatches(thsGnRelSymbols, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 新增同花顺概念与股票代码关联关系
func InsertThsGnRelSymbol(thsGnRelSymbols *model.ThsGnRelSymbol) {
	res := db.Mysql.Create(thsGnRelSymbols).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 删除某个股票代码关联的同花顺概念
func DeleteThsGnRelBySymbol(symbol string) {
	res := db.Mysql.Where("symbol = ?", symbol).Delete(&model.ThsGnRelSymbol{}).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 清除退市股票
func DeleteThsGnRelSymbolNotExist() {
	res := db.Mysql.Where("symbol NOT IN (SELECT symbol FROM stock_info)").Delete(&model.ThsGnRelSymbol{}).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 删除同花顺概念与股票代码关联关系
func DeleteAllThsGnRelSymbol() {
	res := db.Mysql.Where("1 = 1").Delete(&model.ThsGnRelSymbol{}).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 新增同花顺概念行情信息
func InsertThsGnQuote(thsGnQuotes *[]model.ThsGnQuote) {
	res := db.Mysql.CreateInBatches(thsGnQuotes, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}
