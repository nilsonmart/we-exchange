package infra

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteObject(collection string, id primitive.ObjectID) (bool, error) {

	client, shouldReturn, _, error := InitDatabase()
	if shouldReturn {
		return false, error
	}
	coll := client.Database("weschedule-app").Collection(collection)
	filter := bson.D{{Key: "_ID", Value: id}}

	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateObject(collection string, id primitive.ObjectID, data interface{}) (bool, error) {

	client, shouldReturn, _, error := InitDatabase()
	if shouldReturn {
		return false, error
	}

	filter := bson.D{{Key: "_ID", Value: id}}

	coll := client.Database("weschedule-app").Collection(collection)

	_, err := coll.UpdateByID(context.TODO(), filter, data)

	if err != nil {
		return false, err
	}
	return true, nil

}

func InsertObject(collection string, data interface{}) (bool, error) {
	client, shouldReturn, _, error := InitDatabase()
	if shouldReturn {
		return false, error
	}

	coll := client.Database("weschedule-app").Collection(collection)
	_, err := coll.InsertOne(context.TODO(), data)

	if err != nil {
		return false, err
	}
	return true, nil
}

func QueryAllObject(collection string) ([]byte, error) {
	client, shouldReturn, bytes, error := InitDatabase()
	if shouldReturn {
		return bytes, error
	}

	coll := client.Database("weschedule-app").Collection(collection)

	var result bson.M
	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found.")
		return nil, err
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Printf("Error when marshing json: %v", err)
		return nil, err
	}

	return jsonData, err
}

func QueryObject(collection string, filter primitive.D) ([]byte, error) {
	client, shouldReturn, bytes, err := InitDatabase()
	if shouldReturn {
		return bytes, err
	}

	coll := client.Database("weschedule-app").Collection(collection)

	var result bson.M
	err = coll.FindOne(context.TODO(), filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Println("No document was found.")
		return nil, err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Printf("Error when marshing json: %v", err)
		return nil, err
	}

	return jsonData, err
}
