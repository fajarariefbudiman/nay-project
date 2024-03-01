package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root:annaqt2006@tcp(localhost:3306)/jar")
	if err != nil {
		log.Println("Connection to Database Error")
		panic(err)
	}

	return db
}
