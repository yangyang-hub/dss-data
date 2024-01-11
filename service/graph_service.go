package service

import (
	"dss-data/dao/mysql"
	neo4j "dss-data/dao/neo4j"
)

// 创建图
func UpdateGraph(date string) {
	//allSymbols := mysql.GetAllSymbol()

	stockInfos, _ := mysql.GetAllStockInfo()
	neo4j.MergeStockInfo(stockInfos)
	bks, _ := mysql.GetAllBk()
	neo4j.MergeBk(bks)
	bkRels, _ := mysql.GetAllBkRelSymbol()
	neo4j.InsertBkRelSymbol(bkRels)
	bkQuotes, _ := mysql.GetBkQuoteByDate(date)
	neo4j.InsertBkQuote(bkQuotes)
	longHus, _ := mysql.GetLongHuByDate(date)
	neo4j.InsertLongHu(longHus)
	longHuDetails, _ := mysql.GetLongHuDetailByDate(date)
	neo4j.InsertLongHuDetail(longHuDetails)
	stockQuotes, _ := mysql.GetStockQuoteByDate(date)
	neo4j.InsertStockQuote(stockQuotes)

}
