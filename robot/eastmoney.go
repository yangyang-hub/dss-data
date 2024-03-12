package robot

import (
	"dss-data/model"
	"dss-data/util"
	"log"
	"strconv"
	"time"
)

// 查询东财所有板块实时详情
func GetAllBkQuotes() *[]model.BkQuote {
	resultQuotes := []model.BkQuote{}
	for i := 1; i <= 3; i++ {
		_, bkQuotes := getBk(i)
		resultQuotes = append(resultQuotes, *bkQuotes...)
	}
	return &resultQuotes
}

// 查询东财所有板块以及板块关联股票
func GetAllBkAndRelSymbol() (*[]model.Bk, *[]model.BkRelSymbol, *[]model.BkQuote) {
	resultBks := []model.Bk{}
	resultRels := []model.BkRelSymbol{}
	resultQuotes := []model.BkQuote{}
	for i := 1; i <= 3; i++ {
		bks, bkQuotes := getBk(i)
		resultBks = append(resultBks, *bks...)
		resultQuotes = append(resultQuotes, *bkQuotes...)
	}
	for _, bk := range resultBks {
		rels := getBkRelSymbol(bk.Code)
		resultRels = append(resultRels, *rels...)
	}
	return &resultBks, &resultRels, &resultQuotes
}

// 查询板块 1:地域;2:行业;3:概念
func getBk(typeStr int) (*[]model.Bk, *[]model.BkQuote) {
	//https://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=50&pn=1&np=1&fltt=2&invt=2&fs=m%3A90+t%3A1&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152,f124,f107,f104,f105,f140,f141,f207,f208,f209,f222
	return getBkByPage(typeStr, 1, 1000)
}

// 查询板块关联个股
func getBkRelSymbol(bkCode string) *[]model.BkRelSymbol {
	// https://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=1000&pn=1&np=1&fltt=2&invt=2&ut=b2884a393a59ad64002292a3e90d46a5&fs=b%3ABK0988&fields=f12%2Cf14%2Cf2%2Cf3%2Cf62%2Cf184%2Cf66%2Cf69%2Cf72%2Cf75%2Cf78%2Cf81%2Cf84%2Cf87%2Cf204%2Cf205%2Cf124%2Cf1%2Cf13
	return getBkRelSymbolByPage(bkCode, 1, 1000)
}

func getBkByPage(typeStr int, page, size int) (*[]model.Bk, *[]model.BkQuote) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()
	url := "https://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=" + strconv.Itoa(size) + "&pn=" + strconv.Itoa(page) + "&np=1&fltt=2&invt=2&fs=m%3A90+t%3A" + strconv.Itoa(typeStr) + "&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f26,f22,f33,f11,f62,f128,f136,f115,f152,f124,f107,f104,f105,f140,f141,f207,f208,f209,f222"
	respone := visitJson(url, "")
	bks := []model.Bk{}
	bkQuotes := []model.BkQuote{}
	datas, _ := (*respone)["data"].(map[string]interface{})
	if len(datas) > 0 {
		items := datas["diff"].([]interface{})
		if len(items) > 0 {
			for i, ite := range items {
				bk := model.Bk{}
				bk.Type = typeStr
				bkQuote := model.BkQuote{}
				bkQuote.TradeDate = time.Now().Format("20060102")
				bkQuote.Rank = i + 1
				item := ite.(map[string]interface{})
				for key, value := range item {
					switch key {
					case "f12":
						code := value.(string)
						bk.Code = code
						bkQuote.BkCode = code
					case "f14":
						bk.Name = value.(string)
					case "f2":
						bkQuote.Close = value.(float64)
					case "f3":
						bkQuote.PctChg = value.(float64)
					case "f4":
						bkQuote.Change = value.(float64)
					case "f20":
						bkQuote.Total = util.FloatDiv(value.(float64), 100000000)
					case "f8":
						bkQuote.Rate = value.(float64)
					case "f104":
						bkQuote.RiseCount = int(value.(float64))
					case "f105":
						bkQuote.FallCount = int(value.(float64))
					case "f140":
						bkQuote.Lead = value.(string)
					case "f136":
						bkQuote.LeadPctChg = value.(float64)
					case "f17":
						bkQuote.Open = value.(float64)
					case "f15":
						bkQuote.High = value.(float64)
					case "f16":
						bkQuote.Low = value.(float64)
					case "f18":
						bkQuote.PreClose = value.(float64)
					case "f5":
						bkQuote.Vol = util.FloatDiv(value.(float64), 10000)
					case "f6":
						bkQuote.Amount = util.FloatDiv(value.(float64), 100000000)
					}
				}
				bks = append(bks, bk)
				bkQuotes = append(bkQuotes, bkQuote)
			}
		}
	}
	return &bks, &bkQuotes
}

func getBkRelSymbolByPage(bkCode string, page, size int) *[]model.BkRelSymbol {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()
	url := "https://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=" + strconv.Itoa(size) + "&pn=" + strconv.Itoa(page) + "&np=1&fltt=2&invt=2&ut=b2884a393a59ad64002292a3e90d46a5&fs=b%3A" + bkCode + "&fields=f12%2Cf14%2Cf2%2Cf3%2Cf62%2Cf184%2Cf66%2Cf69%2Cf72%2Cf75%2Cf78%2Cf81%2Cf84%2Cf87%2Cf204%2Cf205%2Cf124%2Cf1%2Cf13"
	respone := visitJson(url, "")
	bkRelSymbols := []model.BkRelSymbol{}
	datas, _ := (*respone)["data"].(map[string]interface{})
	if len(datas) > 0 {
		items := datas["diff"].([]interface{})
		if len(items) > 0 {
			for _, ite := range items {
				bkRelSymbol := model.BkRelSymbol{}
				bkRelSymbol.BkCode = bkCode
				item := ite.(map[string]interface{})
				for key, value := range item {
					switch key {
					case "f12":
						code := value.(string)
						bkRelSymbol.Symbol = code
					}
				}
				bkRelSymbols = append(bkRelSymbols, bkRelSymbol)
			}
		}
	}
	return &bkRelSymbols
}
