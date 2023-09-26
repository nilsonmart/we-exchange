package infra

import (
	"database/sql"
	"log"
	"os"
)

const fileName = "weexchange.db"

func ConnectionSQLite() *sql.DB {

	os.Remove(fileName)

	db, err := sql.Open("weexchange", fileName)
	if err != nil {
		//TODO Log error
		log.Fatal(err)
	}

	return db
}
