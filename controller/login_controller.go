package controller

import "net/http"

type LoginController interface {
	LoginPageHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
}