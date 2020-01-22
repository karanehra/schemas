package schemas

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Test defines a test struct
type Test struct {
	Field string
}

//Process defines a task to be sent to a processor
type Process struct {
	Name   string `json:"processName"`
	Status string `json:"status"`
}

//GetAllProcesses fetches all processes from the database
func GetAllProcesses(DB *mongo.Database) ([]bson.M, error) {
	coll := DB.Collection("process")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = cur.All(context.TODO(), &results)

	return results, err
}

//CreateProcess adds a process to db
func CreateProcess(DB *mongo.Database, process Process) (*mongo.InsertOneResult, error) {
	coll := DB.Collection("process")
	return coll.InsertOne(context.TODO(), process)
}
