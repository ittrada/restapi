package configs

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/goblog")

	if err != nil {
		log.Fatal("Could not connect to database")
	}

	return db
}
