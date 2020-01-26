package schemas

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validProcesses []string = []string{
	"UPDATE_FEEDS",
	"CHECK_FOR_FEEDS",
	"CHECK_FOR_PROCESS",
	"DUMP_FEEDS",
}

//Process defines a task to be sent to a processor
type Process struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"processName" bson:"processName"`
	Status    string             `json:"status" bson:"status"`
	Type      string             `json:"type" bson:"type"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64              `json:"updateAt" bson:"updateAt"`
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
	if !isValidProcess(process.Type) {
		return nil, errors.New("Invalid process type")
	}
	process.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	process.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
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

//DeleteProcess removes a process from DB
func DeleteProcess(DB *mongo.Database, processID string) error {
	coll := DB.Collection("process")
	objectID, _ := primitive.ObjectIDFromHex(processID)
	_, err := coll.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}

func isValidProcess(process string) bool {
	for i := range validProcesses {
		if process == validProcesses[i] {
			return true
		}
	}
	return false
}
