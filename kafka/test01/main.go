package main

import (
	"bufio"
	"context"
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

func main()  {
	// to produce messages
	topic := "my-topic-A"
	partition := 0


	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	conn.SetWriteDeadline(time.Now().Add(10*time.Second))

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {

		s := input.Text()
		conn.WriteMessages(
			kafka.Message{Value: []byte(s)},
		)
	}
	//conn.WriteMessages(
	//	kafka.Message{Value: []byte("one2!")},
	//	kafka.Message{Value: []byte("two2!")},
	//	kafka.Message{Value: []byte("three2!")},
	//)

	conn.Close()
}