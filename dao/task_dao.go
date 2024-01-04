package dao

import (
	db "dss-data/db"
	"log"
	"time"

	"dss-data/model"
)

//记录定时任务执行记录
func InsertTaskInfo(taskName string, date string, startTime time.Time) {
	currentTime := time.Now()
	if date == "" {
		date = currentTime.Format("20060102")
	}
	spendTime := time.Since(startTime).String()
	taskInfo := model.TaskInfo{TaskName: taskName, Date: date, SpendTime: spendTime}
	res := db.Mysql.Create(&taskInfo).Error
	if res != nil {
		log.Println(res.Error())
	}
}

//记录定时任务执行记录
func QueryTaskInfo(taskName string) *[]model.TaskInfo {
	taskInfos := []model.TaskInfo{}
	res := db.Mysql.Where("task_name = ?", taskName).Find(&taskInfos).Error
	if res != nil {
		log.Println(res.Error())
	}
	return &taskInfos
}

//查询股票行情数据定时任务执行记录
func QueryStockQuoteTaskInfo(dates []string) *[]model.TaskInfo {
	taskInfos := []model.TaskInfo{}
	res := db.Mysql.Where("task_name = ? AND date IN ?", "taskCreateDailyData", dates).Find(&taskInfos).Error
	if res != nil {
		log.Println(res.Error())
	}
	return &taskInfos
}
