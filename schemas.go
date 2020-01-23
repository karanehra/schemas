package schemas

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Process defines a task to be sent to a processor
type Process struct {
	Name   string `json:"processName"`
	Status string `json:"status"`
	Type   string `json:"type"`
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
	if process.Name == "" {
		return nil, errors.New("Process name is required")
	}
	if process.Status == "" {
		return nil, errors.New("Process status is required")
	}
	if process.Type == "" {
		return nil, errors.New("Process type is required")
	}
	return coll.InsertOne(context.TODO(), process)
}

//GetNewProcess returns an un-intiated process
func GetNewProcess(DB *mongo.Database) Process {
	coll := DB.Collection("process")
	process := Process{}
	document := coll.FindOne(context.Background(), bson.M{"status": "CREATED"})
	document.Decode(&process)
	return process
}

//UpdateProcessStatus changes process status against the process with the provided ID
func UpdateProcessStatus(DB *mongo.Database, status, processID string) (*mongo.UpdateResult, error) {
	coll := DB.Collection("process")
	objectID, _ := primitive.ObjectIDFromHex(processID)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"status": status}}
	return coll.UpdateOne(context.TODO(), filter, update)
}
