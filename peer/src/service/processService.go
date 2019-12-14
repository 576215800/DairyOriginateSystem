package service

import (
	"dao"
	"entity"
)

func AddProcessInfo(processInfo *entity.ProcessInfo) (bool, error) {

	err := dao.InsertProcessInfo(processInfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
