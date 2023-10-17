package infra

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nilsonmart/we-exchange/internal/models"
)

func QueryObject(collection string, q *models.DBQuery) ([]byte, error) {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Printf("Not possible to get Env. Err: %s", err)
		return nil, err
	}
	db := os.Getenv("DATABASE")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db))
	if err != nil {
		fmt.Printf("Not possible to connect db. Err: %s", err)
		return nil, err
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Printf("Some error occured. Err: %s", err)
			panic(err)
		}
	}()

	coll := client.Database("weschedule-app").Collection(collection)

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: q.Key, Value: q.Value}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the value %s\n", q.Value)
		return nil, err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Printf("Error when marshing json: %v", err)
		return nil, err
	}

	return jsonData, err
}

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
