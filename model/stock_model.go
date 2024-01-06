package model

import (
	"fmt"

	util "dss-data/util"
)

type Stock interface {
	StockInfo | StockCompany | StockQuote
}

type Tabler interface {
	TableName() string
}

//交易日历
type TradeCal struct {
	Exchange     string `json:"exchange"`      //交易所 SSE上交所 SZSE深交所
	CalDate      string `json:"cal_date"`      //日历日期
	IsOpen       int64  `json:"is_open"`       //是否交易 0休市 1交易
	PretradeDate string `json:"pretrade_date"` //上一个交易日
}

//基础数据
// type StockInfo struct {
// 	TsCode     string `json:"ts_code" gorm:"column:ts_code;primary_key"` //TS代码
// 	Symbol     string `json:"symbol" gorm:"column:symbol"`               //股票代码
// 	Name       string `json:"name" gorm:"column:name"`                   //股票名称
// 	Area       string `json:"area" gorm:"column:area"`                   //地域
// 	Industry   string `json:"industry" gorm:"column:industry"`           //所属行业
// 	Fullname   string `json:"fullname" gorm:"column:fullname"`           //股票全称
// 	Enname     string `json:"enname" gorm:"column:enname"`               //英文全称
// 	Cnspell    string `json:"cnspell" gorm:"column:cnspell"`             //拼音缩写
// 	Market     string `json:"market" gorm:"column:market"`               //市场类型（主板/创业板/科创板/CDR）
// 	Exchange   string `json:"exchange" gorm:"column:exchange"`           //交易所代码
// 	CurrType   string `json:"curr_type" gorm:"column:curr_type"`         //交易货币
// 	ListStatus string `json:"list_status" gorm:"column:list_status"`     //上市状态 L上市 D退市 P暂停上市
// 	ListDate   string `json:"list_date" gorm:"column:list_date"`         //上市日期
// 	DelistDate string `json:"delist_date" gorm:"column:delist_date"`     //退市日期
// 	IsHs       string `json:"is_hs" gorm:"column:is_hs"`                 //是否沪深港通标的，N否 H沪股通 S深股通
// }
type StockInfo struct {
	TsCode   string `json:"ts_code" gorm:"column:ts_code;primary_key"` //代码
	Symbol   string `json:"symbol" gorm:"column:symbol"`               //股票代码
	Exchange string `json:"exchange" gorm:"column:exchange"`           //交易所代码
	Market   string `json:"market" gorm:"column:market"`               //市场类型（主板/创业板/科创板/CDR）
	Name     string `json:"name" gorm:"column:name"`                   //股票名称
}

func (stockInfo StockInfo) TableName() string {
	//表名
	return "stock_info"
}

//上市公司基本信息
type StockCompany struct {
	TsCode        string  `json:"ts_code" gorm:"column:ts_code;primary_key"`           //TS代码
	Exchange      string  `json:"exchange" gorm:"column:exchange"`                     //交易所代码
	Chairman      string  `json:"chairman" gorm:"column:chairman"`                     //法人代表
	Manager       string  `json:"manager" gorm:"column:manager"`                       //总经理
	Secretary     string  `json:"secretary" gorm:"column:secretary"`                   //董秘
	RegCapital    float64 `json:"reg_capital" gorm:"column:reg_capital;decimal(11,2)"` //注册资本
	SetupDate     string  `json:"setup_date" gorm:"column:setup_date"`                 //注册日期
	Province      string  `json:"province" gorm:"column:province"`                     //所在省份
	City          string  `json:"city" gorm:"column:city"`                             //所在城市
	Introduction  string  `json:"introduction" gorm:"column:introduction"`             //公司介绍
	Website       string  `json:"website" gorm:"column:website"`                       //公司主页
	Email         string  `json:"email" gorm:"column:email"`                           //电子邮件
	Office        string  `json:"office" gorm:"column:office"`                         //办公室
	Employees     int64   `json:"employees" gorm:"column:employees"`                   //员工人数
	MainBusiness  string  `json:"main_business" gorm:"column:main_business"`           //主要业务及产品
	BusinessScope string  `json:"business_scope" gorm:"column:business_scope"`         //经营范围
}

func (stockCompany StockCompany) TableName() string {
	//表名
	return "stock_company"
}

//股票行情
type StockQuote struct {
	TsCode    string  `json:"ts_code" gorm:"column:ts_code;primary_key"`       //股票代码
	TradeDate string  `json:"trade_date" gorm:"column:trade_date;primary_key"` //交易日期
	Open      float64 `json:"open" gorm:"column:open;float(11,2)"`             //开盘价
	High      float64 `json:"high" gorm:"column:high;float(11,2)"`             //最高价
	Low       float64 `json:"low" gorm:"column:low;float(11,2)"`               //最低价
	Close     float64 `json:"close" gorm:"column:close;float(11,2)"`           //收盘价
	PreClose  float64 `json:"pre_close" gorm:"column:pre_close;float(11,2)"`   //昨收价(前复权)
	Change    float64 `json:"change" gorm:"column:change;float(11,2)"`         //涨跌额
	PctChg    float64 `json:"pct_chg" gorm:"column:pct_chg;float(11,2)"`       //涨跌幅(未复权)
	Vol       float64 `json:"vol" gorm:"column:vol;float(11,2)"`               //成交量(万手)
	Amount    float64 `json:"amount" gorm:"column:amount;float(11,2)"`         //成交额(万)
	LimitUp   int64   `json:"limit_up" gorm:"column:limit_up;tinyint(1)"`      //涨停板
}

func (stockQuote StockQuote) TableName() string {
	//表名
	tableName := fmt.Sprintf("%s%s", "stock_quote_", util.Substr(stockQuote.TradeDate, 0, 4))
	return tableName
}

type LongHu struct {
	Id        string  `json:"id" gorm:"column:id;primary_key"`               //id
	Type      string  `json:"type" gorm:"column:type"`                       //类型
	Symbol    string  `json:"symbol" gorm:"column:symbol"`                   //股票代码
	TradeDate string  `json:"trade_date" gorm:"column:trade_date"`           //交易日期
	Name      string  `json:"name" gorm:"column:name"`                       //股票名称
	Close     float64 `json:"close" gorm:"column:close;float(11,2)"`         //收盘价
	PctChg    float64 `json:"pct_chg" gorm:"column:pct_chg;float(11,2)"`     //涨跌幅
	Volume    float64 `json:"volume" gorm:"column:volume;float(11,2)"`       //成交量
	Amount    float64 `json:"amount" gorm:"column:amount;float(11,2)"`       //成交额
	NetWorth  float64 `json:"net_worth" gorm:"column:net_worth;float(11,2)"` //净买入额
	// Detail    []LongHuDetail `json:"detail" gorm:"foreignKey:long_hu_id;references:id"` //详情
}

func (longHu LongHu) TableName() string {
	return "long_hu"
}

type LongHuDetail struct {
	LongHuId string  `json:"long_hu_id" gorm:"column:long_hu_id"`           //龙虎榜id
	Dept     string  `json:"dept" gorm:"column:dept"`                       //营业部
	Buy      float64 `json:"buy" gorm:"column:buy;float(11,2)"`             //买入额
	Sell     float64 `json:"sell" gorm:"column:sell;float(11,2)"`           //卖出额
	Ratio    float64 `json:"ratio" gorm:"column:ratio;float(11,2)"`         //占比
	NetWorth float64 `json:"net_worth" gorm:"column:net_worth;float(11,2)"` //净买入额
}

func (longHuDetail LongHuDetail) TableName() string {
	return "long_hu_detail"
}

// 东方财富网板块列表
type Bk struct {
	Code string `json:"code" gorm:"column:code;primary_key"` //板块代码
	Name string `json:"name" gorm:"column:name"`             //板块名称
	Type int    `json:"type" gorm:"column:type"`             //板块类型 1:地域;2:行业;3:概念
}

func (bk Bk) TableName() string {
	return "bk"
}

//概念关联列表
type BkRelSymbol struct {
	Symbol string `json:"symbol" gorm:"column:symbol"`   //股票代码
	BkCode string `json:"bk_code" gorm:"column:bk_code"` //板块编码
}

func (bkRelSymbol BkRelSymbol) TableName() string {
	return "bk_rel_symbol"
}

//同花顺行业行情
type ThsHyQuote struct {
	Code      string  `json:"code" gorm:"column:code;primary_key"`             //行业代码
	TradeDate string  `json:"trade_date" gorm:"column:trade_date;primary_key"` //交易日期
	Open      float64 `json:"open" gorm:"column:open;float(30,2)"`             //开盘价
	High      float64 `json:"high" gorm:"column:high;float(30,2)"`             //最高价
	Low       float64 `json:"low" gorm:"column:low;float(30,2)"`               //最低价
	Close     float64 `json:"close" gorm:"column:close;float(30,2)"`           //收盘价
	PreClose  float64 `json:"pre_close" gorm:"column:pre_close;float(30,2)"`   //昨收价
	Change    float64 `json:"change" gorm:"column:change;float(30,2)"`         //资金流入(亿)
	PctChg    float64 `json:"pct_chg" gorm:"column:pct_chg;float(30,2)"`       //涨跌幅
	Vol       float64 `json:"vol" gorm:"column:vol;float(30,2)"`               //成交量(万手)
	Amount    float64 `json:"amount" gorm:"column:amount;float(30,2)"`         //成交额(亿)
	Rank      int     `json:"rank" gorm:"column:rank;int(11)"`                 //涨幅排名
	RiseCount int     `json:"rise_count" gorm:"column:rise_count;int(11)"`     //上涨家数
	FallCount int     `json:"fall_count" gorm:"column:fall_count;int(11)"`     //下跌家数
}

func (thsHyQuote ThsHyQuote) TableName() string {
	return "ths_hy_quote"
}

//同花顺概念行情
type ThsGnQuote struct {
	Code      string  `json:"code" gorm:"column:code;primary_key"`             //概念代码
	TradeDate string  `json:"trade_date" gorm:"column:trade_date;primary_key"` //交易日期
	Open      float64 `json:"open" gorm:"column:open;float(30,2)"`             //开盘价
	High      float64 `json:"high" gorm:"column:high;float(30,2)"`             //最高价
	Low       float64 `json:"low" gorm:"column:low;float(30,2)"`               //最低价
	Close     float64 `json:"close" gorm:"column:close;float(30,2)"`           //收盘价
	PreClose  float64 `json:"pre_close" gorm:"column:pre_close;float(30,2)"`   //昨收价
	Change    float64 `json:"change" gorm:"column:change;float(30,2)"`         //资金流入(亿)
	PctChg    float64 `json:"pct_chg" gorm:"column:pct_chg;float(30,2)"`       //涨跌幅
	Vol       float64 `json:"vol" gorm:"column:vol;float(30,2)"`               //成交量(万手)
	Amount    float64 `json:"amount" gorm:"column:amount;float(30,2)"`         //成交额(亿)
	Rank      int     `json:"rank" gorm:"column:rank;int(11)"`                 //涨幅排名
	RiseCount int     `json:"rise_count" gorm:"column:rise_count;int(11)"`     //上涨家数
	FallCount int     `json:"fall_count" gorm:"column:fall_count;int(11)"`     //下跌家数
}

func (thsGnQuote ThsGnQuote) TableName() string {
	return "ths_gn_quote"
}
