package schedule

import (
	"dss-data/dao/mysql"
	"dss-data/service"
	"log"
	"time"

	"dss-data/constant"

	"github.com/go-co-op/gocron"
)

func InitScheduler() {
	defer log.Println("success init scheduler")
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)
	// 每天下午六点18执行 定时任务-插入每日行情数据
	s.Every(1).Days().At("18:18").Do(taskCreateDailyData)
	s.Every(1).Days().At("08:08").Do(taskUpdateTradeCals)
	go s.StartBlocking()
}

func taskUpdateTradeCals() {
	// 更新交易日信息
	service.UpdateTradeCals()
}

// 定时任务-插入每日行情数据
func taskCreateDailyData() {
	trade_date := time.Now().Format(constant.TimeFormatA)
	startTime := time.Now()
	defer mysql.InsertTaskInfo("CreateDailyData", trade_date, startTime)
	// 更新交易日信息
	service.UpdateTradeCals()
	// 查询是否为交易日
	tradeCals := service.QueryNowDateTradeCal()
	if tradeCals != "1" {
		log.Printf("(%v)为非交易日,结束任务", trade_date)
		return
	}
	log.Printf("Start Scheduler CreateDailyData date(%v)", time.Now().Format(constant.TimeFormatA))
	service.CreateDailyData()
	log.Printf("CreateDailyData end,spend time %v,start create graph", time.Since(startTime))
	service.UpdateGraph(trade_date)
	log.Printf("Scheduler CreateDailyData end,spend time %v", time.Since(startTime))
}
