package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("SUBSCRIPTION_HOST")},
		Topic:     os.Getenv("SUBSCRIPTION_TOPIC"),
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	defer func() {
		if err := r.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}()

	for {
		m, err := r.ReadMessage(context.Background())
		log.Println(err)
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}

// groupIDとか？考える
// https://github.com/pachyderm/pachyderm/blob/1b209a3a9f751c5115a6a7624234bf475fcd7cbf/etc/testing/kafka/main.go#L34
