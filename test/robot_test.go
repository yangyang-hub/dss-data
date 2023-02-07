package test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/robertkrimen/otto"
)

func TestGetThsGnBySymbol(t *testing.T) {
	// result := robot.GetThsGnBySymbol("300071")
	// log.Println(result)
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
