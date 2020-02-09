package schemas

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var TestDatabase *mongo.Database
var CreatedID primitive.ObjectID

func TestCreateProcess(t *testing.T) {
	testProcess := Process{}

	_, err := CreateProcess(TestDatabase, testProcess)

	if err == nil || err.Error() != "Process name is required" {
		t.Error("Creation should throw error")
	}
	testProcess.Name = "New Process"
	_, err = CreateProcess(TestDatabase, testProcess)
	if err == nil || err.Error() != "Process type is required" {
		t.Error("Creation should throw error")
	}

	testProcess.Type = "RANDOM_TASK"

	_, err = CreateProcess(TestDatabase, testProcess)
	if err == nil || err.Error() != "Invalid process type" {
		t.Error("Creation should throw error")
	}

	testProcess.Type = ValidProcesses[len(ValidProcesses)*rand.Intn(1)]
	data, err := CreateProcess(TestDatabase, testProcess)
	if err != nil {
		t.Error("Process creation should be success")
	}
	var ok bool
	CreatedID, ok = data.InsertedID.(primitive.ObjectID)
	if !ok {
		t.Error("Incorrect record ID")
	}
}

func TestGetAllProcesses(t *testing.T) {
	data, err := GetAllProcesses(TestDatabase)
	if err != nil {
		t.Error("Error during Find. Check DB call")
	}
	if len(data) != 1 {
		t.Error(fmt.Sprintf("Incorrect number of docs. Got %v, Expected %v", 1, len(data)))
	}
}

func TestGetNewProcess(t *testing.T) {
	data := GetNewProcess(TestDatabase)
	if data.Name != "New Process" {
		t.Error("unexpected process found")
	}
}

func TestUpdateProcessStatus(t *testing.T) {
	_, err := UpdateProcessStatus(TestDatabase, "NEW", CreatedID)
	if err != nil {
		t.Error("Error during update in DB layer")
	}

	data := ProcessExtractor{}

	newRes := TestDatabase.Collection("process").FindOne(context.TODO(), bson.M{"_id": CreatedID})
	newRes.Decode(&data)
	if data.Status != "NEW" {
		t.Error("Update not reflecting in DB")
	}
}

func TestDeleteProcess(t *testing.T) {
	err := DeleteProcess(TestDatabase, CreatedID)
	if err != nil {
		t.Error("Delete problem in DB call")
	}
	data, _ := GetAllProcesses(TestDatabase)
	if len(data) > 0 {
		t.Error("Delete not reflecting in DB")
	}
}
