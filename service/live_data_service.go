package service

import (
	"dss-data/model"
	"dss-data/util"
	"errors"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 批量获取股票实时数据
func GetLiveData(symbols []string) (*[]model.LiveData, error) {
	if len(symbols) <= 100 {
		return getLiveData(symbols)
	}
	datas := []string{}
	result := []model.LiveData{}
	for i, v := range symbols {
		datas = append(datas, v)
		if len(datas) > 100 || i+1 == len(symbols) {
			liveDatas, _ := getLiveData(datas)
			result = append(result, *liveDatas...)
			datas = []string{}
		}
	}
	return &result, nil
}

func getLiveData(symbols []string) (*[]model.LiveData, error) {
	//随机从网易和腾讯获取数据
	// random := rand.Intn(2) //生成0-99之间的随机数
	var result *[]model.LiveData
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

// 从腾讯接口获取实时数据
func TencentLiveData(symbols []string) (*[]model.LiveData, error) {
	if len(symbols) <= 0 {
		return nil, errors.New("symbols is null")
	}
	//http://qt.gtimg.cn/q=sz002603,sz002693,sh603232
	url := "http://qt.gtimg.cn/q="
	param := ""
	for index, symbol := range symbols {
		param += symbol
		if index != len(symbols) {
			param += ","
		}
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
	// gbk转utf-8
	tmp, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(response))
	response = string(tmp)
	result := []model.LiveData{}
	lines := strings.Split(response, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		dataStr := util.Substr(line, 12, len([]rune(line))-2)
		values := strings.Split(dataStr, "~")
		livaData := model.LiveData{}
		livaData.Name = values[1]                    //名称
		livaData.Code = values[2]                    //编码
		tmp1, _ := strconv.ParseFloat(values[3], 64) //当前价
		livaData.Now = tmp1
		tmp4, _ := strconv.ParseFloat(values[4], 64) //昨收价
		livaData.PreClose = tmp4
		tmp5, _ := strconv.ParseFloat(values[5], 64) //开盘价
		livaData.Open = tmp5
		tmp6, _ := strconv.ParseFloat(values[6], 64) //成交量
		livaData.Vol = tmp6
		tmp9, _ := strconv.ParseFloat(values[9], 64) //买1价格
		livaData.Buy1 = tmp9
		tmp10, _ := strconv.ParseFloat(values[10], 64) //买1数量（手）
		livaData.Buy1Count = tmp10
		tmp11, _ := strconv.ParseFloat(values[11], 64) //买2价格
		livaData.Buy2 = tmp11
		tmp12, _ := strconv.ParseFloat(values[12], 64) //买2数量（手）
		livaData.Buy2Count = tmp12
		tmp13, _ := strconv.ParseFloat(values[13], 64) //买3价格
		livaData.Buy3 = tmp13
		tmp14, _ := strconv.ParseFloat(values[14], 64) //买3数量（手）
		livaData.Buy3Count = tmp14
		tmp15, _ := strconv.ParseFloat(values[15], 64) //买4价格
		livaData.Buy4 = tmp15
		tmp16, _ := strconv.ParseFloat(values[16], 64) //买4数量（手）
		livaData.Buy4Count = tmp16
		tmp17, _ := strconv.ParseFloat(values[17], 64) //买5价格
		livaData.Buy5 = tmp17
		tmp18, _ := strconv.ParseFloat(values[18], 64) //买5数量（手）
		livaData.Buy5Count = tmp18
		tmp19, _ := strconv.ParseFloat(values[19], 64) //卖1价格
		livaData.Sell1 = tmp19
		tmp20, _ := strconv.ParseFloat(values[20], 64) //卖1数量（手）
		livaData.Sell1Count = tmp20
		tmp21, _ := strconv.ParseFloat(values[21], 64) //卖2价格
		livaData.Sell2 = tmp21
		tmp22, _ := strconv.ParseFloat(values[22], 64) //卖2数量（手）
		livaData.Sell2Count = tmp22
		tmp23, _ := strconv.ParseFloat(values[23], 64) //卖3价格
		livaData.Sell3 = tmp23
		tmp24, _ := strconv.ParseFloat(values[24], 64) //卖3数量（手）
		livaData.Sell3Count = tmp24
		tmp25, _ := strconv.ParseFloat(values[25], 64) //卖4价格
		livaData.Sell4 = tmp25
		tmp26, _ := strconv.ParseFloat(values[26], 64) //卖4数量（手）
		livaData.Sell4Count = tmp26
		tmp27, _ := strconv.ParseFloat(values[27], 64) //卖5价格
		livaData.Sell5 = tmp27
		tmp28, _ := strconv.ParseFloat(values[28], 64) //买5数量（手）
		livaData.Sell5Count = tmp28
		tmp31, _ := strconv.ParseFloat(values[31], 64) //涨跌
		livaData.Change = tmp31
		tmp32, _ := strconv.ParseFloat(values[32], 64) //涨跌幅
		livaData.PctChg = tmp32
		livaData.Time = values[30]                     //更新时间
		tmp33, _ := strconv.ParseFloat(values[33], 64) //最高价
		livaData.Max = tmp33
		tmp34, _ := strconv.ParseFloat(values[34], 64) //最低价
		livaData.Min = tmp34
		tmp38, _ := strconv.ParseFloat(values[38], 64) //换手
		livaData.Hands = tmp38
		tmp42, _ := strconv.ParseFloat(values[43], 64) //振幅
		livaData.Amplitude = tmp42
		tmp43, _ := strconv.ParseFloat(values[44], 64) //流通市值
		livaData.FloatValue = tmp43
		tmp44, _ := strconv.ParseFloat(values[45], 64) //总市值
		livaData.TotalValue = tmp44
		tmp45, _ := strconv.ParseFloat(values[46], 64) //市净率
		livaData.Ptb = tmp45
		tmp51, _ := strconv.ParseFloat(values[52], 64) //市盈率
		livaData.Pte = tmp51
		tmp56, _ := strconv.ParseFloat(values[57], 64) //成交额
		livaData.Amount = tmp56
		result = append(result, livaData)
	}
	//转换结果
	return &result, nil
}

// 从网易接口获取实时数据
func WangyiLiveData(symbols []string) (*[]model.LiveData, error) {
	//http://api.money.126.net/data/feed/0601398%2c1000001%2c1000881%2cmoney.api
	return nil, nil
}
