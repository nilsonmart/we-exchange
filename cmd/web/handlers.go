package main

import (
	"database/sql/driver"
	"log"

	"github.com/nilsonmart/we-exchange/internal/models"
	"github.com/nilsonmart/we-exchange/internal/repository"
)

var conn = driver.ConnectionSQLite()
var repo = repository.NewSQLiteRepository()

func getSchema() []models.Schema {
	result, err := repo.AllSchema()
	if err != nil {
		//TODO Log error
		log.Fatal(err)
	}
	return result
}

func getActivities() {

}
