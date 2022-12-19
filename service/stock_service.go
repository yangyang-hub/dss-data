package service

import (
	dao "dss-data/dao"
	robot "dss-data/robot"
	tushare "dss-data/tushare"
	"log"
	"sync"
	"time"

	"github.com/yangyang-hub/dss-common/constant"
	"github.com/yangyang-hub/dss-common/model"
	"github.com/yangyang-hub/dss-common/thread"
)

//初始化基本数据（股票基本信息、上市公司基本信息、日线数据、周线数据、月线数据）
func InitData() {
	initDataInfos := dao.QueryTaskInfo("InitData")
	if len(*initDataInfos) == 0 {
		//初始化基本数据
		CreateBaseData("")
		return
	}
	log.Printf("InitData already completed...")
	//检查以前的行情数据是否有遗漏 最近一个月为核查标准
	//获取近一个月的交易日历
	startDate := time.Now().AddDate(0, -1, 0).Format("20060102")
	tradeCals := tushare.GetTradeCal(startDate, time.Now().Format("20060102"))
	tradeDates := []string{}
	for _, item := range *tradeCals {
		tradeDates = append(tradeDates, item.CalDate)
	}
	taskInfos := dao.QueryStockQuoteTaskInfo(tradeDates)
	//丢失的行情数据
	missDate := []string{}
	initData := (*initDataInfos)[0].Date
	if len(*taskInfos) == 0 {
		for _, date := range tradeDates {
			if initData >= date {
				continue
			}
			missDate = append(missDate, date)
		}
	} else {
		for _, taskInfo := range *taskInfos {
			if initData >= taskInfo.Date {
				continue
			}
			flag := false
			for _, date := range tradeDates {
				if date == taskInfo.Date {
					flag = true
					return
				}
			}
			if !flag {
				missDate = append(missDate, taskInfo.Date)
			}
		}
	}
	if len(missDate) == 0 {
		log.Printf("not find miss_data...")
	} else {
		log.Printf("start run miss_data...")
		for _, data := range missDate {
			CreateDailyData(data, false)
		}
	}
	log.Printf("init stock_service end...")
}

//初始化基本数据（股票基本信息、上市公司基本信息、日线数据、周线数据、月线数据）
func CreateBaseData(startDate string) {
	startTime := time.Now()
	//记录定时任务日志
	defer dao.InsertTaskInfo("InitData", "", startTime)
	log.Printf("InitData start...")
	start := time.Now()
	for _, exchange := range constant.ExchangeConst.List() {
		//基础数据
		stockBasics := tushare.GetStockBasicData(map[string]interface{}{"exchange": exchange})
		dao.InsertStockBasic(stockBasics)
		//上市公司基本信息
		stockCompanys := tushare.GetStockCompanyData(map[string]interface{}{"exchange": exchange})
		dao.InsertStockCompany(stockCompanys)
	}
	tsCodes, err := dao.GetAllTsCode()
	if err != nil {
		log.Println("查询ts_code失败，结束初始化进程")
		return
	}
	//初始化同花顺概念数据
	InitThsGn()
	//初始化同花顺行业数据
	InitThsHy()
	//获取当日的同花顺概念及行业行情
	thsHyQuotes := robot.GetAllThsHyQuote(dao.QueryAllThsHy())
	dao.InsertThsHyQuote(thsHyQuotes)
	thsGnQuotes := robot.GetAllThsGnQuote(dao.QueryAllThsGn())
	dao.InsertThsGnQuote(thsGnQuotes)
	//初始化日线、周线、月线数据 默认从2010年1月1日开始
	if startDate == "" {
		startDate = "20100101"
	}
	log.Printf("start init quote data ...")
	//初始化股票行情数据表
	dao.InitCreateStockQuoteTable(startDate)
	// 创建容量为 100 的任务池
	pool, err := thread.NewPool(100)
	if err != nil {
		panic(err)
	}
	wg := new(sync.WaitGroup)
	for _, tsCode := range tsCodes {
		for _, quote := range constant.QuoteConst.List() {
			data := tushare.GetStockQuoteData(map[string]interface{}{"ts_code": tsCode, "start_date": startDate, "end_date": time.Now().Format("20060102")}, quote)
			// 将任务放入任务池
			wg.Add(1)
			pool.Put(&thread.Task{
				Handler: func(v ...interface{}) {
					wg.Done()
					dao.InsertStockQuote(data)
				},
			})
		}
	}
	wg.Wait()
	// 安全关闭任务池（保证已加入池中的任务被消费完）
	pool.Close()
	log.Printf("InitData end,spend time %v", time.Since(start))
}

/**插入日数据
trade_date:日期
includeThs:是否拉取同花顺行情数据
*/
func CreateDailyData(trade_date string, includeThs bool) {
	//trade_date为空则默认查询当日数据
	if trade_date == "" {
		trade_date = time.Now().Format("20060102")
	}
	startTime := time.Now()
	defer dao.InsertTaskInfo("taskCreateDailyData", trade_date, startTime)
	start := time.Now()
	log.Printf("CreateDailyData date(%v) start... ", trade_date)
	log.Println("更新基本信息")
	//更新基本信息
	for _, exchange := range constant.ExchangeConst.List() {
		//基础数据
		stockBasics := tushare.GetStockBasicData(map[string]interface{}{"exchange": exchange})
		dao.MergeStockBasic(stockBasics)
		//上市公司基本信息
		stockCompanys := tushare.GetStockCompanyData(map[string]interface{}{"exchange": exchange})
		dao.MergeStockCompany(stockCompanys)
	}
	//初始化股票行情数据表
	dao.InitCreateStockQuoteTable(trade_date)
	tsCodes, err := dao.GetAllTsCode()
	if err != nil {
		log.Println("查询ts_code失败,结束进程")
		return
	}
	//获取当日的同花顺概念及行业行情
	if includeThs {
		getThsHyQuote(dao.QueryAllThsHy())
		getThsGnQuote(dao.QueryAllThsGn())
	}
	log.Println("更新k线信息")
	// 创建容量为 100 的任务池
	pool, err := thread.NewPool(100)
	if err != nil {
		panic(err)
	}
	wg := new(sync.WaitGroup)
	for _, tsCode := range tsCodes {
		for _, quote := range constant.QuoteConst.List() {
			data := tushare.GetStockQuoteData(map[string]interface{}{"ts_code": tsCode, "trade_date": trade_date}, quote)
			// 将任务放入任务池
			if len(*data) > 0 {
				wg.Add(1)
				pool.Put(&thread.Task{
					Handler: func(v ...interface{}) {
						wg.Done()
						dao.InsertStockQuote(data)
					},
				})
			}
		}
	}
	wg.Wait()
	// 安全关闭任务池（保证已加入池中的任务被消费完）
	pool.Close()
	log.Printf("CreateDailyData end,spend time %v", time.Since(start))
}

//更新同花顺概念和行业的股票关联信息
func UpdateThsGnAndHy() {
	trade_date := time.Now().Format("20060102")
	startTime := time.Now()
	defer dao.InsertTaskInfo("taskUpdateThsGnAndHy", trade_date, startTime)
	start := time.Now()
	log.Printf("UpdateThsGnAndHy date(%v) start... ", trade_date)
	thsGns := robot.GetAllThsGn()
	if len(*thsGns) > 0 {
		dao.DeleteAllThsGn()
		dao.InsertThsGn(thsGns)
	}
	AllThsGns := dao.QueryAllThsGn()
	newThsGns := model.Minus(*AllThsGns, *thsGns)
	//获取当日的同花顺概念及行业行情
	if len(newThsGns) > 0 {
		getThsGnQuote(&newThsGns)
	}
	thsGnRelSymbols := robot.GetAllThsGnRelSymbol(thsGns)
	if len(*thsGnRelSymbols) > 0 {
		dao.DeleteAllThsGnRelSymbol()
		dao.InsertThsGnRelSymbol(thsGnRelSymbols)
	}
	thsHys := robot.GetAllThsHy()
	if len(*thsHys) > 0 {
		dao.DeleteAllThsHy()
		dao.InsertThsHy(thsHys)
	}
	AllThsHys := dao.QueryAllThsHy()
	newThsHys := model.Minus(*AllThsHys, *thsHys)
	//获取当日的同花顺概念及行业行情
	if len(newThsHys) > 0 {
		getThsHyQuote(&newThsHys)
	}
	thsHyRelSymbols := robot.GetAllThsHyRelSymbol(thsHys)
	if len(*thsHyRelSymbols) > 0 {
		dao.DeleteAllThsHyRelSymbol()
		dao.InsertThsHyRelSymbol(thsHyRelSymbols)
	}
	log.Printf("UpdateThsGnAndHy end,spend time %v", time.Since(start))
}

//获取当日的同花顺概念详情数据
func getThsGnQuote(thsGns *[]model.ThsGn) {
	thsGnQuotes := robot.GetAllThsGnQuote(thsGns)
	dao.InsertThsGnQuote(thsGnQuotes)
}

//获取当日的同花顺行业详情数据
func getThsHyQuote(thsHys *[]model.ThsHy) {
	thsHyQuotes := robot.GetAllThsHyQuote(thsHys)
	dao.InsertThsHyQuote(thsHyQuotes)
}

//初始化同花顺概念
func InitThsGn() {
	thsGns := robot.GetAllThsGn()
	dao.InsertThsGn(thsGns)
	thsGnRelSymbols := robot.GetAllThsGnRelSymbol(thsGns)
	dao.InsertThsGnRelSymbol(thsGnRelSymbols)
}

//初始化同花顺行业
func InitThsHy() {
	thsHys := robot.GetAllThsHy()
	dao.InsertThsHy(thsHys)
	thsHyRelSymbols := robot.GetAllThsHyRelSymbol(thsHys)
	dao.InsertThsHyRelSymbol(thsHyRelSymbols)
}
