package router

import (
	"API/controllers"

	"github.com/gorilla/mux"
)

func AddGalleryServiceRoutes(router *mux.Router) {
	router.HandleFunc("/api/gallery/new", controllers.CreateGallery).Methods("POST")
	router.HandleFunc("/api/gallery/first", controllers.FirstGallery).Methods("GET")
	router.HandleFunc("/api/gallery/slug", controllers.GalleryBySlug).Methods("GET")
}
