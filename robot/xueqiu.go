package robot

import (
	"log"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/yangyang-hub/dss-common/util"
)

type LongHu struct {
	Id        string         `json:"id" gorm:"column:id;primary_key"`                   //id
	Type      string         `json:"type" gorm:"column:type"`                           //类型
	Symbol    string         `json:"symbol" gorm:"column:symbol"`                       //股票代码
	TradeDate string         `json:"trade_date" gorm:"column:trade_date"`               //交易日期
	Name      string         `json:"name" gorm:"column:name"`                           //股票名称
	Close     float64        `json:"close" gorm:"column:close;float(11,2)"`             //收盘价
	PctChg    float64        `json:"pct_chg" gorm:"column:pct_chg;float(11,2)"`         //涨跌幅
	Volume    float64        `json:"volume" gorm:"column:volume;float(11,2)"`           //成交量
	Amount    float64        `json:"amount" gorm:"column:amount;float(11,2)"`           //成交额
	NetWorth  float64        `json:"net_worth" gorm:"column:net_worth;float(11,2)"`     //净买入额
	Detail    []LongHuDetail `json:"detail" gorm:"foreignKey:long_hu_id;references:id"` //详情
}

type LongHuDetail struct {
	LongHuId string  `json:"long_hu_id" gorm:"column:long_hu_id"`           //龙虎榜id
	Dept     string  `json:"dept" gorm:"column:dept"`                       //营业部
	Buy      float64 `json:"buy" gorm:"column:buy;float(11,2)"`             //买入额
	Sell     float64 `json:"sell" gorm:"column:sell;float(11,2)"`           //卖出额
	Ratio    float64 `json:"ratio" gorm:"column:ratio;float(11,2)"`         //买入额
	NetWorth float64 `json:"net_worth" gorm:"column:net_worth;float(11,2)"` //净买入额
}

type StockInfo struct {
	Code     string `json:"code" gorm:"column:code;primary_key"` //代码
	Symbol   string `json:"symbol" gorm:"column:symbol"`         //股票代码
	Exchange string `json:"exchange" gorm:"column:exchange"`     //交易所代码
	Market   string `json:"market" gorm:"column:market"`         //市场类型（主板/创业板/科创板/CDR）
	Name     string `json:"name" gorm:"column:name"`             //股票名称
}

var (
	xueqiuCookie     string
	xueqiuCookieTime int64
)

// 查询所有股票代码
func GetAllStock() *[]StockInfo {
	return getStockByPage(1, 60)
}

// 分页查询股票代码
func getStockByPage(page, size int) *[]StockInfo {
	result := []StockInfo{}
	url := "https://stock.xueqiu.com/v5/stock/screener/quote/list.json?page=" + strconv.Itoa(page) + "&size=" + strconv.Itoa(size) + "&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz"
	respone := visitXueQiuJson(url)
	data, _ := (*respone)["data"].(map[string]interface{})
	if len(data) > 0 {
		lists, _ := data["list"].([]interface{})
		if len(lists) > 0 {
			for _, list := range lists {
				item := list.(map[string]interface{})
				stockInfo := StockInfo{}
				for key, value := range item {
					switch key {
					case "symbol":
						symbol := value.(string)
						code := util.Substr(symbol, 2, len(symbol))
						stockInfo.Symbol = symbol
						stockInfo.Code = code
						stockInfo.Exchange = util.Substr(symbol, 0, 2)
						stockInfo.Market = getMarket(code)
					case "name":
						stockInfo.Name = value.(string)
					}
				}
				result = append(result, stockInfo)
			}
		}
		page++
		temp := getStockByPage(page, size)
		result = append(result, *temp...)
	}
	return &result
}

// 访问雪球当日龙虎榜数据
func GetLonghu() *[]LongHu {
	result := []LongHu{}
	date := time.Now().Format("20060102")
	// date := "20230829"
	url := "http://stock.xueqiu.com/v5/stock/hq/longhu.json?date=" + getUnix(date)
	respone := visitXueQiuJson(url)
	data, _ := (*respone)["data"].(map[string]interface{})
	if len(data) > 0 {
		items, _ := data["items"].([]interface{})
		if len(items) > 0 {
			for _, ite := range items {
				item := ite.(map[string]interface{})
				symbol := ""
				longhu := LongHu{}
				longhu.TradeDate = date
				longHuId := uuid.NewV4().String()
				longhu.Id = longHuId
				for key, value := range item {
					switch key {
					case "symbol":
						symbol = value.(string)
						longhu.Symbol = util.Substr(symbol, 2, len(symbol))
					case "name":
						longhu.Name = value.(string)
					case "close":
						longhu.Close = value.(float64)
					case "percent":
						longhu.PctChg = value.(float64)
					case "volume":
						longhu.Volume = value.(float64)
					case "amount":
						longhu.Amount = value.(float64)
					case "type_name":
						typeNames := value.([]interface{})
						str := ""
						for _, typeName := range typeNames {
							str += typeName.(string) + ","
						}
						longhu.Type = util.Substr(str, 0, len(str)-1)
					}
				}
				longhuDetails := getLonghuDetail(symbol, longHuId)
				longhu.Detail = *longhuDetails
				result = append(result, longhu)
			}
		}
	}
	return &result
}

// 访问雪球龙虎榜数据详情
func getLonghuDetail(symbol, longHuId string) *[]LongHuDetail {
	result := []LongHuDetail{}
	url := "http://stock.xueqiu.com/v5/stock/capital/longhu.json?symbol=" + symbol + "&page=1&size=1"
	respone := visitXueQiuJson(url)
	data, _ := (*respone)["data"].(map[string]interface{})
	if len(data) > 0 {
		items, _ := data["items"].([]interface{})
		if len(items) > 0 {
			for _, item1 := range items {
				item := item1.([]interface{})
				if len(item) > 0 {
					for _, ite1 := range item {
						ite := ite1.(map[string]interface{})
						branches := ite["branches"].([]interface{})
						if len(branches) > 0 {
							for _, branch1 := range branches {
								branch := branch1.(map[string]interface{})
								longHuDetail := LongHuDetail{}
								longHuDetail.LongHuId = longHuId
								for key, value := range branch {
									switch key {
									case "branch_name":
										longHuDetail.Dept = value.(string)
									case "buy_amt":
										longHuDetail.Buy = value.(float64)
									case "sell_amt":
										longHuDetail.Sell = value.(float64)
									case "ratio":
										longHuDetail.Ratio = value.(float64)
									case "net_amt":
										longHuDetail.NetWorth = value.(float64)
									}
								}
								result = append(result, longHuDetail)
							}
						}
					}
				}
			}
		}
	}
	return &result
}

// 访问雪球返回json
func visitXueQiuJson(url string) *map[string]interface{} {
	return visitJson(url, getXueQiuCookie())
}

// 访问雪球返回html
func visitXueQiuHtml(url, goquerySelector string, f colly.HTMLCallback) {
	visitHtml(url, getXueQiuCookie(), goquerySelector, f)
}

// 获取雪球cookie
func getXueQiuCookie() string {
	if time.Now().Unix()-xueqiuCookieTime > 600 || xueqiuCookie == "" {
		cookieStr := ""
		c := colly.NewCollector()
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
		c.OnResponse(func(r *colly.Response) {
			cookies := c.Cookies(r.Request.URL.String())
			for _, cookie := range cookies {
				cookieStr += cookie.Name + "=" + cookie.Value + "; "
			}
		})
		c.Visit("http://xueqiu.com")
		c.Wait()
		xueqiuCookieTime = time.Now().Unix()
		xueqiuCookie = cookieStr
	}
	return xueqiuCookie
}

func getUnix(date string) string {
	t, err := time.ParseInLocation("20060102", date, time.Local)
	if err != nil {
		log.Println("GetLonghu 解析日期失败:", err)
		return ""
	}
	timestamp := t.Unix()
	return strconv.FormatInt(timestamp, 10) + "000"
}

func getMarket(code string) string {
	prefix := util.Substr(code, 0, 2)
	if prefix == "60" {
		return "SHA"
	} else if prefix == "90" {
		return "SHB"
	} else if prefix == "00" {
		return "SZA"
	} else if prefix == "20" {
		return "SZB"
	} else if prefix == "30" {
		return "CY"
	} else if prefix == "68" {
		return "KC"
	} else if prefix == "43" || prefix == "83" || prefix == "87" || prefix == "88" {
		return "BJ"
	}
	return ""
}
