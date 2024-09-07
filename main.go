package main

import (
	"net/http"
	"publisher/app"
	"publisher/controller"
	"publisher/repository"
	"publisher/service"
	"github.com/segmentio/kafka-go"
)

func main() {

	const (
		topic = "my-topic"
		brokerAddress = "localhost:9092"
	)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
	})
	defer writer.Close()

	db := app.NewDB()
	userRepository := repository.NewUserRepository(db)
	kafkaService := service.NewKafkaService(writer)
	loginController := controller.NewLoginController(userRepository, kafkaService)
	router := app.NewRouter(loginController)

	// var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodGet {
    //         loginController.LoginPageHandler(w, r)
    //     } else if r.Method == http.MethodPost {
    //         loginController.LoginHandler(w, r)
    //     }
	// }

	server := http.Server{
		Addr: "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
