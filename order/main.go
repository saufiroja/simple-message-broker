package main

import (
	"encoding/json"
	"log"
	"microservice/simple-message-broker/order/producer"
	"microservice/simple-message-broker/utils"
)

func main() {
	data := Order{
		ID:    "1",
		Name:  "Order 1",
		Price: 1000,
		Qty:   2,
	}

	producer, err := producer.NewProducer()
	if err != nil {
		log.Println("Error while creating producer")
		panic(err)
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Error while marshalling data")
		panic(err)
	}

	for i := 0; i < 10; i++ {
		err = producer.SendMessage(utils.OrderTopic, string(bytes))
		if err != nil {
			log.Println("Error while sending message")
			panic(err)
		}
	}
}

type Order struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}
