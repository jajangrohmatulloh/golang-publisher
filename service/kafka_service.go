package service

type KafkaService interface {
	Publish(data string) (string, error)
}