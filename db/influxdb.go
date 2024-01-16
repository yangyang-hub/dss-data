package db

import (
	"context"
	config "dss-data/configs"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
)

var client influxdb2.Client
var writeAPI api.WriteAPIBlocking

func InitInfluxDb() {
	defer log.Println("success init influxdb")
	client = influxdb2.NewClient(config.Config.InfluxDbUrl, config.Config.InfluxDbToken)
	writeAPI = client.WriteAPIBlocking(config.Config.InfluxDbOrg, config.Config.InfluxDbBucket)
	writeAPI.EnableBatching()
}

func CloseClient(c influxdb2.Client) {
	c.Close()
}

func WritePoint(points *[]*write.Point) {
	err := writeAPI.WritePoint(context.Background(), *points...)
	if err != nil {
		log.Println("influxdb write point error: ", err)
		return
	}
}
