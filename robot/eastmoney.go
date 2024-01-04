package robot

// 查询板块 1:地域;2:概念;3:行业
// func getBk(typeStr string) string {
// 	//https://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=50&pn=1&np=1&fltt=2&invt=2&fs=m%3A90+t%3A1&fields=f12%2Cf14%2Cf2%2Cf3%2Cf62%2Cf184%2Cf66%2Cf69%2Cf72%2Cf75%2Cf78%2Cf81%2Cf84%2Cf87%2Cf204%2Cf205%2Cf124%2Cf1%2Cf13
// 	getBkByPage(typeStr, 1, 500)
// }
// func getBkByPage(typeStr string, page, size int) {
// 	url := "https://push2.eastmoney.com/api/qt/clist/get?fid=f62&po=1&pz=" + strconv.Itoa(size) + "&pn=" + strconv.Itoa(page) + "&np=1&fltt=2&invt=2&fs=m%3A90+t%3A" + typeStr + "&fields=f12%2Cf14%2Cf2%2Cf3%2Cf62%2Cf184%2Cf66%2Cf69%2Cf72%2Cf75%2Cf78%2Cf81%2Cf84%2Cf87%2Cf204%2Cf205%2Cf124%2Cf1%2Cf13"
// 	respone := visitJson(url, "")
// 	datas, _ := (*respone)["data"].(map[string]interface{})
// 	if len(datas) > 0 {
// 		items := datas["diff"].([]interface{})
// 		if len(items) > 0 {
// 			for _, ite := range items {
// 				item := ite.(map[string]interface{})
// 				for key, value := range item {
// 					switch key {
// 					case "f12":
// 						code = value.(string)
// 					case "f14":
// 						name = value.(string)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return ""
// }
