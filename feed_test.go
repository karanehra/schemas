package schemas

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestGetFeeds(t *testing.T) {
	data, err := GetFeeds(TestDatabase, bson.D{})
	if err == nil {
		t.Error("Error in Query")
	}
	t.Log(data)
}
