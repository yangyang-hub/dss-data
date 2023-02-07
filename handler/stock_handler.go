package handler

import (
	dao "dss-data/dao"
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
	router.RegisterHandler("Get", party, "/dailyData/:date", createDailyData)
	router.RegisterHandler("Post", party, "/getLiveData", getLiveData)
	router.RegisterHandler("Get", party, "/getAllStockInfo", getAllStockInfo)
}

//获取所有股票信息
func getAllStockInfo(ctx *gin.Context) {
	result, err := dao.GetAllStockInfo()
	if err != nil {
		log.Panicf("getAllStockInfo,error: %v", err)
	}
	ctx.JSON(200, &result)
}

//批量获取股票实时数据
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

func createDailyData(ctx *gin.Context) {
	date := ctx.Param("data")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	go service.CreateDailyData(date)
	ctx.JSON(200, "开始更新("+date+")数据任务")
}
