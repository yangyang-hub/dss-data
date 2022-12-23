package config

import (
	"log"
	"reflect"

	viper "github.com/spf13/viper"
)

var Config *ConfigMap

type ConfigMap struct {
	TushareUrl    string `yml:"tushare.url"`
	TushareToken  string `yml:"tushare.token"`
	ProxyUrl      string `yml:"proxy.url"`
	MysqlHost     string `yml:"mysql.host"`
	MysqlPort     int    `yml:"mysql.port"`
	MysqlDatabase string `yml:"mysql.database"`
	MysqlUsername string `yml:"mysql.username"`
	Mysqlpassword string `yml:"mysql.password"`
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
		ymlTag := field.Tag.Get("yml")
		value := viper.Get(ymlTag)
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
