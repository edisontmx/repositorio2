package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	conxString := "root:@tcp(localhost:3308)/sistema"
	db, err := sql.Open("mysql", conxString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.Close()
}
