package service

import (
	"dao"
	"entity"
)

func AddLogisticsInfo(logisticsInfo *entity.LogisticsInfo) (bool, error) {

	err := dao.InsertLogisticesInfo(logisticsInfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
