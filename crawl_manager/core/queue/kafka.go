package queue

import (
	"context"
	"fmt"

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

	producer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: []byte(fmt.Sprint(data)),
	})

	return nil
}
