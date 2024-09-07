package service

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
)

type KafkaServiceImpl struct {
	writer *kafka.Writer
}

func NewKafkaService(kafkaWriter *kafka.Writer) KafkaService {
	return &KafkaServiceImpl{
		writer: kafkaWriter,
	}
}

func (service *KafkaServiceImpl) Publish(data string) (string, error) {
    eventMessage := kafka.Message{
		Value: []byte(data),
	}
    ctx := context.Background()
	err := service.writer.WriteMessages(ctx, eventMessage)

	var message string
	if err != nil {
        return message, errors.New("Failed to publish event")
	}
	message = "Successfully published event"
	return message, nil
}