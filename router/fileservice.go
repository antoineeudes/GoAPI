package router

import (
	"API/controllers"

	"github.com/gorilla/mux"
)

func AddFileServiceRoutes(router *mux.Router) {
	router.HandleFunc("/api/upload", controllers.Upload).Methods("POST")
}
