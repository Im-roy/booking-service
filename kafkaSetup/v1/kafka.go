package v1

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const KAFKA_BROKER_URL = "pkc-ymrq7.us-east-2.aws.confluent.cloud:9092"

func getWriteConnection() *kafka.Writer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(KAFKA_BROKER_URL),
		Topic:                  "topic-A",
		AllowAutoTopicCreation: true,
	}
	return w
}

func WriteMessageOnTopic() error {
	w := getWriteConnection()
	fmt.Println(w)
	messages := []kafka.Message{
		{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	}

	var err error
	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// attempt to create topic prior to publishing the message
		err = w.WriteMessages(ctx, messages...)
		if errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Second * 2)
			continue
		}

		if err != nil {
			log.Fatalf("unexpected error %v", err)
		}
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	return nil
}
