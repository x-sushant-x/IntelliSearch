package main

import (
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/api"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/cache"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/database"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/queue"
)

const (
	kafkaTopic = "crawl_urls"
)

func main() {
	mongoDB := database.NewMongoDBConnection()
	cache.NewRedisCache()

	kafkaQueue := queue.NewKafkaQueue("localhost:9092", kafkaTopic, mongoDB)

	server := api.NewServer("8080")

	go kafkaQueue.ConsumeCrawledPages()

	server.Start(kafkaQueue)
}
