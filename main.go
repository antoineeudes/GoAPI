package main

import (
	"API/middleware"
	"API/models"
	"API/router"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {

	router := router.InitializeRouter()
	defer models.GetDB().Close()
	router.Use(middleware.JwtAuthentication)
	log.Fatal(http.ListenAndServe(":5000", handlers.LoggingHandler(os.Stdout, router)))
}
