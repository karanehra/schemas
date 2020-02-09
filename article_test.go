package schemas

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMain(t *testing.M) {
	fmt.Println("Setting Up test")
	databaseClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err := mongo.NewClient(databaseClientOptions)
	if err != nil {
		log.Fatal(err)
	}
	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = mongoClient.Connect(context)
	if err != nil {
		log.Fatal(err)
	}
	TestDatabase = mongoClient.Database("schemaTestDB")
	fmt.Println("Database Connection Success")
	fmt.Println("Purging database")
	coll := TestDatabase.Collection("process")
	coll.DeleteMany(context, bson.D{})
	coll = TestDatabase.Collection("feeds")
	coll.DeleteMany(context, bson.D{})
	fmt.Println("Purged database")
}

func TestGetArticles(t *testing.T) {

}
