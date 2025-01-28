package main

import (
	"log"

	"github.com/x-sushant-x/IntelliSearch/crawler/core/queue"
)

func main() {
	kafkaConsumer := queue.NewKafkaQueue("localhost:9092", "IntelliSearch", 0)

	go kafkaConsumer.Consume()

	log.Println("Crawler Started")

	select {}
}
