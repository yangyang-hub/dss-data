package service

import (
	"errors"
	"strings"

	"github.com/yangyang-hub/dss-common/util"
)

//批量获取股票实时数据
func GetLiveData(symbols []string) (*[]LiveData, error) {
	//随机从网易和腾讯获取数据
	// random := rand.Intn(2) //生成0-99之间的随机数
	var result *[]LiveData
	var err error
	// if random == 0 {
	result, err = TencentLiveData(symbols)
	// if err != nil {
	// 	result, err = WangyiLiveData(symbols)
	// }
	// } else {
	// 	result, err = WangyiLiveData(symbols)
	// 	if err != nil {
	// 		result, err = TencentLiveData(symbols)
	// 	}
	// }
	return result, err
}

//从腾讯接口获取实时数据
func TencentLiveData(symbols []string) (*[]LiveData, error) {
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
	response, err := util.SendGetResString(url + param)
	if err != nil {
		return nil, err
	}
	/*
		v_sz002603="51~以岭药业~002603~29.60~29.73~29.72~184862~90592~94271~29.59~431~29.58~311~29.57~337~29.56~711~29.55~1757~29.60~120~29.61~35~29.62~53~29.63~70~29.64~122~~20230113103221~-0.13~-0.44~29.84~29.55~29.60/184862/548722160~184862~54872~1.34~32.23~~29.84~29.55~0.98~407.29~494.53~4.95~32.70~26.76~0.94~3147~29.68~26.22~36.80~~~0.64~54872.2160~0.0000~0~
		~GP-A~-1.20~1.09~1.01~15.37~10.26~53.96~18.82~-3.99~-33.93~22.26~1375990295~1670705376~79.73~-0.24~1375990295~~~30.00~-0.17~";
		v_sz002693="51~双成药业~002693~7.30~7.24~7.20~47357~28846~18511~7.29~83~7.28~315~7.27~248~7.26~216~7.25~299~7.30~186~7.31~273~7.32~1148~7.33~1188~7.34~926~~20230113103221~0.06~0.83~7.34~7.18~7.30/47357/34438056~47357~3444~1.16~-152.05~~7.34~7.18~2.21~29.79~30.28~6.07~7.96~6.52~1.46~-2560~7.27~-627.21~-148.96~~~0.68~3443.8056~0.0000~0~
		~GP-A~0.41~-2.54~0.00~-3.99~-4.23~11.95~3.91~-0.82~-15.41~-23.88~408143821~414737000~-52.44~72.58~408143821~~~47.18~0.14~";
		v_sh603232="1~格尔软件~603232~16.36~16.71~16.41~18437~8505~9932~16.36~34~16.35~247~16.34~82~16.33~30~16.32~66~16.37~135~16.38~38~16.39~2~16.40~163~16.41~1~~20230113103223~-0.35~-2.09~16.69~16.31~16.36/18437/30306018~18437~3031~0.79~73.34~~16.69~16.31~2.27~38.08~38.08~2.86~18.38~15.04~0.79~120~16.44~-56.42~47.78~~~1.15~3030.6018~0.0000~0~
		~GP-A~7.63~-4.33~0.79~3.94~3.08~21.10~8.85~8.63~0.18~33.77~232790328~232790328~15.04~47.12~232790328~~~0.00~-0.18~";
	*/
	result := []LiveData{}
	lines := strings.Split(response, "\n")
	for _, line := range lines {
		code := util.Substr(line, 2, 10)
		dataStr := util.Substr(line, 12, len(line)-2)
		values := strings.Split(dataStr, "~")
		livaData := LiveData{Code: code}
		livaData.Name = values[1]
		livaData.Now = values[3]
		livaData.Change = values[31]
		livaData.ChangePercent = values[32]
		livaData.Time = values[30]
		livaData.Max = values[33]
		livaData.Min = values[34]
	}
	//转换结果
	return &result, nil
}

//从网易接口获取实时数据
func WangyiLiveData(symbols []string) (*[]LiveData, error) {
	//http://api.money.126.net/data/feed/0601398%2c1000001%2c1000881%2cmoney.api
	return nil, nil
}
