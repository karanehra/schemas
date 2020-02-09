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

var testDatabase *mongo.Database

func init() {
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
	testDatabase = mongoClient.Database("schemaTestDB")
	fmt.Println("Database Connection Success")
	fmt.Println("Purging database")
	coll := testDatabase.Collection("feeds")
	coll.DeleteMany(context, bson.D{})
	coll.DeleteMany(context, bson.D{})
	fmt.Println("Purging database")
}

func TestCreateProcess(t *testing.T) {
	testProcess := Process{}

	_, err := CreateProcess(testDatabase, testProcess)

	if err == nil || err.Error() != "Process name is required" {
		t.Error("Creation should throw error")
	}
	testProcess.Name = "New Process"
	_, err = CreateProcess(testDatabase, testProcess)
	if err == nil || err.Error() != "Process type is required" {
		t.Error("Creation should throw error", err)
	}

	testProcess.Type = "RANDOM_TASK"

	_, err = CreateProcess(testDatabase, testProcess)
	if err == nil || err.Error() != "Process type is required" {
		t.Error("Creation should throw error", err)
	}
}
