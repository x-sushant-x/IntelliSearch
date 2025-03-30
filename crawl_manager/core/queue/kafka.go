package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/core/database"
	"github.com/x-sushant-x/IntelliSearch/crawl_manager/models"
	"io"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	crawlURLsKafkaTopic = "crawl_urls"
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
	log.Println("Waiting for crawled pages from crawler...")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "crawled_pages",
		MaxBytes:    10485760,
		GroupID:     "crawled_pages_group",
		StartOffset: kafka.FirstOffset,
	})

	if _, err := r.FetchMessage(context.Background()); err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}

	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("error while consuming: " + err.Error())
			continue
		}

		crawledFilePath := message.Value

		file, err := os.Open(string(crawledFilePath))
		if err != nil {
			log.Println("error while opening file: " + err.Error())
			continue
		}

		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			log.Println("error while reading file: " + err.Error())
			continue
		}

		var crawledData models.CrawledPage

		err = json.Unmarshal(fileBytes, &crawledData)
		if err != nil {
			log.Println("error while marshalling file: " + err.Error())
			continue
		}

		q.mongoDB.SaveCrawledPage(&crawledData)

		for _, newURL := range crawledData.AssociatedURLs {
			err := q.Send(crawlURLsKafkaTopic, "", newURL)
			if err != nil {
				log.Println("unable to send newly discovered url to crawler")
			}
		}
	}
}
