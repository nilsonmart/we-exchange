package main

import (
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/nilsonmart/we-exchange/internal/models"
	"github.com/nilsonmart/we-exchange/internal/repository"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

// Account
func AllAccount() ([]models.Account, error) {

	model, err := repository.AllAccount()
	if err != nil {
		//TODO Log Error
		return nil, ErrNotExists
	}

	return model, nil
}

func GetAccountByID(id primitive.ObjectID) (*models.Account, error) {

	model, err := repository.GetAccountByID(id)
	if err != nil {
		//TODO Log Error
		return nil, ErrNotExists
	}

	return model, nil
}

func ValidateAccount(email, password string) (bool, error) {
	if email == "" || password == "" {
		//TODO Log error
		log.Fatal("Email or Password invalid.")
		return false, errors.New("Email or Password invalid.")
	}

	isValid, err := repository.ValidateAccount(email, password)

	if err != nil {
		return false, err
	}

	if isValid {
		return true, nil
	}

	return false, ErrNotExists
}

// RequestChange

func CreateRequestChange(model models.RequestChange) (bool, error) {

	result, err := repository.CreateRequestChange(model)

	if err != nil {
		//TODO Log error
		log.Fatal(err)
		return false, err
	}

	return result, nil
}

func AllRequestChange() ([]models.RequestChange, error) {
	model, err := repository.AllRequestChange()
	if err != nil {
		//TODO Log Error
		return nil, ErrNotExists
	}

	return model, nil
}

func GetRequestChangeByID(id primitive.ObjectID) (*models.RequestChange, error) {
	query, err := repository.GetRequestChangeByID(id)
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	return query, nil
}

func GetRequestChangeByUserID(userId string) (*models.RequestChange, error) {

	query, err := repository.GetRequestChangeByUserID(userId)

	if err != nil {
		return nil, err
	}

	return query, nil
}

func UpdateRequestChange(id primitive.ObjectID, model models.RequestChange) (bool, error) {
	if id == primitive.NilObjectID {
		return false, errors.New("invalid updated ID")
	}

	result, err := repository.UpdateRequestChange(id, model)

	if err != nil {
		return false, err
	}

	return result, ErrNotExists
}

func DeleteRequestChange(id primitive.ObjectID) (bool, error) {
	if id == primitive.NilObjectID {
		return false, errors.New("invalid updated ID")
	}

	result, err := repository.DeleteRequestChange(id)
	if err != nil {
		//TODO Log error
		log.Fatal(err)
		return false, err
	}

	return result, nil
}

// SCHEMA

func CreateSchema(model models.Schema) (bool, error) {

	result, err := repository.CreateSchema(model)

	if err != nil {
		//TODO Log error
		log.Fatal(err)
		return false, err
	}

	return result, nil

}

func AllSchema() ([]models.Schema, error) {
	query, err := repository.AllSchema()
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	return query, nil
}

func GetSchemaByID(id primitive.ObjectID) (*models.Schema, error) {
	query, err := repository.GetSchemaByID(id)
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	return query, nil
}

func GetSchemaByUserID(id primitive.ObjectID) (*models.Schema, error) {
	query, err := repository.GetSchemaByUserID(id)
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	return query, nil
}

func UpdateSchema(id primitive.ObjectID, model models.Schema) (bool, error) {
	if id == primitive.NilObjectID {
		return false, errors.New("invalid updated ID")
	}

	result, err := repository.UpdateSchema(id, model)

	if err != nil {
		return false, err
	}

	return result, ErrNotExists
}

func DeleteSchema(id primitive.ObjectID) (bool, error) {
	if id == primitive.NilObjectID {
		return false, errors.New("invalid updated ID")
	}

	result, err := repository.DeleteSchema(id)
	if err != nil {
		//TODO Log error
		log.Fatal(err)
		return false, err
	}

	return result, nil
}
