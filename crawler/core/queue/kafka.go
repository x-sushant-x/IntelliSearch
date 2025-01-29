package queue

import (
	"context"
	"fmt"
	"github.com/x-sushant-x/IntelliSearch/crawler/core"
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
			log.Println("error while consuming: " + err.Error())
			break
		}

		url := string(m.Value)

		core.ScrapURL(url)

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
