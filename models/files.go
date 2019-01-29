package models

import (
	u "API/utils"

	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	Url    string `json:"url"`
	Type   string `json:"type"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

func (file *File) Create() map[string]interface{} {

	GetDB().Create(file)

	if file.ID <= 0 {
		return u.Message(false, "Failed to upload file, connection error.")
	}

	response := u.Message(true, "File has been created")
	response["file"] = file
	return response
}
