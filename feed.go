package schemas

import (
	"context"
	"time"

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
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	URL   string             `json:"url" bson:"url"`
	Title string             `json:"title" bson:"title"`
	Tags  []string           `json:"tags" bson:"tags"`
}

// func CreateFeed(DB *mongo.Database, feed Feed)()

//GetFeeds returns feed docs
func GetFeeds(DB *mongo.Database, filter primitive.D) ([]FeedExtractor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := DB.Collection("articles").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []FeedExtractor
	for cur.Next(ctx) {
		var p FeedExtractor
		if err := cur.Decode(&p); err != nil {
			continue
		}
		results = append(results, p)
		break
	}
	return results, err
}
