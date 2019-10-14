package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func main()  {
	// to produce messages
	topic := "my-topic-A"
	partition := 2


	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("one2!")},
		kafka.Message{Value: []byte("two2!")},
		kafka.Message{Value: []byte("three2!")},
	)

	conn.Close()
}