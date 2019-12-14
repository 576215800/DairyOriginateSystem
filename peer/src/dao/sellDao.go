package dao

import (
	"entity"
	"errors"
	"fmt"
	_ "time"
)

func InsertSellInfo(si *entity.SellInfo) error {

	db, err := DbInit()
	if err != nil {
		fmt.Println(err.Error())
		errors.New("OpenDB error")
	}
	conn, _ := db.Begin()
	_, err1 := conn.Exec("INSERT INTO product_sell(id, place,  operator,onshelf_time,admin_operator_id)  values(?,?,?,?,?)", si.SellId, si.SellPlace, si.SellOperator, si.OnShelfTime, si.AdminOperatorId)

	if err1 != nil {
		conn.Rollback()
		fmt.Println(err1.Error())
		return err1
	}
	conn.Commit()
	return nil
}
