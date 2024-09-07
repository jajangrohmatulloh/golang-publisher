package main

import (
	"net/http"
	"publisher/app"
	"publisher/controller"
	"publisher/repository"
	"fmt"
	"publisher/service"
	// "golang.org/x/crypto/bcrypt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {

	// err := bcrypt.CompareHashAndPassword([]byte("$2y$10$8CL01N66k1d11jd/vrgdzOk.oIIenPFZ5RiMt/ac2Rw2rrnmamNU."), []byte("emptysecret"))
	// if err != nil {
	// 	fmt.Println("Password mismatch:", err)
	// }
	// fmt.Println("Password match!")

	// producer, err := kafka.NewProducer(&kafka.ConfigMap{
    //     "bootstrap.servers": "localhost:9092",
    // })

	// if err != nil {
    //     fmt.Println("Failed to create producer: %s", err)
    // }
    // defer producer.Close()


	db := app.NewDB()
	userRepository := repository.NewUserRepository(db)
	// kafkaService := service.NewKafkaService(producer)
	// loginController := controller.NewLoginController(userRepository, kafkaService)
	loginController := controller.NewLoginController(userRepository, kafkaService)

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
            loginController.LoginPageHandler(w, r)
        } else if r.Method == http.MethodPost {
            loginController.LoginHandler(w, r)
        }
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
