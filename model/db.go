package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Setup() {
	dsn := "host=containers-us-west-129.railway.app port=6141 user=postgres password=y1hzTmtClZ3eE6HJDG5O dbname=railway sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}
