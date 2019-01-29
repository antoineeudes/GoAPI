package controllers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var Upload = func(w http.ResponseWriter, r *http.Request) {
	file, handler, err1 := r.FormFile("file_upload")
	if err1 != nil {
		fmt.Println("Could not receive file")
	}
	defer file.Close()
	encodedFileName := base64.StdEncoding.EncodeToString([]byte(handler.Filename+time.Now().String())) + filepath.Ext(handler.Filename)
	f, err2 := os.OpenFile("/go/src/API/uploads/"+encodedFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		fmt.Println("Could not save file on server")
	}
	defer f.Close()
	io.Copy(f, file)
}
