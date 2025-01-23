package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
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
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{k.connAddr},
		Topic:     k.topic,
		Partition: k.partition,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("error while reading")
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
