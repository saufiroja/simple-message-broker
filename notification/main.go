package main

import (
	"fmt"
	"microservice/simple-message-broker/notification/consumer"
	"microservice/simple-message-broker/utils"
)

func main() {
	consume, err := consumer.NewConsumer()
	if err != nil {
		panic(err)
	}

	partitionConsumer, err := consume.Consume(utils.OrderTopic)
	if err != nil {
		panic(err)
	}

	for {
		msg := <-partitionConsumer.Messages()
		fmt.Println(string(msg.Value))
	}
}
