package service

import (
	dao "dss-data/dao"
	"dss-data/robot"
	tushare "dss-data/tushare"
	"sync"
	"time"

	"dss-data/constant"
	"dss-data/thread"
)

func GetLongHuDaily() {
	longHu, longHuDetail := robot.GetLonghu()
	dao.InsertLongHu(longHu)
	dao.InsertLongHuDetail(longHuDetail)
}

// 更新stockInfo
func UpdateStockInfo() {
	stocks := robot.GetAllStock()
	dao.MergeStockInfo(stocks)
}

func GetDailyData(startDate string) {
	allStock := robot.GetAllStock()
	tsCodes := []string{}
	for _, item := range *allStock {
		tsCodes = append(tsCodes, item.Symbol+"."+item.Exchange)
	}
	// 创建容量为 100 的任务池
	pool, err := thread.NewPool(100)
	if err != nil {
		panic(err)
	}
	wg := new(sync.WaitGroup)
	for _, tsCode := range tsCodes {
		data := tushare.GetStockQuoteData(map[string]interface{}{"ts_code": tsCode, "start_date": startDate, "end_date": time.Now().Format(constant.TimeFormatA)}, "daily")
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

//查询最近连板股
func GetConStock() *map[int][]string {
	day := 2
	result := map[int][]string{}
	dates := []string{}
	c := 0
	a := 0
	for {
		for {
			date := dao.QueryXTradeDate(a)
			dates = append(dates, date)
			c++
			a++
			if c == day {
				break
			}
		}
		res, _ := dao.GetConStock(dates)
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
