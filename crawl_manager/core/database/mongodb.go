package database

import (
	"context"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
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

func (db *MongoDBStore) SaveCrawledPage(page *models.CrawledPage) {
	coll := db.client.Database("IntelliSense").Collection("crawled_pages")

	_, err := coll.InsertOne(context.Background(), page)
	if err != nil {
		log.Println("Unable to save page to database: " + err.Error())
	}
}
