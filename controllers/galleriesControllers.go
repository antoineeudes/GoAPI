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
