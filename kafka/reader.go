package kafka

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
	"time"
)

var reader *kafka.Reader

func InitReader(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Reader, err error) {
	config := kafka.ReaderConfig{
		Brokers:         kafkaBrokerUrls,
		GroupID:         clientId,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}
	reader = kafka.NewReader(config)
	go func() {
		for {
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Error("error while receiving message: %s", err.Error())
				continue
			}
			value := m.Value
			if err != nil {
				log.Error("error while receiving message: %s", err.Error())
				continue
			}
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(value))
		}
	}()

	return reader, nil
}
