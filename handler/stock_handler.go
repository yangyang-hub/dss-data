package stock

import (
	router "dss-base-data/router"
	service "dss-base-data/service"
	"log"
	"time"

	"github.com/kataras/iris/v12"
)

const party = "/stock"

func init() {
	log.Printf("init router %v", party)
	router.RegisterHandler("Get", party, "/init", initData)
	router.RegisterHandler("Get", party, "/dailyData/{date:string}", createDailyData)
}

func initData(ctx iris.Context) {
	params := ctx.Params()
	startDate := params.GetString("start_date")
	go service.CreateBaseData(startDate)
	ctx.Write([]byte("开始初始化数据任务"))
}

func createDailyData(ctx iris.Context) {
	date := ctx.Params().GetString("date")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	go service.CreateDailyData(date, true)
	ctx.Write([]byte("开始更新(" + date + ")数据任务"))
}
