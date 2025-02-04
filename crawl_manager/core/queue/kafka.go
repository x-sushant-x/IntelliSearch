package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/database"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/models"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaQueue struct {
	topic    string
	connAddr string
	mongoDB  database.DB
}

func NewKafkaQueue(connAddr, topic string, db database.DB) *KafkaQueue {
	return &KafkaQueue{
		topic:    topic,
		connAddr: connAddr,
		mongoDB:  db,
	}
}

func (q KafkaQueue) Send(topic, key string, data interface{}) error {
	producer := &kafka.Writer{
		Addr:  kafka.TCP(q.connAddr),
		Topic: q.topic,
	}

	err := producer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: []byte(fmt.Sprint(data)),
	})
	if err != nil {
		log.Printf("Failed to send message to topic %s: %v", topic, err)
		return err
	}

	return nil
}

func (q KafkaQueue) ConsumeCrawledPages() {
	fmt.Println("Waiting for crawled pages from crawler...")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "crawled_pages",
	})

	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("error while consuming: " + err.Error())
			continue
		}

		msgBytes := message.Value
		crawledPage := models.CrawledPage{}
		err = json.Unmarshal(msgBytes, &crawledPage)
		if err != nil {
			log.Println("error while un-marshalling crawled page: " + err.Error())
			continue
		}

		q.mongoDB.SaveCrawledPage(&crawledPage)
	}
}
