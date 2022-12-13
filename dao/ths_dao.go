package dao

import (
	db "dss-base-data/db"
	"log"

	"github.com/yangyang-hub/dss-common/constant"
	"github.com/yangyang-hub/dss-common/model"
)

//新增同花顺行业行情信息
func InsertThsHyQuote(thsHyQuotes *[]model.ThsHyQuote) {
	res := db.Mysql.CreateInBatches(thsHyQuotes, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//新增同花顺概念行情信息
func InsertThsGnQuote(thsGnQuotes *[]model.ThsGnQuote) {
	res := db.Mysql.CreateInBatches(thsGnQuotes, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//查询所有同花顺行业
func QueryAllThsHy() *[]model.ThsHy {
	thsHys := []model.ThsHy{}
	res := db.Mysql.Find(&thsHys).Error
	if res != nil {
		log.Println(res.Error())
	}
	return &thsHys
}

//查询所有同花顺概念
func QueryAllThsGn() *[]model.ThsGn {
	thsGns := []model.ThsGn{}
	res := db.Mysql.Find(&thsGns).Error
	if res != nil {
		log.Println(res.Error())
	}
	return &thsGns
}

//新增同花顺概念
func InsertThsGn(thsGns *[]model.ThsGn) {
	res := db.Mysql.CreateInBatches(thsGns, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//删除所有同花顺概念
func DeleteAllThsGn() {
	res := db.Mysql.Where("1 = 1").Delete(&model.ThsGn{}).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//新增同花顺概念与股票代码关联关系
func InsertThsGnRelSymbol(thsGnRelSymbols *[]model.ThsGnRelSymbol) {
	res := db.Mysql.CreateInBatches(thsGnRelSymbols, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//删除同花顺概念与股票代码关联关系
func DeleteAllThsGnRelSymbol() {
	res := db.Mysql.Where("1 = 1").Delete(&model.ThsGnRelSymbol{}).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//新增同花顺行业
func InsertThsHy(thsHys *[]model.ThsHy) {
	res := db.Mysql.CreateInBatches(thsHys, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//删除所有同花顺行业
func DeleteAllThsHy() {
	res := db.Mysql.Where("1 = 1").Delete(&model.ThsHy{}).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//新增同花顺行业与股票代码关联关系
func InsertThsHyRelSymbol(thsHyRelSymbols *[]model.ThsHyRelSymbol) {
	res := db.Mysql.CreateInBatches(thsHyRelSymbols, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//删除同花顺行业与股票代码关联关系
func DeleteAllThsHyRelSymbol() {
	res := db.Mysql.Where("1 = 1").Delete(&model.ThsHyRelSymbol{}).Error
	if res != nil {
		log.Println(res.Error())
	}
}
