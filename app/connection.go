package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	db,err := sql.Open("mysql", "root:root@tcp(localhost:3306)/restful_api_ecommerce?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
