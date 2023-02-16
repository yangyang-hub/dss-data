package schedule

import (
	service "dss-data/service"
	"dss-data/tushare"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/yangyang-hub/dss-common/constant"
)

func InitScheduler() {
	defer log.Println("success init scheduler")
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)
	// 每天下午六点执行 定时任务-插入每日行情数据
	s.Every(1).Days().At("18:00").Do(taskCreateDailyData)
	// 每天上午1.30点执行 定时任务-刷新同花顺概念
	s.Every(1).Days().At("01:30").Do(taskRefreshThsGn)
	go s.StartBlocking()
}

//定时任务-插入每日行情数据
func taskCreateDailyData() {
	log.Printf("Start Scheduler CreateDailyData date(%v)", time.Now().Format(constant.TimeFormatA))
	service.CreateDailyData("")
}

//定时任务-刷新同花顺概念
func taskRefreshThsGn() {
	trade_date := time.Now().Format(constant.TimeFormatA)
	//查询是否为交易日
	tradeCals := tushare.GetTradeCal(trade_date, trade_date)
	if len(*tradeCals) < 1 {
		log.Printf("(%v)为非交易日,结束任务 RefreshThsGn", trade_date)
		return
	}
	log.Printf("Start Scheduler RefreshThsGn date(%v)", trade_date)
	service.RobotAllThsGnBySymbols()
}
