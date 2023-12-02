package consumer

import "github.com/IBM/sarama"

type Consumer struct {
	Consumer sarama.Consumer
}

func NewConsumer() (*Consumer, error) {
	consumer, err := sarama.NewConsumer([]string{"localhost:29092"}, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{Consumer: consumer}, nil
}

func (c *Consumer) Consume(topic string) (sarama.PartitionConsumer, error) {
	return c.Consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
}
