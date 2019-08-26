package main

import (
	"fmt"
	"kafka"
	"github.com/bsm/sarama-cluster"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8085"}
var topic = "schematry"

func main() {
	
	consumerCallbacks := kafka.ConsumerCallbacks{
		OnDataReceived: func(msg kafka.Message) {
			fmt.Println(msg)
		},
		OnError: func(err error) {
			fmt.Println("Consumer error", err)
		},
		OnNotification: func(notification *cluster.Notification) {
			fmt.Println(notification)
		},
	}
	
	consumer, err := kafka.NewAvroConsumer(kafkaServers, schemaRegistryServers, topic, "consumer-group", consumerCallbacks)
	
	if err != nil {
		fmt.Println(err)
	}
	
	consumer.Consume()
	
}
