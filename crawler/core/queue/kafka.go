package queue

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/x-sushant-x/IntelliSearch/crawler/core"
	"log"
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
		Brokers:   []string{k.connAddr},
		Topic:     k.topic,
		Partition: k.partition,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("error while consuming: " + err.Error())
			continue
		}

		url := string(m.Value)

		htmlContent := core.ScrapURL(url)

		crawledPage, err := core.ExtractContent(htmlContent)
		if err != nil {
			log.Println("error while extracting page content: " + err.Error())
			continue
		}

		k.Send("crawled_pages", "", crawledPage)
	}
}

func (k *KafkaQueue) Send(topic, key string, data interface{}) {
	producer := &kafka.Writer{
		Addr:  kafka.TCP(k.connAddr),
		Topic: topic,
	}

	err := producer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: []byte(fmt.Sprint(data)),
	})

	if err != nil {
		log.Printf("Failed to send crawled page back to crawl manager: %v", err)
	}
}
