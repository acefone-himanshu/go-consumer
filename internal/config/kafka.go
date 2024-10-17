package config

import (
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

func GetKafkaReader() *kafka.Reader {

	kfakaBrokers := os.Getenv("KAFKA_BROKERS")
	topic := os.Getenv("KAFKA_CONSUMER_TOPIC")
	groupID := os.Getenv("KAFKA_GROUP_ID")

	brokers := strings.Split(kfakaBrokers, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
