package service

//实时数据
type LiveData struct {
	Code          string `json:"code"`          //编码
	Name          string `json:"name"`          //名称
	Now           string `json:"now"`           //当前价
	Change        string `json:"change"`        //涨跌
	ChangePercent string `json:"changePercent"` //涨跌幅
	Time          string `json:"time"`          //更新时间
	Max           string `json:"max"`           //最高价
	Min           string `json:"min"`           //最低价
}
