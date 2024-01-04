package model

//实时数据
type LiveData struct {
	Code       string  `json:"code"`        //编码
	Name       string  `json:"name"`        //名称
	Time       string  `json:"time"`        //更新时间
	Now        float64 `json:"now"`         //当前价
	Change     float64 `json:"change"`      //涨跌
	PctChg     float64 `json:"pct_chg"`     //涨跌幅
	Open       float64 `json:"open"`        //开盘价
	PreClose   float64 `json:"pre_close"`   //昨收价
	Max        float64 `json:"max"`         //最高价
	Min        float64 `json:"min"`         //最低价
	Amplitude  float64 `json:"amplitude"`   //振幅
	Hands      float64 `json:"hands"`       //换手
	Vol        float64 `json:"vol"`         //成交量
	Amount     float64 `json:"amount"`      //成交额
	FloatValue float64 `json:"float_value"` //流通市值
	TotalValue float64 `json:"total_value"` //总市值
	Ptb        float64 `json:"ptb"`         //市净率
	Pte        float64 `json:"pte"`         //市盈率
	Buy1       float64 `json:"buy1"`        //买1价格
	Buy2       float64 `json:"buy2"`        //买2价格
	Buy3       float64 `json:"buy3"`        //买3价格
	Buy4       float64 `json:"buy4"`        //买4价格
	Buy5       float64 `json:"buy5"`        //买5价格
	Sell1      float64 `json:"sell1"`       //卖1价格
	Sell2      float64 `json:"sell2"`       //卖2价格
	Sell3      float64 `json:"sell3"`       //卖3价格
	Sell4      float64 `json:"sell4"`       //卖4价格
	Sell5      float64 `json:"sell5"`       //卖5价格
	Buy1Count  float64 `json:"buy1_count"`  //买1数量（手）
	Buy2Count  float64 `json:"buy2_count"`  //买2数量（手）
	Buy3Count  float64 `json:"buy3_count"`  //买3数量（手）
	Buy4Count  float64 `json:"buy4_count"`  //买4数量（手）
	Buy5Count  float64 `json:"buy5_count"`  //买5数量（手）
	Sell1Count float64 `json:"sell1_count"` //卖1数量（手）
	Sell2Count float64 `json:"sell2_count"` //卖2数量（手）
	Sell3Count float64 `json:"sell3_count"` //卖3数量（手）
	Sell4Count float64 `json:"sell4_count"` //卖4数量（手）
	Sell5Count float64 `json:"sell5_count"` //卖5数量（手）
}
