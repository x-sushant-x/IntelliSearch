package main

import (
	"log"

	"github.com/x-sushant-x/IntelliSearch/crawler/core/queue"
)

const (
	kafkaTopic = "crawl_urls"
)

func main() {
	kafkaConsumer := queue.NewKafkaQueue("localhost:9092", kafkaTopic, 0)

	log.Println("Crawler Started")

	kafkaConsumer.Consume()

	select {}
}
