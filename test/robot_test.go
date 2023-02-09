package test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/robertkrimen/otto"
)

func TestGetThsGnBySymbol(t *testing.T) {
	// reg := regexp.MustCompile(`　+.*`)
	// flag := reg.MatchString("　　一般项目：软件开发；软件的销售；电子产品销售；科普宣传服务；教学专用仪器制造；教学专用仪器销售；信息系统集成服务；光通信设备销售；光通信设备制造；通信设备制造；通信设备销售；第一类医疗器械销售；第一类医疗器械生产；第二类医疗器械销售；社会经济咨询服务；信息咨询服务；（不含许可类信息咨询服务）；电子专用设备制造；教学用模型及教具制造；教学用模型及教具销售；仪器仪表制造；技术服务；技术开发；技术咨询；技术交流；技术转让；技术推广；机械设备租赁；住房租赁；汽车新车销售；办公用品销售；建筑材料销售；电子元器件零售；电子元器件批发；计算机软硬件及辅助设备批发；计算机软硬件及辅助设备零售；（除依法须经批准的项目外")
	// log.Println(flag)
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
