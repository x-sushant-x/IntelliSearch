package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/api"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/cache"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/database"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/queue"
)

const (
	kafkaTopic = "crawl_urls"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load env file")
	}

	log.Println("ENV Loaded")
}

func main() {
	mongoDB := database.NewMongoDBConnection()
	redisCache := cache.NewRedisCache()

	kafkaQueue := queue.NewKafkaQueue("localhost:9092", kafkaTopic, mongoDB)

	core.NewElasticClient()

	server := api.NewServer("8080")

	go kafkaQueue.ConsumeCrawledPages()

	server.Start(kafkaQueue, redisCache)
}
