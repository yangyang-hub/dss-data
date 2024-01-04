package util

import (
	"strconv"
)

// 计算涨停板
func CalLimitUp(code string, preClose, change float64, stStock map[string]string) int64 {
	prefix := Substr(code, 0, 2)
	var divValue float64 = 10
	if stStock[code] != "" {
		divValue = 0.05
	} else if prefix == "60" {
		divValue = 0.1
	} else if prefix == "00" {
		divValue = 0.1
	} else if prefix == "68" {
		divValue = 0.2
	} else if prefix == "30" {
		divValue = 0.2
	} else if Substr(code, 0, 1) == "8" {
		divValue = 0.3
	}
	limit := FloatMul(preClose, divValue)
	res := FloatCmp(change, limit)
	if res >= 0 {
		return 1
	} else {
		return 0
	}
}

//单位转换为万
func UnitConversion(v string) float64 {
	l := len([]rune(v))
	value := Substr(v, 0, (l - 2))
	f, _ := strconv.ParseFloat(value, 64)
	unit := Substr(v, l-1, l)
	var result float64 = f
	if unit == "亿" {
		result = FloatMul(f, 10000)
	}
	return result
}
