package service

import (
	dao "dss-data/dao/mysql"
	"dss-data/model"
	"dss-data/robot"
	"dss-data/tushare"
	"dss-data/util"
	"sync"
	"time"

	"dss-data/constant"
	"dss-data/thread"
)

func QueryNowDateTradeCal() string {
	return dao.QueryNowDateTradeCal()
}

// 更新交易日历
func UpdateTradeCals() {
	cals := robot.GetTradeCals()
	dao.MergeTradeCals(cals)
}

// 更新stockInfo&获取当日股票详情数据入库&获取当日龙虎榜信息
func CreateDailyData() {
	UpdateStockInfo()
	dao.InitCreateStockQuoteTable()
	symbols, _ := dao.GetAllTsCode()
	datas := []string{}
	for i, v := range *symbols {
		datas = append(datas, v)
		if len(datas) > 50 || i+1 == len(*symbols) {
			liveDatas, _ := GetLiveData(datas)
			quotes := liveToQuote(liveDatas)
			dao.InsertStockQuote(quotes)
			datas = []string{}
			time.Sleep(1000)
		}
	}
	GetLongHuDaily()
	UpdateAllBk()
}

// 更新板块数据
func UpdateAllBk() {
	bks, rels, quotes := robot.GetAllBkAndRelSymbol()
	dao.DeleteBk()
	dao.DeleteBkRelSymbol()
	dao.InsertBk(bks)
	dao.InsertBkRelSymbol(rels)
	dao.InsertBkQuote(quotes)
}

// 获取当日龙虎榜数据入库
func GetLongHuDaily() {
	longHu, longHuDetail := robot.GetLonghu()
	dao.InsertLongHu(longHu)
	dao.InsertLongHuDetail(longHuDetail)
}

// 更新stockInfo
func UpdateStockInfo() {
	stocks := robot.GetAllStock()
	dao.DeleteStockInfo()
	dao.MergeStockInfo(stocks)
}

// 获取从xx日开始至今的历史数据
func GetDailyData(startDate, endDate string) {
	allStock := robot.GetAllStock()
	tsCodes := []string{}
	for _, item := range *allStock {
		tsCodes = append(tsCodes, item.Symbol+"."+item.Exchange)
	}
	if "" == endDate {
		endDate = time.Now().Format(constant.TimeFormatA)
	}
	// 创建容量为 100 的任务池
	pool, err := thread.NewPool(100)
	if err != nil {
		panic(err)
	}
	wg := new(sync.WaitGroup)
	for _, tsCode := range tsCodes {
		data := tushare.GetStockQuoteData(map[string]interface{}{"ts_code": tsCode, "start_date": startDate, "end_date": endDate}, "daily")
		// 将任务放入任务池
		wg.Add(1)
		pool.Put(&thread.Task{
			Handler: func(v ...interface{}) {
				dao.InsertStockQuote(data)
				wg.Done()
			},
		})
	}
	wg.Wait()
	// 安全关闭任务池（保证已加入池中的任务被消费完）
	pool.Close()
}

// 查询最近连板股
func GetConStock() *map[int][]string {
	day := 1
	result := map[int][]string{}
	for {
		dates, _ := dao.GetXDayTradeDate(day)
		res, _ := dao.GetConStock(*dates)
		if len(*res) > 0 {
			dup := result[day-1]
			if len(dup) > 0 {
				same := []string{}
				for _, v := range *res {
					for _, d := range dup {
						if v == d {
							same = append(same, d)
						}
					}
				}
				newDup := []string{}
				m := make(map[string]int)
				for _, v := range same {
					m[v]++
				}
				for _, value := range dup {
					times := m[value]
					if times == 0 {
						newDup = append(newDup, value)
					}
				}
				result[day-1] = newDup
			}
			result[day] = *res
			day++
		} else {
			break
		}
	}
	return &result
}

func GetLimitUpXDayStock(day int) *[]string {
	dates, _ := dao.GetXDayTradeDate(day)
	res, _ := dao.GetLimitUpXDayStock(*dates)
	return res
}

func GetXDayUpYStock(day, percentage int) *[]string {
	dates, _ := dao.GetXDayTradeDate(day)
	res, _ := dao.GetXDayUpYStock(*dates, percentage)
	return res
}

// 实时行情转换为每日记录行情
func liveToQuote(liveDatas *[]model.LiveData) *[]model.StockQuote {
	stockQuotes := []model.StockQuote{}
	stStock := map[string]string{}
	sts, _ := dao.GetAllSTStock()
	for _, v := range *sts {
		stStock[v.TsCode] = v.Name
	}
	for _, data := range *liveDatas {
		stockQuote := model.StockQuote{}
		stockQuote.TsCode = data.Code
		stockQuote.TradeDate = util.Substr(data.Time, 0, 8)
		stockQuote.Open = data.Open
		stockQuote.High = data.Max
		stockQuote.Low = data.Min
		stockQuote.Close = data.Now
		stockQuote.PreClose = data.PreClose
		stockQuote.Change = data.Change
		stockQuote.PctChg = data.PctChg
		stockQuote.Vol = util.FloatDiv(data.Vol, 10000)
		stockQuote.Amount = data.Amount
		//计算涨停板
		stockQuote.LimitUp = util.CalLimitUp(data.Code, data.PreClose, data.Change, stStock)
		stockQuotes = append(stockQuotes, stockQuote)
	}
	return &stockQuotes
}
