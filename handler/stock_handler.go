package stock

import (
	router "dss-data/router"
	service "dss-data/service"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

const party = "/stock"

func init() {
	log.Printf("init router %v", party)
	router.RegisterHandler("Get", party, "/dailyData/{date:string}", createDailyData)
	router.RegisterHandler("Get", party, "/liveData/{symbols:string}", getLiveData)
}

func createDailyData(ctx iris.Context) {
	date := ctx.Params().GetString("date")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	go service.CreateDailyData(date, false, false)
	ctx.Write([]byte("开始更新(" + date + ")数据任务"))
}

//批量获取股票实时数据
func getLiveData(ctx iris.Context) {
	start := time.Now()
	symbols := ctx.Params().GetString("symbols")
	url := ctx.Request().URL.Path
	log.Printf("%v param(%v)", url, symbols)
	if symbols == "" {
		ctx.Write([]byte("请传入股票编码"))
	}
	result, err := service.GetLiveData(strings.Split(symbols, ","))
	if err != nil {
		log.Panicf("%v 获取异常,error: %v", url, err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		log.Panicf("%v json转换异常, result: %v", url, result)
	}
	ctx.Write(data)
	log.Printf("%v end,spend time %v", url, time.Since(start))
}
