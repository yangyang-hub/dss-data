package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleErrorStr(err error, format string, a ...any) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func SendGetResJson(url string) (map[string]interface{}, error) {
	res, err := http.Get(url)
	if HandleErrorStr(err, "请求失败;url: %v\n", url) {
		return nil, err
	}
	defer res.Body.Close()
	contentBytes, err := ioutil.ReadAll(res.Body)
	HandleErrorStr(err, "response读取异常;url: %v\n", url)
	var result map[string]interface{}
	err = json.Unmarshal(contentBytes, &result)
	if HandleErrorStr(err, "response返回数据json转换异常(%v)", string(contentBytes)) {
		return nil, err
	}
	return result, nil
}

func SendGetResString(url string) (string, error) {
	res, err := http.Get(url)
	if HandleErrorStr(err, "请求失败;url: %v\n", url) {
		return "", err
	}
	defer res.Body.Close()
	contentBytes, err := ioutil.ReadAll(res.Body)
	if HandleErrorStr(err, "response读取异常;url: %v\n", url) {
		return "", err
	}
	result := string(contentBytes)
	return result, nil
}
