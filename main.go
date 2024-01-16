package main

import (
	configs "dss-data/configs"
	db "dss-data/db"
	_ "dss-data/handler"
	_ "dss-data/robot"
	router "dss-data/router"
	"dss-data/schedule"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	//时区设置
	os.Setenv("TZ", "Asia/Shanghai")
	flag.Parse()
	//初始化配置文件
	configs.ConfigRead()
	//初始化mysql
	db.InitMysql()
	//初始化neo4j
	db.InitNeo4j()
	//初始化influxdb
	db.InitInfluxDb()
	//初始化定时任务
	schedule.InitScheduler()
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln("读取配置文件失败")
		return
	}
	g := gin.Default()
	router.RegisterRoutes(g)
	g.Run(viper.GetString("server.port"))

}
