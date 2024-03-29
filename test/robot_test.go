package test

import (
	"dss-data/robot"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"testing"

	"dss-data/util"

	"github.com/robertkrimen/otto"
	"github.com/shopspring/decimal"
)

func TestGetAllStock(t *testing.T) {
	// res := robot.GetAllStock()
	// fmt.Println(res)
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
}

func TestGetLongHu(t *testing.T) {
	res, resd := robot.GetLonghu()
	fmt.Println(*res)
	fmt.Println(*resd)
}

func TestGetTradeCal(t *testing.T) {
	res := robot.GetTradeCal()
	fmt.Println(res)
}

func TestGetThsGnBySymbol(t *testing.T) {
	v := "1.97亿"
	l := len([]rune(v))
	value := util.Substr(v, 0, (l - 2))
	f, _ := strconv.ParseFloat(value, 64)
	d := decimal.NewFromFloat(f)
	unit := util.Substr(v, l-1, l)
	var result float64 = f
	if unit == "亿" {
		result, _ = d.Mul(decimal.NewFromFloat(10000)).Round(2).Float64()
	}
	log.Println(result)
}

func TestGetGn(t *testing.T) {
	// result := robot.GetAllThsGn()
	// fmt.Println(*result)
}
func TestGetThsCookie(t *testing.T) {
	filePath := "../robot/js/ths_cookie_v.js"
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file error")
		panic(err)
	}
	vm := otto.New()
	str := string(bytes)
	_, err = vm.Run(str)
	if err != nil {
		fmt.Println("run js error")
		panic(err)
	}
	value, err := vm.Call("get_v", nil)
	if err != nil {
		fmt.Println("run func error")
		panic(err)
	}
	fmt.Println(value.String())
}
