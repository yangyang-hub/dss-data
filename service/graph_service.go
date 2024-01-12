package service

import (
	"dss-data/dao/mysql"
	neo4j "dss-data/dao/neo4j"
)

// 创建图
func UpdateGraph(date string) {
	// 更新 stockInfo
	//allSymbols, _ := mysql.GetAllSymbol()
	//nSymbols, _ := neo4j.GetAllSymbol()
	//symbols := util.Difference(*nSymbols, *allSymbols)
	//if len(symbols) > 1 {
	//	neo4j.DeleteStockInfo(&symbols)
	//}
	//stockInfos, _ := mysql.GetAllStockInfo()
	//neo4j.MergeStockInfo(stockInfos)
	//// 更新板块
	//mbkCodes, _ := mysql.GetAllBkCode()
	//nbkCodes, _ := neo4j.GetAllBkCode()
	//bkCodes := util.Difference(*nbkCodes, *mbkCodes)
	//if len(bkCodes) > 1 {
	//	neo4j.DeleteBk(&bkCodes)
	//}
	//bks, _ := mysql.GetAllBk()
	//neo4j.MergeBk(bks)
	//// 更新板块-股票关联
	//neo4j.DeleteBkRelSymbol()
	//bkRels, _ := mysql.GetAllBkRelSymbol()
	//neo4j.InsertBkRelSymbol(bkRels)
	bkQuotes, _ := mysql.GetBkQuoteByDate(date)
	// 插入板块行情
	neo4j.InsertBkQuote(bkQuotes)
	longHus, _ := mysql.GetLongHuByDate(date)
	// 插入龙虎榜
	neo4j.InsertLongHu(longHus)
	longHuDetails, _ := mysql.GetLongHuDetailByDate(date)
	neo4j.InsertLongHuDetail(longHuDetails)
	// 插入股票行情
	stockQuotes, _ := mysql.GetStockQuoteByDate(date)
	neo4j.InsertStockQuote(stockQuotes)
	// 保留最近30个交易日的行情数据
	dates, _ := mysql.GetXDayTradeDate(30)
	neo4j.DeleteStockQuoteByDate(*dates)
	neo4j.DeleteBkQuoteByDate(*dates)
	neo4j.DeleteLongHuByDate(*dates)
}
