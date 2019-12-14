package service

import (
	"dao"
	"entity"
)

func AddProduceInfo(produceInfo *entity.ProduceInfo) (bool, error) {

	err := dao.InsertProductInfo(produceInfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
