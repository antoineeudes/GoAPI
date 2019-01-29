package router

import (
	"API/controllers"

	"github.com/gorilla/mux"
)

// import (
// 	"API/controllers"

// 	"github.com/gorilla/mux"
// )

// func AddCarsServiceRoutes(router *mux.Router) {
// 	router.Methods("GET").Path("/cars").Name("Index").HandlerFunc(controllers.CarsIndex)
// 	router.Methods("POST").Path("/cars").Name("Create").HandlerFunc(controllers.CarsCreate)
// 	router.Methods("GET").Path("/cars/{id}").Name("Show").HandlerFunc(controllers.CarsShow)
// 	router.Methods("PUT").Path("/cars/{id}").Name("Update").HandlerFunc(controllers.CarsUpdate)
// 	router.Methods("DELETE").Path("/cars/{id}").Name("DELETE").HandlerFunc(controllers.CarsDelete)
// }

func AddGalleryServiceRoutes(router *mux.Router) {
	router.HandleFunc("/api/gallery/new", controllers.CreateGallery).Methods("POST")
}
