package dao

import (
	"database/sql"
	configs "dss-data/configs"
	db "dss-data/db"
	"fmt"
	"log"
	"reflect"
	"time"
	"unsafe"

	"github.com/yangyang-hub/dss-common/constant"
	"github.com/yangyang-hub/dss-common/model"
)

//新增tushare stock_basic数据
func InsertStockBasic(stockInfos *[]model.StockInfo) bool {
	res := db.Mysql.CreateInBatches(stockInfos, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
		return false
	}
	return true
}

//新增tushare stock_company数据
func InsertStockCompany(stockCompanys *[]model.StockCompany) bool {
	res := db.Mysql.CreateInBatches(stockCompanys, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
		return false
	}
	return true
}

//更新（新增）tushare stock_basic数据
func MergeStockBasic(stockInfos *[]model.StockInfo) error {
	for _, stockInfo := range *stockInfos {
		var count int64
		db.Mysql.Table(stockInfo.TableName()).Where("ts_code = ?", stockInfo.TsCode).Count(&count)
		if count == 0 {
			results := db.Mysql.Create(stockInfo)
			if results.Error != nil {
				return results.Error
			}
		} else {
			results := db.Mysql.Model(stockInfo).Updates(stockInfo)
			if results.Error != nil {
				return results.Error
			}
		}
	}
	return nil
}

//更新（新增）tushare stock_company数据
func MergeStockCompany(stockCompanys *[]model.StockCompany) error {
	for _, stockCompany := range *stockCompanys {
		var count int64
		db.Mysql.Table(stockCompany.TableName()).Where("ts_code = ?", stockCompany.TsCode).Count(&count)
		if count == 0 {
			results := db.Mysql.Create(stockCompany)
			if results.Error != nil {
				return results.Error
			}
		} else {
			results := db.Mysql.Model(stockCompany).Updates(stockCompany)
			if results.Error != nil {
				return results.Error
			}
		}
	}
	return nil
}

//创建StockQuote数据库表
func InitCreateStockQuoteTable(startDate string) bool {
	start, _ := time.Parse("20060102", startDate)
	startYear := start.Year()
	nowYear := time.Now().Year()
	//初始化数据时从开始日期开始按年创建数据库行情表
	for i := startYear; i <= nowYear; i++ {
		tableName := fmt.Sprintf("%s%d", "stock_quote_", i)
		var count int64
		db.Mysql.Raw("SELECT * FROM information_schema.TABLES t, information_schema.SCHEMATA n \n"+
			"WHERE t.table_name = ? AND n.SCHEMA_NAME = ?", tableName, configs.Config.MysqlDatabase).Count(&count)
		if count == 0 {
			sql := "CREATE TABLE `" + tableName + "`  (\n" +
				"  `ts_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'TS代码',\n" +
				"  `type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'D：日线行情；W：周线行情；M：月线行情',\n" +
				"  `trade_date` varchar(8) NOT NULL COMMENT '交易日期',\n" +
				"  `open` float(30, 4) NULL DEFAULT NULL COMMENT '开盘价',\n" +
				"  `high` float(30, 4) NULL DEFAULT NULL COMMENT '最高价',\n" +
				"  `low` float(30, 4) NULL DEFAULT NULL COMMENT '最低价',\n" +
				"  `close` float(30, 4) NULL DEFAULT NULL COMMENT '收盘价',\n" +
				"  `pre_close` float(30, 4) NULL DEFAULT NULL COMMENT '昨收价(前复权)',\n" +
				"  `change` float(30, 4) NULL DEFAULT NULL COMMENT '涨跌额',\n" +
				"  `pct_chg` float(30, 4) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',\n" +
				"  `vol` float(30, 4) NULL DEFAULT NULL COMMENT '成交量(手)',\n" +
				"  `amount` float(30, 4) NULL DEFAULT NULL COMMENT '成交额(千元)',\n" +
				"  PRIMARY KEY (`ts_code`, `type`, `trade_date`) USING BTREE\n)" +
				" ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;"
			res := db.Mysql.Exec(sql)
			if res.Error != nil {
				log.Println(res.Error.Error())
			}
		}
	}
	return true
}

//新增StockQuote数据
func InsertStockQuote(stockQuotes *[]model.StockQuote) bool {
	quotes := map[string][]model.StockQuote{}
	for _, quote := range *stockQuotes {
		tableName := quote.TableName()
		temps := quotes[tableName]
		if len(temps) > 0 {
			temps = append(temps, quote)
		} else {
			temps = []model.StockQuote{quote}
		}
		quotes[tableName] = temps
	}
	for key, value := range quotes {
		res := db.Mysql.Table(key).CreateInBatches(value, constant.InsertBatchSize).Error
		if res != nil {
			log.Printf("method: %s, errpr: %s", "InsertStockQuote", res.Error())
		}
	}
	return true
}

//查询所有的股票编码数据（ts_code）
func GetAllTsCode() ([]string, error) {
	rows, _ := db.Mysql.Raw("SELECT ts_code FROM stock_info").Rows()
	res := scanRows2List(rows)
	return res, nil
}

func scanRows2List(rows *sql.Rows) []string {
	res := make([]string, 0)                                //  定义结果 map
	colTypes, _ := rows.ColumnTypes()                       // 列信息
	var rowParam = make([]interface{}, len(colTypes))       // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes))       // 接收数据一行列的数组
	rowValue[0] = reflect.New(colTypes[0].ScanType())       // 跟据数据库参数类型，创建默认值 和类型
	rowParam[0] = reflect.ValueOf(&rowValue[0]).Interface() // 跟据接收的数据的类型反射出值的地址
	// 遍历
	for rows.Next() {
		rows.Scan(rowParam...) // 赋值到 rowValue 中
		record := ""
		if rowValue[0] != nil {
			record = Byte2Str(rowValue[0].([]byte))
		}
		res = append(res, record)
	}
	return res
}

func scanRows2map(rows *sql.Rows) []map[string]string {
	res := make([]map[string]string, 0)               //  定义结果 map
	colTypes, _ := rows.ColumnTypes()                 // 列信息
	var rowParam = make([]interface{}, len(colTypes)) // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes)) // 接收数据一行列的数组

	for i, colType := range colTypes {
		rowValue[i] = reflect.New(colType.ScanType())           // 跟据数据库参数类型，创建默认值 和类型
		rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface() // 跟据接收的数据的类型反射出值的地址

	}
	// 遍历
	for rows.Next() {
		rows.Scan(rowParam...) // 赋值到 rowValue 中
		record := make(map[string]string)
		for i, colType := range colTypes {

			if rowValue[i] == nil {
				record[colType.Name()] = ""
			} else {
				record[colType.Name()] = Byte2Str(rowValue[i].([]byte))
			}
		}
		res = append(res, record)
	}
	return res
}

// []byte to string
func Byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
