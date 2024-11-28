package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_products?charset=utf8mb4&parseTime=True&loc=Local")
	// db, err := sql.Open("mysql", "root: @/go_products")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
	DB = db
}
