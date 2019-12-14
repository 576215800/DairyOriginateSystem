package testFile

import (
	"dao"
	"entity"
	"time"
)

func TestProducePut() {
	produceInfo := &entity.ProduceInfo{}
	produceInfo.AdminOperatorId = 1
	produceInfo.AnimalDiseaseInfo = "无病"
	produceInfo.AnimalFeedInfo = "饲料B"
	produceInfo.BrandName = "伊利"
	produceInfo.ProduceCompany = "伊利加工厂"
	produceInfo.ProduceId = 11112
	produceInfo.ProduceOperator = "王某"
	produceInfo.ProduceTime = time.Now().Format("2006-1-1 15:04:05")
	produceInfo.ProducePlace = "新疆牧场"
	dao.InsertProductInfo(produceInfo)
}

func TestProcessPut() {
	processInfo := &entity.ProcessInfo{}
	processInfo.AdminOperatorId = 1
	processInfo.ExpiryDate = "2006-1-1 15:03:04"
	processInfo.ProcessBeginTime = "2006-1-1 15:03:04"
	processInfo.ProcessEndTime = "2006-1-2 15:03:04 "
	processInfo.ProcessId = 11112
	processInfo.ProcessMode = "腌制"
	processInfo.ProcessOperator = "小张"
	processInfo.ProcessPlace = "成都菜市场"
	processInfo.ProcessCompany = "成都处理厂"
	dao.InsertProcessInfo(processInfo)
}
func TestLogisticsPut() {
	logisticsInfo := &entity.LogisticsInfo{}
	logisticsInfo.AdminOperatorId = 5
	logisticsInfo.LogisticsBeginTime = "2006-1-1 15:03:04"
	logisticsInfo.LogisticsCompany = "菜鸟驿站"
	logisticsInfo.LogisticsEndTime = "2006-1-1 15:03:04"
	logisticsInfo.LogisticsFrom = "北京"
	logisticsInfo.LogisticsTo = "上海"
	logisticsInfo.LogisticsId = 113
	logisticsInfo.LogisticsMode = "空运"
	logisticsInfo.LogisticsOperator = "小李"
	dao.InsertLogisticesInfo(logisticsInfo)
}
func TestSellPut() {
	sellInfo := &entity.SellInfo{}
	sellInfo.AdminOperatorId = 1
	sellInfo.OnShelfTime = "2006-2-4 15:04:08"
	sellInfo.SellId = 11114
	sellInfo.SellOperator = "小王"
	sellInfo.SellPlace = "电子科技大学校园超市"
	dao.InsertSellInfo(sellInfo)
}
