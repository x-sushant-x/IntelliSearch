package database

import (
	"context"
	"log"

	"github.com/x-sushant-x/IntelliSearch/search_engine/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	MongoDbUri = "mongodb://localhost:27017/IntelliSense"
)

type MongoDBStore struct {
	client *mongo.Client
}

func NewMongoDBConnection() *MongoDBStore {
	client, err := mongo.Connect(options.Client().ApplyURI(MongoDbUri))
	if err != nil {
		panic("MongoDB Error: " + err.Error())
	}

	return &MongoDBStore{
		client: client,
	}
}

func (db *MongoDBStore) CreateIndexes() {
	coll := db.client.Database("IntelliSense").Collection("crawled_pages")

	titleModel := mongo.IndexModel{
		Keys: bson.D{{Key: "title", Value: "text"}},
	}

	metaDescriptionModel := mongo.IndexModel{
		Keys: bson.D{{Key: "metaData", Value: "text"}},
	}

	_, err := coll.Indexes().CreateOne(context.TODO(), titleModel)

	if err != nil {
		panic("unable to create index in database: " + err.Error())
	}

	_, err = coll.Indexes().CreateOne(context.TODO(), metaDescriptionModel)

	if err != nil {
		panic("unable to create index in database" + err.Error())
	}

}

func (db *MongoDBStore) SaveCrawledPage(query string) (*[]models.SearchResponse, error) {
	coll := db.client.Database("IntelliSense").Collection("crawled_pages")

	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: query}}}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Println("unable to get result from database: " + err.Error())
		return nil, err
	}

	var results []models.CrawledPage

	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Println("unable to parse result from database: " + err.Error())
		return nil, err
	}

	var response []models.SearchResponse

	for _, doc := range results {
		resp := models.SearchResponse{
			Title:           doc.Title,
			MetaDescription: doc.MetaData,
			Url:             doc.Url,
		}
		response = append(response, resp)
	}

	return &response, nil
}
