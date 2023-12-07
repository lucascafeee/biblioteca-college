package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:postgres@tcp(localhost:3306)/biblioteca")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, nil
}
