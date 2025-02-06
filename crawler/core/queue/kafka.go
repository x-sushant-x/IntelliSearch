package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/x-sushant-x/IntelliSearch/crawler/core"
	"log"
	"os"
	"strings"
)

type KafkaQueue struct {
	topic     string
	partition int
	connAddr  string
}

func NewKafkaQueue(connAddr, topic string, partition int) *KafkaQueue {
	return &KafkaQueue{
		topic:     topic,
		partition: partition,
		connAddr:  connAddr,
	}
}

func (k *KafkaQueue) Consume() {
	log.Println("Kafka queue consuming...")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{k.connAddr},
		Topic:       k.topic,
		Partition:   k.partition,
		GroupID:     "crawler_group",
		StartOffset: kafka.LastOffset,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("error while consuming: " + err.Error())
			continue
		}

		url := string(m.Value)

		htmlContent := core.ScrapURL(url)

		crawledPage, err := core.ExtractContent(htmlContent, url)
		if err != nil {
			log.Println("error while extracting page content: " + err.Error())
			continue
		}

		marshal, err := json.Marshal(crawledPage)
		if err != nil {
			log.Println("error while marshalling page content: " + err.Error())
			continue
		}

		urlFilePath := strings.Replace(url, "/", "_", len(url))
		filePath := "/Users/sushantdhiman/GoLang/IntelliSearch/crawled_pages/" + urlFilePath + ".txt"

		f, err := os.Create(filePath)
		if err != nil {
			log.Println("error while creating file for page content: " + err.Error())
			continue
		}

		_, err = f.WriteString(string(marshal))
		if err != nil {
			log.Println("error while writing file for page content: " + err.Error())
			continue
		}

		k.Send("crawled_pages", "", filePath)

		log.Println("Page: " + url + " crawled successfully.")

	}
}

func (k *KafkaQueue) Send(topic, key string, data interface{}) {
	producer := &kafka.Writer{
		Addr:       kafka.TCP(k.connAddr),
		Topic:      topic,
		BatchBytes: 10485760,
	}

	err := producer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: []byte(fmt.Sprint(data)),
	})

	if err != nil {
		log.Printf("Failed to send crawled page back to crawl manager: %v", err)
	}
}
