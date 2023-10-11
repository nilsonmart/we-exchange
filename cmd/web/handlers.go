package main

import (
	"log"

	"github.com/nilsonmart/we-exchange/internal/infra"
	"github.com/nilsonmart/we-exchange/internal/models"
	"github.com/nilsonmart/we-exchange/internal/repository"
)

var conn = infra.ConnectionSQLite()
var repo = repository.NewSQLiteRepository(conn)

func ValidateAccount(email, password string) (bool, error) {
	result, err := repo.ValidateAccount(email, password)
	if err != nil {
		return false, err
	}

	return result, err
}

func GetAccountByEmail(email string) *models.Account {
	result, err := repo.GetAccountByEmail(email)
	if err != nil {
		//TODO Log error
		log.Fatal(err)
	}

	return result
}

func getSchema() []models.Schema {
	result, err := repo.AllSchema()
	if err != nil {
		//TODO Log error
		log.Fatal(err)
	}
	return result
}
