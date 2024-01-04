package test

import (
	"dss-data/service"
	"fmt"
	"testing"
)

func TestGetLiveData(t *testing.T) {
	param := []string{"sh603232"}
	res, _ := service.GetLiveData(param)
	fmt.Println(res)
}
func TestGetDailyData(t *testing.T) {
	service.GetDailyData("20230101")
}
func TestGetLongHuDaily(t *testing.T) {
	service.GetLongHuDaily()
}
