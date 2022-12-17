package tushare

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func HandleErrorStr(err error, format string, a ...any) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

type httpClient[T any] struct {
	url         string
	contentType string
	param       map[string]interface{}
}

type response[T any] struct {
	Code      int             `json:"code"`
	RequestId string          `json:"request_id"`
	Msg       string          `json:"msg"`
	Data      responseData[T] `json:"data"`
}

type responseData[T any] struct {
	Node   []T
	Fields []string        `json:"fields"`
	Items  [][]interface{} `json:"items"`
}

func (client httpClient[T]) sendPost() *response[T] {
	bytesData, err := json.Marshal(client.param)
	if HandleErrorStr(err, "参数异常;url: %v,param:%v\n", client.url, client.param) {
		return nil
	}
	res, err := http.Post(client.url, client.contentType, bytes.NewBuffer(bytesData))
	if HandleErrorStr(err, "请求失败;url: %v,param:%v\n", client.url, client.param) {
		return nil
	}
	defer res.Body.Close()
	contentBytes, err := ioutil.ReadAll(res.Body)
	HandleErrorStr(err, "response读取异常;url: %v,param:%v\n", client.url, client.param)
	//var mapResult map[string]interface{}
	var result response[T]
	err = json.Unmarshal(contentBytes, &result)
	if HandleErrorStr(err, "response返回数据json转换异常(%v)", string(contentBytes)) {
		return nil
	}
	result.Data.resultToStruct(getQuoteType(client.param))
	return &result
}

//获取股票行情 设置行情类型
func getQuoteType(param map[string]interface{}) string {
	if param["api_name"] == "daily" {
		return "D"
	} else if param["api_name"] == "weekly" {
		return "W"
	} else if param["api_name"] == "monthly" {
		return "M"
	}
	return ""
}

func (res *responseData[T]) resultToStruct(quoteType string) {
	items := res.Items
	fields := res.Fields
	var result []T
	fieldsMap := make(map[int]string)
	for i := 0; i < len(fields); i++ {
		fieldsMap[i] = fields[i]
	}
	for _, item := range items {
		entity := new(T)
		t := reflect.TypeOf(*entity)
		if quoteType != "" && t.Name() == "StockQuote" {
			reflect.ValueOf(&*entity).Elem().FieldByName("Type").SetString(quoteType)
		}
		for j, value := range item {
			if value == nil {
				continue
			}
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				jsonName := field.Tag.Get("json")
				if jsonName == fieldsMap[j] {
					fieldName := field.Name
					fieldType := field.Type.Name()
					switch fieldType {
					case "string":
						reflect.ValueOf(&*entity).Elem().FieldByName(fieldName).SetString(value.(string))
					case "float64":
						reflect.ValueOf(&*entity).Elem().FieldByName(fieldName).SetFloat(value.(float64))
					case "int64":
						switch value.(type) {
						case float64:
							reflect.ValueOf(&*entity).Elem().FieldByName(fieldName).SetInt(int64(value.(float64)))
						default:
							reflect.ValueOf(&*entity).Elem().FieldByName(fieldName).SetInt(int64(value.(int)))
						}
					}
					break
				}
			}
		}
		result = append(result, *entity)
	}
	res.Node = result
}
