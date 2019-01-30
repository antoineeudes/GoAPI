package controllers

import (
	"API/models"
	u "API/utils"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

var Upload = func(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		u.Respond(w, u.Message(false, "Could not parse request. Please retry"))
	}

	params := mux.Vars(r)
	slug := params["gallery_slug"]

	associatedGallery, errDB := models.GetGalleryBySlug(slug)
	if errDB != nil {
		u.Respond(w, u.Message(false, "Gallery not found"))
		return
	}

	uploadedFile, handler, err := r.FormFile("file_upload")
	if err != nil {
		u.Respond(w, u.Message(false, "Could not receive file"))
		return
	}
	defer uploadedFile.Close()

	extension := filepath.Ext(handler.Filename)
	encodedFileName := base64.StdEncoding.EncodeToString([]byte(handler.Filename+time.Now().String())) + extension
	path := "/go/src/API/uploads/" + encodedFileName
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		u.Respond(w, u.Message(false, "Could not open file on server"))
		return
	}
	defer f.Close()

	_, errCopy := io.Copy(f, uploadedFile)
	if errCopy != nil {
		u.Respond(w, u.Message(false, "Could not copy file on server"))
		return
	}

	file := models.File{
		Type:      extension,
		GalleryID: associatedGallery.ID,
		Url:       "/files/" + encodedFileName,
	}

	fmt.Println(extension)
	resp := file.Create()

	u.Respond(w, resp)
}

var FirstFile = func(w http.ResponseWriter, r *http.Request) {
	resp := models.GetFirstFile()
	u.Respond(w, resp)
}

var AllFilesFromFirstGallery = func(w http.ResponseWriter, r *http.Request) {
	resp := models.GetImagesFromFirstGallery()
	u.Respond(w, resp)
}

var AllFilesFromGalleryBySlug = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["gallery_slug"]
	resp := models.GetImagesFromGalleryBySlug(slug)
	u.Respond(w, resp)
}
