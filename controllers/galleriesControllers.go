package controllers

import (
	"API/models"
	u "API/utils"
	"encoding/json"
	"net/http"
)

var CreateGallery = func(w http.ResponseWriter, r *http.Request) {

	gallery := &models.Gallery{}
	err := json.NewDecoder(r.Body).Decode(gallery) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := gallery.Create() //Create account
	u.Respond(w, resp)
}

var FirstGallery = func(w http.ResponseWriter, r *http.Request) {
	resp := models.GetFirstGallery()
	u.Respond(w, resp)
}

var GalleryBySlug = func(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["slug"]

	if !ok || len(keys[0]) < 1 {
		u.Respond(w, u.Message(false, "Url Param 'slug' is missing"))
		return
	}

	slug := keys[0]
	gallery, errDB := models.GetGalleryBySlug(slug)

	if errDB != nil {
		u.Respond(w, u.Message(false, "Gallery not found"))
		return
	}

	response := u.Message(true, "Gallery found")
	response["gallery"] = gallery
	u.Respond(w, response)
}
