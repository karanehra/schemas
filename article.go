package schemas

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Article acts as a container for feed items. Each item comes from a
//single Feed parent which must be defined
type Article struct {
	Title           string `json:"title" bson:"title"`
	Content         string `json:"content" bson:"content"`
	Description     string `json:"description" bson:"description"`
	URL             string `json:"url" bson:"url"`
	FeedTitle       string `json:"feedTitle" bson:"feedTitle"`
	FeedDescription string `json:"feedDescription" bson:"feedDescription"`
	FeedURL         string `json:"feedURL" bson:"feedURL"`
	CreatedAt       int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt       int64  `json:"updatedAt" bson:"updatedAt"`
}

//ArticleExtractor with bson tags for extracting _id
type ArticleExtractor struct {
	Title           string `json:"title" bson:"title"`
	Content         string `json:"content" bson:"content"`
	Description     string `json:"description" bson:"description"`
	URL             string `json:"url" bson:"url"`
	FeedTitle       string `json:"feedTitle" bson:"feedTitle"`
	FeedDescription string `json:"feedDescription" bson:"feedDescription"`
	FeedURL         string `json:"feedURL" bson:"feedURL"`
	CreatedAt       int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt       int64  `json:"updatedAt" bson:"updatedAt"`
	ID              string `json:"_id" bson:"_id"`
}

//GetArticles returns feed docs
func GetArticles(DB *mongo.Database, filter primitive.D) ([]ArticleExtractor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := DB.Collection("articles").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []ArticleExtractor
	for cur.Next(ctx) {
		var p ArticleExtractor
		if err := cur.Decode(&p); err != nil {
			continue
		}
		results = append(results, p)
	}
	return results, err
}
