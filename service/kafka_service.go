package service

type KafkaService interface {
	Publish(eventMessage string, topic string)
}