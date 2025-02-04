package main

import (
	"log"

	"github.com/x-sushant-x/IntelliSearch/crawler/core/queue"
)

func main() {
	kafkaConsumer := queue.NewKafkaQueue("localhost:9092", "crawl_urls", 0)

	log.Println("Crawler Started")

	kafkaConsumer.Consume()

	select {}
}
