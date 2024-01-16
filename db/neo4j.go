package db

import (
	"context"
	. "dss-data/configs"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

// 节点标签枚举
type NodeEnum int

func (n NodeEnum) String() string {
	return [...]string{"NT_STOCK_INFO", "NT_STOCK_QUOTE", "NT_BK", "NT_BK_QUOTE", "NT_LONG_HU", "NT_LONG_HU_DETAIL", "NT_STOCK_CON"}[n]
}

const (
	StockInfo NodeEnum = iota
	StockQuote
	Bk
	BkQuote
	LongHu
	LongHuDetail
	StockCon
)

// 边标签枚举
type EdgeEnum int

func (e EdgeEnum) String() string {
	return [...]string{"RT_STOCK_QUOTE", "RT_STOCK_BK", "RT_BK_QUOTE", "RT_STOCK_LONG_HU", "RT_LONG_HU_DETAIL", "RT_STOCK_CON"}[e]
}

const (
	RelStockQuote EdgeEnum = iota
	RelStockBk
	RelBkQuote
	RelStockLongHu
	RelLongHuDetail
	RelStockCon
)

var (
	Neo4j *neo4j.DriverWithContext
	ctx   context.Context
)

func CreateDriver() (neo4j.DriverWithContext, error) {
	ctx = context.Background()
	driver, err := neo4j.NewDriverWithContext(
		Config.Neo4jUrl,
		neo4j.BasicAuth(Config.Neo4jUsername, Config.Neo4jPassword, ""))
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
	return driver, nil
}

func CloseDriver(driver neo4j.DriverWithContext) error {
	return driver.Close(ctx)
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
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: Config.Neo4jDatabase})
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			for _, maps := range cyphers {
				cypher := maps["cypher"]
				param := maps["param"]
				_, err := tx.Run(ctx, cypher.(string), param.(map[string]interface{}))
				if err != nil {
					log.Println("exec to neo4j with error:", err)
					log.Println("exec to neo4j with error cypher:", cypher)
					return err, nil
				}
			}
			return nil, nil
		})
	return err
}

// 执行cypher无返回
func CypherExec(cypher string, param map[string]interface{}) error {
	driver := *Neo4j
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: Config.Neo4jDatabase})
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, cypher, param)
			if err != nil {
				log.Println("exec to neo4j with error:", err)
				log.Println("exec to neo4j with error cypher:", cypher)
				return err, nil
			}
			return result, nil
		})
	return err
}

// 执行cypher返回string数组
func CypherExecReturnStringList(cypher string, param map[string]interface{}) (*[]string, error) {
	var list []string
	driver := *Neo4j
	result, _ := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		param, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(Config.Neo4jDatabase))
	for _, record := range result.Records {
		list = append(list, record.Values[0].(string))
	}
	return &list, nil
}

// 执行cypher返回map数组
func CypherExecReturnMapList(cypher string, param map[string]interface{}) (*[]map[string]interface{}, error) {
	var list []map[string]interface{}
	driver := *Neo4j
	result, _ := neo4j.ExecuteQuery(ctx, driver,
		cypher,
		param, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(Config.Neo4jDatabase))
	for _, record := range result.Records {
		list = append(list, record.AsMap())
	}
	return &list, nil
}
