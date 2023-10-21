package main

import (
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/nilsonmart/we-exchange/internal/models"
	"github.com/nilsonmart/we-exchange/internal/repository"
)

// Account
func AllAccount() ([]models.Account, error) {

	model, err := repository.AllAccount()
	if err != nil {
		//TODO Log Error
		return nil, err
	}

	return model, nil
}

func GetAccountByID(id primitive.ObjectID) (*models.Account, error) {

	model, err := repository.GetAccountByID(id)
	if err != nil {
		//TODO Log Error
		return nil, err
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

	return false, err
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
		return nil, err
	}

	return model, nil
}

func GetRequestChangeByID(id string) (*models.RequestChange, error) {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	query, err := repository.GetRequestChangeByID(objID)
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, err
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

func UpdateRequestChange(id string, model models.RequestChange) (bool, error) {
	if id == "" {
		return false, errors.New("invalid updated ID")
	}

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return false, err
	}

	result, err := repository.UpdateRequestChange(objID, model)

	if err != nil {
		return false, err
	}

	return result, err
}

func DeleteRequestChange(id string) (bool, error) {
	if id == "" {
		return false, errors.New("invalid updated ID")
	}

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return false, err
	}

	result, err := repository.DeleteRequestChange(objID)
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
		return nil, err
	}

	return query, nil
}

func GetSchemaByUserID(id string) (*models.Schema, error) {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	query, err := repository.GetSchemaByUserID(objID)
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, err
	}

	return query, nil
}

func UpdateSchema(id string, model models.Schema) (bool, error) {
	if id == "" {
		return false, errors.New("invalid updated ID")
	}

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return false, err
	}

	result, err := repository.UpdateSchema(objID, model)

	if err != nil {
		return false, err
	}

	return result, err
}
