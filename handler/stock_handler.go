package handler

import (
	router "dss-data/router"
	service "dss-data/service"
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
}

// 更新股票信息
func updateStockInfo(ctx *gin.Context) {
	service.UpdateStockInfo()
	ctx.JSON(200, "更新股票信息完成")
}

// 从**日开始获取每日数据
func getDailyData(ctx *gin.Context) {
	date := ctx.Query("data")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	go service.GetDailyData(date)
	ctx.JSON(200, "开始更新("+date+")至今的数据任务")
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
