package util

import (
	"github.com/shopspring/decimal"
)

// 乘
func FloatMul(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2)).Round(2).Float64()
	return res
}

// 除
func FloatDiv(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Div(decimal.NewFromFloat(f2)).Round(2).Float64()
	return res
}

// 比较大小
func FloatCmp(f1, f2 float64) int {
	res := decimal.NewFromFloat(f1).Cmp(decimal.NewFromFloat(f2))
	return res
}

func FloatRound(f float64, r int32) float64 {
	res, _ := decimal.NewFromFloat(f).Round(r).Float64()
	return res
}
