package robot

import (
	configs "dss-data/configs"
	"dss-data/util"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// 访问url返回json
func visitJson(url, cookie string) *map[string]interface{} {
	var resp map[string]interface{}
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	// proxy := getProxy()
	// c.SetProxy(proxy)
	if cookie != "" {
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Add("Cookie", cookie)
		})
	}
	c.OnResponse(func(r *colly.Response) {
		// 判断响应的内容类型是否为JSON
		if r.Headers.Get("Content-Type") == "application/json;charset=UTF-8" {
			err := json.Unmarshal(r.Body, &resp)
			if err != nil {
				log.Println("visitXueQiuJson 解析JSON失败:", err)
				return
			}
		}
	})
	err := c.Visit(url)
	c.Wait()
	if err != nil {
		//重试 并删除无效代理
		// deleteProxy(proxy)
		visitXueQiuJson(url)
	}
	return &resp
}

// 访问url返回html
func visitHtml(url, cookie, goquerySelector string, f colly.HTMLCallback) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	proxy := getProxy()
	c.SetProxy(proxy)
	if cookie != "" {
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Add("Cookie", cookie)
		})
	}
	c.OnHTML(goquerySelector, f)
	err := c.Visit(url)
	c.Wait()
	if err != nil {
		//重试 并删除无效代理
		deleteProxy(proxy)
		visit(url, goquerySelector, f)
	}
}
func visit(url string, goquerySelector string, f colly.HTMLCallback) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"
	proxy := getProxy()
	c.SetProxy(proxy)
	// c.OnError(func(r *colly.Response, err error) {
	// 	log.Printf("Something went wrong: %v, Proxy Address: %v\n", err, proxy)
	// })
	c.OnHTML(goquerySelector, f)
	err := c.Visit(url)
	c.Wait()
	if err != nil {
		//重试 并删除无效代理
		deleteProxy(proxy)
		visit(url, goquerySelector, f)
	}
}

func getProxy() string {
	result, err := util.SendGetResJson(configs.Config.ProxyUrl + "/get/")
	// result, err := util.SendGetResJson("http://127.0.0.1:5010/get/")
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
	// log.Println("delete proxy", proxy)
	util.SendGetResJson(fmt.Sprintf(configs.Config.ProxyUrl+"/delete/?proxy=%v", proxy))
	// util.SendGetResJson(fmt.Sprintf("http://127.0.0.1:5010/delete/?proxy=%v", proxy))
}
