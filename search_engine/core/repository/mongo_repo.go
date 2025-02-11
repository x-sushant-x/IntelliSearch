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

	index := mongo.IndexModel{
		Keys: bson.D{{Key: "metaData", Value: "text"}, {Key: "title", Value: "text"}},
	}

	_, err := coll.Indexes().CreateOne(context.TODO(), index)

	if err != nil {
		panic("unable to create index in database" + err.Error())
	}

}

func (db *MongoDBStore) Search(query string) (*[]models.SearchResponse, error) {
	var response []models.SearchResponse

	coll := db.client.Database("IntelliSense").Collection("crawled_pages")

	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: query}}}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Println("unable to get result from database: " + err.Error())
		return &response, err
	}

	var results []models.CrawledPage

	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Println("unable to parse result from database: " + err.Error())
		return &response, err
	}

	for _, doc := range results {
		resp := models.SearchResponse{
			Title:           doc.Title,
			MetaDescription: doc.MetaData,
			Url:             doc.Url,
		}
		response = append(response, resp)
	}

	if response == nil {
		response = []models.SearchResponse{}
	}

	return &response, nil
}
