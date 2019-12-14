package dao

import (
	"entity"
	"errors"
	"fmt"
	_ "time"
)

func InsertProductInfo(pi *entity.ProduceInfo) error {

	db, err := DbInit()
	if err != nil {
		fmt.Println(err.Error())
		errors.New("OpenDB error")
	}
	conn, _ := db.Begin()
	_, err1 := conn.Exec("INSERT INTO produce(id, place, produce_time, operator, company, animal_feed_info, animal_disease_info, brand_name, admin_operator_id)  values(?,?,?,?,?,?,?,?,?)", pi.ProduceId, pi.ProducePlace, pi.ProduceTime, pi.ProduceOperator, pi.ProduceCompany, pi.AnimalFeedInfo, pi.AnimalDiseaseInfo, pi.BrandName, pi.AdminOperatorId)

	if err1 != nil {
		conn.Rollback()
		fmt.Println(err1.Error())
		return err1
	}
	conn.Commit()
	return nil
}
