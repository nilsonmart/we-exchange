package infra

import (
	"database/sql"
	"fmt"
	"log"
)

const fileName = "weexchange.db"

func ConnectionSQLite() *sql.DB {

	//os.Remove(fileName)

	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		//TODO Log error
		log.Fatal(err)
	}

	//defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)

	return db
}
