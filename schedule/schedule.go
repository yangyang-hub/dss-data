package schedule

import (
	service "dss-data/service"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func InitScheduler() {
	defer log.Println("success init scheduler")
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)
	// 每天下午六点执行 定时任务-插入每日行情数据
	s.Every(1).Days().At("18:00").Do(taskCreateDailyData)
	// 每天上午五点执行 定时任务-刷新同花顺概念
	s.Every(1).Days().At("05:00").Do(taskRefreshThsGn)
	go s.StartBlocking()
}

//定时任务-插入每日行情数据
func taskCreateDailyData() {
	log.Printf("Start Scheduler CreateDailyData date(%v)", time.Now().Format("20060102"))
	service.CreateDailyData("")
}

//定时任务-刷新同花顺概念
func taskRefreshThsGn() {
	log.Printf("Start Scheduler RefreshThsGn date(%v)", time.Now().Format("20060102"))
	service.RobotAllThsGnBySymbols()
}
