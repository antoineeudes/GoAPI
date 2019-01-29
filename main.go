package main

import (
	"API/config"
	"API/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/cars").Name("Index").HandlerFunc(controllers.CarsIndex)
	router.Methods("POST").Path("/cars").Name("Create").HandlerFunc(controllers.CarsCreate)
	router.Methods("GET").Path("/cars/{id}").Name("Show").HandlerFunc(controllers.CarsShow)
	router.Methods("PUT").Path("/cars/{id}").Name("Update").HandlerFunc(controllers.CarsUpdate)
	router.Methods("DELETE").Path("/cars/{id}").Name("DELETE").HandlerFunc(controllers.CarsDelete)
	return router
}

func main() {
	config.DatabaseInit()
	router := InitializeRouter()
	defer config.Db().Close()
	log.Fatal(http.ListenAndServe(":5000", router))
}
