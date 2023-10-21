package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDatabase() (*mongo.Client, bool, []byte, error) {
	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Printf("Not possible to get Env. Err: %s", err)
		return nil, true, nil, err
	}
	db := os.Getenv("DATABASE")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db))
	if err != nil {
		fmt.Printf("Not possible to connect db. Err: %s", err)
		return nil, true, nil, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Printf("Some error occured. Err: %s", err)
			panic(err)
		}
	}()
	return client, false, nil, nil
}
