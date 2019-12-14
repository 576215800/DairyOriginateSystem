package dao

import (
	"entity"
	"errors"
	"fmt"
)

func InsertLogisticesInfo(li *entity.LogisticsInfo) error {
	db, err := DbInit()
	if err != nil {
		fmt.Println(err.Error())
		errors.New("OpenDB error")
	}
	fmt.Println(li.LogisticsId, li.LogisticsFrom, li.LogisticsTo, li.LogisticsCompany, li.LogisticsMode, li.LogisticsOperator, li.LogisticsBeginTime, li.LogisticsEndTime, li.AdminOperatorId)
	conn, _ := db.Begin()
	//_, err1 := conn.Exec("INSERT INTO logistics(id, from, to, company, mode, operator, begin_time, end_time, admin_operator_id)  values(?,?,?,?,?,?,?,?,?)", li.LogisticsId, li.LogisticsFrom, li.LogisticsTo, li.LogisticsCompany, li.LogisticsMode, li.LogisticsOperator, li.LogisticsBeginTime, li.LogisticsEndTime, li.AdminOperatorId)
	_, err1 := conn.Exec("INSERT INTO logistics(id,from_place,to_place, company, mode, operator, begin_time, end_time,admin_operator_id)  values(?,?,?,?,?,?,?,?,?)", li.LogisticsId, li.LogisticsFrom, li.LogisticsTo, li.LogisticsCompany, li.LogisticsMode, li.LogisticsOperator, li.LogisticsBeginTime, li.LogisticsEndTime, li.AdminOperatorId)

	if err1 != nil {
		conn.Rollback()
		fmt.Println(err1.Error())
		return err1
	}
	conn.Commit()
	return nil
}
