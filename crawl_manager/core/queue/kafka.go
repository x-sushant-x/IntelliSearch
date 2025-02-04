package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaQueue struct {
	topic    string
	connAddr string
}

func NewKafkaQueue(connAddr, topic string) *KafkaQueue {
	return &KafkaQueue{
		topic:    topic,
		connAddr: connAddr,
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

		fmt.Println("Page Crawled: " + string(message.Value))
	}
}
