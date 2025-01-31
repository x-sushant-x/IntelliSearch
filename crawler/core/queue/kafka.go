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
		Brokers:     []string{k.connAddr},
		Topic:       k.topic,
		Partition:   k.partition,
		StartOffset: kafka.FirstOffset,
	})

	for {
		fmt.Println("Waiting for message...")
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("error while consuming: " + err.Error())
			continue
		}

		url := string(m.Value)

		htmlContent := core.ScrapURL(url)

		_, _, err = core.ExtractTitleAndMetaData(htmlContent)

		if err != nil {
			log.Println("error while extracting title and metadata: " + err.Error())
			continue
		}
	}
}
