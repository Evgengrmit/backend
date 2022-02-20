package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase(uri string) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}
