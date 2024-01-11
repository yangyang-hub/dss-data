package db

import (
	. "dss-data/configs"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

// 节点标签枚举
type NodeEnum int

func (n NodeEnum) String() string {
	return [...]string{"NT_STOCK_INFO", "NT_STOCK_QUOTE", "NT_BK", "NT_BK_QUOTE", "NT_LONG_HU", "NT_LONG_HU_DETAIL"}[n]
}

const (
	StockInfo NodeEnum = iota
	StockQuote
	Bk
	BkQuote
	LongHu
	LongHuDetail
)

// 边标签枚举
type EdgeEnum int

func (e EdgeEnum) String() string {
	return [...]string{"RT_STOCK_QUOTE", "RT_STOCK_BK", "RT_BK_QUOTE", "RT_STOCK_LONG_HU", "RT_LONG_HU_DETAIL"}[e]
}

const (
	RelStockQuote EdgeEnum = iota
	RelStockBk
	RelBkQuote
	RelStockLongHu
	RelLongHuDetail
)

var Neo4j *neo4j.Driver

func CreateDriver() (neo4j.Driver, error) {
	return neo4j.NewDriver(Config.Neo4jUrl, neo4j.BasicAuth(Config.Neo4jUsername, Config.Neo4jPassword, ""))
}

func CloseDriver(driver neo4j.Driver) error {
	return driver.Close()
}

func InitNeo4j() {
	defer log.Println("success init neo4j")
	neo4jDriver, err := CreateDriver()
	if err != nil {
		log.Println(err)
		panic("error connecting to neo4j")
	}
	Neo4j = &neo4jDriver
}

// cypher批量执行
func CypherBatchExec(cyphers []map[string]interface{}) error {
	driver := *Neo4j
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	transaction, err := session.BeginTransaction()
	if err != nil {
		log.Println("neo4j beginTransaction with error:", err)
		return err
	}
	for _, maps := range cyphers {
		cypher := maps["cypher"]
		param := maps["param"]
		_, err := transaction.Run(cypher.(string), param.(map[string]interface{}))
		if err != nil {
			log.Println("exec to neo4j with error:", err)
			transaction.Rollback()
			return err
		}
	}
	transaction.Commit()
	return err
}

// 执行cypher无返回
func CypherExec(cypher string, param map[string]interface{}) error {
	driver := *Neo4j
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	transaction, err := session.BeginTransaction()
	if err != nil {
		log.Println("neo4j beginTransaction with error:", err)
		return err
	}
	_, err = transaction.Run(cypher, param)
	if err != nil {
		log.Println("exec neo4j with error:", err)
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return err
}

// 执行cypher返回string数组
func CypherExecReturnStringList(cypher string, param map[string]interface{}) ([]string, error) {
	driver := *Neo4j
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	results, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string
		result, err := tx.Run(cypher, param)
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

// 执行cypher返回map数组
//func CypherExecReturnMapList(cypher string, param map[string]interface{}) (*[]map[string]interface{}, error) {
//	driver := *Neo4j
//	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
//	defer session.Close()
//	results, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
//		var list []map[string]interface{}
//		result, err := tx.Run(cypher, param)
//		if err != nil {
//			return nil, err
//		}
//		for result.Next() {
//			list = append(list, result.Record())
//		}
//		if err = result.Err(); err != nil {
//			return nil, err
//		}
//		return list, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	return &results.([]map[string]interface{}), nil
//}
