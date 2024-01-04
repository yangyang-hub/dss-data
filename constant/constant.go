package constant

//批量新增每批次数量
const InsertBatchSize int = 1000

//交易所编码枚举
type Exchange int

const ExchangeConst Exchange = 0

func (n Exchange) String() string {
	return [...]string{"SSE", "SZSE", "BSE"}[n]
}

func (n Exchange) List() []string {
	return []string{"SSE", "SZSE", "BSE"}
}

func (n Exchange) Exclude(i Exchange) []string {
	exchange := ExchangeConst.List()
	exchange = append(exchange[:i], exchange[i+1:]...)
	return exchange
}

const (
	SSE  Exchange = iota //上交所
	SZSE                 //深交所
	BSE                  //北交所
)

// 时间格式化
const (
	TimeFormatA string = "20060102"
	TimeFormatB string = "2006-01-02"
)
