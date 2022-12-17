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
	// 每天下午五点半执行 定时任务-插入每日行情数据
	s.Every(1).Days().At("17:30").Do(taskCreateDailyData)
	// 每天早上7点执行 定时任务-更新同花顺概念和行业的股票关联信息
	s.Every(1).Days().At("7:00").Do(taskUpdateThsGnAndHy)
	go s.StartBlocking()
}

//定时任务-插入每日行情数据
func taskCreateDailyData() {
	log.Printf("Start Scheduler CreateDailyData date(%v)", time.Now().Format("20060102"))
	go service.CreateDailyData("", true)
}

//定时任务-更新同花顺概念和行业的股票关联信息
func taskUpdateThsGnAndHy() {
	log.Printf("Start Scheduler UpdateThsGnAndHy date(%v)", time.Now().Format("20060102"))
	go service.UpdateThsGnAndHy()
}
