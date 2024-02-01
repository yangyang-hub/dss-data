package mysql

import (
	"database/sql"
	configs "dss-data/configs"
	"dss-data/db"
	"dss-data/model"
	"dss-data/util"
	"fmt"
	"log"
	"reflect"
	"time"
	"unsafe"

	"dss-data/constant"
)

// 查询前第X个交易日
func GetXDayUpYStock(dates []string, percentage int) (*[]string, error) {
	rows, _ := db.Mysql.Raw("SELECT ts_code FROM stock_quote_2023 WHERE trade_date in ? GROUP BY ts_code HAVING SUM(pct_chg) > ?", dates, percentage).Rows()
	res := scanRows2List(rows)
	return &res, nil
}

// 查询某天是否为交易日
func QueryTradeDate(date string) bool {
	var count int64
	db.Mysql.Table("stock_quote_2023").Where("trade_date = ?", date).Count(&count)
	if count == 0 {
		return false
	} else {
		return true
	}
}

// 查询近X交易日
func GetXDayTradeDate(day int) (*[]string, error) {
	nowYear := time.Now().Format("2006")
	rows, _ := db.Mysql.Raw("SELECT DISTINCT trade_date FROM stock_quote_"+nowYear+" ORDER BY trade_date DESC LIMIT 0,?", day).Rows()
	res := scanRows2List(rows)
	if len(res) >= day {
		return &res, nil
	}
	lastYear := time.Now().AddDate(-1, 0, 0).Format("2006")
	lastday := day - len(res)
	rowslast, _ := db.Mysql.Raw("SELECT DISTINCT trade_date FROM stock_quote_"+lastYear+" ORDER BY trade_date DESC LIMIT 0,?", lastday).Rows()
	reslast := scanRows2List(rowslast)
	res = append(res, reslast...)
	return &res, nil
}

// 查询近X日有涨停的股
func GetLimitUpXDayStock(dates []string) (*[]string, error) {
	rows, _ := db.Mysql.Raw("SELECT DISTINCT ts_code FROM stock_quote_2023 WHERE trade_date IN ? AND limit_up = '1'", dates).Rows()
	res := scanRows2List(rows)
	return &res, nil
}

// 查询近X日连板股
func GetConStock(dates []string) (*[]string, error) {
	nowYear := time.Now().Format("2006")
	rows, _ := db.Mysql.Raw("SELECT ts_code FROM stock_quote_"+nowYear+" WHERE trade_date IN ? AND limit_up = '1' GROUP BY ts_code HAVING COUNT(1)  = ?", dates, len(dates)).Rows()
	res := scanRows2List(rows)
	if len(res) > 0 {
		return &res, nil
	}
	lastYear := time.Now().AddDate(-1, 0, 0).Format("2006")
	rows, _ = db.Mysql.Raw("SELECT ts_code FROM (SELECT * FROM stock_quote_"+lastYear+" union all SELECT * FROM stock_quote_"+nowYear+") s WHERE trade_date IN ? AND limit_up = '1' GROUP BY ts_code HAVING COUNT(1)  = ?", dates, len(dates)).Rows()
	res = scanRows2List(rows)
	return &res, nil
}

// 查询所有ST股票
func GetAllSTStock() (*[]model.StockInfo, error) {
	stockInfos := []model.StockInfo{}
	res := db.Mysql.Where("name like '%ST%'").Find(&stockInfos).Error
	return &stockInfos, res
}

// 查询所有板块
func GetAllBk() (*[]model.Bk, error) {
	data := []model.Bk{}
	res := db.Mysql.Find(&data).Error
	return &data, res
}

// 查询所有板块编码
func GetAllBkCode() (*[]string, error) {
	rows, _ := db.Mysql.Raw("SELECT code FROM bk").Rows()
	res := scanRows2List(rows)
	return &res, nil
}

// 查询所有板块-股票关联
func GetAllBkRelSymbol() (*[]model.BkRelSymbol, error) {
	data := []model.BkRelSymbol{}
	res := db.Mysql.Find(&data).Error
	return &data, res
}

// 查询指定日期的板块行情
func GetBkQuoteByDate(date string) (*[]model.BkQuote, error) {
	data := []model.BkQuote{}
	res := db.Mysql.Where("trade_date = ?", date).Find(&data).Error
	return &data, res
}

// 查询指定日期的龙虎榜
func GetLongHuByDate(date string) (*[]model.LongHu, error) {
	data := []model.LongHu{}
	res := db.Mysql.Where("trade_date = ?", date).Find(&data).Error
	return &data, res
}

// 查询指定日期的龙虎榜详情
func GetLongHuDetailByDate(date string) (*[]model.LongHuDetail, error) {
	data := []model.LongHuDetail{}
	res := db.Mysql.Where("long_hu_id in (SELECT id FROM long_hu WHERE trade_date = ?)", date).Find(&data).Error
	return &data, res
}

// 查询指定日期的股票行情
func GetStockQuoteByDate(date string) (*[]model.StockQuote, error) {
	data := []model.StockQuote{}
	tableName := "stock_quote_" + util.Substr(date, 0, 4)
	res := db.Mysql.Table(tableName).Where("trade_date = ?", date).Find(&data).Error
	return &data, res
}

// 查询所有股票信息
func GetAllStockInfo() (*[]model.StockInfo, error) {
	stockInfos := []model.StockInfo{}
	res := db.Mysql.Find(&stockInfos).Error
	return &stockInfos, res
}

// 查询所有的股票symbol
func GetAllSymbol() (*[]string, error) {
	rows, _ := db.Mysql.Raw("SELECT symbol FROM stock_info").Rows()
	res := scanRows2List(rows)
	return &res, nil
}

// 查询所有的股票编码数据（ts_code）
func GetAllTsCode() (*[]string, error) {
	rows, _ := db.Mysql.Raw("SELECT ts_code FROM stock_info").Rows()
	res := scanRows2List(rows)
	return &res, nil
}

// 删除板块数据
func DeleteBk() {
	res := db.Mysql.Exec("DELETE FROM bk").Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 删除板块-股票关联数据
func DeleteBkRelSymbol() {
	res := db.Mysql.Exec("DELETE FROM bk_rel_symbol").Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 新增板块数据
func InsertBk(datas *[]model.Bk) {
	res := db.Mysql.CreateInBatches(datas, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 新增板块-股票关联数据
func InsertBkRelSymbol(datas *[]model.BkRelSymbol) {
	res := db.Mysql.CreateInBatches(datas, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 新增板块行情数据
func InsertBkQuote(datas *[]model.BkQuote) {
	res := db.Mysql.CreateInBatches(datas, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 新增龙虎榜数据
func InsertLongHu(longHus *[]model.LongHu) {
	res := db.Mysql.CreateInBatches(longHus, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 新增龙虎榜数据详情
func InsertLongHuDetail(longHuDetails *[]model.LongHuDetail) {
	res := db.Mysql.CreateInBatches(longHuDetails, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 新增 stock_info数据
func InsertStockInfo(stockInfos *model.StockInfo) bool {
	res := db.Mysql.Create(stockInfos).Error
	if res != nil {
		log.Println(res.Error())
		return false
	}
	return true
}

// 删除stock_info数据
func DeleteStockInfo() {
	res := db.Mysql.Exec("DELETE FROM stock_info").Error
	if res != nil {
		log.Println(res.Error())
	}
}

// 新增tushare stock_company数据
func InsertStockCompany(stockCompanys *[]model.StockCompany) bool {
	res := db.Mysql.CreateInBatches(stockCompanys, constant.InsertBatchSize).Error
	if res != nil {
		log.Println(res.Error())
		return false
	}
	return true
}

// 更新（新增） stock_info数据
func MergeStockInfo(stockInfos *[]model.StockInfo) error {
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

// 更新（新增）tushare stock_company数据
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

// 创建StockQuote数据库表
func InitCreateStockQuoteTable() bool {
	nowYear := time.Now().Year()
	tableName := fmt.Sprintf("%s%d", "stock_quote_", nowYear)
	var count int64
	db.Mysql.Raw("SELECT COUNT(1) FROM information_schema.TABLES t \n"+
		"WHERE t.TABLE_NAME = ? AND t.TABLE_SCHEMA = ?", tableName, configs.Config.MysqlDatabase).Count(&count)
	if count == 0 {
		sql := "CREATE TABLE `" + tableName + "`  (\n" +
			"  `ts_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'TS代码',\n" +
			"  `trade_date` varchar(8) NOT NULL COMMENT '交易日期',\n" +
			"  `open` float(30, 2) NULL DEFAULT NULL COMMENT '开盘价',\n" +
			"  `high` float(30, 2) NULL DEFAULT NULL COMMENT '最高价',\n" +
			"  `low` float(30, 2) NULL DEFAULT NULL COMMENT '最低价',\n" +
			"  `close` float(30, 2) NULL DEFAULT NULL COMMENT '收盘价',\n" +
			"  `pre_close` float(30, 2) NULL DEFAULT NULL COMMENT '昨收价(前复权)',\n" +
			"  `change` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌额',\n" +
			"  `pct_chg` float(30, 2) NULL DEFAULT NULL COMMENT '涨跌幅(未复权)',\n" +
			"  `vol` float(30, 2) NULL DEFAULT NULL COMMENT '成交量(万手)',\n" +
			"  `amount` float(30, 2) NULL DEFAULT NULL COMMENT '成交额(万元)',\n" +
			"  `limit_up` tinyint(1) NULL DEFAULT NULL COMMENT '涨停板',\n" +
			"  PRIMARY KEY (`ts_code`, `trade_date`) USING BTREE,\n" +
			"  INDEX `trade_date_index`(`trade_date`) USING BTREE,\n" +
			"  INDEX `ts_code_index`(`ts_code`) USING BTREE\n)" +
			" ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '股票行情' ROW_FORMAT = Dynamic;"
		res := db.Mysql.Exec(sql)
		if res.Error != nil {
			log.Println(res.Error.Error())
		}
	}
	return true
}

// 新增StockQuote数据
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
