package tushare

import (
	configs "dss-data/configs"

	model "dss-data/model"
)

//交易日历
func GetTradeCal(startDate string, endDate string) *[]model.TradeCal {
	client := httpClient[model.TradeCal]{
		url:         configs.Config.TushareUrl,
		contentType: "application/json;charset=utf-8",
		param: map[string]interface{}{
			"api_name": "trade_cal",
			"params":   map[string]interface{}{"exchange": "SSE", "start_date": startDate, "end_date": endDate, "is_open": "1"},
			"fields":   "exchange,cal_date,is_open,pretrade_date",
			"token":    configs.Config.TushareToken,
		},
	}
	response := client.sendPost()
	datas := response.Data.Node
	return &datas
}

//股票基本信息
func GetStockBasicData(params map[string]interface{}) *[]model.StockInfo {
	client := httpClient[model.StockInfo]{
		url:         configs.Config.TushareUrl,
		contentType: "application/json;charset=utf-8",
		param: map[string]interface{}{
			"api_name": "stock_basic",
			"params":   params,
			"fields":   "ts_code,symbol,name,area,industry,fullname,enname,cnspell,market,exchange,curr_type,list_status,list_date,delist_date,is_hs",
			"token":    configs.Config.TushareToken,
		},
	}
	response := client.sendPost()
	datas := response.Data.Node
	return &datas
}

//上市公司基本信息
func GetStockCompanyData(params map[string]interface{}) *[]model.StockCompany {
	client := httpClient[model.StockCompany]{
		url:         configs.Config.TushareUrl,
		contentType: "application/json;charset=utf-8",
		param: map[string]interface{}{
			"api_name": "stock_company",
			"params":   params,
			"fields":   "ts_code,exchange,chairman,manager,secretary,reg_capital,setup_date,province,city,introduction,website,email,office,employees,main_business,business_scope",
			"token":    configs.Config.TushareToken,
		},
	}
	response := client.sendPost()
	datas := response.Data.Node
	return &datas
}

//获取股票行情
func GetStockQuoteData(params map[string]interface{}, quoteType string) *[]model.StockQuote {
	client := httpClient[model.StockQuote]{
		url:         configs.Config.TushareUrl,
		contentType: "application/json;charset=utf-8",
		param: map[string]interface{}{
			"api_name": quoteType,
			"params":   params,
			"fields":   "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount",
			"token":    configs.Config.TushareToken,
		},
	}
	response := client.sendPost()
	datas := response.Data.Node
	return &datas
}
