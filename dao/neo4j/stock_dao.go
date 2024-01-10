package dao

import (
	"bytes"
	"dss-data/db"
	"dss-data/model"
	"dss-data/util"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// 新增stock_info数据
func InsertStockInfo(stockInfos *[]db.Node[model.StockInfo]) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *stockInfos {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("CREATE (n")
		for _, label := range node.Label {
			buffer.WriteString(":")
			buffer.WriteString(label)
		}
		properties, param := util.BuildNodeProperties(node.Properties, ":", "")
		buffer.WriteString("{")
		buffer.WriteString(properties)
		buffer.WriteString("}")
		buffer.WriteString(")")
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherWrite(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherWrite(cyphers)
}

// 更新（新增）stock_info数据
func MergeStockInfo(stockInfos *[]db.Node[model.StockInfo]) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *stockInfos {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MERGE (n")
		for _, label := range node.Label {
			buffer.WriteString(":")
			buffer.WriteString(label)
		}
		buffer.WriteString("{ ts_code: $ts_code" + "}) ")
		buffer.WriteString(" ON CREATE SET ")
		properties, param := util.BuildNodeProperties(node.Properties, "=", "n.")
		buffer.WriteString(properties)
		buffer.WriteString("\n ON MATCH  SET ")
		buffer.WriteString(properties)
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherWrite(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherWrite(cyphers)
}

// 新增quote数据（）
func InsertStockQuote(stockInfos *[]db.Node[model.StockQuote]) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *stockInfos {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MATCH (n:")
		buffer.WriteString(db.StockInfo.String())
		buffer.WriteString("{symbol: $ts_code})\n CREATE (c")
		for _, label := range node.Label {
			buffer.WriteString(":")
			buffer.WriteString(label)
		}
		properties, param := util.BuildNodeProperties(node.Properties, ":", "")
		buffer.WriteString("{")
		buffer.WriteString(properties)
		buffer.WriteString("})-[r:")
		buffer.WriteString(db.RelStockQuote.String())
		buffer.WriteString("]->(n)")
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherWrite(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherWrite(cyphers)
}

// 新增板块数据（）
func InsertBk(bks *[]db.Node[model.Bk]) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *bks {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MERGE (n")
		for _, label := range node.Label {
			buffer.WriteString(":")
			buffer.WriteString(label)
		}
		buffer.WriteString("{ code: code" + "}) ")
		buffer.WriteString(" ON CREATE SET ")
		properties, param := util.BuildNodeProperties(node.Properties, "=", "n.")
		buffer.WriteString(properties)
		buffer.WriteString("\n ON MATCH  SET ")
		buffer.WriteString(properties)
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherWrite(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherWrite(cyphers)
}

// 关联板块-股票
func InsertBkRelSymbol(bkRelSymbols *[]db.Edge[model.BkRelSymbol]) {
	var cyphers []map[string]interface{}
	step := 100
	for i, node := range *bkRelSymbols {
		maps := make(map[string]interface{})
		buffer := bytes.Buffer{}
		buffer.WriteString("MATCH (a:")
		buffer.WriteString(db.StockInfo.String())
		buffer.WriteString(")-[r:")
		buffer.WriteString(db.RelStockBk.String())
		buffer.WriteString("]->(b:")
		buffer.WriteString(db.Bk.String())
		buffer.WriteString(")\n SET ")
		properties, param := util.BuildNodeProperties(node.Properties, "=", "r.")
		buffer.WriteString(properties)
		buffer.WriteString("\n WHERE a.symbol = $symbol AND b.code = $bk_code")
		maps["cypher"] = buffer.String()
		maps["param"] = param
		cyphers = append(cyphers, maps)
		if (i+1)%step == 0 {
			db.CypherWrite(cyphers)
			cyphers = cyphers[:0]
		}
	}
	db.CypherWrite(cyphers)
}

// 查询所有的股票编码数据（ts_code）
func GetAllTsCode() ([]string, error) {
	driver := *db.Neo4j
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	results, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string
		result, err := tx.Run("MATCH (n:"+db.StockInfo.String()+") RETURN n.ts_code", nil)
		if err != nil {
			return nil, err
		}
		for result.Next() {
			list = append(list, result.Record().Values[0].(string))
		}
		if err = result.Err(); err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return results.([]string), nil
}
