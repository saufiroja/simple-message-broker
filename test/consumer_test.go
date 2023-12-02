package test

import (
	"fmt"
	"microservice/simple-message-broker/notification/consumer"
	"testing"

	"github.com/IBM/sarama"
	"github.com/IBM/sarama/mocks"
)

func TestConsume(t *testing.T) {
	t.Run("Consume OK", func(t *testing.T) {
		// membuat consumer mock
		mockedConsumer := mocks.NewConsumer(t, nil)
		// membuat expect consumer success atau berhasil mengambil pesan
		mockedConsumer.ExpectConsumePartition("test_topic", 0, sarama.OffsetOldest).YieldMessage(&sarama.ConsumerMessage{Value: []byte("hello world")})
		kafka := &consumer.Consumer{
			Consumer: mockedConsumer,
		}

		partitionConsumer, err := kafka.Consume("test_topic")
		if err != nil {
			t.Error(err)
		}

		msg := <-partitionConsumer.Messages()
		if string(msg.Value) != "hello world" {
			t.Errorf("Message should be 'hello world' but have: %s", string(msg.Value))
		}

		if err := partitionConsumer.Close(); err != nil {
			t.Error(err)
		}

		if err := mockedConsumer.Close(); err != nil {
			t.Error(err)
		}
	})

	t.Run("Consume NOK", func(t *testing.T) {
		mockedConsumer := mocks.NewConsumer(t, nil)
		// membuat expect consumer gagal mengambil pesan
		mockedConsumer.ExpectConsumePartition("test_topic", 0, sarama.OffsetOldest).YieldError(fmt.Errorf("Error"))
		kafka := &consumer.Consumer{
			Consumer: mockedConsumer,
		}

		_, err := kafka.Consume("test_topic")
		if err != nil {
			t.Error("this should be error")
		}

		if err := mockedConsumer.Close(); err != nil {
			t.Error(err)
		}
	})
}
