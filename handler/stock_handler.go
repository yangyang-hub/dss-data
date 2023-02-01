package stock

import (
	router "dss-data/router"
	service "dss-data/service"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const party = "/stock"

func init() {
	log.Printf("init router %v", party)
	router.RegisterHandler("Get", party, "/dailyData/:date", createDailyData)
	router.RegisterHandler("Get", party, "/liveData/:symbols", getLiveData)
}

func createDailyData(ctx *gin.Context) {
	date := ctx.Param("data")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	go service.CreateDailyData(date, false, false)
	ctx.JSON(200, "开始更新("+date+")数据任务")
}

//批量获取股票实时数据
func getLiveData(ctx *gin.Context) {
	url := ctx.Request.URL.Path
	symbols := ctx.Param("symbols")
	if symbols == "" {
		ctx.JSON(200, "请传入股票编码")
	}
	result, err := service.GetLiveData(strings.Split(symbols, ","))
	if err != nil {
		log.Panicf("%v 获取异常,error: %v", url, err)
	}
	ctx.JSON(200, &result)
}
