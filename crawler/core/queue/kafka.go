package queue

import (
	"context"
	"fmt"
	"github.com/x-sushant-x/IntelliSearch/crawler/core"
	"log"
	"os"
	"path/filepath"
	"strings"

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

		wd, err := os.Getwd()
		if err != nil {
			log.Println("error while getting working directory: " + err.Error())
			continue
		}

		safeFilename := strings.ReplaceAll(url, "://", "_")
		safeFilename = strings.ReplaceAll(safeFilename, "/", "_")

		filePath := filepath.Join(wd, safeFilename+"_.html")

		file, err := os.Create(filePath)
		if err != nil {
			log.Println("error while creating html content file: " + err.Error())
			continue
		}

		_, err = file.Write([]byte(htmlContent))
		if err != nil {
			log.Println("error while saving html content: " + err.Error())
			continue
		}

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
