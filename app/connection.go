package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/restful_api_ecommerce?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Duration(1) * time.Hour)
	db.SetConnMaxIdleTime(time.Duration(5) * time.Minute)
	return db
}
