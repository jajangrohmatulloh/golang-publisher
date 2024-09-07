package app

import (
	"net/http"
	"publisher/controller"
)

func NewRouter(loginController controller.LoginController) *http.ServeMux { 
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
				w.Header().Set("Location", "/login")
				w.WriteHeader(http.StatusSeeOther)
		}
			loginController.LoginHandler(w, r)
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			loginController.LoginPageHandler(w, r)
		}
	})

	return mux
}

