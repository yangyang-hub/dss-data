package service

import (
	"errors"
	"fmt"

	"github.com/yangyang-hub/dss-common/util"
)

//批量获取股票实时数据
func GetLiveData(symbols []string) {
	//随机从网易和腾讯获取数据

}

//从腾讯接口获取实时数据
func TencentLiveData(symbols []string) (interface{}, error) {
	if len(symbols) <= 0 {
		return nil, errors.New("symbols is null")
	}
	//http://qt.gtimg.cn/q=sz002603,sz002693,sh603232
	url := "http://qt.gtimg.cn/q="
	param := ""
	isFirst := true
	for _, symbol := range symbols {
		if !isFirst {
			param += ","
			isFirst = false
		}
		param += symbol
	}
	result, err := util.SendGetResString(url + param)
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	//转换结果
	return nil, nil
}

//从网易接口获取实时数据
func WangyiLiveData(symbols []string) {
	//http://api.money.126.net/data/feed/0601398%2c1000001%2c1000881%2cmoney.api
}
