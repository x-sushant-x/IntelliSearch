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
