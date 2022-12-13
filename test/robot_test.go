package test

import (
	robot "dss-base-data/robot"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/robertkrimen/otto"
)

func TestGetGn(t *testing.T) {
	robot.GetAllThsGn()
}
func TestGetGnDetail(t *testing.T) {
	robot.GetThsGnDetail("301558")
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
