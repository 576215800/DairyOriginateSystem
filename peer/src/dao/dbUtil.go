package dao

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// var db *sql.DB

func DbInit() (s *sql.DB, err error) {

	//db, err := sql.Open("mysql", "root:root@tcp(192.168.56.101:3306)/traceability?charset=utf8")
	db, err := sql.Open("mysql", "root:030511@tcp(39.108.208.124:3306)/traceability?charset=utf8")

	return db, err
}
