package service

import (
	"log"
	"strings"
	"time"

	dao "dss-data/dao"
	"dss-data/robot"

	"github.com/yangyang-hub/dss-common/model"
	"github.com/yangyang-hub/dss-common/util"
)

// 爬取龙虎榜数据
func RobotLongHu(dates string) {
	start := time.Now()
	for _, date := range strings.Split(dates, ",") {
		longHu, longHuDetail := robot.GetLongHu(date)
		dao.InsertLongHu(longHu)
		dao.InsertLongHuDetail(longHuDetail)
		log.Printf("RobotLongHu %v end", date)
	}
	log.Printf("RobotLongHu end,spend time %v", time.Since(start))
}

// 查询龙虎榜数据
func GetLongHu(date string) *[]model.LongHu {
	dates := strings.Split(date, ",")
	result := dao.QueryLongHu(dates)
	return result
}

// 根据股票编码查询所属概念
func GetThsGnBySymbols(symbols string) *map[string][]string {
	params := []string{}
	for _, symbol := range strings.Split(symbols, ",") {
		code := symbol
		if len(symbol) > 6 {
			code = util.Substr(symbol, 2, 8)
		}
		params = append(params, code)
	}
	rels := dao.QueryThsGnBySymbols(params)
	result := map[string][]string{}
	for _, rel := range *rels {
		if result[rel.Symbol] == nil {
			result[rel.Symbol] = []string{rel.GnName}
		} else {
			result[rel.Symbol] = append(result[rel.Symbol], rel.GnName)
		}
	}
	return &result
}

// 根据概念名称查询所有股票
func GetSymbolsByThsGn(gns string) *map[string][]string {
	params := strings.Split(gns, ",")
	rels := dao.QuerySymbolsByThsGn(params)
	result := map[string][]string{}
	for _, rel := range *rels {
		if result[rel.GnName] == nil {
			result[rel.GnName] = []string{rel.Symbol}
		} else {
			result[rel.GnName] = append(result[rel.GnName], rel.Symbol)
		}
	}
	return &result
}

// 爬取所有股票编码查询所属概念保存到数据库
func RobotAllThsGnBySymbols() {
	start := time.Now()
	thsGns := robot.GetAllThsGn()
	dao.DeleteAllThsGn()
	dao.InsertThsGn(thsGns)
	dao.DeleteThsGnRelSymbolNotExist()
	symbols, _ := dao.GetAllSymbol()
	for _, symbol := range symbols {
		gns := robot.GetThsGnBySymbol(symbol)
		if len(*gns) > 0 {
			dao.DeleteThsGnRelBySymbol(symbol)
			for _, gn := range *gns {
				rel := model.ThsGnRelSymbol{}
				rel.Symbol = symbol
				rel.GnName = gn
				dao.InsertThsGnRelSymbol(&rel)
			}
		}
		// time.Sleep(2 * time.Second)
	}
	log.Printf("RobotAllThsGnBySymbols end,spend time %v", time.Since(start))
}
