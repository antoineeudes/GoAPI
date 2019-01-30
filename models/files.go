package models

import (
	u "API/utils"

	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	Url       string `json:"url"`
	Type      string `json:"type"`
	GalleryID uint
}

func (file *File) Create() map[string]interface{} {
	associatedGallery, _ := GetGalleryById(file.GalleryID)
	associatedGallery.AddFile(file)
	GetDB().Create(file)
	GetDB().Model(associatedGallery).Association("Images").Append(file)

	if file.ID < 0 {
		return u.Message(false, "Failed to upload file, connection error.")
	}

	response := u.Message(true, "File has been created")
	response["file"] = file
	return response
}

func GetFirstFile() map[string]interface{} {

	file := &File{}
	GetDB().Table("files").Last(file)
	response := u.Message(true, "file found")
	response["file"] = file
	return response
}
