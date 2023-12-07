package robot

// 查询当天是否为交易日 1:是;0:不是
func GetTradeCal() string {
	respone := visitJson("http://www.szse.cn/api/report/exchange/onepersistenthour/monthList", "")
	datas, _ := (*respone)["data"].([]interface{})
	nowdate, _ := (*respone)["nowdate"].(string)
	if len(datas) > 0 {
		for _, data := range datas {
			items := data.(map[string]interface{})
			jyrq := ""
			jybz := ""
			for key, value := range items {
				switch key {
				case "jyrq":
					jyrq = value.(string)
				case "jybz":
					jybz = value.(string)
				}
			}
			if jyrq == nowdate {
				return jybz
			}
		}
	}
	return ""
}
