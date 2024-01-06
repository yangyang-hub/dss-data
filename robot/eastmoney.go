package robot

import (
	"dss-data/model"
	"strconv"
)

// 查询东财所有板块以及板块关联股票
// func GetAllBkAndRelSymbol() (*[]model.Bk, *[]model.BkRelSymbol) {
// 	for i := 1; i <= 3; i++ {

// 	}
// }

// 查询板块 1:地域;2:行业;3:概念
func getBk(typeStr int) *[]model.Bk {
	//https://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=50&pn=1&np=1&fltt=2&invt=2&fs=m%3A90+t%3A1&fields=f12%2Cf14%2Cf2%2Cf3%2Cf62%2Cf184%2Cf66%2Cf69%2Cf72%2Cf75%2Cf78%2Cf81%2Cf84%2Cf87%2Cf204%2Cf205%2Cf124%2Cf1%2Cf13
	return getBkByPage(typeStr, 1, 1000)
}

// 查询板块关联个股 1:地域;2:行业;3:概念
func getBkRelSymbol(bkCode string) *[]model.BkRelSymbol {
	// https://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=1000&pn=1&np=1&fltt=2&invt=2&ut=b2884a393a59ad64002292a3e90d46a5&fs=b%3ABK0988&fields=f12%2Cf14%2Cf2%2Cf3%2Cf62%2Cf184%2Cf66%2Cf69%2Cf72%2Cf75%2Cf78%2Cf81%2Cf84%2Cf87%2Cf204%2Cf205%2Cf124%2Cf1%2Cf13
	return getBkRelSymbolByPage(bkCode, 1, 1000)
}

func getBkByPage(typeStr int, page, size int) *[]model.Bk {
	url := "https://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=" + strconv.Itoa(size) + "&pn=" + strconv.Itoa(page) + "&np=1&fltt=2&invt=2&fs=m%3A90+t%3A" + strconv.Itoa(typeStr) + "&fields=f12%2Cf14%2Cf2%2Cf3%2Cf62%2Cf184%2Cf66%2Cf69%2Cf72%2Cf75%2Cf78%2Cf81%2Cf84%2Cf87%2Cf204%2Cf205%2Cf124%2Cf1%2Cf13"
	respone := visitJson(url, "")
	bks := []model.Bk{}
	datas, _ := (*respone)["data"].(map[string]interface{})
	if len(datas) > 0 {
		items := datas["diff"].([]interface{})
		if len(items) > 0 {
			for _, ite := range items {
				bk := model.Bk{}
				bk.Type = typeStr
				item := ite.(map[string]interface{})
				for key, value := range item {
					switch key {
					case "f12":
						code := value.(string)
						bk.Code = code
					case "f14":
						name := value.(string)
						bk.Name = name
					}
				}
				bks = append(bks, bk)
			}
		}
	}
	return &bks
}

func getBkRelSymbolByPage(bkCode string, page, size int) *[]model.BkRelSymbol {
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
