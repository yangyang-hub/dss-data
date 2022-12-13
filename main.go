package main

import (
	configs "dss-base-data/configs"
	db "dss-base-data/db"
	_ "dss-base-data/handler"
	_ "dss-base-data/robot"
	router "dss-base-data/router"
	schedule "dss-base-data/schedule"
	service "dss-base-data/service"
	"flag"
	"os"

	"github.com/kataras/iris/v12"
)

func init() {
	//时区设置
	os.Setenv("TZ", "Asia/Shanghai")
	flag.Parse()
	//初始化配置文件
	configs.ConfigRead()
	//初始化mysql
	db.InitMysql()
	//初始化定时任务
	schedule.InitScheduler()
	//初始化基本数据
	service.InitData()
}

func newApp() *iris.Application {
	app := iris.New()
	//初始化路由
	router.RegisterRoutes(app)
	return app
}

func main() {
	app := newApp()
	cfg := iris.YAML("configs/config.yml")
	addr := cfg.Other["Addr"].(string)
	app.Listen(addr, iris.WithConfiguration(cfg))
}
