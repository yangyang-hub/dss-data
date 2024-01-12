package dao

import (
	"bytes"
	"dss-data/db"
	"dss-data/model"
	"dss-data/util"
)

// 新增stock_info数据
func InsertStockInfo(stockInfos *[]model.StockInfo) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *stockInfos {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("CREATE (n:")
		buffer.WriteString(db.StockInfo.String())
		properties, param := util.BuildNodeProperties(node, ":", "")
		buffer.WriteString("{")
		buffer.WriteString(properties)
		buffer.WriteString("}")
		buffer.WriteString(")")
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherBatchExec(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherBatchExec(cyphers)
}

// 更新（新增）stock_info数据
func MergeStockInfo(stockInfos *[]model.StockInfo) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *stockInfos {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MERGE (n:")
		buffer.WriteString(db.StockInfo.String())
		buffer.WriteString("{ ts_code: $ts_code" + "}) ")
		buffer.WriteString(" ON CREATE SET ")
		properties, param := util.BuildNodeProperties(node, "=", "n.")
		buffer.WriteString(properties)
		buffer.WriteString("\n ON MATCH  SET ")
		buffer.WriteString(properties)
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherBatchExec(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherBatchExec(cyphers)
}

// 删除不存在了的股票
func DeleteStockInfo(symbols *[]string) {
	buffer := bytes.Buffer{}
	buffer.WriteString("MATCH (n:")
	buffer.WriteString(db.StockInfo.String())
	buffer.WriteString(")\nWHERE n.symbol IN $symbol")
	buffer.WriteString("\nDELETE n")
	param := make(map[string]interface{})
	param["symbol"] = symbols
	db.CypherExec(buffer.String(), param)
}

// 新增quote数据（）
func InsertStockQuote(stockInfos *[]model.StockQuote) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *stockInfos {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MATCH (n:")
		buffer.WriteString(db.StockInfo.String())
		buffer.WriteString("{symbol: $ts_code})\n CREATE (c:")
		buffer.WriteString(db.StockQuote.String())
		properties, param := util.BuildNodeProperties(node, ":", "")
		buffer.WriteString("{")
		buffer.WriteString(properties)
		buffer.WriteString("})-[r:")
		buffer.WriteString(db.RelStockQuote.String())
		buffer.WriteString("]->(n)")
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherBatchExec(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherBatchExec(cyphers)
}

// 通过日期删除股票行情
func DeleteStockQuoteByDate(dates []string) {
	buffer := bytes.Buffer{}
	buffer.WriteString("MATCH (a:")
	buffer.WriteString(db.StockQuote.String())
	buffer.WriteString(")")
	buffer.WriteString("\nWHERE NOT a.trade_date IN $trade_date")
	buffer.WriteString("\nDELETE a")
	cypher := buffer.String()
	param := make(map[string]interface{})
	param["trade_date"] = dates
	db.CypherExec(cypher, param)
}

// 新增板块数据（）
func MergeBk(bks *[]model.Bk) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *bks {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MERGE (n:")
		buffer.WriteString(db.Bk.String())
		buffer.WriteString("{ code: $code" + "}) ")
		buffer.WriteString(" ON CREATE SET ")
		properties, param := util.BuildNodeProperties(node, "=", "n.")
		buffer.WriteString(properties)
		buffer.WriteString("\n ON MATCH  SET ")
		buffer.WriteString(properties)
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherBatchExec(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherBatchExec(cyphers)
}

// 删除不存在了的板块
func DeleteBk(bkCodes *[]string) {
	buffer := bytes.Buffer{}
	buffer.WriteString("MATCH (n:")
	buffer.WriteString(db.Bk.String())
	buffer.WriteString(")\nWHERE n.code IN $code")
	buffer.WriteString("\nDELETE n")
	param := make(map[string]interface{})
	param["code"] = bkCodes
	db.CypherExec(buffer.String(), param)
}

// 删除板块-股票关联
func DeleteBkRelSymbol() {
	buffer := bytes.Buffer{}
	buffer.WriteString("MATCH (a:")
	buffer.WriteString(db.StockInfo.String())
	buffer.WriteString(")-[r:")
	buffer.WriteString(db.RelStockBk.String())
	buffer.WriteString("]->(b:")
	buffer.WriteString(db.Bk.String())
	buffer.WriteString(")")
	buffer.WriteString("\nDELETE r")
	cypher := buffer.String()
	db.CypherExec(cypher, nil)
}

// 关联板块-股票
func InsertBkRelSymbol(bkRelSymbols *[]model.BkRelSymbol) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *bkRelSymbols {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MATCH (a:")
		buffer.WriteString(db.StockInfo.String())
		buffer.WriteString("),(b:")
		buffer.WriteString(db.Bk.String())
		buffer.WriteString(")\n WHERE a.symbol = $symbol AND b.code = $bk_code")
		buffer.WriteString("\nCREATE (a)-[r:")
		buffer.WriteString(db.RelStockBk.String())
		buffer.WriteString("{")
		properties, param := util.BuildNodeProperties(node, ":", "")
		buffer.WriteString(properties)
		buffer.WriteString("}")
		buffer.WriteString("]->(b)")
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherBatchExec(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherBatchExec(cyphers)
}

// 关联板块行情
func InsertBkQuote(bkQuotes *[]model.BkQuote) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *bkQuotes {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MATCH (n:")
		buffer.WriteString(db.Bk.String())
		buffer.WriteString("{code: $bk_code})\n CREATE (c:")
		buffer.WriteString(db.BkQuote.String())
		properties, param := util.BuildNodeProperties(node, ":", "")
		buffer.WriteString("{")
		buffer.WriteString(properties)
		buffer.WriteString("})-[r:")
		buffer.WriteString(db.RelBkQuote.String())
		buffer.WriteString("]->(n)")
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherBatchExec(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherBatchExec(cyphers)
}

// 通过日期删除板块行情
func DeleteBkQuoteByDate(dates []string) {
	buffer := bytes.Buffer{}
	buffer.WriteString("MATCH (a:")
	buffer.WriteString(db.BkQuote.String())
	buffer.WriteString(")")
	buffer.WriteString("\nWHERE NOT a.trade_date IN $trade_date")
	buffer.WriteString("\nDELETE a")
	cypher := buffer.String()
	param := make(map[string]interface{})
	param["trade_date"] = dates
	db.CypherExec(cypher, param)
}

// 新增龙虎榜
func InsertLongHu(datas *[]model.LongHu) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *datas {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MATCH (n:")
		buffer.WriteString(db.StockInfo.String())
		buffer.WriteString("{symbol: $symbol})\n CREATE (c:")
		buffer.WriteString(db.LongHu.String())
		properties, param := util.BuildNodeProperties(node, ":", "")
		buffer.WriteString("{")
		buffer.WriteString(properties)
		buffer.WriteString("})-[r:")
		buffer.WriteString(db.RelStockLongHu.String())
		buffer.WriteString("]->(n)")
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherBatchExec(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherBatchExec(cyphers)
}

// 通过日期删除龙虎榜
func DeleteLongHuByDate(dates []string) {
	buffer := bytes.Buffer{}
	buffer.WriteString("MATCH (a:")
	buffer.WriteString(db.LongHu.String())
	buffer.WriteString(")-[r:")
	buffer.WriteString(db.RelLongHuDetail.String())
	buffer.WriteString("]-(b:")
	buffer.WriteString(db.LongHuDetail.String())
	buffer.WriteString(")\nWHERE NOT a.trade_date IN $trade_date")
	buffer.WriteString("\nDELETE a,b")
	cypher := buffer.String()
	param := make(map[string]interface{})
	param["trade_date"] = dates
	db.CypherExec(cypher, param)
}

// 新增龙虎榜详情
func InsertLongHuDetail(datas *[]model.LongHuDetail) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *datas {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MATCH (n:")
		buffer.WriteString(db.LongHu.String())
		buffer.WriteString("{id: $long_hu_id})\n CREATE (c:")
		buffer.WriteString(db.LongHuDetail.String())
		properties, param := util.BuildNodeProperties(node, ":", "")
		buffer.WriteString("{")
		buffer.WriteString(properties)
		buffer.WriteString("})-[r:")
		buffer.WriteString(db.RelLongHuDetail.String())
		buffer.WriteString("]->(n)")
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherBatchExec(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherBatchExec(cyphers)
}

// 查询所有的股票编码数据（ts_code）
func GetAllSymbol() (*[]string, error) {
	cypher := "MATCH (n:" + db.StockInfo.String() + ") RETURN n.symbol"
	return db.CypherExecReturnStringList(cypher, nil)
}

// 查询所有的板块编码
func GetAllBkCode() (*[]string, error) {
	cypher := "MATCH (n:" + db.Bk.String() + ") RETURN n.code"
	return db.CypherExecReturnStringList(cypher, nil)
}

// 查询所有的板块-股票关联
func GetAllBkRelSymbol() (*[]model.BkRelSymbol, error) {
	cypher := "MATCH (n:" + db.StockInfo.String() + ")-[r:" + db.RelStockBk.String() + "]->(t:" + db.Bk.String() + ") RETURN r.bk_code,symbol"
	results, _ := db.CypherExecReturnMapList(cypher, nil)
	var brss []model.BkRelSymbol
	for _, result := range *results {
		brs := model.BkRelSymbol{
			Symbol: result["symbol"].(string),
			BkCode: result["bk_code"].(string),
		}
		brss = append(brss, brs)
	}
	return &brss, nil
}
