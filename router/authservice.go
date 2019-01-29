package router

import (
	"API/controllers"

	"github.com/gorilla/mux"
)

func AddAuthServiceRoutes(router *mux.Router) {
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
}
