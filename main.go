package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

var(
	kafkaTopic = "test"
	groupId = "groupid"
)

func main() {
	fmt.Println(`============================
===== Kafka Go Project =====
============================`)
	goReader()
	time.Sleep(200 * time.Millisecond)

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	})

	stdinReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a message to be sent to brokers topic:")
		str, _ := stdinReader.ReadString('\n')
		if str == "exit" { break }
		fmt.Println("sending... " + str)
		writer.WriteMessages(context.Background(), kafka.Message{Key: []byte("keyA"), Value: []byte(str)})
	}

	fmt.Println("=== finished ===")
}

func goReader(){
	reader := *kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  groupId,
		Topic:    kafkaTopic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer reader.Close()
	fmt.Println("start consuming in another thread...\n")
	go startReader(reader)
	fmt.Println("thread started...\n")
}

func startReader(reader kafka.Reader) {
	fmt.Println("  \t\t\tTopic\t\tPartition\tOffset\t\tKey\t\tValue")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil && err.Error() != "EOF" {
			fmt.Println(err.Error())
		}
		//if msg.Value != nil {
			fmt.Printf("> (%v)\t (%v)\t\t (%v)\t\t (%v)\t\t (%v)\t\t (%v)\n",
				time.Now().Format("15:04:05.000"), msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		//}
		time.Sleep(1000 * time.Millisecond)
	}
}
