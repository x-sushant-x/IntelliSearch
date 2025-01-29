package main

import (
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/api"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/queue"
	"log"
)

func main() {

	kafkaQueue := queue.NewKafkaQueue("localhost:9092", "crawl_urls")

	server := api.NewServer("8080")

	log.Println("Crawl Manager Started")

	server.Start(kafkaQueue)
}
