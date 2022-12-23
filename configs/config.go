package config

import (
	"log"
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
}

func ConfigRead() {
	defer log.Println("success init config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")
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
			value = viper.Get(envTag)
		}
		if envTag == "" || value == nil {
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
