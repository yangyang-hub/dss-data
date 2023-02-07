package handler

import (
	router "dss-data/router"
	"dss-data/service"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	party := "/ths"
	log.Printf("init router %v", party)
	router.RegisterHandler("Post", party, "/getGnBySymbols", getGnBySymbols)
	router.RegisterHandler("Post", party, "/getSymbolsByGn", getSymbolsByGn)
	router.RegisterHandler("Get", party, "/refreshThsGn", refreshThsGn)
}

// 根据股票编码查询所属概念
func getGnBySymbols(ctx *gin.Context) {
	var params map[string]string
	b, _ := ctx.GetRawData()
	json.Unmarshal(b, &params)
	symbols := params["symbols"]
	if symbols == "" {
		ctx.JSON(200, "请传入股票编码")
	}
	result := service.GetThsGnBySymbols(symbols)
	ctx.JSON(200, result)
}

// 根据概念名称查询所有股票
func getSymbolsByGn(ctx *gin.Context) {
	var params map[string]string
	b, _ := ctx.GetRawData()
	json.Unmarshal(b, &params)
	gns := params["gns"]
	if gns == "" {
		ctx.JSON(200, "请传入概念名称")
	}
	result := service.GetSymbolsByThsGn(gns)
	ctx.JSON(200, result)
}

// 刷新同花顺概念
func refreshThsGn(ctx *gin.Context) {
	go service.RobotAllThsGnBySymbols()
	ctx.JSON(200, "开始刷新同花顺概念")
}
