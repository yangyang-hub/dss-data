package db

import (
	. "dss-data/configs"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

// 节点标签枚举
type NodeEnum int

func (n NodeEnum) String() string {
	return [...]string{"NT_STOCK_INFO", "NT_STOCK_QUOTE", "NT_BK", "NT_BK_QUOTE"}[n]
}

const (
	StockInfo NodeEnum = iota
	StockQuote
	Bk
	BkQuote
)

// 边标签枚举
type EdgeEnum int

func (e EdgeEnum) String() string {
	return [...]string{"RT_STOCK_QUOTE", "RT_STOCK_BK", "RT_BK_QUOTE"}[e]
}

const (
	RelStockQuote EdgeEnum = iota
	RelStockBk
	RelBkQuote
)

type Node[T any] struct {
	Label      []string `json:"label"`
	Id         string   `json:"id"`
	Properties T        `json:"properties"`
}

type Edge[T any] struct {
	Label      string `json:"label"`
	Id         string `json:"id"`
	Properties T      `json:"properties"`
}

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

// cypher批量写入
func CypherWrite(cyphers []map[string]interface{}) error {
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
			log.Println("wirte to neo4j with error:", err)
			transaction.Rollback()
			return err
		}
	}
	transaction.Commit()
	return err
}
