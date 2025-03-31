package main

import (
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/api"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/database"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/queue"
	"log"
)

const (
	kafkaTopic = "crawl_urls"
)

func main() {
	mongoDB := database.NewMongoDBConnection()

	kafkaQueue := queue.NewKafkaQueue("localhost:9092", kafkaTopic, mongoDB)

	server := api.NewServer("8080")

	log.Println("Crawl Manager Started")

	go kafkaQueue.ConsumeCrawledPages()

	server.Start(kafkaQueue)
}
