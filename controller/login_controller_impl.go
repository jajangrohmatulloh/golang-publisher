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
		fmt.Fprintln(w, "Username did not exist")
		return
	}

	isMatch := helper.CheckPassword(password, userData.Password)
	if isMatch {
		fmt.Fprintln(w, "Login successfull ")
	} else {
		fmt.Fprintln(w, "Password did not match")
		return
	}

	// cookie := http.Cookie{
    //     Name:     "my-cookie",
    //     Value:    "my-cookie",
    //     Path:     "/",
    //     Expires:  time.Now().Add(365 * 24 * time.Hour),
    //     HttpOnly: true,
    // }
    // http.SetCookie(w, &cookie)

	userEvent := &event.UserEvent{}
	userEvent.Id = userData.Id
	userEvent.Nama = userData.FirstName + " " + userData.LastName
	userEvent.LoginDatetime = now
	userEvent.Agents = r.UserAgent()

	eventData, err := helper.ToString(userEvent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, eventData)

	message, err := controller.KafkaService.Publish(eventData)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, message)

}

