package controller

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"publisher/helper"
	"publisher/model/event"
	"publisher/repository"
	"publisher/service"
	"time"
)

type LoginControllerImpl struct {
	UserRepository repository.UserRepository
	KafkaService service.KafkaService
}

func NewLoginController(userRepository repository.UserRepository, kafkaService service.KafkaService) LoginController {
	return &LoginControllerImpl{
		UserRepository: userRepository,
		KafkaService: kafkaService,
	}
}

func (controller *LoginControllerImpl) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.gohtml"))
    tmpl.Execute(w, nil)
}

func (controller *LoginControllerImpl) LoginHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	username := r.FormValue("username")
    password := r.FormValue("password")
	
	ctx := context.Background()
	userData, err := controller.UserRepository.FindByUsername(ctx, username)
	if err != nil {
		fmt.Fprintf(w, "username did not exist")
	}

	isMatch := helper.CheckPassword(password, userData.Password)
	if isMatch != true {
		fmt.Fprintf(w, "Login unsuccessfull")
	}

	userEvent := &event.UserEvent{}
	userEvent.Id = userData.Id
	userEvent.Nama = userData.FirstName + " " + userData.LastName
	userEvent.LoginDatetime = now
	userEvent.Agent = r.UserAgent()

	eventMessage, err := helper.ToString(userEvent)
	if err != nil {
		fmt.Println(err)
	}

	controller.KafkaService.Publish(eventMessage, "my-topic")

	fmt.Fprintf(w, "Login successfull")
}

