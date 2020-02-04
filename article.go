package schemas

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
	Article
	ID string `json:"_id" bson:"_id"`
}
