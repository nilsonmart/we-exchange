package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/nilsonmart/we-exchange/internal/infra"
	"github.com/nilsonmart/we-exchange/internal/models"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

func queryAllObjectData(collection string) ([]byte, error) {
	query, err := infra.QueryAllObject(collection)
	if err != nil {
		return nil, err
	}
	return query, err
}

func queryObjectData(collection, key, value string) ([]byte, error) {
	filter := bson.D{{Key: key, Value: value}}
	query, err := infra.QueryObject(collection, filter)
	if err != nil {
		return nil, err
	}
	return query, err
}

func queryObjectDataByID(collection string, id primitive.ObjectID) ([]byte, error) {

	filter := bson.D{{Key: "_ID", Value: id}}
	query, err := infra.QueryObject(collection, filter)
	if err != nil {
		return nil, err
	}
	return query, err
}

// Account
func AllAccount() ([]models.Account, error) {

	query, err := queryAllObjectData("account")
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	var account []models.Account
	json.Unmarshal(query, &account)

	return account, nil
}

func GetAccountByID(id primitive.ObjectID) (*models.Account, error) {

	query, err := queryObjectDataByID("account", id)
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	var account models.Account
	json.Unmarshal(query, &account)

	return &account, nil
}

func ValidateAccount(email, password string) (bool, error) {
	if email == "" || password == "" {
		//TODO Log error
		log.Fatal("Email or Password invalid.")
		return false, errors.New("Email or Password invalid.")
	}

	query, err := queryObjectData("account", "_email", email)

	if err != nil {
		return false, err
	}

	var account models.Account
	json.Unmarshal(query, &account)

	if account.Email == email && account.Password == password {
		return true, nil
	}

	return false, ErrNotExists
}

// RequestChange

func CreateRequestChange(model models.RequestChange) (bool, error) {

	result, err := infra.InsertObject("requestchange", model)

	if err != nil {
		//TODO Log error
		log.Fatal(err)
		return false, err
	}

	return result, nil
}

func AllRequestChange() ([]models.RequestChange, error) {
	query, err := queryAllObjectData("requestchange")
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	var requestChange []models.RequestChange
	json.Unmarshal(query, &requestChange)

	return requestChange, nil
}

func GetRequestChangeByID(id primitive.ObjectID) (*models.RequestChange, error) {
	query, err := queryObjectDataByID("requestchange", id)
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	var requestChange models.RequestChange
	json.Unmarshal(query, &requestChange)

	return &requestChange, nil
}

func GetRequestChangeByUserID(userId string) (*models.RequestChange, error) {

	query, err := queryObjectData("requestChange", "_requestedfor", userId)

	if err != nil {
		return nil, err
	}

	var requestChange models.RequestChange
	json.Unmarshal(query, &requestChange)

	return &requestChange, nil
}

func UpdateRequestChange(id primitive.ObjectID, model models.RequestChange) (bool, error) {
	if id == primitive.NilObjectID {
		return false, errors.New("invalid updated ID")
	}

	result, err := infra.UpdateObject("requestchange", id, model)

	if err != nil {
		return false, err
	}

	return result, ErrNotExists
}

func DeleteRequestChange(id primitive.ObjectID) (bool, error) {
	if id == primitive.NilObjectID {
		return false, errors.New("invalid updated ID")
	}

	result, err := infra.DeleteObject("requestchange", id)
	if err != nil {
		//TODO Log error
		log.Fatal(err)
		return false, err
	}

	return result, nil
}

// SCHEMA

func CreateSchema(model models.Schema) (bool, error) {

	result, err := infra.InsertObject("schema", model)

	if err != nil {
		//TODO Log error
		log.Fatal(err)
		return false, err
	}

	return result, nil

}

func AllSchema() ([]models.Schema, error) {
	query, err := queryAllObjectData("schema")
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	var schema []models.Schema
	json.Unmarshal(query, &schema)

	return schema, nil
}

func GetSchemaByID(id primitive.ObjectID) (*models.Schema, error) {
	query, err := queryObjectDataByID("schema", id)
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	var schema models.Schema
	json.Unmarshal(query, &schema)

	return &schema, nil
}

func GetSchemaByUserID(id primitive.ObjectID) (*models.Schema, error) {
	query, err := queryObjectDataByID("schema", id)
	if err != nil {
		//TODO Log Error
		fmt.Printf("Result for query: %v and error: %v", query, err)
		return nil, ErrNotExists
	}

	var schema models.Schema
	json.Unmarshal(query, &schema)

	return &schema, nil
}

func UpdateSchema(id primitive.ObjectID, model models.Schema) (bool, error) {
	if id == primitive.NilObjectID {
		return false, errors.New("invalid updated ID")
	}

	result, err := infra.UpdateObject("account", id, model)

	if err != nil {
		return false, err
	}

	return result, ErrNotExists
}

func DeleteSchema(id primitive.ObjectID) (bool, error) {
	if id == primitive.NilObjectID {
		return false, errors.New("invalid updated ID")
	}

	result, err := infra.DeleteObject("schema", id)
	if err != nil {
		//TODO Log error
		log.Fatal(err)
		return false, err
	}

	return result, nil
}
