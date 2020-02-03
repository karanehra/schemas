package schemas

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Feed acts as a container for interacting with feed objects.
//Acts as a parent to articles. Doesnt contain ID so that mongo can
//attach ID's on it's own to not impede indexing perf
type Feed struct {
	URL   string   `json:"url" bson:"url"`
	Title string   `json:"title" bson:"title"`
	Tags  []string `json:"tags" bson:"tags"`
}

//FeedExtractor adds an ID to the feed. Use for manipulating feeds
type FeedExtractor struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Feed
}

//GetFeeds returns feed docs
func GetFeeds(DB *mongo.Database, filter primitive.D) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := DB.Collection("feeds").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = cur.All(ctx, results)
	return results, err
}
