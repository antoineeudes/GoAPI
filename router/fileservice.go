package router

import (
	"API/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func AddFileServiceRoutes(router *mux.Router) {
	router.HandleFunc("/api/upload/{gallery_slug}", controllers.Upload).Methods("POST")
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("/go/src/API/uploads/")))).Methods("GET")
	router.HandleFunc("/api/files/first", controllers.FirstFile).Methods("GET")
	router.HandleFunc("/api/files/all", controllers.AllFilesFromFirstGallery).Methods("GET")
	router.HandleFunc("/api/gallery/getallfiles/{gallery_slug}", controllers.AllFilesFromGalleryBySlug).Methods("GET")
}
