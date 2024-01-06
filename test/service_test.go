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
func TestGetLongHuDaily(t *testing.T) {
	service.GetLongHuDaily()
}
