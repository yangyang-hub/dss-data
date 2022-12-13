package robot

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/robertkrimen/otto"
	"github.com/yangyang-hub/dss-common/model"
	"github.com/yangyang-hub/dss-common/util"
)

var proxyUrls []string
var timeStamp int64

// 爬取时间间隔（毫秒）
var step time.Duration = 500

// 上次爬取请求的时间戳
var preStepTime int64

/**
同花顺行业 	start==========================================================================================================================
*/

//获取所有同花顺行业
func GetAllThsHy() *[]model.ThsHy {
	c := getCollyCollector()
	thsHys := []model.ThsHy{}
	url := "http://q.10jqka.com.cn/thshy/"
	c.OnHTML("div[class='cate_group'] > div[class='cate_items'] > a", func(e *colly.HTMLElement) {
		// 获取概念连接
		link := e.Attr("href")
		code := util.Substr(link, len(link)-7, 6)
		name := e.Text
		thsHys = append(thsHys, model.ThsHy{Code: code, Name: name})
	})
	err := collyVisit(c, url)
	if err != nil {
		log.Println("err", err)
	}
	log.Printf("hy found size: %v\n", len(thsHys))
	return &thsHys
}

//获取所有同花顺行业所关联的股票代码
func GetAllThsHyRelSymbol(thsHys *[]model.ThsHy) *[]model.ThsHyRelSymbol {
	thsHyRelSymbol := []model.ThsHyRelSymbol{}
	if len(*thsHys) > 0 {
		for index, thsHy := range *thsHys {
			log.Printf("start get hy %v %v %v", index, thsHy.Name, thsHy.Code)
			thsHyRelSymbol = append(thsHyRelSymbol, *GetThsHyDetail(thsHy.Code)...)
		}
	}
	return &thsHyRelSymbol
}

//获取单个同花顺行业所关联的股票代码
func GetThsHyDetail(code string) *[]model.ThsHyRelSymbol {
	thsHyRelSymbol := []model.ThsHyRelSymbol{}
	//获取总页码
	totalPage := 0
	url := "http://q.10jqka.com.cn/thshy/detail/field/199112/order/desc/page/2/ajax/1/code/" + code
	c := getCollyCollector()
	c.OnHTML("div[class='m-pager'] > a[class='changePage']", func(e *colly.HTMLElement) {
		if e.Text == "尾页" {
			totalPage, _ = strconv.Atoi(e.Attr("page"))
			log.Printf("total page found: %v\n", totalPage)
		}
	})
	err := collyVisit(c, url)
	if err != nil {
		log.Println("err", err)
	}
	//循环获取概念所有代码
	if totalPage > 0 {
		for i := 1; i <= totalPage; i++ {
			thsHyRelSymbol = append(thsHyRelSymbol, *getThsHyDetailByPage(code, i)...)
		}
	}
	return &thsHyRelSymbol
}

//读取单页同花顺行业所关联的股票代码
func getThsHyDetailByPage(code string, page int) *[]model.ThsHyRelSymbol {
	thsHyRelSymbol := []model.ThsHyRelSymbol{}
	url := "http://q.10jqka.com.cn/thshy/detail/field/199112/order/desc/page/" + strconv.Itoa(page) + "/ajax/1/code/" + code
	c := getCollyCollector()
	//获取当页的股票代码
	c.OnHTML("table[class='m-table m-pager-table'] > tbody > tr", func(e *colly.HTMLElement) {
		symbol := e.ChildText("td:nth-child(2)")
		if symbol != "" {
			thsHyRelSymbol = append(thsHyRelSymbol, model.ThsHyRelSymbol{HyCode: code, Symbol: symbol})
		}
	})
	err := collyVisit(c, url)
	if err != nil {
		log.Println("err", err)
	}
	return &thsHyRelSymbol
}

//获取当日所有同花顺行业的行情信息
func GetAllThsHyQuote(thsHys *[]model.ThsHy) *[]model.ThsHyQuote {
	thsHyQuote := []model.ThsHyQuote{}
	//迭代获取行业行情
	if len(*thsHys) > 0 {
		for index, thsHy := range *thsHys {
			log.Printf("start get hy %v %v %v", index, thsHy.Name, thsHy.Code)
			thsHyQuote = append(thsHyQuote, *GetThsHyQuote(thsHy.Code))
		}
	}
	return &thsHyQuote
}

//获取单个当日同花顺行业的行情信息
func GetThsHyQuote(code string) *model.ThsHyQuote {
	thsHyQuote := model.ThsHyQuote{}
	thsHyQuote.Code = code
	thsHyQuote.TradeDate = time.Now().Format("20060102")
	url := "http://q.10jqka.com.cn/thshy/detail/code/" + code
	c := getCollyCollector()
	c.OnHTML("div[class='board-infos'] > dl", func(e *colly.HTMLElement) {
		title := e.ChildText("dt")
		value := e.ChildText("dd")
		switch title {
		case "今开":
			data, _ := strconv.ParseFloat(value, 64)
			thsHyQuote.Open = data
		case "昨收":
			data, _ := strconv.ParseFloat(value, 64)
			thsHyQuote.PreClose = data
		case "最低":
			data, _ := strconv.ParseFloat(value, 64)
			thsHyQuote.Low = data
		case "最高":
			data, _ := strconv.ParseFloat(value, 64)
			thsHyQuote.High = data
		case "成交量(万手)":
			data, _ := strconv.ParseFloat(value, 64)
			thsHyQuote.Vol = data
		case "板块涨幅":
			data, _ := strconv.ParseFloat(util.Substr(value, 0, len(value)-1), 64)
			thsHyQuote.PctChg = data
		case "涨幅排名":
			data, _ := strconv.Atoi(strings.Split(value, "/")[0])
			thsHyQuote.Rank = data
		case "涨跌家数":
			rise, _ := strconv.Atoi(e.ChildText("dd > span[class='arr-rise-s']"))
			fall, _ := strconv.Atoi(e.ChildText("dd > span[class='arr-fall-s']"))
			thsHyQuote.RiseCount = rise
			thsHyQuote.FallCount = fall
		case "资金净流入(亿)":
			data, _ := strconv.ParseFloat(value, 64)
			thsHyQuote.Change = data
		case "成交额(亿)":
			data, _ := strconv.ParseFloat(value, 64)
			thsHyQuote.Amount = data
		}

	})
	err := collyVisit(c, url)
	if err != nil {
		log.Println("err", err)
	}
	return &thsHyQuote
}

/**
同花顺行业 	end==========================================================================================================================
*/

/**
同花顺概念 	start==========================================================================================================================
*/

//获取所有同花顺概念
func GetAllThsGn() *[]model.ThsGn {
	c := getCollyCollector()
	thsGns := []model.ThsGn{}
	url := "http://q.10jqka.com.cn/gn/"
	c.OnHTML("div[class='cate_group'] > div[class='cate_items'] > a", func(e *colly.HTMLElement) {
		// 获取概念连接
		link := e.Attr("href")
		code := util.Substr(link, len(link)-7, 6)
		name := e.Text
		thsGns = append(thsGns, model.ThsGn{Code: code, Name: name})
	})
	err := collyVisit(c, url)
	if err != nil {
		log.Println("err", err)
	}
	log.Printf("gn found size: %v\n", len(thsGns))
	return &thsGns
}

//获取所有同花顺概念所关联的股票代码
func GetAllThsGnRelSymbol(thsGns *[]model.ThsGn) *[]model.ThsGnRelSymbol {
	thsGnRelSymbol := []model.ThsGnRelSymbol{}
	if len(*thsGns) > 0 {
		for index, thsGn := range *thsGns {
			log.Printf("start get gn %v %v %v total %v", index, thsGn.Name, thsGn.Code, len(*thsGns))
			thsGnRelSymbol = append(thsGnRelSymbol, *GetThsGnDetail(thsGn.Code)...)
		}
	}
	return &thsGnRelSymbol
}

//获取单个同花顺概念所关联的股票代码
func GetThsGnDetail(code string) *[]model.ThsGnRelSymbol {
	thsGnRelSymbol := []model.ThsGnRelSymbol{}
	//获取总页码
	totalPage := 0
	url := "http://q.10jqka.com.cn/gn/detail/field/199112/order/desc/page/1/ajax/1/code/" + code
	c := getCollyCollector()
	c.OnHTML("div[class='m-pager'] > a[class='changePage']", func(e *colly.HTMLElement) {
		if e.Text == "尾页" {
			totalPage, _ = strconv.Atoi(e.Attr("page"))
			log.Printf("total page found: %v\n", totalPage)
		}
	})
	err := collyVisit(c, url)
	if err != nil {
		log.Println("err", err)
	}
	//循环获取概念所有代码
	if totalPage > 0 {
		for i := 1; i <= totalPage; i++ {
			thsGnRelSymbol = append(thsGnRelSymbol, *getThsGnDetailByPage(code, i)...)
		}
	}
	return &thsGnRelSymbol
}

//读取单页同花顺概念所关联的股票代码
func getThsGnDetailByPage(code string, page int) *[]model.ThsGnRelSymbol {
	thsGnRelSymbol := []model.ThsGnRelSymbol{}
	url := "http://q.10jqka.com.cn/gn/detail/field/199112/order/desc/page/" + strconv.Itoa(page) + "/ajax/1/code/" + code
	c := getCollyCollector()
	//获取当页的股票代码
	c.OnHTML("table[class='m-table m-pager-table'] > tbody > tr", func(e *colly.HTMLElement) {
		symbol := e.ChildText("td:nth-child(2)")
		if symbol != "" {
			thsGnRelSymbol = append(thsGnRelSymbol, model.ThsGnRelSymbol{GnCode: code, Symbol: symbol})
		}
	})
	err := collyVisit(c, url)
	if err != nil {
		log.Println("err", err)
	}
	return &thsGnRelSymbol
}

//获取当日所有同花顺概念的行情信息
func GetAllThsGnQuote(thsGns *[]model.ThsGn) *[]model.ThsGnQuote {
	thsGnQuote := []model.ThsGnQuote{}
	//迭代获取概念当日行情
	if len(*thsGns) > 0 {
		for index, thsGn := range *thsGns {
			log.Printf("start get hy %v %v %v", index, thsGn.Name, thsGn.Code)
			thsGnQuote = append(thsGnQuote, *GetThsGnQuote(thsGn.Code))
		}
	}
	return &thsGnQuote
}

//获取单个当日同花顺概念的行情信息
func GetThsGnQuote(code string) *model.ThsGnQuote {
	thsGnQuote := model.ThsGnQuote{}
	thsGnQuote.Code = code
	thsGnQuote.TradeDate = time.Now().Format("20060102")
	url := "http://q.10jqka.com.cn/gn/detail/code/" + code
	c := getCollyCollector()
	c.OnHTML("div[class='board-infos'] > dl", func(e *colly.HTMLElement) {
		title := e.ChildText("dt")
		value := e.ChildText("dd")
		switch title {
		case "今开":
			data, _ := strconv.ParseFloat(value, 64)
			thsGnQuote.Open = data
		case "昨收":
			data, _ := strconv.ParseFloat(value, 64)
			thsGnQuote.PreClose = data
		case "最低":
			data, _ := strconv.ParseFloat(value, 64)
			thsGnQuote.Low = data
		case "最高":
			data, _ := strconv.ParseFloat(value, 64)
			thsGnQuote.High = data
		case "成交量(万手)":
			data, _ := strconv.ParseFloat(value, 64)
			thsGnQuote.Vol = data
		case "板块涨幅":
			data, _ := strconv.ParseFloat(util.Substr(value, 0, len(value)-1), 64)
			thsGnQuote.PctChg = data
		case "涨幅排名":
			data, _ := strconv.Atoi(strings.Split(value, "/")[0])
			thsGnQuote.Rank = data
		case "涨跌家数":
			rise, _ := strconv.Atoi(e.ChildText("dd > span[class='arr-rise-s']"))
			fall, _ := strconv.Atoi(e.ChildText("dd > span[class='arr-fall-s']"))
			thsGnQuote.RiseCount = rise
			thsGnQuote.FallCount = fall
		case "资金净流入(亿)":
			data, _ := strconv.ParseFloat(value, 64)
			thsGnQuote.Change = data
		case "成交额(亿)":
			data, _ := strconv.ParseFloat(value, 64)
			thsGnQuote.Amount = data
		}

	})
	err := collyVisit(c, url)
	if err != nil {
		log.Println("err", err)
	}
	return &thsGnQuote
}

/**
同花顺概念 	end==========================================================================================================================
*/

//生成同花顺cookie中的v值
func getThsCookie() string {
	filePath := "./robot/js/ths_cookie_v.js"
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("read file error")
		panic(err)
	}
	vm := otto.New()
	str := string(bytes)
	_, err = vm.Run(str)
	if err != nil {
		log.Println("run js error")
		panic(err)
	}
	value, err := vm.Call("get_v", nil)
	if err != nil {
		log.Println("run func get_v error")
		panic(err)
	}
	v := fmt.Sprintf("v=%v", value.String())
	// log.Printf("success get cookie %v", v)
	return v
}

//func init() {
//	GetProxy()
//}

//获取代理ip
func GetProxy() []string {
	//查询时间戳，设置代理过期时间为一个小时
	currentTime := time.Now().Unix() //秒为单位的时间戳
	if (timeStamp + 3600) > currentTime {
		return proxyUrls
	}
	timeStamp = currentTime
	//重新设置代理池
	//默认获取100个代理ip
	newProxyUrls := []string{}
	for i := 1; i <= 1; i++ {
		c := colly.NewCollector()
		url := "https://www.kuaidaili.com/free/inha/" + strconv.Itoa(i) + "/"
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})
		c.OnHTML("table > tbody > tr", func(e *colly.HTMLElement) {
			ip := e.ChildText("td[data-title='IP']")
			port := e.ChildText("td[data-title='PORT']")
			proxyUrl := fmt.Sprintf("//%v:%v", ip, port)
			newProxyUrls = append(newProxyUrls, proxyUrl)
		})
		c.OnScraped(func(r *colly.Response) {
			log.Println("Finished", r.Request.URL)
		})
		err := collyVisit(c, url)
		if err != nil {
			log.Println("err", err)
		}
	}
	proxyUrls = newProxyUrls
	return proxyUrls
}

func getCollyCollector() *colly.Collector {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	//proxyUrls := GetProxy()
	//rp, err := proxy.RoundRobinProxySwitcher(proxyUrls...)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////设置代理IP使用轮询ip方式
	//c.SetProxyFunc(rp)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Cookie", getThsCookie())
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	// c.OnScraped(func(r *colly.Response) {
	// 	log.Println("Finished", r.Request.URL)
	// 	//log.Printf("Proxy Address: %s\n", r.Request.ProxyURL)
	// })
	return c
}

func collyVisit(c *colly.Collector, url string) error {
	stepSleep()
	err := c.Visit(url)
	if err != nil {
		return err
	}
	// 采集等待结束
	c.Wait()
	return nil
}

func stepSleep() {
	//查询时间戳 保持每次爬取的间隔大于 step
	currentTime := time.Now().Unix() //秒为单位的时间戳
	if (preStepTime + int64(step)) < currentTime {
		preStepTime = time.Now().Unix()
		return
	}
	time.Sleep(step * time.Millisecond)
	preStepTime = time.Now().Unix()
}
