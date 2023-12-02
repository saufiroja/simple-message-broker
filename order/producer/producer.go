package producer

import (
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	Producer sarama.SyncProducer
}

func NewProducer() (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:29092"}, config)
	if err != nil {
		return nil, err
	}

	return &Producer{Producer: producer}, nil
}

func (p *Producer) SendMessage(topic, msg string) error {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		log.Printf("Send message failed, err: %v\n", err)
		return err
	}

	log.Printf("Send message to partition %d, offset %d\n", partition, offset)
	return nil
}
