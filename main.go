package main

import (
	"embed"
	"net/http"
	"publisher/app"
	"publisher/controller"
	"publisher/repository"
	"publisher/service"

	"github.com/segmentio/kafka-go"
)

//go:embed templates
var templates embed.FS
func main() {

	const (
		topic = "my-topic"
		brokerAddress = "kafka:9092"
	)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
	})
	defer writer.Close()

	db := app.NewDB()
	userRepository := repository.NewUserRepository(db)
	kafkaService := service.NewKafkaService(writer)
	loginController := controller.NewLoginController(userRepository, kafkaService, &templates)
	router := app.NewRouter(loginController)

	server := http.Server{
		Addr: "0.0.0.0:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
