package handler

import (
	"dss-data/robot"
	"dss-data/router"
	"dss-data/service"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	party := "/stock"
	log.Printf("init router %v", party)
	router.RegisterHandler("Get", party, "/dailyData", getDailyData)
	router.RegisterHandler("Post", party, "/getLiveData", getLiveData)
	router.RegisterHandler("Get", party, "/updateStockInfo", updateStockInfo)
	router.RegisterHandler("Get", party, "/getLongHuDaily", getLongHuDaily)
	router.RegisterHandler("Get", party, "/updateAllBk", updateAllBk)
	router.RegisterHandler("Get", party, "/createDailyData", createDailyData)
	router.RegisterHandler("Get", party, "/updateGraph", updateGraph)
	router.RegisterHandler("Get", party, "/getConStock", getConStock)
	router.RegisterHandler("Get", party, "/getTradeCal", getTradeCal)
}

// 查询是否为交易日
func getTradeCal(ctx *gin.Context) {
	tradeCals := robot.GetTradeCal()
	ctx.JSON(200, tradeCals)
}

// 获取连板股票
func getConStock(ctx *gin.Context) {
	res := service.GetConStock()
	ctx.JSON(200, res)
}

// 更新图
func updateGraph(ctx *gin.Context) {
	date := ctx.Query("date")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	service.UpdateGraph(date)
	ctx.JSON(200, "更新("+date+")的图形任务完成")
}

// 更新股票信息
func updateStockInfo(ctx *gin.Context) {
	service.UpdateStockInfo()
	ctx.JSON(200, "更新股票信息完成")
}

// 获取当日行情数据
func createDailyData(ctx *gin.Context) {
	service.CreateDailyData()
	ctx.JSON(200, "获取当日行情数据")
}

// 更新板块信息
func updateAllBk(ctx *gin.Context) {
	service.UpdateAllBk()
	ctx.JSON(200, "更新板块信息完成")
}

// 从**日开始获取每日数据
func getDailyData(ctx *gin.Context) {
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")
	if startDate == "" {
		startDate = time.Now().Format("20060102")
	}
	if endDate == "" {
		endDate = time.Now().Format("20060102")
	}
	go service.GetDailyData(startDate, endDate)
	ctx.JSON(200, "开始更新("+startDate+"-"+endDate+")的数据任务")
}

// 获取当日龙虎榜
func getLongHuDaily(ctx *gin.Context) {
	go service.GetLongHuDaily()
	ctx.JSON(200, "获取当日龙虎榜完成")
}

// 批量获取股票实时数据
func getLiveData(ctx *gin.Context) {
	var params map[string]string
	b, _ := ctx.GetRawData()
	json.Unmarshal(b, &params)
	symbols := params["symbols"]
	if symbols == "" {
		ctx.JSON(200, "请传入股票编码")
	}
	result, err := service.GetLiveData(strings.Split(symbols, ","))
	if err != nil {
		log.Panicf("getLiveData,error: %v", err)
	}
	ctx.JSON(200, &result)
}
