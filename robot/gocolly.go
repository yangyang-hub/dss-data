package robot

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/robertkrimen/otto"
	"github.com/yangyang-hub/dss-common/model"
	"github.com/yangyang-hub/dss-common/util"
)

/**
同花顺行业 	start==========================================================================================================================
*/

//获取所有同花顺行业
func GetAllThsHy() *[]model.ThsHy {
	thsHys := []model.ThsHy{}
	url := "http://q.10jqka.com.cn/thshy/"
	visit(url, "div[class='cate_group'] > div[class='cate_items'] > a", func(e *colly.HTMLElement) {
		// 获取概念连接
		link := e.Attr("href")
		code := util.Substr(link, len(link)-7, 6)
		name := e.Text
		thsHys = append(thsHys, model.ThsHy{Code: code, Name: name})
	})
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
	visit(url, "div[class='m-pager'] > a[class='changePage']", func(e *colly.HTMLElement) {
		if e.Text == "尾页" {
			totalPage, _ = strconv.Atoi(e.Attr("page"))
			log.Printf("total page found: %v\n", totalPage)
		}
	})
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
	//获取当页的股票代码
	visit(url, "table[class='m-table m-pager-table'] > tbody > tr", func(e *colly.HTMLElement) {
		symbol := e.ChildText("td:nth-child(2)")
		if symbol != "" {
			thsHyRelSymbol = append(thsHyRelSymbol, model.ThsHyRelSymbol{HyCode: code, Symbol: symbol})
		}
	})
	return &thsHyRelSymbol
}

//获取当日所有同花顺行业的行情信息
func GetAllThsHyQuote(thsHys *[]model.ThsHy) *[]model.ThsHyQuote {
	thsHyQuote := []model.ThsHyQuote{}
	//迭代获取行业行情
	if len(*thsHys) > 0 {
		for index, thsHy := range *thsHys {
			log.Printf("start get hy quote %v %v %v", index, thsHy.Name, thsHy.Code)
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
	visit(url, "div[class='board-infos'] > dl", func(e *colly.HTMLElement) {
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
	thsGns := []model.ThsGn{}
	url := "http://q.10jqka.com.cn/gn/"
	visit(url, "div[class='cate_group'] > div[class='cate_items'] > a", func(e *colly.HTMLElement) {
		// 获取概念连接
		link := e.Attr("href")
		code := util.Substr(link, len(link)-7, 6)
		name := e.Text
		thsGns = append(thsGns, model.ThsGn{Code: code, Name: name})
	})
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
	visit(url, "div[class='m-pager'] > a[class='changePage']", func(e *colly.HTMLElement) {
		if e.Text == "尾页" {
			totalPage, _ = strconv.Atoi(e.Attr("page"))
			log.Printf("total page found: %v\n", totalPage)
		}
	})
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
	//获取当页的股票代码
	visit(url, "table[class='m-table m-pager-table'] > tbody > tr", func(e *colly.HTMLElement) {
		symbol := e.ChildText("td:nth-child(2)")
		if symbol != "" {
			thsGnRelSymbol = append(thsGnRelSymbol, model.ThsGnRelSymbol{GnCode: code, Symbol: symbol})
		}
	})
	return &thsGnRelSymbol
}

//获取当日所有同花顺概念的行情信息
func GetAllThsGnQuote(thsGns *[]model.ThsGn) *[]model.ThsGnQuote {
	thsGnQuote := []model.ThsGnQuote{}
	//迭代获取概念当日行情
	if len(*thsGns) > 0 {
		for index, thsGn := range *thsGns {
			log.Printf("start get gn quote %v %v %v", index, thsGn.Name, thsGn.Code)
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
	visit(url, "div[class='board-infos'] > dl", func(e *colly.HTMLElement) {
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
	return v
}

func visit(url string, goquerySelector string, f colly.HTMLCallback) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	proxy := getProxy()
	c.SetProxy(proxy)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Cookie", getThsCookie())
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Something went wrong: %v, Proxy Address: %v\n", err, proxy)
	})
	c.OnHTML(goquerySelector, f)
	err := c.Visit(url)
	if err != nil {
		//重试 并删除无效代理
		deleteProxy(proxy)
		visit(url, goquerySelector, f)
	}
	// 采集等待结束
	c.Wait()
}

func getProxy() string {
	result, err := util.SendGet("http://127.0.0.1:5010/get")
	if err != nil {
		log.Println("get proxy wrong:", err)
	}
	var proxyUrl string
	b, _ := strconv.ParseBool(fmt.Sprint(result["https"]))
	if bool(b) {
		proxyUrl = fmt.Sprintf("https://%v", fmt.Sprint(result["proxy"]))
	} else {
		proxyUrl = fmt.Sprintf("http://%v", fmt.Sprint(result["proxy"]))
	}
	return proxyUrl
}

func deleteProxy(proxyUrl string) {
	proxy := strings.Split(proxyUrl, "//")[1]
	log.Println("delete proxy", proxy)
	util.SendGet(fmt.Sprintf("http://127.0.0.1:5010/delete/?proxy=%v", proxy))
}
