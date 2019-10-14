package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main()  {

	defer func() {
		fmt.Println(recover())
	}()
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "my-topic-A",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	err := r.SetOffset(0)
	if err != nil {
		fmt.Println(err)
	}

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
}