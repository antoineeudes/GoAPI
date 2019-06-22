package router

import (
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	AddAuthServiceRoutes(router)
	AddGalleryServiceRoutes(router)
	AddFileServiceRoutes(router)
	return router
}
