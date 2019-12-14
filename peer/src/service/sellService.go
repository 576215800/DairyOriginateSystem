package service

import (
	"dao"
	"entity"
)

func AddSellInfo(sellInfo *entity.SellInfo) (bool, error) {

	err := dao.InsertSellInfo(sellInfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
