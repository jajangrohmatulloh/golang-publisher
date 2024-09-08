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
	"embed"
)

type LoginControllerImpl struct {
	UserRepository repository.UserRepository
	KafkaService service.KafkaService
	tmpl *template.Template
}

func NewLoginController(userRepository repository.UserRepository, kafkaService service.KafkaService, templates *embed.FS) LoginController {
	return &LoginControllerImpl{
		UserRepository: userRepository,
		KafkaService: kafkaService,
		tmpl: template.Must(template.ParseFS(templates, "templates/*.gohtml")),
	}
}

func (controller *LoginControllerImpl) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
    controller.tmpl.Execute(w, "login.gohtml")
}

func (controller *LoginControllerImpl) LoginHandler(w http.ResponseWriter, r *http.Request) {
	wibTimeZone, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(wibTimeZone).Format("2006-01-02 15:04:05 -0700")
	
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

