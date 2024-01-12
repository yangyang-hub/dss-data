package config

import (
	"log"
	"os"
	"reflect"

	viper "github.com/spf13/viper"
)

var Config *ConfigMap

type ConfigMap struct {
	TushareUrl    string `yml:"tushare.url" env:"TUSHARE_URL"`
	TushareToken  string `yml:"tushare.token" env:"TUSHARE_TOKEN"`
	ProxyUrl      string `yml:"proxy.url" env:"PROXY_URL"`
	MysqlUrl      string `yml:"mysql.url" env:"MYSQL_URL"`
	MysqlDatabase string `yml:"mysql.database" env:"MYSQL_DATABASE"`
	Neo4jUrl      string `yml:"neo4j.url" env:"NEO4J_URL"`
	Neo4jDatabase string `yml:"neo4j.database" env:"NEO4J_DATABASE"`
	Neo4jUsername string `yml:"neo4j.username" env:"NEO4J_USERNAME"`
	Neo4jPassword string `yml:"neo4j.password" env:"NEO4J_PASSWORD"`
}

func ConfigRead() {
	defer log.Println("success init config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err.Error())
		panic("init config fail")
	}
	config := new(ConfigMap)
	config.readConfigFromYaml()
	Config = config
}

func (this *ConfigMap) readConfigFromYaml() {
	t := reflect.TypeOf(*this)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		var value interface{}
		envTag := field.Tag.Get("env")
		if envTag != "" {
			value = os.Getenv(envTag)
		}
		if envTag == "" || value == "" {
			ymlTag := field.Tag.Get("yml")
			value = viper.Get(ymlTag)
		}
		fieldName := field.Name
		fieldType := field.Type.Name()
		switch fieldType {
		case "string":
			reflect.ValueOf(&*this).Elem().FieldByName(fieldName).SetString(value.(string))
		case "int64":
			reflect.ValueOf(&*this).Elem().FieldByName(fieldName).SetInt(int64(value.(int)))
		case "int":
			reflect.ValueOf(&*this).Elem().FieldByName(fieldName).SetInt(int64(value.(int)))
		}
	}
}
