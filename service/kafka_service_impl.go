package service

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaServiceImpl struct {
	producer *kafka.Producer
}

func NewKafkaService(kafkaProducer *kafka.Producer) KafkaService {
	return &KafkaServiceImpl{
		producer: kafkaProducer,
	}
}

func (service *KafkaServiceImpl) Publish(eventMessage string, topic string) {
	message := &kafka.Message{
        // TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        TopicPartition: kafka.TopicPartition{Topic: &topic},
        Value:          []byte(eventMessage),
    }

    // Publish the message to the Kafka topic
    err := service.producer.Produce(message, nil)
    if err != nil {
        fmt.Println("Failed to produce message: %s", err)
    }

	// Wait for the delivery report
	event := <-service.producer.Events()
    msg := event.(*kafka.Message)

    if msg.TopicPartition.Error != nil {
        fmt.Printf("Failed to deliver message: %v\n", msg.TopicPartition.Error)
    } else {
        fmt.Printf("Message delivered to topic %s [%d] at offset %v\n",
            *msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
    }

    // Flush the producer to ensure all events are delivered
    service.producer.Flush(15 * 1000)
}