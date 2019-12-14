package dao

import (
	"entity"
	"errors"
	"fmt"
)

func InsertProcessInfo(pi *entity.ProcessInfo) error {

	db, err := DbInit()
	if err != nil {
		fmt.Println(err.Error())
		errors.New("OpenDB error")
	}
	conn, _ := db.Begin()
	_, err1 := conn.Exec("INSERT INTO process(id, place, begin_time, end_time, operator,mode,company,expiry_date, admin_operator_id)  values(?,?,?,?,?,?,?,?,?)", pi.ProcessId, pi.ProcessPlace, pi.ProcessBeginTime, pi.ProcessEndTime, pi.ProcessOperator, pi.ProcessMode, pi.ProcessCompany, pi.ExpiryDate, pi.AdminOperatorId)

	if err1 != nil {
		conn.Rollback()
		fmt.Println(err1.Error())
		return err1
	}
	conn.Commit()
	return nil
}
